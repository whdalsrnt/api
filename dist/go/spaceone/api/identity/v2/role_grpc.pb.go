// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/identity/v2/role.proto

package v2

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
	Role_Create_FullMethodName        = "/spaceone.api.identity.v2.Role/create"
	Role_Update_FullMethodName        = "/spaceone.api.identity.v2.Role/update"
	Role_Enable_FullMethodName        = "/spaceone.api.identity.v2.Role/enable"
	Role_Disable_FullMethodName       = "/spaceone.api.identity.v2.Role/disable"
	Role_Delete_FullMethodName        = "/spaceone.api.identity.v2.Role/delete"
	Role_Get_FullMethodName           = "/spaceone.api.identity.v2.Role/get"
	Role_List_FullMethodName          = "/spaceone.api.identity.v2.Role/list"
	Role_ListBasicRole_FullMethodName = "/spaceone.api.identity.v2.Role/list_basic_role"
	Role_Stat_FullMethodName          = "/spaceone.api.identity.v2.Role/stat"
)

// RoleClient is the client API for Role service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleClient interface {
	Create(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*RoleInfo, error)
	Update(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*RoleInfo, error)
	Enable(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*RoleInfo, error)
	Disable(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*RoleInfo, error)
	Delete(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*RoleInfo, error)
	List(ctx context.Context, in *RoleSearchQuery, opts ...grpc.CallOption) (*RolesInfo, error)
	ListBasicRole(ctx context.Context, in *RoleSearchQuery, opts ...grpc.CallOption) (*RolesInfo, error)
	Stat(ctx context.Context, in *RoleStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type roleClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleClient(cc grpc.ClientConnInterface) RoleClient {
	return &roleClient{cc}
}

func (c *roleClient) Create(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*RoleInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleInfo)
	err := c.cc.Invoke(ctx, Role_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) Update(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*RoleInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleInfo)
	err := c.cc.Invoke(ctx, Role_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) Enable(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*RoleInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleInfo)
	err := c.cc.Invoke(ctx, Role_Enable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) Disable(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*RoleInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleInfo)
	err := c.cc.Invoke(ctx, Role_Disable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) Delete(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Role_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) Get(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*RoleInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleInfo)
	err := c.cc.Invoke(ctx, Role_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) List(ctx context.Context, in *RoleSearchQuery, opts ...grpc.CallOption) (*RolesInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RolesInfo)
	err := c.cc.Invoke(ctx, Role_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) ListBasicRole(ctx context.Context, in *RoleSearchQuery, opts ...grpc.CallOption) (*RolesInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RolesInfo)
	err := c.cc.Invoke(ctx, Role_ListBasicRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) Stat(ctx context.Context, in *RoleStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Role_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServer is the server API for Role service.
// All implementations must embed UnimplementedRoleServer
// for forward compatibility.
type RoleServer interface {
	Create(context.Context, *CreateRoleRequest) (*RoleInfo, error)
	Update(context.Context, *UpdateRoleRequest) (*RoleInfo, error)
	Enable(context.Context, *RoleRequest) (*RoleInfo, error)
	Disable(context.Context, *RoleRequest) (*RoleInfo, error)
	Delete(context.Context, *RoleRequest) (*empty.Empty, error)
	Get(context.Context, *RoleRequest) (*RoleInfo, error)
	List(context.Context, *RoleSearchQuery) (*RolesInfo, error)
	ListBasicRole(context.Context, *RoleSearchQuery) (*RolesInfo, error)
	Stat(context.Context, *RoleStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedRoleServer()
}

// UnimplementedRoleServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRoleServer struct{}

func (UnimplementedRoleServer) Create(context.Context, *CreateRoleRequest) (*RoleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRoleServer) Update(context.Context, *UpdateRoleRequest) (*RoleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRoleServer) Enable(context.Context, *RoleRequest) (*RoleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedRoleServer) Disable(context.Context, *RoleRequest) (*RoleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedRoleServer) Delete(context.Context, *RoleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRoleServer) Get(context.Context, *RoleRequest) (*RoleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRoleServer) List(context.Context, *RoleSearchQuery) (*RolesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRoleServer) ListBasicRole(context.Context, *RoleSearchQuery) (*RolesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBasicRole not implemented")
}
func (UnimplementedRoleServer) Stat(context.Context, *RoleStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedRoleServer) mustEmbedUnimplementedRoleServer() {}
func (UnimplementedRoleServer) testEmbeddedByValue()              {}

// UnsafeRoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServer will
// result in compilation errors.
type UnsafeRoleServer interface {
	mustEmbedUnimplementedRoleServer()
}

func RegisterRoleServer(s grpc.ServiceRegistrar, srv RoleServer) {
	// If the following call pancis, it indicates UnimplementedRoleServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Role_ServiceDesc, srv)
}

func _Role_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).Create(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).Update(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).Enable(ctx, req.(*RoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).Disable(ctx, req.(*RoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).Delete(ctx, req.(*RoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).Get(ctx, req.(*RoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).List(ctx, req.(*RoleSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_ListBasicRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).ListBasicRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_ListBasicRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).ListBasicRole(ctx, req.(*RoleSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).Stat(ctx, req.(*RoleStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Role_ServiceDesc is the grpc.ServiceDesc for Role service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Role_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.identity.v2.Role",
	HandlerType: (*RoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Role_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Role_Update_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _Role_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _Role_Disable_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Role_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Role_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Role_List_Handler,
		},
		{
			MethodName: "list_basic_role",
			Handler:    _Role_ListBasicRole_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Role_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/identity/v2/role.proto",
}
