// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.6.1
// source: spaceone/api/cost_analysis/v1/cost_report_data.proto

package v1

import (
	context "context"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CostReportData_List_FullMethodName    = "/spaceone.api.cost_analysis.v1.CostReportData/list"
	CostReportData_Analyze_FullMethodName = "/spaceone.api.cost_analysis.v1.CostReportData/analyze"
	CostReportData_Stat_FullMethodName    = "/spaceone.api.cost_analysis.v1.CostReportData/stat"
)

// CostReportDataClient is the client API for CostReportData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CostReportDataClient interface {
	List(ctx context.Context, in *CostReportDataQuery, opts ...grpc.CallOption) (*CostReportsDataInfo, error)
	Analyze(ctx context.Context, in *CostReportDataAnalyzeQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
	Stat(ctx context.Context, in *CostReportDataStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type costReportDataClient struct {
	cc grpc.ClientConnInterface
}

func NewCostReportDataClient(cc grpc.ClientConnInterface) CostReportDataClient {
	return &costReportDataClient{cc}
}

func (c *costReportDataClient) List(ctx context.Context, in *CostReportDataQuery, opts ...grpc.CallOption) (*CostReportsDataInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CostReportsDataInfo)
	err := c.cc.Invoke(ctx, CostReportData_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costReportDataClient) Analyze(ctx context.Context, in *CostReportDataAnalyzeQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, CostReportData_Analyze_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costReportDataClient) Stat(ctx context.Context, in *CostReportDataStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, CostReportData_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CostReportDataServer is the server API for CostReportData service.
// All implementations must embed UnimplementedCostReportDataServer
// for forward compatibility
type CostReportDataServer interface {
	List(context.Context, *CostReportDataQuery) (*CostReportsDataInfo, error)
	Analyze(context.Context, *CostReportDataAnalyzeQuery) (*_struct.Struct, error)
	Stat(context.Context, *CostReportDataStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedCostReportDataServer()
}

// UnimplementedCostReportDataServer must be embedded to have forward compatible implementations.
type UnimplementedCostReportDataServer struct {
}

func (UnimplementedCostReportDataServer) List(context.Context, *CostReportDataQuery) (*CostReportsDataInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCostReportDataServer) Analyze(context.Context, *CostReportDataAnalyzeQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedCostReportDataServer) Stat(context.Context, *CostReportDataStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedCostReportDataServer) mustEmbedUnimplementedCostReportDataServer() {}

// UnsafeCostReportDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CostReportDataServer will
// result in compilation errors.
type UnsafeCostReportDataServer interface {
	mustEmbedUnimplementedCostReportDataServer()
}

func RegisterCostReportDataServer(s grpc.ServiceRegistrar, srv CostReportDataServer) {
	s.RegisterService(&CostReportData_ServiceDesc, srv)
}

func _CostReportData_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CostReportDataQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostReportDataServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostReportData_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostReportDataServer).List(ctx, req.(*CostReportDataQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostReportData_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CostReportDataAnalyzeQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostReportDataServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostReportData_Analyze_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostReportDataServer).Analyze(ctx, req.(*CostReportDataAnalyzeQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostReportData_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CostReportDataStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostReportDataServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostReportData_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostReportDataServer).Stat(ctx, req.(*CostReportDataStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// CostReportData_ServiceDesc is the grpc.ServiceDesc for CostReportData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CostReportData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.cost_analysis.v1.CostReportData",
	HandlerType: (*CostReportDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _CostReportData_List_Handler,
		},
		{
			MethodName: "analyze",
			Handler:    _CostReportData_Analyze_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _CostReportData_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/cost_analysis/v1/cost_report_data.proto",
}
