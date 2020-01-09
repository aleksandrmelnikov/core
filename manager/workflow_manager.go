package manager

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/onepanelio/core/kube"
	"github.com/onepanelio/core/model"
	"github.com/onepanelio/core/util"
	"google.golang.org/grpc/codes"
)

var (
	labelKeyPrefix                  = os.Getenv("KUBE_LABEL_KEY_PREFIX")
	workflowTemplateUIDLabelKey     = labelKeyPrefix + "workflow-template-uid"
	workflowTemplateVersionLabelKey = labelKeyPrefix + "workflow-template-version"
)

func (r *ResourceManager) CreateWorkflow(namespace string, workflow *model.Workflow) (*model.Workflow, error) {
	workflowTemplate, err := r.GetWorkflowTemplate(namespace, workflow.WorkflowTemplate.UID, workflow.WorkflowTemplate.Version)
	if err != nil {
		return nil, err
	}

	// TODO: Need to pull system parameters from k8s config/secret here, example: HOST
	opts := &kube.WorkflowOptions{}
	for _, param := range workflow.Parameters {
		opts.Parameters = append(opts.Parameters, kube.WorkflowParameter{
			Name:  param.Name,
			Value: param.Value,
		})
	}
	if opts.Labels == nil {
		opts.Labels = &map[string]string{}
	}
	(*opts.Labels)[workflowTemplateUIDLabelKey] = workflowTemplate.UID
	(*opts.Labels)[workflowTemplateVersionLabelKey] = fmt.Sprint(workflowTemplate.Version)
	createdWorkflows, err := r.kubeClient.CreateWorkflow(namespace, workflowTemplate.GetManifestBytes(), opts)
	if err != nil {
		return nil, err
	}

	workflow.Name = createdWorkflows[0].Name
	workflow.UID = string(createdWorkflows[0].ObjectMeta.UID)
	workflow.WorkflowTemplate = workflowTemplate
	// Manifests could get big, don't return them in this case.
	workflow.WorkflowTemplate.Manifest = ""

	return workflow, nil
}

func (r *ResourceManager) GetWorkflow(namespace, name string) (workflow *model.Workflow, err error) {
	wf, err := r.kubeClient.GetWorkflow(namespace, name)
	if err != nil {
		return nil, util.NewUserError(codes.NotFound, "Workflow not found.")
	}

	uid := wf.ObjectMeta.Labels[workflowTemplateUIDLabelKey]
	version, err := strconv.ParseInt(
		wf.ObjectMeta.Labels[workflowTemplateVersionLabelKey],
		10,
		32,
	)
	if err != nil {
		return nil, util.NewUserError(codes.InvalidArgument, "Invalid version number.")
	}
	workflowTemplate, err := r.GetWorkflowTemplate(namespace, uid, int32(version))
	if err != nil {
		return
	}

	// TODO: Do we need to parse parameters into workflow.Parameters?
	status, err := json.Marshal(wf.Status)
	if err != nil {
		return nil, util.NewUserError(codes.InvalidArgument, "Invalid status.")
	}
	workflow = &model.Workflow{
		UID:              string(wf.UID),
		Name:             wf.Name,
		Status:           string(status),
		WorkflowTemplate: workflowTemplate,
	}

	return
}

func (r *ResourceManager) WatchWorkflow(namespace, name string) (<-chan *model.Workflow, error) {
	wf, err := r.GetWorkflow(namespace, name)
	if err != nil {
		return nil, util.NewUserError(codes.NotFound, "Workflow template not found.")
	}

	watcher, err := r.kubeClient.WatchWorkflow(namespace, name)
	if err != nil {
		return nil, util.NewUserError(codes.Unknown, "Unknown error.")
	}

	var workflow *kube.Workflow
	workflowWatcher := make(chan *model.Workflow)
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case next := <-watcher.ResultChan():
				workflow, _ = next.Object.(*kube.Workflow)
			case <-ticker.C:
			}

			if workflow == nil {
				continue
			}
			status, err := json.Marshal(workflow.Status)
			if err != nil {
				continue
			}
			workflowWatcher <- &model.Workflow{
				UID:              string(workflow.UID),
				Name:             workflow.Name,
				Status:           string(status),
				WorkflowTemplate: wf.WorkflowTemplate,
			}

			if !workflow.Status.FinishedAt.IsZero() {
				break
			}
		}
		close(workflowWatcher)
		watcher.Stop()
	}()

	return workflowWatcher, nil
}

