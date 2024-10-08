// An Alert, a set of Events, is the smallest unit to manage incidents.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/monitoring/v1/alert.proto

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
	Alert_Create_FullMethodName      = "/spaceone.api.monitoring.v1.Alert/create"
	Alert_Update_FullMethodName      = "/spaceone.api.monitoring.v1.Alert/update"
	Alert_AssignUser_FullMethodName  = "/spaceone.api.monitoring.v1.Alert/assign_user"
	Alert_UpdateState_FullMethodName = "/spaceone.api.monitoring.v1.Alert/update_state"
	Alert_Delete_FullMethodName      = "/spaceone.api.monitoring.v1.Alert/delete"
	Alert_Get_FullMethodName         = "/spaceone.api.monitoring.v1.Alert/get"
	Alert_List_FullMethodName        = "/spaceone.api.monitoring.v1.Alert/list"
	Alert_Stat_FullMethodName        = "/spaceone.api.monitoring.v1.Alert/stat"
)

// AlertClient is the client API for Alert service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertClient interface {
	// Creates a new Alert. Alerts generated with `create` method are made in a manual way. Manually made Alerts can be used for Notifications.
	Create(ctx context.Context, in *CreateAlertRequest, opts ...grpc.CallOption) (*AlertInfo, error)
	// Updates a specific Alert. You can make changes in Alert settings, including the `title`, `description`, `responder`, `state`, and `urgency`. The `responder` of the Alert is a User who is assigned to respond to the Alert.
	Update(ctx context.Context, in *UpdateAlertRequest, opts ...grpc.CallOption) (*AlertInfo, error)
	AssignUser(ctx context.Context, in *AssignUserRequest, opts ...grpc.CallOption) (*AlertInfo, error)
	// Updates the state of an Alert via callback URL by creating a temporary `access_key` while generating a Notification about the Alert.
	// +noauth
	UpdateState(ctx context.Context, in *UpdateAlertStateRequest, opts ...grpc.CallOption) (*AlertInfo, error)
	// Deletes a specific Alert and remove it from the list of Alerts. You must specify the `alert_id` of the Alert to delete.
	Delete(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific Alert. Prints detailed information about the Alert.
	Get(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (*AlertInfo, error)
	// Gets a list of all Alerts. You can use a query to get a filtered list of Alerts.
	List(ctx context.Context, in *AlertQuery, opts ...grpc.CallOption) (*AlertsInfo, error)
	Stat(ctx context.Context, in *AlertStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type alertClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertClient(cc grpc.ClientConnInterface) AlertClient {
	return &alertClient{cc}
}

func (c *alertClient) Create(ctx context.Context, in *CreateAlertRequest, opts ...grpc.CallOption) (*AlertInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlertInfo)
	err := c.cc.Invoke(ctx, Alert_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertClient) Update(ctx context.Context, in *UpdateAlertRequest, opts ...grpc.CallOption) (*AlertInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlertInfo)
	err := c.cc.Invoke(ctx, Alert_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertClient) AssignUser(ctx context.Context, in *AssignUserRequest, opts ...grpc.CallOption) (*AlertInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlertInfo)
	err := c.cc.Invoke(ctx, Alert_AssignUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertClient) UpdateState(ctx context.Context, in *UpdateAlertStateRequest, opts ...grpc.CallOption) (*AlertInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlertInfo)
	err := c.cc.Invoke(ctx, Alert_UpdateState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertClient) Delete(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Alert_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertClient) Get(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (*AlertInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlertInfo)
	err := c.cc.Invoke(ctx, Alert_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertClient) List(ctx context.Context, in *AlertQuery, opts ...grpc.CallOption) (*AlertsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlertsInfo)
	err := c.cc.Invoke(ctx, Alert_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertClient) Stat(ctx context.Context, in *AlertStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Alert_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertServer is the server API for Alert service.
// All implementations must embed UnimplementedAlertServer
// for forward compatibility.
type AlertServer interface {
	// Creates a new Alert. Alerts generated with `create` method are made in a manual way. Manually made Alerts can be used for Notifications.
	Create(context.Context, *CreateAlertRequest) (*AlertInfo, error)
	// Updates a specific Alert. You can make changes in Alert settings, including the `title`, `description`, `responder`, `state`, and `urgency`. The `responder` of the Alert is a User who is assigned to respond to the Alert.
	Update(context.Context, *UpdateAlertRequest) (*AlertInfo, error)
	AssignUser(context.Context, *AssignUserRequest) (*AlertInfo, error)
	// Updates the state of an Alert via callback URL by creating a temporary `access_key` while generating a Notification about the Alert.
	// +noauth
	UpdateState(context.Context, *UpdateAlertStateRequest) (*AlertInfo, error)
	// Deletes a specific Alert and remove it from the list of Alerts. You must specify the `alert_id` of the Alert to delete.
	Delete(context.Context, *AlertRequest) (*empty.Empty, error)
	// Gets a specific Alert. Prints detailed information about the Alert.
	Get(context.Context, *AlertRequest) (*AlertInfo, error)
	// Gets a list of all Alerts. You can use a query to get a filtered list of Alerts.
	List(context.Context, *AlertQuery) (*AlertsInfo, error)
	Stat(context.Context, *AlertStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedAlertServer()
}

// UnimplementedAlertServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAlertServer struct{}

func (UnimplementedAlertServer) Create(context.Context, *CreateAlertRequest) (*AlertInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAlertServer) Update(context.Context, *UpdateAlertRequest) (*AlertInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAlertServer) AssignUser(context.Context, *AssignUserRequest) (*AlertInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignUser not implemented")
}
func (UnimplementedAlertServer) UpdateState(context.Context, *UpdateAlertStateRequest) (*AlertInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateState not implemented")
}
func (UnimplementedAlertServer) Delete(context.Context, *AlertRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAlertServer) Get(context.Context, *AlertRequest) (*AlertInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAlertServer) List(context.Context, *AlertQuery) (*AlertsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAlertServer) Stat(context.Context, *AlertStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedAlertServer) mustEmbedUnimplementedAlertServer() {}
func (UnimplementedAlertServer) testEmbeddedByValue()               {}

// UnsafeAlertServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertServer will
// result in compilation errors.
type UnsafeAlertServer interface {
	mustEmbedUnimplementedAlertServer()
}

func RegisterAlertServer(s grpc.ServiceRegistrar, srv AlertServer) {
	// If the following call pancis, it indicates UnimplementedAlertServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Alert_ServiceDesc, srv)
}

func _Alert_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alert_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).Create(ctx, req.(*CreateAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alert_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alert_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).Update(ctx, req.(*UpdateAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alert_AssignUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).AssignUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alert_AssignUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).AssignUser(ctx, req.(*AssignUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alert_UpdateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).UpdateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alert_UpdateState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).UpdateState(ctx, req.(*UpdateAlertStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alert_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alert_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).Delete(ctx, req.(*AlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alert_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alert_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).Get(ctx, req.(*AlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alert_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alert_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).List(ctx, req.(*AlertQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alert_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Alert_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).Stat(ctx, req.(*AlertStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Alert_ServiceDesc is the grpc.ServiceDesc for Alert service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Alert_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.monitoring.v1.Alert",
	HandlerType: (*AlertServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Alert_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Alert_Update_Handler,
		},
		{
			MethodName: "assign_user",
			Handler:    _Alert_AssignUser_Handler,
		},
		{
			MethodName: "update_state",
			Handler:    _Alert_UpdateState_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Alert_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Alert_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Alert_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Alert_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/monitoring/v1/alert.proto",
}
