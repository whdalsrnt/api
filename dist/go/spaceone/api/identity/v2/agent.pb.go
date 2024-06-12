// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: spaceone/api/identity/v2/agent.proto

package v2

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

type AgentInfo_State int32

const (
	AgentInfo_STATE_NONE AgentInfo_State = 0
	AgentInfo_ENABLED    AgentInfo_State = 1
	AgentInfo_DISABLED   AgentInfo_State = 2
	AgentInfo_EXPIRED    AgentInfo_State = 3
)

// Enum value maps for AgentInfo_State.
var (
	AgentInfo_State_name = map[int32]string{
		0: "STATE_NONE",
		1: "ENABLED",
		2: "DISABLED",
		3: "EXPIRED",
	}
	AgentInfo_State_value = map[string]int32{
		"STATE_NONE": 0,
		"ENABLED":    1,
		"DISABLED":   2,
		"EXPIRED":    3,
	}
)

func (x AgentInfo_State) Enum() *AgentInfo_State {
	p := new(AgentInfo_State)
	*p = x
	return p
}

func (x AgentInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_agent_proto_enumTypes[0].Descriptor()
}

func (AgentInfo_State) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_agent_proto_enumTypes[0]
}

func (x AgentInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentInfo_State.Descriptor instead.
func (AgentInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{2, 0}
}

type AgentInfo_RoleType int32

const (
	AgentInfo_ROLE_TYPE_NONE  AgentInfo_RoleType = 0
	AgentInfo_DOMAIN_ADMIN    AgentInfo_RoleType = 1
	AgentInfo_WORKSPACE_OWNER AgentInfo_RoleType = 2
)

// Enum value maps for AgentInfo_RoleType.
var (
	AgentInfo_RoleType_name = map[int32]string{
		0: "ROLE_TYPE_NONE",
		1: "DOMAIN_ADMIN",
		2: "WORKSPACE_OWNER",
	}
	AgentInfo_RoleType_value = map[string]int32{
		"ROLE_TYPE_NONE":  0,
		"DOMAIN_ADMIN":    1,
		"WORKSPACE_OWNER": 2,
	}
)

func (x AgentInfo_RoleType) Enum() *AgentInfo_RoleType {
	p := new(AgentInfo_RoleType)
	*p = x
	return p
}

func (x AgentInfo_RoleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentInfo_RoleType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_agent_proto_enumTypes[1].Descriptor()
}

func (AgentInfo_RoleType) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_agent_proto_enumTypes[1]
}

func (x AgentInfo_RoleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentInfo_RoleType.Descriptor instead.
func (AgentInfo_RoleType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{2, 1}
}

type AgentInfo_ResourceGroup int32

const (
	AgentInfo_GROUP_NONE AgentInfo_ResourceGroup = 0
	AgentInfo_DOMAIN     AgentInfo_ResourceGroup = 1
	AgentInfo_WORKSPACE  AgentInfo_ResourceGroup = 2
)

// Enum value maps for AgentInfo_ResourceGroup.
var (
	AgentInfo_ResourceGroup_name = map[int32]string{
		0: "GROUP_NONE",
		1: "DOMAIN",
		2: "WORKSPACE",
	}
	AgentInfo_ResourceGroup_value = map[string]int32{
		"GROUP_NONE": 0,
		"DOMAIN":     1,
		"WORKSPACE":  2,
	}
)

func (x AgentInfo_ResourceGroup) Enum() *AgentInfo_ResourceGroup {
	p := new(AgentInfo_ResourceGroup)
	*p = x
	return p
}

func (x AgentInfo_ResourceGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentInfo_ResourceGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_agent_proto_enumTypes[2].Descriptor()
}

func (AgentInfo_ResourceGroup) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_agent_proto_enumTypes[2]
}

func (x AgentInfo_ResourceGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentInfo_ResourceGroup.Descriptor instead.
func (AgentInfo_ResourceGroup) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{2, 2}
}

type AgentSearchQuery_State int32

