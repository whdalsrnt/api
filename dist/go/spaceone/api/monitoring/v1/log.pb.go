// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: spaceone/api/monitoring/v1/log.proto

package v1

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSourceId string `protobuf:"bytes,1,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	ResourceId   string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// +optional
	Keyword string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	// +optional
	Start string `protobuf:"bytes,10,opt,name=start,proto3" json:"start,omitempty"`
	// +optional
	End string `protobuf:"bytes,11,opt,name=end,proto3" json:"end,omitempty"`
	// +optional
	Sort *_struct.Struct `protobuf:"bytes,15,opt,name=sort,proto3" json:"sort,omitempty"`
	// +optional
	Limit int32 `protobuf:"varint,16,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_monitoring_v1_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_v1_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogRequest) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *LogRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *LogRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *LogRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *LogRequest) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *LogRequest) GetSort() *_struct.Struct {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *LogRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type LogDataInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results  *_struct.ListValue `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
	DomainId string             `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *LogDataInfo) Reset() {
	*x = LogDataInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_monitoring_v1_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogDataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogDataInfo) ProtoMessage() {}

func (x *LogDataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_v1_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogDataInfo.ProtoReflect.Descriptor instead.
func (*LogDataInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogDataInfo) GetResults() *_struct.ListValue {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *LogDataInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_monitoring_v1_log_proto protoreflect.FileDescriptor

var file_spaceone_api_monitoring_v1_log_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8,
	0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x60, 0x0a, 0x0b, 0x4c, 0x6f, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x32, 0x82, 0x01, 0x0a, 0x03,
	0x4c, 0x6f, 0x67, 0x12, 0x7b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_monitoring_v1_log_proto_rawDescOnce sync.Once
	file_spaceone_api_monitoring_v1_log_proto_rawDescData = file_spaceone_api_monitoring_v1_log_proto_rawDesc
)

func file_spaceone_api_monitoring_v1_log_proto_rawDescGZIP() []byte {
	file_spaceone_api_monitoring_v1_log_proto_rawDescOnce.Do(func() {
		file_spaceone_api_monitoring_v1_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_monitoring_v1_log_proto_rawDescData)
	})
	return file_spaceone_api_monitoring_v1_log_proto_rawDescData
}

var file_spaceone_api_monitoring_v1_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spaceone_api_monitoring_v1_log_proto_goTypes = []any{
	(*LogRequest)(nil),        // 0: spaceone.api.monitoring.v1.LogRequest
	(*LogDataInfo)(nil),       // 1: spaceone.api.monitoring.v1.LogDataInfo
	(*_struct.Struct)(nil),    // 2: google.protobuf.Struct
	(*_struct.ListValue)(nil), // 3: google.protobuf.ListValue
}
var file_spaceone_api_monitoring_v1_log_proto_depIdxs = []int32{
	2, // 0: spaceone.api.monitoring.v1.LogRequest.sort:type_name -> google.protobuf.Struct
	3, // 1: spaceone.api.monitoring.v1.LogDataInfo.results:type_name -> google.protobuf.ListValue
	0, // 2: spaceone.api.monitoring.v1.Log.list:input_type -> spaceone.api.monitoring.v1.LogRequest
	1, // 3: spaceone.api.monitoring.v1.Log.list:output_type -> spaceone.api.monitoring.v1.LogDataInfo
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spaceone_api_monitoring_v1_log_proto_init() }
func file_spaceone_api_monitoring_v1_log_proto_init() {
	if File_spaceone_api_monitoring_v1_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_monitoring_v1_log_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LogRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_monitoring_v1_log_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LogDataInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_monitoring_v1_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_monitoring_v1_log_proto_goTypes,
		DependencyIndexes: file_spaceone_api_monitoring_v1_log_proto_depIdxs,
		MessageInfos:      file_spaceone_api_monitoring_v1_log_proto_msgTypes,
	}.Build()
	File_spaceone_api_monitoring_v1_log_proto = out.File
	file_spaceone_api_monitoring_v1_log_proto_rawDesc = nil
	file_spaceone_api_monitoring_v1_log_proto_goTypes = nil
	file_spaceone_api_monitoring_v1_log_proto_depIdxs = nil
}