func (r *ResourceManager) ListWorkflows(namespace, workflowTemplateUID string) (workflows []*model.Workflow, err error) {
	opts := &kube.WorkflowOptions{}
	if workflowTemplateUID != "" {
		opts.ListOptions = &kube.ListOptions{
			LabelSelector: fmt.Sprintf("%sworkflow-template-uid=%s", labelKeyPrefix, workflowTemplateUID),
		}
	}
	wfs, err := r.kubeClient.ListWorkflows(namespace, opts)
	if err != nil {
		return nil, util.NewUserError(codes.NotFound, "Workflows not found.")
	}

	for _, wf := range wfs {
		workflows = append(workflows, &model.Workflow{
			Name: wf.ObjectMeta.Name,
			UID:  string(wf.ObjectMeta.UID),
		})
	}

	return
}

func (r *ResourceManager) CreateWorkflowTemplate(namespace string, workflowTemplate *model.WorkflowTemplate) (*model.WorkflowTemplate, error) {
	// validate workflow template
	if err := r.kubeClient.ValidateWorkflow(workflowTemplate.GetManifestBytes()); err != nil {
		return nil, util.NewUserError(codes.InvalidArgument, err.Error())
	}

	workflowTemplate, err := r.workflowRepository.CreateWorkflowTemplate(namespace, workflowTemplate)
	if err != nil {
		return nil, util.NewUserErrorWrap(err, "Workflow template")
	}

	return workflowTemplate, nil
}

func (r *ResourceManager) CreateWorkflowTemplateVersion(namespace string, workflowTemplate *model.WorkflowTemplate) (*model.WorkflowTemplate, error) {
	// validate workflow template
	if err := r.kubeClient.ValidateWorkflow(workflowTemplate.GetManifestBytes()); err != nil {
		return nil, util.NewUserError(codes.InvalidArgument, err.Error())
	}

	workflowTemplate, err := r.workflowRepository.CreateWorkflowTemplateVersion(namespace, workflowTemplate)
	if err != nil {
		return nil, util.NewUserErrorWrap(err, "Workflow template")
	}
	if err == nil && workflowTemplate == nil {
		return nil, util.NewUserError(codes.NotFound, "Workflow template not found.")
	}

	return workflowTemplate, nil
}

func (r *ResourceManager) GetWorkflowTemplate(namespace, uid string, version int32) (workflowTemplate *model.WorkflowTemplate, err error) {
	workflowTemplate, err = r.workflowRepository.GetWorkflowTemplate(namespace, uid, version)
	if err != nil {
		return nil, util.NewUserError(codes.Unknown, "Unknown error.")
	}
	if err == nil && workflowTemplate == nil {
		return nil, util.NewUserError(codes.NotFound, "Workflow template not found.")
	}

	return
}

func (r *ResourceManager) ListWorkflowTemplateVersions(namespace, uid string) (workflowTemplateVersions []*model.WorkflowTemplate, err error) {
	workflowTemplateVersions, err = r.workflowRepository.ListWorkflowTemplateVersions(namespace, uid)
	if err != nil {
		return nil, util.NewUserError(codes.NotFound, "Workflow template versions not found.")
	}

	return
}

func (r *ResourceManager) ListWorkflowTemplates(namespace string) (workflowTemplateVersions []*model.WorkflowTemplate, err error) {
	workflowTemplateVersions, err = r.workflowRepository.ListWorkflowTemplates(namespace)
	if err != nil {
		return nil, util.NewUserError(codes.NotFound, "Workflow templates not found.")
	}

	return
}