const (
	AgentSearchQuery_STATE_NONE AgentSearchQuery_State = 0
	AgentSearchQuery_ENABLED    AgentSearchQuery_State = 1
	AgentSearchQuery_DISABLED   AgentSearchQuery_State = 2
	AgentSearchQuery_EXPIRED    AgentSearchQuery_State = 3
)

// Enum value maps for AgentSearchQuery_State.
var (
	AgentSearchQuery_State_name = map[int32]string{
		0: "STATE_NONE",
		1: "ENABLED",
		2: "DISABLED",
		3: "EXPIRED",
	}
	AgentSearchQuery_State_value = map[string]int32{
		"STATE_NONE": 0,
		"ENABLED":    1,
		"DISABLED":   2,
		"EXPIRED":    3,
	}
)

func (x AgentSearchQuery_State) Enum() *AgentSearchQuery_State {
	p := new(AgentSearchQuery_State)
	*p = x
	return p
}

func (x AgentSearchQuery_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentSearchQuery_State) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_agent_proto_enumTypes[3].Descriptor()
}

func (AgentSearchQuery_State) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_agent_proto_enumTypes[3]
}

func (x AgentSearchQuery_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentSearchQuery_State.Descriptor instead.
func (AgentSearchQuery_State) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{3, 0}
}

type CreateAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceAccountId string          `protobuf:"bytes,1,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	Options          *_struct.Struct `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *CreateAgentRequest) Reset() {
	*x = CreateAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAgentRequest) ProtoMessage() {}

func (x *CreateAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAgentRequest.ProtoReflect.Descriptor instead.
func (*CreateAgentRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAgentRequest) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *CreateAgentRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

type AgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceAccountId string `protobuf:"bytes,1,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
}

func (x *AgentRequest) Reset() {
	*x = AgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRequest) ProtoMessage() {}

func (x *AgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRequest.ProtoReflect.Descriptor instead.
func (*AgentRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{1}
}

func (x *AgentRequest) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

type AgentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId          string             `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Options          *_struct.Struct    `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	ClientSecret     string             `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	State            AgentInfo_State    `protobuf:"varint,4,opt,name=state,proto3,enum=spaceone.api.identity.v2.AgentInfo_State" json:"state,omitempty"`
	IsManaged        bool               `protobuf:"varint,5,opt,name=is_managed,json=isManaged,proto3" json:"is_managed,omitempty"`
	RoleType         AgentInfo_RoleType `protobuf:"varint,6,opt,name=role_type,json=roleType,proto3,enum=spaceone.api.identity.v2.AgentInfo_RoleType" json:"role_type,omitempty"`
	DomainId         string             `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId      string             `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ProjectId        string             `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceAccountId string             `protobuf:"bytes,24,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	AppId            string             `protobuf:"bytes,25,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	RoleId           string             `protobuf:"bytes,26,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	ClientId         string             `protobuf:"bytes,27,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	CreatedAt        string             `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiredAt        string             `protobuf:"bytes,32,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	LastAccessedAt   string             `protobuf:"bytes,33,opt,name=last_accessed_at,json=lastAccessedAt,proto3" json:"last_accessed_at,omitempty"`
}

func (x *AgentInfo) Reset() {
	*x = AgentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfo) ProtoMessage() {}

func (x *AgentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfo.ProtoReflect.Descriptor instead.
func (*AgentInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{2}
}

func (x *AgentInfo) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *AgentInfo) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *AgentInfo) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *AgentInfo) GetState() AgentInfo_State {
	if x != nil {
		return x.State
	}
	return AgentInfo_STATE_NONE
}

func (x *AgentInfo) GetIsManaged() bool {
	if x != nil {
		return x.IsManaged
	}
	return false
}

func (x *AgentInfo) GetRoleType() AgentInfo_RoleType {
	if x != nil {
		return x.RoleType
	}
	return AgentInfo_ROLE_TYPE_NONE
}

func (x *AgentInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *AgentInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *AgentInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *AgentInfo) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *AgentInfo) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *AgentInfo) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *AgentInfo) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AgentInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AgentInfo) GetExpiredAt() string {
	if x != nil {
		return x.ExpiredAt
	}
	return ""
}

