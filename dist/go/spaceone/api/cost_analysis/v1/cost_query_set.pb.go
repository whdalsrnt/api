// A CostQuerySet is a set of saved queries a User frequently uses as a setting.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: spaceone/api/cost_analysis/v1/cost_query_set.proto

package v1

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
	empty "github.com/golang/protobuf/ptypes/empty"
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

//	{
//	       "data_source_id": "ds-085d1e872789",
//	       "name": "project_provider_region",
//	       "options": {}
//	}
type CreateCostQuerySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSourceId string          `protobuf:"bytes,1,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	Name         string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Options      *_struct.Struct `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreateCostQuerySetRequest) Reset() {
	*x = CreateCostQuerySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCostQuerySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCostQuerySetRequest) ProtoMessage() {}

func (x *CreateCostQuerySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCostQuerySetRequest.ProtoReflect.Descriptor instead.
func (*CreateCostQuerySetRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCostQuerySetRequest) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *CreateCostQuerySetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCostQuerySetRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *CreateCostQuerySetRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

//	{
//	       "cost_query_set_id": "query-76a58ea5d02c",
//	       "name": "project_provider_region",
//	       "options": {},
//	       "tags": {},
//	       "user_id": "test@cloudforet.io",
//	       "domain_id": "domain-58010aa2e451",
//	       "created_at": "2022-07-19T06:11:03.701Z",
//	       "updated_at": "2022-07-19T06:11:03.701Z"
//	}
type UpdateCostQuerySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostQuerySetId string `protobuf:"bytes,1,opt,name=cost_query_set_id,json=costQuerySetId,proto3" json:"cost_query_set_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Options *_struct.Struct `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UpdateCostQuerySetRequest) Reset() {
	*x = UpdateCostQuerySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCostQuerySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCostQuerySetRequest) ProtoMessage() {}

func (x *UpdateCostQuerySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCostQuerySetRequest.ProtoReflect.Descriptor instead.
func (*UpdateCostQuerySetRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCostQuerySetRequest) GetCostQuerySetId() string {
	if x != nil {
		return x.CostQuerySetId
	}
	return ""
}

func (x *UpdateCostQuerySetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCostQuerySetRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *UpdateCostQuerySetRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

//	{
//	   "cost_query_set_id": "query-16ae671dc8fb",
//	    "domain_id": "domain-58010aa2e451"
//	}
type CostQuerySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostQuerySetId string `protobuf:"bytes,1,opt,name=cost_query_set_id,json=costQuerySetId,proto3" json:"cost_query_set_id,omitempty"`
}

func (x *CostQuerySetRequest) Reset() {
	*x = CostQuerySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostQuerySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostQuerySetRequest) ProtoMessage() {}

func (x *CostQuerySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostQuerySetRequest.ProtoReflect.Descriptor instead.
func (*CostQuerySetRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescGZIP(), []int{2}
}

func (x *CostQuerySetRequest) GetCostQuerySetId() string {
	if x != nil {
		return x.CostQuerySetId
	}
	return ""
}

//	{
//	   "query": {}
//	}
type CostQuerySetQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query        *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DataSourceId string    `protobuf:"bytes,2,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CostQuerySetQuery) Reset() {
	*x = CostQuerySetQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostQuerySetQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostQuerySetQuery) ProtoMessage() {}

func (x *CostQuerySetQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostQuerySetQuery.ProtoReflect.Descriptor instead.
func (*CostQuerySetQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescGZIP(), []int{3}
}

func (x *CostQuerySetQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CostQuerySetQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *CostQuerySetQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//	{
//	       "cost_query_set_id": "query-76a58ea5d02c",
//	       "name": "project_provider_region",
//	       "options": {},
//	       "tags": {},
//	       "user_id": "test@cloudforet.io",
//	       "data_source_id": "ds-085d1e872789",
//	       "domain_id": "domain-58010aa2e451",
//	       "created_at": "2022-07-19T06:11:03.701Z",
//	       "updated_at": "2022-07-19T06:11:03.701Z"
//	}
type CostQuerySetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostQuerySetId string          `protobuf:"bytes,1,opt,name=cost_query_set_id,json=costQuerySetId,proto3" json:"cost_query_set_id,omitempty"`
	Name           string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Options        *_struct.Struct `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	Tags           *_struct.Struct `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId       string          `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId    string          `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	UserId         string          `protobuf:"bytes,23,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DataSourceId   string          `protobuf:"bytes,24,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	CreatedAt      string          `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      string          `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CostQuerySetInfo) Reset() {
	*x = CostQuerySetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostQuerySetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostQuerySetInfo) ProtoMessage() {}

func (x *CostQuerySetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostQuerySetInfo.ProtoReflect.Descriptor instead.
func (*CostQuerySetInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescGZIP(), []int{4}
}

func (x *CostQuerySetInfo) GetCostQuerySetId() string {
	if x != nil {
		return x.CostQuerySetId
	}
	return ""
}

func (x *CostQuerySetInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CostQuerySetInfo) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *CostQuerySetInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CostQuerySetInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *CostQuerySetInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CostQuerySetInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CostQuerySetInfo) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *CostQuerySetInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CostQuerySetInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//	{
//	       "results": [
//	           {
//	               "cost_query_set_id": "query-16ae671dc8fb",
//	               "name": "3 month product pie chart",
//	               "options": {},
//	               "tags": {},
//	               "user_id": "yuda@mz.co.kr",
//	               "data_source_id": "ds-085d1e872789",
//	               "domain_id": "domain-58010aa2e451",
//	               "created_at": "2022-03-08T03:37:31.404Z",
//	               "updated_at": "2022-03-08T03:37:31.404Z"
//	           },
//	           {
//	               "cost_query_set_id": "query-d90addf25e4b",
//	               "name": "6 month project group",
//	               "options": {},
//	               "tags": {},
//	               "user_id": "yuda@mz.co.kr",
//	               "data_source_id": "ds-085d1e872789",
//	               "domain_id": "domain-58010aa2e451",
//	               "created_at": "2022-03-14T09:29:54.306Z",
//	               "updated_at": "2022-03-14T09:29:54.306Z"
//	           }
//	       ],
//	       "total_count": 34
//	}
type CostQuerySetsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*CostQuerySetInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32               `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *CostQuerySetsInfo) Reset() {
	*x = CostQuerySetsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostQuerySetsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostQuerySetsInfo) ProtoMessage() {}

func (x *CostQuerySetsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostQuerySetsInfo.ProtoReflect.Descriptor instead.
func (*CostQuerySetsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescGZIP(), []int{5}
}

func (x *CostQuerySetsInfo) GetResults() []*CostQuerySetInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *CostQuerySetsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CostQuerySetStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query        *v2.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DataSourceId string              `protobuf:"bytes,2,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
}

