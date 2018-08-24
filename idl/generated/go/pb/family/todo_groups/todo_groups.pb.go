// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/family/todo_groups/todo_groups.proto

package todo_groups // import "github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TodoGroupPermitType int32

const (
	TodoGroupPermitType_OWNER  TodoGroupPermitType = 0
	TodoGroupPermitType_EDITOR TodoGroupPermitType = 1
	TodoGroupPermitType_VIEWER TodoGroupPermitType = 2
)

var TodoGroupPermitType_name = map[int32]string{
	0: "OWNER",
	1: "EDITOR",
	2: "VIEWER",
}
var TodoGroupPermitType_value = map[string]int32{
	"OWNER":  0,
	"EDITOR": 1,
	"VIEWER": 2,
}

func (x TodoGroupPermitType) String() string {
	return proto.EnumName(TodoGroupPermitType_name, int32(x))
}
func (TodoGroupPermitType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{0}
}

type TodoGroup struct {
	TodoGroupId          string   `protobuf:"bytes,1,opt,name=todo_group_id,json=todoGroupId" json:"todo_group_id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	CreatedBy            string   `protobuf:"bytes,4,opt,name=created_by,json=createdBy" json:"created_by,omitempty"`
	CreatedAt            int64    `protobuf:"varint,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,6,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Order                string   `protobuf:"bytes,7,opt,name=order" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoGroup) Reset()         { *m = TodoGroup{} }
func (m *TodoGroup) String() string { return proto.CompactTextString(m) }
func (*TodoGroup) ProtoMessage()    {}
func (*TodoGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{0}
}
func (m *TodoGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoGroup.Unmarshal(m, b)
}
func (m *TodoGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoGroup.Marshal(b, m, deterministic)
}
func (dst *TodoGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoGroup.Merge(dst, src)
}
func (m *TodoGroup) XXX_Size() int {
	return xxx_messageInfo_TodoGroup.Size(m)
}
func (m *TodoGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoGroup.DiscardUnknown(m)
}

var xxx_messageInfo_TodoGroup proto.InternalMessageInfo

func (m *TodoGroup) GetTodoGroupId() string {
	if m != nil {
		return m.TodoGroupId
	}
	return ""
}

func (m *TodoGroup) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TodoGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TodoGroup) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *TodoGroup) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *TodoGroup) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *TodoGroup) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