func (x *AgentInfo) GetLastAccessedAt() string {
	if x != nil {
		return x.LastAccessedAt
	}
	return ""
}

type AgentSearchQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	AgentId string `protobuf:"bytes,2,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	// +optional
	ServiceAccountId string `protobuf:"bytes,3,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// +optional
	State AgentSearchQuery_State `protobuf:"varint,4,opt,name=state,proto3,enum=spaceone.api.identity.v2.AgentSearchQuery_State" json:"state,omitempty"`
}

func (x *AgentSearchQuery) Reset() {
	*x = AgentSearchQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentSearchQuery) ProtoMessage() {}

func (x *AgentSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentSearchQuery.ProtoReflect.Descriptor instead.
func (*AgentSearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{3}
}

func (x *AgentSearchQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *AgentSearchQuery) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *AgentSearchQuery) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *AgentSearchQuery) GetState() AgentSearchQuery_State {
	if x != nil {
		return x.State
	}
	return AgentSearchQuery_STATE_NONE
}

type AgentsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*AgentInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32        `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *AgentsInfo) Reset() {
	*x = AgentsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentsInfo) ProtoMessage() {}

func (x *AgentsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentsInfo.ProtoReflect.Descriptor instead.
func (*AgentsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{4}
}

func (x *AgentsInfo) GetResults() []*AgentInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *AgentsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_spaceone_api_identity_v2_agent_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_v2_agent_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x75, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3c, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0xaf, 0x06, 0x0a, 0x09, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x12, 0x49, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x21, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x3f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45,
	0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41,
	0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45,
	0x44, 0x10, 0x03, 0x22, 0x45, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x41, 0x44,
	0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41,
	0x43, 0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x02, 0x22, 0x3a, 0x0a, 0x0d, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x0a, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x4f, 0x52, 0x4b, 0x53,
	0x50, 0x41, 0x43, 0x45, 0x10, 0x02, 0x22, 0x97, 0x02, 0x0a, 0x10, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x3f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42,
	0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03,
	0x22, 0x6c, 0x0a, 0x0a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xf2,
	0x06, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x81, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01,
	0x2a, 0x22, 0x19, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x7b, 0x0a, 0x06,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x7d, 0x0a, 0x07, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x0a, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22,
	0x1d, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x6e,
	0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x75,
	0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x7c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_v2_agent_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v2_agent_proto_rawDescData = file_spaceone_api_identity_v2_agent_proto_rawDesc
)

func file_spaceone_api_identity_v2_agent_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v2_agent_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v2_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_v2_agent_proto_rawDescData)
	})
	return file_spaceone_api_identity_v2_agent_proto_rawDescData
}

var file_spaceone_api_identity_v2_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_spaceone_api_identity_v2_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_identity_v2_agent_proto_goTypes = []any{
	(AgentInfo_State)(0),         // 0: spaceone.api.identity.v2.AgentInfo.State
	(AgentInfo_RoleType)(0),      // 1: spaceone.api.identity.v2.AgentInfo.RoleType
	(AgentInfo_ResourceGroup)(0), // 2: spaceone.api.identity.v2.AgentInfo.ResourceGroup
	(AgentSearchQuery_State)(0),  // 3: spaceone.api.identity.v2.AgentSearchQuery.State
	(*CreateAgentRequest)(nil),   // 4: spaceone.api.identity.v2.CreateAgentRequest
	(*AgentRequest)(nil),         // 5: spaceone.api.identity.v2.AgentRequest
	(*AgentInfo)(nil),            // 6: spaceone.api.identity.v2.AgentInfo
	(*AgentSearchQuery)(nil),     // 7: spaceone.api.identity.v2.AgentSearchQuery
	(*AgentsInfo)(nil),           // 8: spaceone.api.identity.v2.AgentsInfo
	(*_struct.Struct)(nil),       // 9: google.protobuf.Struct
	(*v2.Query)(nil),             // 10: spaceone.api.core.v2.Query
	(*empty.Empty)(nil),          // 11: google.protobuf.Empty
}
var file_spaceone_api_identity_v2_agent_proto_depIdxs = []int32{
	9,  // 0: spaceone.api.identity.v2.CreateAgentRequest.options:type_name -> google.protobuf.Struct
	9,  // 1: spaceone.api.identity.v2.AgentInfo.options:type_name -> google.protobuf.Struct
	0,  // 2: spaceone.api.identity.v2.AgentInfo.state:type_name -> spaceone.api.identity.v2.AgentInfo.State
	1,  // 3: spaceone.api.identity.v2.AgentInfo.role_type:type_name -> spaceone.api.identity.v2.AgentInfo.RoleType
	10, // 4: spaceone.api.identity.v2.AgentSearchQuery.query:type_name -> spaceone.api.core.v2.Query
	3,  // 5: spaceone.api.identity.v2.AgentSearchQuery.state:type_name -> spaceone.api.identity.v2.AgentSearchQuery.State
	6,  // 6: spaceone.api.identity.v2.AgentsInfo.results:type_name -> spaceone.api.identity.v2.AgentInfo
	4,  // 7: spaceone.api.identity.v2.Agent.create:input_type -> spaceone.api.identity.v2.CreateAgentRequest
	5,  // 8: spaceone.api.identity.v2.Agent.enable:input_type -> spaceone.api.identity.v2.AgentRequest
	5,  // 9: spaceone.api.identity.v2.Agent.disable:input_type -> spaceone.api.identity.v2.AgentRequest
	5,  // 10: spaceone.api.identity.v2.Agent.regenerate:input_type -> spaceone.api.identity.v2.AgentRequest
	5,  // 11: spaceone.api.identity.v2.Agent.delete:input_type -> spaceone.api.identity.v2.AgentRequest
	5,  // 12: spaceone.api.identity.v2.Agent.get:input_type -> spaceone.api.identity.v2.AgentRequest
	7,  // 13: spaceone.api.identity.v2.Agent.list:input_type -> spaceone.api.identity.v2.AgentSearchQuery
	6,  // 14: spaceone.api.identity.v2.Agent.create:output_type -> spaceone.api.identity.v2.AgentInfo
	6,  // 15: spaceone.api.identity.v2.Agent.enable:output_type -> spaceone.api.identity.v2.AgentInfo
	6,  // 16: spaceone.api.identity.v2.Agent.disable:output_type -> spaceone.api.identity.v2.AgentInfo
	6,  // 17: spaceone.api.identity.v2.Agent.regenerate:output_type -> spaceone.api.identity.v2.AgentInfo
	11, // 18: spaceone.api.identity.v2.Agent.delete:output_type -> google.protobuf.Empty
	6,  // 19: spaceone.api.identity.v2.Agent.get:output_type -> spaceone.api.identity.v2.AgentInfo
	8,  // 20: spaceone.api.identity.v2.Agent.list:output_type -> spaceone.api.identity.v2.AgentsInfo
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v2_agent_proto_init() }
func file_spaceone_api_identity_v2_agent_proto_init() {
	if File_spaceone_api_identity_v2_agent_proto != nil {
		return
	}
	file_spaceone_api_identity_v2_app_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_identity_v2_agent_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAgentRequest); i {
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
		file_spaceone_api_identity_v2_agent_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AgentRequest); i {
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
		file_spaceone_api_identity_v2_agent_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AgentInfo); i {
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
		file_spaceone_api_identity_v2_agent_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AgentSearchQuery); i {
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
		file_spaceone_api_identity_v2_agent_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AgentsInfo); i {
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
			RawDescriptor: file_spaceone_api_identity_v2_agent_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v2_agent_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v2_agent_proto_depIdxs,
		EnumInfos:         file_spaceone_api_identity_v2_agent_proto_enumTypes,
		MessageInfos:      file_spaceone_api_identity_v2_agent_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v2_agent_proto = out.File
	file_spaceone_api_identity_v2_agent_proto_rawDesc = nil
	file_spaceone_api_identity_v2_agent_proto_goTypes = nil
	file_spaceone_api_identity_v2_agent_proto_depIdxs = nil
}