func (x *CostQuerySetStatQuery) Reset() {
	*x = CostQuerySetStatQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostQuerySetStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostQuerySetStatQuery) ProtoMessage() {}

func (x *CostQuerySetStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostQuerySetStatQuery.ProtoReflect.Descriptor instead.
func (*CostQuerySetStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescGZIP(), []int{6}
}

func (x *CostQuerySetStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CostQuerySetStatQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

var File_spaceone_api_cost_analysis_v1_cost_query_set_proto protoreflect.FileDescriptor

var file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5,
	0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xba, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x11, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x22, 0x40, 0x0a, 0x13, 0x43, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x11, 0x63, 0x6f,
	0x73, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x74, 0x49, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x73, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x24,
	0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xee, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x73,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a,
	0x11, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7f, 0x0a, 0x11, 0x43, 0x6f, 0x73,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7a, 0x0a, 0x15, 0x43, 0x6f,
	0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x32, 0xb4, 0x07, 0x0a, 0x0c, 0x43, 0x6f, 0x73, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x12, 0xa7, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x38, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x73,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x32, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x73, 0x74,
	0x2d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0xa7, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x38, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a,
	0x01, 0x2a, 0x22, 0x27, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2d, 0x73, 0x65, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x06,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f,
	0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x74, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x9b, 0x01, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x32,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24,
	0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x74,
	0x2f, 0x67, 0x65, 0x74, 0x12, 0x9c, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x30, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x73, 0x74, 0x2d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x74, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x87, 0x01, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x34, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x73,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x30, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x44, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescData = file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDesc
)

func file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescData)
	})
	return file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDescData
}

