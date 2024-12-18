// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/alert_manager/plugin/webhook.proto

package plugin

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Webhook_Init_FullMethodName   = "/spaceone.api.alert_manager.plugin.Webhook/init"
	Webhook_Verify_FullMethodName = "/spaceone.api.alert_manager.plugin.Webhook/verify"
)

// WebhookClient is the client API for Webhook service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebhookClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*PluginInfo, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type webhookClient struct {
	cc grpc.ClientConnInterface
}

func NewWebhookClient(cc grpc.ClientConnInterface) WebhookClient {
	return &webhookClient{cc}
}

func (c *webhookClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*PluginInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, Webhook_Init_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Webhook_Verify_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebhookServer is the server API for Webhook service.
// All implementations must embed UnimplementedWebhookServer
// for forward compatibility.
type WebhookServer interface {
	Init(context.Context, *InitRequest) (*PluginInfo, error)
	Verify(context.Context, *VerifyRequest) (*empty.Empty, error)
	mustEmbedUnimplementedWebhookServer()
}

// UnimplementedWebhookServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWebhookServer struct{}

func (UnimplementedWebhookServer) Init(context.Context, *InitRequest) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedWebhookServer) Verify(context.Context, *VerifyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedWebhookServer) mustEmbedUnimplementedWebhookServer() {}
func (UnimplementedWebhookServer) testEmbeddedByValue()                 {}

// UnsafeWebhookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebhookServer will
// result in compilation errors.
type UnsafeWebhookServer interface {
	mustEmbedUnimplementedWebhookServer()
}

func RegisterWebhookServer(s grpc.ServiceRegistrar, srv WebhookServer) {
	// If the following call pancis, it indicates UnimplementedWebhookServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Webhook_ServiceDesc, srv)
}

func _Webhook_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_Verify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Webhook_ServiceDesc is the grpc.ServiceDesc for Webhook service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Webhook_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.alert_manager.plugin.Webhook",
	HandlerType: (*WebhookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "init",
			Handler:    _Webhook_Init_Handler,
		},
		{
			MethodName: "verify",
			Handler:    _Webhook_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/alert_manager/plugin/webhook.proto",
}
