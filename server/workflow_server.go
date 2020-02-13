package server

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/onepanelio/core/api"
	"github.com/onepanelio/core/manager"
	"github.com/onepanelio/core/model"
	"github.com/onepanelio/core/util/ptr"
)

type WorkflowServer struct {
	resourceManager *manager.ResourceManager
}

func NewWorkflowServer(resourceManager *manager.ResourceManager) *WorkflowServer {
	return &WorkflowServer{resourceManager: resourceManager}
}

func apiWorkflow(wf *model.Workflow) (workflow *api.Workflow) {
	workflow = &api.Workflow{
		CreatedAt:  wf.CreatedAt.Format(time.RFC3339),
		Name:       wf.Name,
		Uid:        wf.UID,
		Phase:      string(wf.Phase),
		StartedAt:  wf.CreatedAt.Format(time.RFC3339),
		FinishedAt: wf.FinishedAt.Format(time.RFC3339),
		Manifest:   wf.Manifest,
	}

	if wf.WorkflowTemplate != nil {
		workflow.WorkflowTemplate = &api.WorkflowTemplate{
			Uid:        wf.WorkflowTemplate.UID,
			CreatedAt:  wf.WorkflowTemplate.CreatedAt.UTC().Format(time.RFC3339),
			Name:       wf.WorkflowTemplate.Name,
			Version:    wf.WorkflowTemplate.Version,
			Manifest:   wf.WorkflowTemplate.Manifest,
			IsLatest:   wf.WorkflowTemplate.IsLatest,
			IsArchived: wf.WorkflowTemplate.IsArchived,
		}
	}

	return
}

func apiWorkflowTemplate(wft *model.WorkflowTemplate) *api.WorkflowTemplate {
	return &api.WorkflowTemplate{
		Uid:        wft.UID,
		CreatedAt:  wft.CreatedAt.UTC().Format(time.RFC3339),
		Name:       wft.Name,
		Version:    wft.Version,
		Manifest:   wft.Manifest,
		IsLatest:   wft.IsLatest,
		IsArchived: wft.IsArchived,
	}
}

func (s *WorkflowServer) CreateWorkflow(ctx context.Context, req *api.CreateWorkflowRequest) (*api.Workflow, error) {
	workflow := &model.Workflow{
		WorkflowTemplate: &model.WorkflowTemplate{
			UID:     req.Workflow.WorkflowTemplate.Uid,
			Version: req.Workflow.WorkflowTemplate.Version,
		},
	}
	for _, param := range req.Workflow.Parameters {
		workflow.Parameters = append(workflow.Parameters, model.Parameter{
			Name:  param.Name,
			Value: ptr.String(param.Value),
		})
	}

	wf, err := s.resourceManager.CreateWorkflow(req.Namespace, workflow)
	if err != nil {
		if errors.As(err, &userError) {
			return nil, userError.GRPCError()
		}
	}

	return apiWorkflow(wf), nil
}

func (s *WorkflowServer) GetWorkflow(ctx context.Context, req *api.GetWorkflowRequest) (*api.Workflow, error) {
	wf, err := s.resourceManager.GetWorkflow(req.Namespace, req.Name)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}

	return apiWorkflow(wf), nil
}

func (s *WorkflowServer) WatchWorkflow(req *api.WatchWorkflowRequest, stream api.WorkflowService_WatchWorkflowServer) error {
	watcher, err := s.resourceManager.WatchWorkflow(req.Namespace, req.Name)
	if errors.As(err, &userError) {
		return userError.GRPCError()
	}

	wf := &model.Workflow{}
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case wf = <-watcher:
		case <-ticker.C:
		}

		if wf == nil {
			break
		}
		if err := stream.Send(apiWorkflow(wf)); err != nil {
			return err
		}
	}

	return nil
}

func (s *WorkflowServer) GetWorkflowLogs(req *api.GetWorkflowLogsRequest, stream api.WorkflowService_GetWorkflowLogsServer) error {
	watcher, err := s.resourceManager.GetWorkflowLogs(req.Namespace, req.Name, req.PodName, req.ContainerName)
	if errors.As(err, &userError) {
		return userError.GRPCError()
	}

	le := &model.LogEntry{}
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case le = <-watcher:
		case <-ticker.C:
		}

		if le == nil {
			break
		}
		if err := stream.Send(&api.LogEntry{
			Timestamp: le.Timestamp.String(),
			Content:   le.Content,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (s *WorkflowServer) GetWorkflowMetrics(ctx context.Context, req *api.GetWorkflowMetricsRequest) (*api.GetWorkflowMetricsResponse, error) {
	metrics, err := s.resourceManager.GetWorkflowMetrics(req.Namespace, req.Name, req.PodName)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}

	return &api.GetWorkflowMetricsResponse{Metrics: *metrics}, nil
}

func (s *WorkflowServer) ListWorkflows(ctx context.Context, req *api.ListWorkflowsRequest) (*api.ListWorkflowsResponse, error) {
	if req.PageSize <= 0 {
		req.PageSize = 15
	}

	if req.Page <= 0 {
		req.Page = 1
	}

	workflows, err := s.resourceManager.ListWorkflows(req.Namespace, req.WorkflowTemplateUid, req.WorkflowTemplateVersion)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}

	apiWorkflows := make([]*api.Workflow, 0)
	for _, wf := range workflows {
		apiWorkflows = append(apiWorkflows, apiWorkflow(wf))
	}

	pages := int32(math.Ceil(float64(len(apiWorkflows)) / float64(req.PageSize)))
	if req.Page > pages {
		req.Page = pages
	}

	start := (req.Page - 1) * req.PageSize
	end := start + req.PageSize
	if end >= int32(len(apiWorkflows)) {
		end = int32(len(apiWorkflows)) - 1
	}

	return &api.ListWorkflowsResponse{
		Count:      end - start,
		Workflows:  apiWorkflows[start:end],
		Page:       req.Page,
		Pages:      pages,
		TotalCount: int32(len(apiWorkflows)),
	}, nil
}

