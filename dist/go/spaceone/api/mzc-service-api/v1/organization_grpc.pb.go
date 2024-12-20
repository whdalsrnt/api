// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/mzc_service_api/v1/organization.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Organization_Create_FullMethodName  = "/spaceone.api.mzc_service_api.v1.Organization/create"
	Organization_Update_FullMethodName  = "/spaceone.api.mzc_service_api.v1.Organization/update"
	Organization_Enable_FullMethodName  = "/spaceone.api.mzc_service_api.v1.Organization/enable"
	Organization_Disable_FullMethodName = "/spaceone.api.mzc_service_api.v1.Organization/disable"
	Organization_Delete_FullMethodName  = "/spaceone.api.mzc_service_api.v1.Organization/delete"
	Organization_Get_FullMethodName     = "/spaceone.api.mzc_service_api.v1.Organization/get"
	Organization_List_FullMethodName    = "/spaceone.api.mzc_service_api.v1.Organization/list"
	Organization_Stat_FullMethodName    = "/spaceone.api.mzc_service_api.v1.Organization/stat"
)

// OrganizationClient is the client API for Organization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationClient interface {
	Create(ctx context.Context, in *OrganizationCreateRequest, opts ...grpc.CallOption) (*OrganizationInfo, error)
	Update(ctx context.Context, in *OrganizationUpdateRequest, opts ...grpc.CallOption) (*OrganizationInfo, error)
	Enable(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*OrganizationInfo, error)
	Disable(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*OrganizationInfo, error)
	Delete(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*OrganizationInfo, error)
	List(ctx context.Context, in *OrganizationSearchQuery, opts ...grpc.CallOption) (*OrganizationsInfo, error)
	Stat(ctx context.Context, in *OrganizationStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type organizationClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationClient(cc grpc.ClientConnInterface) OrganizationClient {
	return &organizationClient{cc}
}

func (c *organizationClient) Create(ctx context.Context, in *OrganizationCreateRequest, opts ...grpc.CallOption) (*OrganizationInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrganizationInfo)
	err := c.cc.Invoke(ctx, Organization_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) Update(ctx context.Context, in *OrganizationUpdateRequest, opts ...grpc.CallOption) (*OrganizationInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrganizationInfo)
	err := c.cc.Invoke(ctx, Organization_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) Enable(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*OrganizationInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrganizationInfo)
	err := c.cc.Invoke(ctx, Organization_Enable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) Disable(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*OrganizationInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrganizationInfo)
	err := c.cc.Invoke(ctx, Organization_Disable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) Delete(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Organization_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) Get(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*OrganizationInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrganizationInfo)
	err := c.cc.Invoke(ctx, Organization_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) List(ctx context.Context, in *OrganizationSearchQuery, opts ...grpc.CallOption) (*OrganizationsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrganizationsInfo)
	err := c.cc.Invoke(ctx, Organization_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) Stat(ctx context.Context, in *OrganizationStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Organization_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationServer is the server API for Organization service.
// All implementations must embed UnimplementedOrganizationServer
// for forward compatibility.
type OrganizationServer interface {
	Create(context.Context, *OrganizationCreateRequest) (*OrganizationInfo, error)
	Update(context.Context, *OrganizationUpdateRequest) (*OrganizationInfo, error)
	Enable(context.Context, *OrganizationRequest) (*OrganizationInfo, error)
	Disable(context.Context, *OrganizationRequest) (*OrganizationInfo, error)
	Delete(context.Context, *OrganizationRequest) (*empty.Empty, error)
	Get(context.Context, *OrganizationRequest) (*OrganizationInfo, error)
	List(context.Context, *OrganizationSearchQuery) (*OrganizationsInfo, error)
	Stat(context.Context, *OrganizationStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedOrganizationServer()
}

// UnimplementedOrganizationServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrganizationServer struct{}

func (UnimplementedOrganizationServer) Create(context.Context, *OrganizationCreateRequest) (*OrganizationInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrganizationServer) Update(context.Context, *OrganizationUpdateRequest) (*OrganizationInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrganizationServer) Enable(context.Context, *OrganizationRequest) (*OrganizationInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedOrganizationServer) Disable(context.Context, *OrganizationRequest) (*OrganizationInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedOrganizationServer) Delete(context.Context, *OrganizationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOrganizationServer) Get(context.Context, *OrganizationRequest) (*OrganizationInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOrganizationServer) List(context.Context, *OrganizationSearchQuery) (*OrganizationsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrganizationServer) Stat(context.Context, *OrganizationStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedOrganizationServer) mustEmbedUnimplementedOrganizationServer() {}
func (UnimplementedOrganizationServer) testEmbeddedByValue()                      {}

// UnsafeOrganizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationServer will
// result in compilation errors.
type UnsafeOrganizationServer interface {
	mustEmbedUnimplementedOrganizationServer()
}

func RegisterOrganizationServer(s grpc.ServiceRegistrar, srv OrganizationServer) {
	// If the following call pancis, it indicates UnimplementedOrganizationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Organization_ServiceDesc, srv)
}

func _Organization_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).Create(ctx, req.(*OrganizationCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).Update(ctx, req.(*OrganizationUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).Enable(ctx, req.(*OrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).Disable(ctx, req.(*OrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).Delete(ctx, req.(*OrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).Get(ctx, req.(*OrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).List(ctx, req.(*OrganizationSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).Stat(ctx, req.(*OrganizationStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Organization_ServiceDesc is the grpc.ServiceDesc for Organization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Organization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.mzc_service_api.v1.Organization",
	HandlerType: (*OrganizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Organization_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Organization_Update_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _Organization_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _Organization_Disable_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Organization_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Organization_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Organization_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Organization_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/mzc_service_api/v1/organization.proto",
}
