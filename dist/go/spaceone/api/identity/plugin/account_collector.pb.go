// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.6.1
// source: spaceone/api/identity/plugin/account_collector.proto

package plugin

import (
	_ "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options  *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	DomainId string          `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_plugin_account_collector_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *InitRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	// +optional
	SchemaId string `protobuf:"bytes,2,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	// +optional
	SecretData *_struct.Struct `protobuf:"bytes,3,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	DomainId   string          `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_plugin_account_collector_proto_rawDescGZIP(), []int{1}
}

func (x *SyncRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *SyncRequest) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

func (x *SyncRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *SyncRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data           *_struct.Struct   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	ResourceId     string            `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	SecretSchemaId string            `protobuf:"bytes,4,opt,name=secret_schema_id,json=secretSchemaId,proto3" json:"secret_schema_id,omitempty"`
	SecretData     *_struct.Struct   `protobuf:"bytes,5,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	Tags           *_struct.Struct   `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	Location       []*_struct.Struct `protobuf:"bytes,7,rep,name=location,proto3" json:"location,omitempty"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfo.ProtoReflect.Descriptor instead.
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_plugin_account_collector_proto_rawDescGZIP(), []int{2}
}

func (x *AccountInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountInfo) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AccountInfo) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *AccountInfo) GetSecretSchemaId() string {
	if x != nil {
		return x.SecretSchemaId
	}
	return ""
}

func (x *AccountInfo) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *AccountInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *AccountInfo) GetLocation() []*_struct.Struct {
	if x != nil {
		return x.Location
	}
	return nil
}

type AccountsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*AccountInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *AccountsInfo) Reset() {
	*x = AccountsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsInfo) ProtoMessage() {}

func (x *AccountsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsInfo.ProtoReflect.Descriptor instead.
func (*AccountsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_plugin_account_collector_proto_rawDescGZIP(), []int{3}
}

func (x *AccountsInfo) GetResults() []*AccountInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

type PluginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *_struct.Struct `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *PluginInfo) Reset() {
	*x = PluginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfo) ProtoMessage() {}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfo.ProtoReflect.Descriptor instead.
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_plugin_account_collector_proto_rawDescGZIP(), []int{4}
}

func (x *PluginInfo) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_spaceone_api_identity_plugin_account_collector_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_plugin_account_collector_proto_rawDesc = []byte{
	0x0a, 0x34, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5d, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xb4,
	0x01, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x12, 0x38,
	0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xb5, 0x02, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49,
	0x64, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a,
	0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x43, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x22, 0x41, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0xd2, 0x01, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x5d, 0x0a, 0x04, 0x69, 0x6e,
	0x69, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x04, 0x73, 0x79, 0x6e,
	0x63, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f,
	0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_plugin_account_collector_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_plugin_account_collector_proto_rawDescData = file_spaceone_api_identity_plugin_account_collector_proto_rawDesc
)

func file_spaceone_api_identity_plugin_account_collector_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_plugin_account_collector_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_plugin_account_collector_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_plugin_account_collector_proto_rawDescData)
	})
	return file_spaceone_api_identity_plugin_account_collector_proto_rawDescData
}

var file_spaceone_api_identity_plugin_account_collector_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_identity_plugin_account_collector_proto_goTypes = []interface{}{
	(*InitRequest)(nil),    // 0: spaceone.api.identity.plugin.InitRequest
	(*SyncRequest)(nil),    // 1: spaceone.api.identity.plugin.SyncRequest
	(*AccountInfo)(nil),    // 2: spaceone.api.identity.plugin.AccountInfo
	(*AccountsInfo)(nil),   // 3: spaceone.api.identity.plugin.AccountsInfo
	(*PluginInfo)(nil),     // 4: spaceone.api.identity.plugin.PluginInfo
	(*_struct.Struct)(nil), // 5: google.protobuf.Struct
}
var file_spaceone_api_identity_plugin_account_collector_proto_depIdxs = []int32{
	5,  // 0: spaceone.api.identity.plugin.InitRequest.options:type_name -> google.protobuf.Struct
	5,  // 1: spaceone.api.identity.plugin.SyncRequest.options:type_name -> google.protobuf.Struct
	5,  // 2: spaceone.api.identity.plugin.SyncRequest.secret_data:type_name -> google.protobuf.Struct
	5,  // 3: spaceone.api.identity.plugin.AccountInfo.data:type_name -> google.protobuf.Struct
	5,  // 4: spaceone.api.identity.plugin.AccountInfo.secret_data:type_name -> google.protobuf.Struct
	5,  // 5: spaceone.api.identity.plugin.AccountInfo.tags:type_name -> google.protobuf.Struct
	5,  // 6: spaceone.api.identity.plugin.AccountInfo.location:type_name -> google.protobuf.Struct
	2,  // 7: spaceone.api.identity.plugin.AccountsInfo.results:type_name -> spaceone.api.identity.plugin.AccountInfo
	5,  // 8: spaceone.api.identity.plugin.PluginInfo.metadata:type_name -> google.protobuf.Struct
	0,  // 9: spaceone.api.identity.plugin.AccountCollector.init:input_type -> spaceone.api.identity.plugin.InitRequest
	1,  // 10: spaceone.api.identity.plugin.AccountCollector.sync:input_type -> spaceone.api.identity.plugin.SyncRequest
	4,  // 11: spaceone.api.identity.plugin.AccountCollector.init:output_type -> spaceone.api.identity.plugin.PluginInfo
	3,  // 12: spaceone.api.identity.plugin.AccountCollector.sync:output_type -> spaceone.api.identity.plugin.AccountsInfo
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_plugin_account_collector_proto_init() }
func file_spaceone_api_identity_plugin_account_collector_proto_init() {
	if File_spaceone_api_identity_plugin_account_collector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfo); i {
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
		file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountsInfo); i {
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
		file_spaceone_api_identity_plugin_account_collector_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginInfo); i {
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
			RawDescriptor: file_spaceone_api_identity_plugin_account_collector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_plugin_account_collector_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_plugin_account_collector_proto_depIdxs,
		MessageInfos:      file_spaceone_api_identity_plugin_account_collector_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_plugin_account_collector_proto = out.File
	file_spaceone_api_identity_plugin_account_collector_proto_rawDesc = nil
	file_spaceone_api_identity_plugin_account_collector_proto_goTypes = nil
	file_spaceone_api_identity_plugin_account_collector_proto_depIdxs = nil
}
