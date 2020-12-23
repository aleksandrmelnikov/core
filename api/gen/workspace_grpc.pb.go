// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WorkspaceServiceClient is the client API for WorkspaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceServiceClient interface {
	CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*Workspace, error)
	GetWorkspaceStatisticsForNamespace(ctx context.Context, in *GetWorkspaceStatisticsForNamespaceRequest, opts ...grpc.CallOption) (*GetWorkspaceStatisticsForNamespaceResponse, error)
	GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*Workspace, error)
	ListWorkspaces(ctx context.Context, in *ListWorkspaceRequest, opts ...grpc.CallOption) (*ListWorkspaceResponse, error)
	UpdateWorkspaceStatus(ctx context.Context, in *UpdateWorkspaceStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PauseWorkspace(ctx context.Context, in *PauseWorkspaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ResumeWorkspace(ctx context.Context, in *ResumeWorkspaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RetryLastWorkspaceAction(ctx context.Context, in *RetryActionWorkspaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type workspaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceServiceClient(cc grpc.ClientConnInterface) WorkspaceServiceClient {
	return &workspaceServiceClient{cc}
}

func (c *workspaceServiceClient) CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*Workspace, error) {
	out := new(Workspace)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/CreateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetWorkspaceStatisticsForNamespace(ctx context.Context, in *GetWorkspaceStatisticsForNamespaceRequest, opts ...grpc.CallOption) (*GetWorkspaceStatisticsForNamespaceResponse, error) {
	out := new(GetWorkspaceStatisticsForNamespaceResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/GetWorkspaceStatisticsForNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*Workspace, error) {
	out := new(Workspace)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/GetWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ListWorkspaces(ctx context.Context, in *ListWorkspaceRequest, opts ...grpc.CallOption) (*ListWorkspaceResponse, error) {
	out := new(ListWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/ListWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) UpdateWorkspaceStatus(ctx context.Context, in *UpdateWorkspaceStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/UpdateWorkspaceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/UpdateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) PauseWorkspace(ctx context.Context, in *PauseWorkspaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/PauseWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ResumeWorkspace(ctx context.Context, in *ResumeWorkspaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/ResumeWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/DeleteWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) RetryLastWorkspaceAction(ctx context.Context, in *RetryActionWorkspaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/RetryLastWorkspaceAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServiceServer is the server API for WorkspaceService service.
// All implementations must embed UnimplementedWorkspaceServiceServer
// for forward compatibility
type WorkspaceServiceServer interface {
	CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*Workspace, error)
	GetWorkspaceStatisticsForNamespace(context.Context, *GetWorkspaceStatisticsForNamespaceRequest) (*GetWorkspaceStatisticsForNamespaceResponse, error)
	GetWorkspace(context.Context, *GetWorkspaceRequest) (*Workspace, error)
	ListWorkspaces(context.Context, *ListWorkspaceRequest) (*ListWorkspaceResponse, error)
	UpdateWorkspaceStatus(context.Context, *UpdateWorkspaceStatusRequest) (*emptypb.Empty, error)
	UpdateWorkspace(context.Context, *UpdateWorkspaceRequest) (*emptypb.Empty, error)
	PauseWorkspace(context.Context, *PauseWorkspaceRequest) (*emptypb.Empty, error)
	ResumeWorkspace(context.Context, *ResumeWorkspaceRequest) (*emptypb.Empty, error)
	DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*emptypb.Empty, error)
	RetryLastWorkspaceAction(context.Context, *RetryActionWorkspaceRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedWorkspaceServiceServer()
}

// UnimplementedWorkspaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServiceServer struct {
}

func (UnimplementedWorkspaceServiceServer) CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*Workspace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetWorkspaceStatisticsForNamespace(context.Context, *GetWorkspaceStatisticsForNamespaceRequest) (*GetWorkspaceStatisticsForNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceStatisticsForNamespace not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetWorkspace(context.Context, *GetWorkspaceRequest) (*Workspace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) ListWorkspaces(context.Context, *ListWorkspaceRequest) (*ListWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaces not implemented")
}
func (UnimplementedWorkspaceServiceServer) UpdateWorkspaceStatus(context.Context, *UpdateWorkspaceStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspaceStatus not implemented")
}
func (UnimplementedWorkspaceServiceServer) UpdateWorkspace(context.Context, *UpdateWorkspaceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) PauseWorkspace(context.Context, *PauseWorkspaceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) ResumeWorkspace(context.Context, *ResumeWorkspaceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) RetryLastWorkspaceAction(context.Context, *RetryActionWorkspaceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetryLastWorkspaceAction not implemented")
}
func (UnimplementedWorkspaceServiceServer) mustEmbedUnimplementedWorkspaceServiceServer() {}

// UnsafeWorkspaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServiceServer will
// result in compilation errors.
type UnsafeWorkspaceServiceServer interface {
	mustEmbedUnimplementedWorkspaceServiceServer()
}

func RegisterWorkspaceServiceServer(s grpc.ServiceRegistrar, srv WorkspaceServiceServer) {
	s.RegisterService(&_WorkspaceService_serviceDesc, srv)
}

func _WorkspaceService_CreateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/CreateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, req.(*CreateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetWorkspaceStatisticsForNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceStatisticsForNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspaceStatisticsForNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/GetWorkspaceStatisticsForNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspaceStatisticsForNamespace(ctx, req.(*GetWorkspaceStatisticsForNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/GetWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspace(ctx, req.(*GetWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ListWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ListWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/ListWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ListWorkspaces(ctx, req.(*ListWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_UpdateWorkspaceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspaceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).UpdateWorkspaceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/UpdateWorkspaceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).UpdateWorkspaceStatus(ctx, req.(*UpdateWorkspaceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_UpdateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).UpdateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/UpdateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).UpdateWorkspace(ctx, req.(*UpdateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_PauseWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).PauseWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/PauseWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).PauseWorkspace(ctx, req.(*PauseWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ResumeWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ResumeWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/ResumeWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ResumeWorkspace(ctx, req.(*ResumeWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/DeleteWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, req.(*DeleteWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_RetryLastWorkspaceAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetryActionWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).RetryLastWorkspaceAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/RetryLastWorkspaceAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).RetryLastWorkspaceAction(ctx, req.(*RetryActionWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkspaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.WorkspaceService",
	HandlerType: (*WorkspaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkspace",
			Handler:    _WorkspaceService_CreateWorkspace_Handler,
		},
		{
			MethodName: "GetWorkspaceStatisticsForNamespace",
			Handler:    _WorkspaceService_GetWorkspaceStatisticsForNamespace_Handler,
		},
		{
			MethodName: "GetWorkspace",
			Handler:    _WorkspaceService_GetWorkspace_Handler,
		},
		{
			MethodName: "ListWorkspaces",
			Handler:    _WorkspaceService_ListWorkspaces_Handler,
		},
		{
			MethodName: "UpdateWorkspaceStatus",
			Handler:    _WorkspaceService_UpdateWorkspaceStatus_Handler,
		},
		{
			MethodName: "UpdateWorkspace",
			Handler:    _WorkspaceService_UpdateWorkspace_Handler,
		},
		{
			MethodName: "PauseWorkspace",
			Handler:    _WorkspaceService_PauseWorkspace_Handler,
		},
		{
			MethodName: "ResumeWorkspace",
			Handler:    _WorkspaceService_ResumeWorkspace_Handler,
		},
		{
			MethodName: "DeleteWorkspace",
			Handler:    _WorkspaceService_DeleteWorkspace_Handler,
		},
		{
			MethodName: "RetryLastWorkspaceAction",
			Handler:    _WorkspaceService_RetryLastWorkspaceAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workspace.proto",
}