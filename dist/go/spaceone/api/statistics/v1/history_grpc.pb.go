// A History is a record of data collection based on a Schedule.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/statistics/v1/history.proto

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
	History_Create_FullMethodName = "/spaceone.api.statistics.v1.History/create"
	History_List_FullMethodName   = "/spaceone.api.statistics.v1.History/list"
	History_Stat_FullMethodName   = "/spaceone.api.statistics.v1.History/stat"
)

// HistoryClient is the client API for History service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HistoryClient interface {
	// Creates a new History. Gets a Schedule as an input and creates a History as an output. You can use this method to manually run a specific Schedule.
	Create(ctx context.Context, in *CreateHistoryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a list of all Histories. You can use a query to get a filtered list of Histories.
	List(ctx context.Context, in *QueryHistoryRequest, opts ...grpc.CallOption) (*HistoryInfo, error)
	Stat(ctx context.Context, in *HistoryStatRequest, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type historyClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoryClient(cc grpc.ClientConnInterface) HistoryClient {
	return &historyClient{cc}
}

func (c *historyClient) Create(ctx context.Context, in *CreateHistoryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, History_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyClient) List(ctx context.Context, in *QueryHistoryRequest, opts ...grpc.CallOption) (*HistoryInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HistoryInfo)
	err := c.cc.Invoke(ctx, History_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyClient) Stat(ctx context.Context, in *HistoryStatRequest, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, History_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoryServer is the server API for History service.
// All implementations must embed UnimplementedHistoryServer
// for forward compatibility.
type HistoryServer interface {
	// Creates a new History. Gets a Schedule as an input and creates a History as an output. You can use this method to manually run a specific Schedule.
	Create(context.Context, *CreateHistoryRequest) (*empty.Empty, error)
	// Gets a list of all Histories. You can use a query to get a filtered list of Histories.
	List(context.Context, *QueryHistoryRequest) (*HistoryInfo, error)
	Stat(context.Context, *HistoryStatRequest) (*_struct.Struct, error)
	mustEmbedUnimplementedHistoryServer()
}

// UnimplementedHistoryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHistoryServer struct{}

func (UnimplementedHistoryServer) Create(context.Context, *CreateHistoryRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedHistoryServer) List(context.Context, *QueryHistoryRequest) (*HistoryInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedHistoryServer) Stat(context.Context, *HistoryStatRequest) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedHistoryServer) mustEmbedUnimplementedHistoryServer() {}
func (UnimplementedHistoryServer) testEmbeddedByValue()                 {}

// UnsafeHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HistoryServer will
// result in compilation errors.
type UnsafeHistoryServer interface {
	mustEmbedUnimplementedHistoryServer()
}

func RegisterHistoryServer(s grpc.ServiceRegistrar, srv HistoryServer) {
	// If the following call pancis, it indicates UnimplementedHistoryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&History_ServiceDesc, srv)
}

func _History_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: History_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).Create(ctx, req.(*CreateHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _History_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: History_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).List(ctx, req.(*QueryHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _History_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: History_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).Stat(ctx, req.(*HistoryStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// History_ServiceDesc is the grpc.ServiceDesc for History service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var History_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.statistics.v1.History",
	HandlerType: (*HistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _History_Create_Handler,
		},
		{
			MethodName: "list",
			Handler:    _History_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _History_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/statistics/v1/history.proto",
}
