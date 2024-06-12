// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: spaceone/api/cost_analysis/plugin/cost.proto

package plugin

import (
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

type GetLinkedAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options    *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	Schema     string          `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	SecretData *_struct.Struct `protobuf:"bytes,3,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	DomainId   string          `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *GetLinkedAccountsRequest) Reset() {
	*x = GetLinkedAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkedAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkedAccountsRequest) ProtoMessage() {}

func (x *GetLinkedAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkedAccountsRequest.ProtoReflect.Descriptor instead.
func (*GetLinkedAccountsRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescGZIP(), []int{0}
}

func (x *GetLinkedAccountsRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *GetLinkedAccountsRequest) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *GetLinkedAccountsRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *GetLinkedAccountsRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type GetDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options    *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	SecretData *_struct.Struct `protobuf:"bytes,2,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	// +optional
	Schema string `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	// +optional
	TaskOptions *_struct.Struct `protobuf:"bytes,4,opt,name=task_options,json=taskOptions,proto3" json:"task_options,omitempty"`
	DomainId    string          `protobuf:"bytes,5,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *GetDataRequest) Reset() {
	*x = GetDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataRequest) ProtoMessage() {}

func (x *GetDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataRequest.ProtoReflect.Descriptor instead.
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescGZIP(), []int{1}
}

func (x *GetDataRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *GetDataRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *GetDataRequest) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *GetDataRequest) GetTaskOptions() *_struct.Struct {
	if x != nil {
		return x.TaskOptions
	}
	return nil
}

func (x *GetDataRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[2]
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
	return file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescGZIP(), []int{2}
}

func (x *AccountInfo) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *AccountInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CostInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cost           float64         `protobuf:"fixed64,1,opt,name=cost,proto3" json:"cost,omitempty"`
	UsageQuantity  float64         `protobuf:"fixed64,2,opt,name=usage_quantity,json=usageQuantity,proto3" json:"usage_quantity,omitempty"`
	UsageUnit      string          `protobuf:"bytes,3,opt,name=usage_unit,json=usageUnit,proto3" json:"usage_unit,omitempty"`
	Provider       string          `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	RegionCode     string          `protobuf:"bytes,5,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	Product        string          `protobuf:"bytes,6,opt,name=product,proto3" json:"product,omitempty"`
	UsageType      string          `protobuf:"bytes,7,opt,name=usage_type,json=usageType,proto3" json:"usage_type,omitempty"`
	Resource       string          `protobuf:"bytes,8,opt,name=resource,proto3" json:"resource,omitempty"`
	Tags           *_struct.Struct `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
	AdditionalInfo *_struct.Struct `protobuf:"bytes,13,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	Data           *_struct.Struct `protobuf:"bytes,14,opt,name=data,proto3" json:"data,omitempty"`
	BilledDate     string          `protobuf:"bytes,21,opt,name=billed_date,json=billedDate,proto3" json:"billed_date,omitempty"`
}

func (x *CostInfo) Reset() {
	*x = CostInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostInfo) ProtoMessage() {}

func (x *CostInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostInfo.ProtoReflect.Descriptor instead.
func (*CostInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescGZIP(), []int{3}
}

func (x *CostInfo) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *CostInfo) GetUsageQuantity() float64 {
	if x != nil {
		return x.UsageQuantity
	}
	return 0
}

func (x *CostInfo) GetUsageUnit() string {
	if x != nil {
		return x.UsageUnit
	}
	return ""
}

func (x *CostInfo) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CostInfo) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *CostInfo) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *CostInfo) GetUsageType() string {
	if x != nil {
		return x.UsageType
	}
	return ""
}

func (x *CostInfo) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *CostInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CostInfo) GetAdditionalInfo() *_struct.Struct {
	if x != nil {
		return x.AdditionalInfo
	}
	return nil
}

func (x *CostInfo) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CostInfo) GetBilledDate() string {
	if x != nil {
		return x.BilledDate
	}
	return ""
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
		mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsInfo) ProtoMessage() {}

func (x *AccountsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[4]
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
	return file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescGZIP(), []int{4}
}

func (x *AccountsInfo) GetResults() []*AccountInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

type CostsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*CostInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *CostsInfo) Reset() {
	*x = CostsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostsInfo) ProtoMessage() {}

func (x *CostsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostsInfo.ProtoReflect.Descriptor instead.
func (*CostsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescGZIP(), []int{5}
}

func (x *CostsInfo) GetResults() []*CostInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_spaceone_api_cost_analysis_plugin_cost_proto protoreflect.FileDescriptor

var file_spaceone_api_cost_analysis_plugin_cost_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbc, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xee,
	0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x3a, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22,
	0x40, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xb3, 0x03, 0x0a, 0x08, 0x43, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x40, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x69, 0x6c, 0x6c, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x69, 0x6c,
	0x6c, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x58, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x22, 0x52, 0x0a, 0x09, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0xff, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x85,
	0x01, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x3b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74,
	0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescData = file_spaceone_api_cost_analysis_plugin_cost_proto_rawDesc
)

func file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescData)
	})
	return file_spaceone_api_cost_analysis_plugin_cost_proto_rawDescData
}

var file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_spaceone_api_cost_analysis_plugin_cost_proto_goTypes = []any{
	(*GetLinkedAccountsRequest)(nil), // 0: spaceone.api.cost_analysis.plugin.GetLinkedAccountsRequest
	(*GetDataRequest)(nil),           // 1: spaceone.api.cost_analysis.plugin.GetDataRequest
	(*AccountInfo)(nil),              // 2: spaceone.api.cost_analysis.plugin.AccountInfo
	(*CostInfo)(nil),                 // 3: spaceone.api.cost_analysis.plugin.CostInfo
	(*AccountsInfo)(nil),             // 4: spaceone.api.cost_analysis.plugin.AccountsInfo
	(*CostsInfo)(nil),                // 5: spaceone.api.cost_analysis.plugin.CostsInfo
	(*_struct.Struct)(nil),           // 6: google.protobuf.Struct
}
var file_spaceone_api_cost_analysis_plugin_cost_proto_depIdxs = []int32{
	6,  // 0: spaceone.api.cost_analysis.plugin.GetLinkedAccountsRequest.options:type_name -> google.protobuf.Struct
	6,  // 1: spaceone.api.cost_analysis.plugin.GetLinkedAccountsRequest.secret_data:type_name -> google.protobuf.Struct
	6,  // 2: spaceone.api.cost_analysis.plugin.GetDataRequest.options:type_name -> google.protobuf.Struct
	6,  // 3: spaceone.api.cost_analysis.plugin.GetDataRequest.secret_data:type_name -> google.protobuf.Struct
	6,  // 4: spaceone.api.cost_analysis.plugin.GetDataRequest.task_options:type_name -> google.protobuf.Struct
	6,  // 5: spaceone.api.cost_analysis.plugin.CostInfo.tags:type_name -> google.protobuf.Struct
	6,  // 6: spaceone.api.cost_analysis.plugin.CostInfo.additional_info:type_name -> google.protobuf.Struct
	6,  // 7: spaceone.api.cost_analysis.plugin.CostInfo.data:type_name -> google.protobuf.Struct
	2,  // 8: spaceone.api.cost_analysis.plugin.AccountsInfo.results:type_name -> spaceone.api.cost_analysis.plugin.AccountInfo
	3,  // 9: spaceone.api.cost_analysis.plugin.CostsInfo.results:type_name -> spaceone.api.cost_analysis.plugin.CostInfo
	0,  // 10: spaceone.api.cost_analysis.plugin.Cost.get_linked_accounts:input_type -> spaceone.api.cost_analysis.plugin.GetLinkedAccountsRequest
	1,  // 11: spaceone.api.cost_analysis.plugin.Cost.get_data:input_type -> spaceone.api.cost_analysis.plugin.GetDataRequest
	4,  // 12: spaceone.api.cost_analysis.plugin.Cost.get_linked_accounts:output_type -> spaceone.api.cost_analysis.plugin.AccountsInfo
	5,  // 13: spaceone.api.cost_analysis.plugin.Cost.get_data:output_type -> spaceone.api.cost_analysis.plugin.CostsInfo
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_plugin_cost_proto_init() }
func file_spaceone_api_cost_analysis_plugin_cost_proto_init() {
	if File_spaceone_api_cost_analysis_plugin_cost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetLinkedAccountsRequest); i {
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
		file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetDataRequest); i {
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
		file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CostInfo); i {
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
		file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[4].Exporter = func(v any, i int) any {
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
		file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CostsInfo); i {
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
			RawDescriptor: file_spaceone_api_cost_analysis_plugin_cost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_plugin_cost_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_plugin_cost_proto_depIdxs,
		MessageInfos:      file_spaceone_api_cost_analysis_plugin_cost_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_plugin_cost_proto = out.File
	file_spaceone_api_cost_analysis_plugin_cost_proto_rawDesc = nil
	file_spaceone_api_cost_analysis_plugin_cost_proto_goTypes = nil
	file_spaceone_api_cost_analysis_plugin_cost_proto_depIdxs = nil
}