var file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_cost_analysis_v1_cost_query_set_proto_goTypes = []any{
	(*CreateCostQuerySetRequest)(nil), // 0: spaceone.api.cost_analysis.v1.CreateCostQuerySetRequest
	(*UpdateCostQuerySetRequest)(nil), // 1: spaceone.api.cost_analysis.v1.UpdateCostQuerySetRequest
	(*CostQuerySetRequest)(nil),       // 2: spaceone.api.cost_analysis.v1.CostQuerySetRequest
	(*CostQuerySetQuery)(nil),         // 3: spaceone.api.cost_analysis.v1.CostQuerySetQuery
	(*CostQuerySetInfo)(nil),          // 4: spaceone.api.cost_analysis.v1.CostQuerySetInfo
	(*CostQuerySetsInfo)(nil),         // 5: spaceone.api.cost_analysis.v1.CostQuerySetsInfo
	(*CostQuerySetStatQuery)(nil),     // 6: spaceone.api.cost_analysis.v1.CostQuerySetStatQuery
	(*_struct.Struct)(nil),            // 7: google.protobuf.Struct
	(*v2.Query)(nil),                  // 8: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),        // 9: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),               // 10: google.protobuf.Empty
}
var file_spaceone_api_cost_analysis_v1_cost_query_set_proto_depIdxs = []int32{
	7,  // 0: spaceone.api.cost_analysis.v1.CreateCostQuerySetRequest.options:type_name -> google.protobuf.Struct
	7,  // 1: spaceone.api.cost_analysis.v1.CreateCostQuerySetRequest.tags:type_name -> google.protobuf.Struct
	7,  // 2: spaceone.api.cost_analysis.v1.UpdateCostQuerySetRequest.options:type_name -> google.protobuf.Struct
	7,  // 3: spaceone.api.cost_analysis.v1.UpdateCostQuerySetRequest.tags:type_name -> google.protobuf.Struct
	8,  // 4: spaceone.api.cost_analysis.v1.CostQuerySetQuery.query:type_name -> spaceone.api.core.v2.Query
	7,  // 5: spaceone.api.cost_analysis.v1.CostQuerySetInfo.options:type_name -> google.protobuf.Struct
	7,  // 6: spaceone.api.cost_analysis.v1.CostQuerySetInfo.tags:type_name -> google.protobuf.Struct
	4,  // 7: spaceone.api.cost_analysis.v1.CostQuerySetsInfo.results:type_name -> spaceone.api.cost_analysis.v1.CostQuerySetInfo
	9,  // 8: spaceone.api.cost_analysis.v1.CostQuerySetStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 9: spaceone.api.cost_analysis.v1.CostQuerySet.create:input_type -> spaceone.api.cost_analysis.v1.CreateCostQuerySetRequest
	1,  // 10: spaceone.api.cost_analysis.v1.CostQuerySet.update:input_type -> spaceone.api.cost_analysis.v1.UpdateCostQuerySetRequest
	2,  // 11: spaceone.api.cost_analysis.v1.CostQuerySet.delete:input_type -> spaceone.api.cost_analysis.v1.CostQuerySetRequest
	2,  // 12: spaceone.api.cost_analysis.v1.CostQuerySet.get:input_type -> spaceone.api.cost_analysis.v1.CostQuerySetRequest
	3,  // 13: spaceone.api.cost_analysis.v1.CostQuerySet.list:input_type -> spaceone.api.cost_analysis.v1.CostQuerySetQuery
	6,  // 14: spaceone.api.cost_analysis.v1.CostQuerySet.stat:input_type -> spaceone.api.cost_analysis.v1.CostQuerySetStatQuery
	4,  // 15: spaceone.api.cost_analysis.v1.CostQuerySet.create:output_type -> spaceone.api.cost_analysis.v1.CostQuerySetInfo
	4,  // 16: spaceone.api.cost_analysis.v1.CostQuerySet.update:output_type -> spaceone.api.cost_analysis.v1.CostQuerySetInfo
	10, // 17: spaceone.api.cost_analysis.v1.CostQuerySet.delete:output_type -> google.protobuf.Empty
	4,  // 18: spaceone.api.cost_analysis.v1.CostQuerySet.get:output_type -> spaceone.api.cost_analysis.v1.CostQuerySetInfo
	5,  // 19: spaceone.api.cost_analysis.v1.CostQuerySet.list:output_type -> spaceone.api.cost_analysis.v1.CostQuerySetsInfo
	7,  // 20: spaceone.api.cost_analysis.v1.CostQuerySet.stat:output_type -> google.protobuf.Struct
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_v1_cost_query_set_proto_init() }
func file_spaceone_api_cost_analysis_v1_cost_query_set_proto_init() {
	if File_spaceone_api_cost_analysis_v1_cost_query_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCostQuerySetRequest); i {
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
		file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCostQuerySetRequest); i {
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
		file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CostQuerySetRequest); i {
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
		file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CostQuerySetQuery); i {
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
		file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CostQuerySetInfo); i {
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
		file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CostQuerySetsInfo); i {
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
		file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CostQuerySetStatQuery); i {
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
			RawDescriptor: file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_v1_cost_query_set_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_v1_cost_query_set_proto_depIdxs,
		MessageInfos:      file_spaceone_api_cost_analysis_v1_cost_query_set_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_v1_cost_query_set_proto = out.File
	file_spaceone_api_cost_analysis_v1_cost_query_set_proto_rawDesc = nil
	file_spaceone_api_cost_analysis_v1_cost_query_set_proto_goTypes = nil
	file_spaceone_api_cost_analysis_v1_cost_query_set_proto_depIdxs = nil
}