type CreateTodoGroupRequest struct {
	AccountId            string     `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	TodoGroup            *TodoGroup `protobuf:"bytes,2,opt,name=todo_group,json=todoGroup" json:"todo_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateTodoGroupRequest) Reset()         { *m = CreateTodoGroupRequest{} }
func (m *CreateTodoGroupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoGroupRequest) ProtoMessage()    {}
func (*CreateTodoGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{1}
}
func (m *CreateTodoGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoGroupRequest.Unmarshal(m, b)
}
func (m *CreateTodoGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoGroupRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTodoGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoGroupRequest.Merge(dst, src)
}
func (m *CreateTodoGroupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoGroupRequest.Size(m)
}
func (m *CreateTodoGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoGroupRequest proto.InternalMessageInfo

func (m *CreateTodoGroupRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *CreateTodoGroupRequest) GetTodoGroup() *TodoGroup {
	if m != nil {
		return m.TodoGroup
	}
	return nil
}

type CreateTodoGroupResponse struct {
	TodoGroup            *TodoGroup `protobuf:"bytes,1,opt,name=todo_group,json=todoGroup" json:"todo_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateTodoGroupResponse) Reset()         { *m = CreateTodoGroupResponse{} }
func (m *CreateTodoGroupResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoGroupResponse) ProtoMessage()    {}
func (*CreateTodoGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{2}
}
func (m *CreateTodoGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoGroupResponse.Unmarshal(m, b)
}
func (m *CreateTodoGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoGroupResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTodoGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoGroupResponse.Merge(dst, src)
}
func (m *CreateTodoGroupResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoGroupResponse.Size(m)
}
func (m *CreateTodoGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoGroupResponse proto.InternalMessageInfo

func (m *CreateTodoGroupResponse) GetTodoGroup() *TodoGroup {
	if m != nil {
		return m.TodoGroup
	}
	return nil
}

type ListTodoGroupsRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoGroupsRequest) Reset()         { *m = ListTodoGroupsRequest{} }
func (m *ListTodoGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTodoGroupsRequest) ProtoMessage()    {}
func (*ListTodoGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{3}
}
func (m *ListTodoGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoGroupsRequest.Unmarshal(m, b)
}
func (m *ListTodoGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoGroupsRequest.Marshal(b, m, deterministic)
}
func (dst *ListTodoGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoGroupsRequest.Merge(dst, src)
}
func (m *ListTodoGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_ListTodoGroupsRequest.Size(m)
}
func (m *ListTodoGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoGroupsRequest proto.InternalMessageInfo

func (m *ListTodoGroupsRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type ListTodoGroupsResponse struct {
	TodoGroups           []*TodoGroup `protobuf:"bytes,1,rep,name=todo_groups,json=todoGroups" json:"todo_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListTodoGroupsResponse) Reset()         { *m = ListTodoGroupsResponse{} }
func (m *ListTodoGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodoGroupsResponse) ProtoMessage()    {}
func (*ListTodoGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{4}
}
func (m *ListTodoGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoGroupsResponse.Unmarshal(m, b)
}
func (m *ListTodoGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoGroupsResponse.Marshal(b, m, deterministic)
}
func (dst *ListTodoGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoGroupsResponse.Merge(dst, src)
}
func (m *ListTodoGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodoGroupsResponse.Size(m)
}
func (m *ListTodoGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoGroupsResponse proto.InternalMessageInfo

func (m *ListTodoGroupsResponse) GetTodoGroups() []*TodoGroup {
	if m != nil {
		return m.TodoGroups
	}
	return nil
}

type GetTodoGroupRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	TodoGroupId          string   `protobuf:"bytes,2,opt,name=todo_group_id,json=todoGroupId" json:"todo_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoGroupRequest) Reset()         { *m = GetTodoGroupRequest{} }
func (m *GetTodoGroupRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodoGroupRequest) ProtoMessage()    {}
func (*GetTodoGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{5}
}
func (m *GetTodoGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoGroupRequest.Unmarshal(m, b)
}
func (m *GetTodoGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoGroupRequest.Marshal(b, m, deterministic)
}
func (dst *GetTodoGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoGroupRequest.Merge(dst, src)
}
func (m *GetTodoGroupRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodoGroupRequest.Size(m)
}
func (m *GetTodoGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoGroupRequest proto.InternalMessageInfo

func (m *GetTodoGroupRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *GetTodoGroupRequest) GetTodoGroupId() string {
	if m != nil {
		return m.TodoGroupId
	}
	return ""
}

type GetTodoGroupResponse struct {
	TodoGroup            *TodoGroup `protobuf:"bytes,1,opt,name=todo_group,json=todoGroup" json:"todo_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetTodoGroupResponse) Reset()         { *m = GetTodoGroupResponse{} }
func (m *GetTodoGroupResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodoGroupResponse) ProtoMessage()    {}
func (*GetTodoGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{6}
}
func (m *GetTodoGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoGroupResponse.Unmarshal(m, b)
}
func (m *GetTodoGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoGroupResponse.Marshal(b, m, deterministic)
}
func (dst *GetTodoGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoGroupResponse.Merge(dst, src)
}
func (m *GetTodoGroupResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodoGroupResponse.Size(m)
}
func (m *GetTodoGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoGroupResponse proto.InternalMessageInfo

func (m *GetTodoGroupResponse) GetTodoGroup() *TodoGroup {
	if m != nil {
		return m.TodoGroup
	}
	return nil
}

type UpdateTodoGroupRequest struct {
	AccountId            string     `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	TodoGroupId          string     `protobuf:"bytes,2,opt,name=todo_group_id,json=todoGroupId" json:"todo_group_id,omitempty"`
	TodoGroup            *TodoGroup `protobuf:"bytes,3,opt,name=todo_group,json=todoGroup" json:"todo_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateTodoGroupRequest) Reset()         { *m = UpdateTodoGroupRequest{} }
func (m *UpdateTodoGroupRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoGroupRequest) ProtoMessage()    {}
func (*UpdateTodoGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{7}
}
func (m *UpdateTodoGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoGroupRequest.Unmarshal(m, b)
}
func (m *UpdateTodoGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoGroupRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateTodoGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoGroupRequest.Merge(dst, src)
}
func (m *UpdateTodoGroupRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoGroupRequest.Size(m)
}
func (m *UpdateTodoGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoGroupRequest proto.InternalMessageInfo

func (m *UpdateTodoGroupRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *UpdateTodoGroupRequest) GetTodoGroupId() string {
	if m != nil {
		return m.TodoGroupId
	}
	return ""
}

func (m *UpdateTodoGroupRequest) GetTodoGroup() *TodoGroup {
	if m != nil {
		return m.TodoGroup
	}
	return nil
}

type UpdateTodoGroupResponse struct {
	TodoGroup            *TodoGroup `protobuf:"bytes,1,opt,name=todo_group,json=todoGroup" json:"todo_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateTodoGroupResponse) Reset()         { *m = UpdateTodoGroupResponse{} }
func (m *UpdateTodoGroupResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoGroupResponse) ProtoMessage()    {}
func (*UpdateTodoGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{8}
}
func (m *UpdateTodoGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoGroupResponse.Unmarshal(m, b)
}
func (m *UpdateTodoGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoGroupResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateTodoGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoGroupResponse.Merge(dst, src)
}
func (m *UpdateTodoGroupResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoGroupResponse.Size(m)
}
func (m *UpdateTodoGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoGroupResponse proto.InternalMessageInfo

func (m *UpdateTodoGroupResponse) GetTodoGroup() *TodoGroup {
	if m != nil {
		return m.TodoGroup
	}
	return nil
}

type DeleteTodoGroupRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	TodoGroupId          string   `protobuf:"bytes,2,opt,name=todo_group_id,json=todoGroupId" json:"todo_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoGroupRequest) Reset()         { *m = DeleteTodoGroupRequest{} }
func (m *DeleteTodoGroupRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoGroupRequest) ProtoMessage()    {}
func (*DeleteTodoGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{9}
}
func (m *DeleteTodoGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoGroupRequest.Unmarshal(m, b)
}
func (m *DeleteTodoGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoGroupRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteTodoGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoGroupRequest.Merge(dst, src)
}
func (m *DeleteTodoGroupRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoGroupRequest.Size(m)
}
func (m *DeleteTodoGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoGroupRequest proto.InternalMessageInfo

func (m *DeleteTodoGroupRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *DeleteTodoGroupRequest) GetTodoGroupId() string {
	if m != nil {
		return m.TodoGroupId
	}
	return ""
}

type DeleteTodoGroupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoGroupResponse) Reset()         { *m = DeleteTodoGroupResponse{} }
func (m *DeleteTodoGroupResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoGroupResponse) ProtoMessage()    {}
func (*DeleteTodoGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_groups_07f62f25c3f1148b, []int{10}
}
func (m *DeleteTodoGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoGroupResponse.Unmarshal(m, b)
}
func (m *DeleteTodoGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoGroupResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteTodoGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoGroupResponse.Merge(dst, src)
}
func (m *DeleteTodoGroupResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoGroupResponse.Size(m)
}
func (m *DeleteTodoGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoGroupResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TodoGroup)(nil), "pb.family.todos.TodoGroup")
	proto.RegisterType((*CreateTodoGroupRequest)(nil), "pb.family.todos.CreateTodoGroupRequest")
	proto.RegisterType((*CreateTodoGroupResponse)(nil), "pb.family.todos.CreateTodoGroupResponse")
	proto.RegisterType((*ListTodoGroupsRequest)(nil), "pb.family.todos.ListTodoGroupsRequest")
	proto.RegisterType((*ListTodoGroupsResponse)(nil), "pb.family.todos.ListTodoGroupsResponse")
	proto.RegisterType((*GetTodoGroupRequest)(nil), "pb.family.todos.GetTodoGroupRequest")
	proto.RegisterType((*GetTodoGroupResponse)(nil), "pb.family.todos.GetTodoGroupResponse")
	proto.RegisterType((*UpdateTodoGroupRequest)(nil), "pb.family.todos.UpdateTodoGroupRequest")
	proto.RegisterType((*UpdateTodoGroupResponse)(nil), "pb.family.todos.UpdateTodoGroupResponse")
	proto.RegisterType((*DeleteTodoGroupRequest)(nil), "pb.family.todos.DeleteTodoGroupRequest")
	proto.RegisterType((*DeleteTodoGroupResponse)(nil), "pb.family.todos.DeleteTodoGroupResponse")
	proto.RegisterEnum("pb.family.todos.TodoGroupPermitType", TodoGroupPermitType_name, TodoGroupPermitType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TodoGroupsService service

type TodoGroupsServiceClient interface {
	CreateTodoGroup(ctx context.Context, in *CreateTodoGroupRequest, opts ...grpc.CallOption) (*CreateTodoGroupResponse, error)
	ListTodoGroups(ctx context.Context, in *ListTodoGroupsRequest, opts ...grpc.CallOption) (*ListTodoGroupsResponse, error)
	GetTodoGroup(ctx context.Context, in *GetTodoGroupRequest, opts ...grpc.CallOption) (*GetTodoGroupResponse, error)
	UpdateTodoGroup(ctx context.Context, in *UpdateTodoGroupRequest, opts ...grpc.CallOption) (*UpdateTodoGroupResponse, error)
	DeleteTodoGroup(ctx context.Context, in *DeleteTodoGroupRequest, opts ...grpc.CallOption) (*DeleteTodoGroupResponse, error)
}

type todoGroupsServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoGroupsServiceClient(cc *grpc.ClientConn) TodoGroupsServiceClient {
	return &todoGroupsServiceClient{cc}
}

func (c *todoGroupsServiceClient) CreateTodoGroup(ctx context.Context, in *CreateTodoGroupRequest, opts ...grpc.CallOption) (*CreateTodoGroupResponse, error) {
	out := new(CreateTodoGroupResponse)
	err := grpc.Invoke(ctx, "/pb.family.todos.TodoGroupsService/CreateTodoGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGroupsServiceClient) ListTodoGroups(ctx context.Context, in *ListTodoGroupsRequest, opts ...grpc.CallOption) (*ListTodoGroupsResponse, error) {
	out := new(ListTodoGroupsResponse)
	err := grpc.Invoke(ctx, "/pb.family.todos.TodoGroupsService/ListTodoGroups", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGroupsServiceClient) GetTodoGroup(ctx context.Context, in *GetTodoGroupRequest, opts ...grpc.CallOption) (*GetTodoGroupResponse, error) {
	out := new(GetTodoGroupResponse)
	err := grpc.Invoke(ctx, "/pb.family.todos.TodoGroupsService/GetTodoGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGroupsServiceClient) UpdateTodoGroup(ctx context.Context, in *UpdateTodoGroupRequest, opts ...grpc.CallOption) (*UpdateTodoGroupResponse, error) {
	out := new(UpdateTodoGroupResponse)
	err := grpc.Invoke(ctx, "/pb.family.todos.TodoGroupsService/UpdateTodoGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGroupsServiceClient) DeleteTodoGroup(ctx context.Context, in *DeleteTodoGroupRequest, opts ...grpc.CallOption) (*DeleteTodoGroupResponse, error) {
	out := new(DeleteTodoGroupResponse)
	err := grpc.Invoke(ctx, "/pb.family.todos.TodoGroupsService/DeleteTodoGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TodoGroupsService service

type TodoGroupsServiceServer interface {
	CreateTodoGroup(context.Context, *CreateTodoGroupRequest) (*CreateTodoGroupResponse, error)
	ListTodoGroups(context.Context, *ListTodoGroupsRequest) (*ListTodoGroupsResponse, error)
	GetTodoGroup(context.Context, *GetTodoGroupRequest) (*GetTodoGroupResponse, error)
	UpdateTodoGroup(context.Context, *UpdateTodoGroupRequest) (*UpdateTodoGroupResponse, error)
	DeleteTodoGroup(context.Context, *DeleteTodoGroupRequest) (*DeleteTodoGroupResponse, error)
}

func RegisterTodoGroupsServiceServer(s *grpc.Server, srv TodoGroupsServiceServer) {
	s.RegisterService(&_TodoGroupsService_serviceDesc, srv)
}

func _TodoGroupsService_CreateTodoGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGroupsServiceServer).CreateTodoGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.todos.TodoGroupsService/CreateTodoGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGroupsServiceServer).CreateTodoGroup(ctx, req.(*CreateTodoGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGroupsService_ListTodoGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodoGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGroupsServiceServer).ListTodoGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.todos.TodoGroupsService/ListTodoGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGroupsServiceServer).ListTodoGroups(ctx, req.(*ListTodoGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGroupsService_GetTodoGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGroupsServiceServer).GetTodoGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.todos.TodoGroupsService/GetTodoGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGroupsServiceServer).GetTodoGroup(ctx, req.(*GetTodoGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGroupsService_UpdateTodoGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGroupsServiceServer).UpdateTodoGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.todos.TodoGroupsService/UpdateTodoGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGroupsServiceServer).UpdateTodoGroup(ctx, req.(*UpdateTodoGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGroupsService_DeleteTodoGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGroupsServiceServer).DeleteTodoGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.todos.TodoGroupsService/DeleteTodoGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGroupsServiceServer).DeleteTodoGroup(ctx, req.(*DeleteTodoGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoGroupsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.family.todos.TodoGroupsService",
	HandlerType: (*TodoGroupsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodoGroup",
			Handler:    _TodoGroupsService_CreateTodoGroup_Handler,
		},
		{
			MethodName: "ListTodoGroups",
			Handler:    _TodoGroupsService_ListTodoGroups_Handler,
		},
		{
			MethodName: "GetTodoGroup",
			Handler:    _TodoGroupsService_GetTodoGroup_Handler,
		},
		{
			MethodName: "UpdateTodoGroup",
			Handler:    _TodoGroupsService_UpdateTodoGroup_Handler,
		},
		{
			MethodName: "DeleteTodoGroup",
			Handler:    _TodoGroupsService_DeleteTodoGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/family/todo_groups/todo_groups.proto",
}

func init() {
	proto.RegisterFile("pb/family/todo_groups/todo_groups.proto", fileDescriptor_todo_groups_07f62f25c3f1148b)
}

var fileDescriptor_todo_groups_07f62f25c3f1148b = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xd1, 0x6b, 0xd3, 0x5e,
	0x14, 0xfe, 0xdd, 0xf6, 0xd7, 0x49, 0x4e, 0xd5, 0x6e, 0x77, 0xb3, 0x8b, 0x41, 0xb1, 0x86, 0x69,
	0x47, 0x61, 0x09, 0x4e, 0x10, 0x75, 0x4f, 0x9b, 0x2b, 0xa3, 0xa0, 0x4e, 0x6b, 0xe7, 0x44, 0x1f,
	0x46, 0x9a, 0x5c, 0xb3, 0x0b, 0x6d, 0x6e, 0x4c, 0x6e, 0x07, 0x55, 0x04, 0x11, 0x5f, 0x05, 0x41,
	0xf0, 0xaf, 0xf1, 0xbf, 0xf0, 0xcd, 0x67, 0xff, 0x10, 0xc9, 0x4d, 0x9a, 0xa6, 0x49, 0x46, 0x2c,
	0x63, 0x6f, 0xb7, 0xe7, 0x9c, 0x7c, 0xe7, 0x3b, 0xe7, 0x7e, 0x5f, 0x2f, 0x34, 0xdd, 0xbe, 0xfe,
	0xd6, 0x18, 0xd2, 0xc1, 0x58, 0xe7, 0xcc, 0x62, 0x47, 0xb6, 0xc7, 0x46, 0xae, 0x9f, 0x3c, 0x6b,
	0xae, 0xc7, 0x38, 0xc3, 0x35, 0xb7, 0xaf, 0x85, 0x85, 0x5a, 0x90, 0xf4, 0x95, 0x6b, 0x36, 0x63,
	0xf6, 0x80, 0xe8, 0x86, 0x4b, 0x75, 0xc3, 0x71, 0x18, 0x37, 0x38, 0x65, 0x4e, 0x54, 0xae, 0xfe,
	0x46, 0x20, 0xf5, 0x98, 0xc5, 0xf6, 0x02, 0x0c, 0xac, 0xc2, 0xa5, 0x29, 0xe2, 0x11, 0xb5, 0x64,
	0xd4, 0x40, 0xeb, 0x52, 0xb7, 0xca, 0x27, 0x15, 0x1d, 0x0b, 0xaf, 0x40, 0x85, 0x53, 0x3e, 0x20,
	0x72, 0x49, 0xe4, 0xc2, 0x1f, 0xb8, 0x01, 0x55, 0x8b, 0xf8, 0xa6, 0x47, 0xdd, 0x00, 0x5d, 0x2e,
	0x87, 0xdf, 0x25, 0x42, 0xf8, 0x3a, 0x80, 0xe9, 0x11, 0x83, 0x13, 0xeb, 0xa8, 0x3f, 0x96, 0xff,
	0x17, 0x05, 0x52, 0x14, 0xd9, 0x19, 0x27, 0xd3, 0x06, 0x97, 0x2b, 0x0d, 0xb4, 0x5e, 0x8e, 0xd3,
	0xdb, 0x3c, 0x48, 0x8f, 0x5c, 0x6b, 0x92, 0x5e, 0x08, 0xd3, 0x51, 0x64, 0x9b, 0x07, 0xa4, 0x98,
	0x67, 0x11, 0x4f, 0xbe, 0x10, 0x92, 0x12, 0x3f, 0x54, 0x0f, 0xea, 0x8f, 0x04, 0x42, 0x3c, 0x61,
	0x97, 0xbc, 0x1b, 0x11, 0x5f, 0xc0, 0x19, 0xa6, 0xc9, 0x46, 0x0e, 0x9f, 0x4e, 0x29, 0x45, 0x91,
	0x8e, 0x85, 0x1f, 0x00, 0x4c, 0xf7, 0x20, 0x06, 0xad, 0x6e, 0x2a, 0x5a, 0x6a, 0xb3, 0xda, 0x14,
	0x55, 0x8a, 0x17, 0xa4, 0xf6, 0x60, 0x35, 0xd3, 0xd3, 0x77, 0x99, 0xe3, 0x93, 0x14, 0x2a, 0x9a,
	0x07, 0xf5, 0x1e, 0x5c, 0x79, 0x4c, 0x7d, 0x1e, 0xe7, 0xfc, 0x7f, 0x1b, 0x44, 0x3d, 0x80, 0x7a,
	0xfa, 0xbb, 0x88, 0xcc, 0x16, 0x54, 0x13, 0xe2, 0x91, 0x51, 0xa3, 0x5c, 0xc0, 0x06, 0x62, 0x36,
	0xbe, 0xfa, 0x0a, 0x96, 0xf7, 0x08, 0x9f, 0x77, 0xab, 0x19, 0x75, 0x95, 0x32, 0xea, 0x52, 0x9f,
	0xc3, 0xca, 0x2c, 0xf2, 0xd9, 0x77, 0xf7, 0x03, 0x41, 0xfd, 0x40, 0x28, 0xe5, 0x1c, 0x08, 0xa7,
	0x88, 0x95, 0xe7, 0x94, 0x4a, 0x86, 0xd7, 0xd9, 0xc7, 0x7d, 0x03, 0xf5, 0x5d, 0x32, 0x20, 0xe7,
	0x32, 0xad, 0x7a, 0x15, 0x56, 0x33, 0xe0, 0x21, 0xe5, 0xd6, 0x7d, 0x58, 0x8e, 0x83, 0xcf, 0x88,
	0x37, 0xa4, 0xbc, 0x37, 0x76, 0x09, 0x96, 0xa0, 0xb2, 0x7f, 0xf8, 0xb4, 0xdd, 0x5d, 0xfc, 0x0f,
	0x03, 0x2c, 0xb4, 0x77, 0x3b, 0xbd, 0xfd, 0xee, 0x22, 0x0a, 0xce, 0x2f, 0x3b, 0xed, 0xc3, 0x76,
	0x77, 0xb1, 0xb4, 0xf9, 0xb3, 0x02, 0x4b, 0x53, 0x85, 0xbe, 0x20, 0xde, 0x09, 0x35, 0x09, 0xfe,
	0x84, 0xa0, 0x96, 0x72, 0x12, 0x6e, 0x66, 0x56, 0x90, 0xef, 0x6f, 0x65, 0xbd, 0xb8, 0x30, 0xa4,
	0xad, 0x2a, 0x9f, 0x7f, 0xfd, 0xf9, 0x5e, 0x5a, 0x51, 0x6b, 0xfa, 0xc9, 0x1d, 0xf1, 0x77, 0xba,
	0x11, 0x3a, 0xe2, 0x21, 0x6a, 0xe1, 0xf7, 0x70, 0x79, 0xd6, 0x3d, 0xf8, 0x76, 0x06, 0x37, 0xd7,
	0x96, 0x4a, 0xb3, 0xb0, 0x2e, 0x6a, 0xbf, 0x2a, 0xda, 0x2f, 0xe1, 0x74, 0x7b, 0xfc, 0x05, 0xc1,
	0xc5, 0xa4, 0x13, 0xf0, 0x5a, 0x06, 0x32, 0xc7, 0x82, 0xca, 0xad, 0x82, 0xaa, 0xa8, 0x6d, 0x53,
	0xb4, 0xbd, 0x89, 0x6f, 0xa4, 0xda, 0xea, 0x1f, 0x66, 0x24, 0xf0, 0x11, 0x7f, 0x43, 0x50, 0x4b,
	0x89, 0x34, 0xe7, 0x16, 0xf2, 0xed, 0x95, 0x73, 0x0b, 0xa7, 0xe8, 0x5d, 0x6d, 0x09, 0x3e, 0x6b,
	0x4a, 0x11, 0x9f, 0xe0, 0x56, 0xbe, 0x22, 0xa8, 0xa5, 0x44, 0x98, 0x43, 0x29, 0xdf, 0x03, 0x39,
	0x94, 0x4e, 0xd1, 0xf3, 0x64, 0x45, 0xad, 0x22, 0x4a, 0x3b, 0xfb, 0xaf, 0x9f, 0xd8, 0x94, 0x1f,
	0x8f, 0xfa, 0x9a, 0xc9, 0x86, 0x3a, 0x37, 0xc8, 0x31, 0xdb, 0xa0, 0x6c, 0xf2, 0x5a, 0x53, 0x6b,
	0xa0, 0xdb, 0xc4, 0x21, 0x5e, 0xf0, 0x52, 0xe9, 0x36, 0xd3, 0x73, 0x1f, 0xf2, 0xad, 0xc4, 0xb9,
	0xbf, 0x20, 0x9e, 0xe6, 0xbb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xe8, 0x32, 0x20, 0xf4,
	0x07, 0x00, 0x00,
}