func (s *WorkflowServer) ResubmitWorkflow(ctx context.Context, req *api.ResubmitWorkflowRequest) (*api.Workflow, error) {
	wf, err := s.resourceManager.ResubmitWorkflow(req.Namespace, req.Name)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}

	return apiWorkflow(wf), nil
}

func (s *WorkflowServer) TerminateWorkflow(ctx context.Context, req *api.TerminateWorkflowRequest) (*empty.Empty, error) {
	err := s.resourceManager.TerminateWorkflow(req.Namespace, req.Name)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}

	return &empty.Empty{}, nil
}

func (s *WorkflowServer) CreateWorkflowTemplate(ctx context.Context, req *api.CreateWorkflowTemplateRequest) (*api.WorkflowTemplate, error) {
	workflowTemplate := &model.WorkflowTemplate{
		Name:     req.WorkflowTemplate.Name,
		Manifest: req.WorkflowTemplate.Manifest,
	}
	workflowTemplate, err := s.resourceManager.CreateWorkflowTemplate(req.Namespace, workflowTemplate)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}
	req.WorkflowTemplate.Uid = workflowTemplate.UID
	req.WorkflowTemplate.Version = workflowTemplate.Version

	return req.WorkflowTemplate, nil
}

func (s *WorkflowServer) CreateWorkflowTemplateVersion(ctx context.Context, req *api.CreateWorkflowTemplateRequest) (*api.WorkflowTemplate, error) {
	workflowTemplate := &model.WorkflowTemplate{
		UID:      req.WorkflowTemplate.Uid,
		Name:     req.WorkflowTemplate.Name,
		Manifest: req.WorkflowTemplate.Manifest,
	}
	workflowTemplate, err := s.resourceManager.CreateWorkflowTemplateVersion(req.Namespace, workflowTemplate)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}
	req.WorkflowTemplate.Uid = workflowTemplate.UID
	req.WorkflowTemplate.Name = workflowTemplate.Name
	req.WorkflowTemplate.Version = workflowTemplate.Version

	return req.WorkflowTemplate, nil
}

func (s *WorkflowServer) UpdateWorkflowTemplateVersion(ctx context.Context, req *api.UpdateWorkflowTemplateVersionRequest) (*api.WorkflowTemplate, error) {
	workflowTemplate := &model.WorkflowTemplate{
		UID:      req.WorkflowTemplate.Uid,
		Name:     req.WorkflowTemplate.Name,
		Manifest: req.WorkflowTemplate.Manifest,
		Version:  req.WorkflowTemplate.Version,
	}

	workflowTemplate, err := s.resourceManager.UpdateWorkflowTemplateVersion(req.Namespace, workflowTemplate)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}
	req.WorkflowTemplate.Uid = workflowTemplate.UID
	req.WorkflowTemplate.Name = workflowTemplate.Name
	req.WorkflowTemplate.Version = workflowTemplate.Version

	return req.WorkflowTemplate, nil
}

func (s *WorkflowServer) GetWorkflowTemplate(ctx context.Context, req *api.GetWorkflowTemplateRequest) (*api.WorkflowTemplate, error) {
	workflowTemplate, err := s.resourceManager.GetWorkflowTemplate(req.Namespace, req.Uid, req.Version)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}

	return apiWorkflowTemplate(workflowTemplate), nil
}

func (s *WorkflowServer) ListWorkflowTemplateVersions(ctx context.Context, req *api.ListWorkflowTemplateVersionsRequest) (*api.ListWorkflowTemplateVersionsResponse, error) {
	workflowTemplateVersions, err := s.resourceManager.ListWorkflowTemplateVersions(req.Namespace, req.Uid)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}

	workflowTemplates := []*api.WorkflowTemplate{}
	for _, wtv := range workflowTemplateVersions {
		workflowTemplates = append(workflowTemplates, apiWorkflowTemplate(wtv))
	}

	return &api.ListWorkflowTemplateVersionsResponse{
		Count:             int32(len(workflowTemplateVersions)),
		WorkflowTemplates: workflowTemplates,
	}, nil
}

func (s *WorkflowServer) ListWorkflowTemplates(ctx context.Context, req *api.ListWorkflowTemplatesRequest) (*api.ListWorkflowTemplatesResponse, error) {
	workflowTemplates, err := s.resourceManager.ListWorkflowTemplates(req.Namespace)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}

	apiWorkflowTemplates := []*api.WorkflowTemplate{}
	for _, wtv := range workflowTemplates {
		apiWorkflowTemplates = append(apiWorkflowTemplates, apiWorkflowTemplate(wtv))
	}

	return &api.ListWorkflowTemplatesResponse{
		Count:             int32(len(apiWorkflowTemplates)),
		WorkflowTemplates: apiWorkflowTemplates,
	}, nil
}

func (s *WorkflowServer) ArchiveWorkflowTemplate(ctx context.Context, req *api.ArchiveWorkflowTemplateRequest) (*api.ArchiveWorkflowTemplateResponse, error) {
	archived, err := s.resourceManager.ArchiveWorkflowTemplate(req.Namespace, req.Uid)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}

	return &api.ArchiveWorkflowTemplateResponse{
		WorkflowTemplate: &api.WorkflowTemplate{
			IsArchived: archived,
		},
	}, nil
}
