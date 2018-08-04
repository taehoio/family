// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/family/auth/account.proto

package account // import "github.com/taeho-io/family/idl/generated/go/pb/family/account"

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

type AuthType int32

const (
	AuthType_NONE  AuthType = 0
	AuthType_EMAIL AuthType = 1
)

var AuthType_name = map[int32]string{
	0: "NONE",
	1: "EMAIL",
}
var AuthType_value = map[string]int32{
	"NONE":  0,
	"EMAIL": 1,
}

func (x AuthType) String() string {
	return proto.EnumName(AuthType_name, int32(x))
}
func (AuthType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_account_8f58006907f0bacd, []int{0}
}

type Account struct {
	AuthType             AuthType `protobuf:"varint,1,opt,name=auth_type,json=authType,enum=pb.family.account.AuthType" json:"auth_type,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	HashedPassword       string   `protobuf:"bytes,4,opt,name=hashed_password,json=hashedPassword" json:"hashed_password,omitempty"`
	FullName             string   `protobuf:"bytes,5,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	CreatedAt            int64    `protobuf:"varint,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,7,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_8f58006907f0bacd, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAuthType() AuthType {
	if m != nil {
		return m.AuthType
	}
	return AuthType_NONE
}

func (m *Account) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetHashedPassword() string {
	if m != nil {
		return m.HashedPassword
	}
	return ""
}

func (m *Account) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Account) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Account) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type RegisterRequest struct {
	AuthType             AuthType `protobuf:"varint,1,opt,name=auth_type,json=authType,enum=pb.family.account.AuthType" json:"auth_type,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	FullName             string   `protobuf:"bytes,4,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_8f58006907f0bacd, []int{1}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetAuthType() AuthType {
	if m != nil {
		return m.AuthType
	}
	return AuthType_NONE
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

type RegisterResponse struct {
	AuthType             AuthType `protobuf:"varint,1,opt,name=auth_type,json=authType,enum=pb.family.account.AuthType" json:"auth_type,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_8f58006907f0bacd, []int{2}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(dst, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetAuthType() AuthType {
	if m != nil {
		return m.AuthType
	}
	return AuthType_NONE
}

func (m *RegisterResponse) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type LogInRequest struct {
	AuthType             AuthType `protobuf:"varint,1,opt,name=auth_type,json=authType,enum=pb.family.account.AuthType" json:"auth_type,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInRequest) Reset()         { *m = LogInRequest{} }
func (m *LogInRequest) String() string { return proto.CompactTextString(m) }
func (*LogInRequest) ProtoMessage()    {}
func (*LogInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_8f58006907f0bacd, []int{3}
}
func (m *LogInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInRequest.Unmarshal(m, b)
}
func (m *LogInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInRequest.Marshal(b, m, deterministic)
}
func (dst *LogInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInRequest.Merge(dst, src)
}
func (m *LogInRequest) XXX_Size() int {
	return xxx_messageInfo_LogInRequest.Size(m)
}
func (m *LogInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogInRequest proto.InternalMessageInfo

func (m *LogInRequest) GetAuthType() AuthType {
	if m != nil {
		return m.AuthType
	}
	return AuthType_NONE
}

func (m *LogInRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LogInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LogInResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	AccountId            string   `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	TokenType            string   `protobuf:"bytes,4,opt,name=token_type,json=tokenType" json:"token_type,omitempty"`
	ExpiresIn            int64    `protobuf:"varint,5,opt,name=expires_in,json=expiresIn" json:"expires_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInResponse) Reset()         { *m = LogInResponse{} }
func (m *LogInResponse) String() string { return proto.CompactTextString(m) }
func (*LogInResponse) ProtoMessage()    {}
func (*LogInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_8f58006907f0bacd, []int{4}
}
func (m *LogInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInResponse.Unmarshal(m, b)
}
func (m *LogInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInResponse.Marshal(b, m, deterministic)
}
func (dst *LogInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInResponse.Merge(dst, src)
}
func (m *LogInResponse) XXX_Size() int {
	return xxx_messageInfo_LogInResponse.Size(m)
}
func (m *LogInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogInResponse proto.InternalMessageInfo

func (m *LogInResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *LogInResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *LogInResponse) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *LogInResponse) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *LogInResponse) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func init() {
	proto.RegisterType((*Account)(nil), "pb.family.account.Account")
	proto.RegisterType((*RegisterRequest)(nil), "pb.family.account.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "pb.family.account.RegisterResponse")
	proto.RegisterType((*LogInRequest)(nil), "pb.family.account.LogInRequest")
	proto.RegisterType((*LogInResponse)(nil), "pb.family.account.LogInResponse")
	proto.RegisterEnum("pb.family.account.AuthType", AuthType_name, AuthType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AccountService service

type AccountServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/pb.family.account.AccountService/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	out := new(LogInResponse)
	err := grpc.Invoke(ctx, "/pb.family.account.AccountService/LogIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountService service

type AccountServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	LogIn(context.Context, *LogInRequest) (*LogInResponse, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.account.AccountService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.account.AccountService/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.family.account.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AccountService_Register_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _AccountService_LogIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/family/auth/account.proto",
}

func init() {
	proto.RegisterFile("pb/family/auth/account.proto", fileDescriptor_account_8f58006907f0bacd)
}

var fileDescriptor_account_8f58006907f0bacd = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0x4d, 0x6e, 0xd3, 0x40,
	0x18, 0xc5, 0xf9, 0x69, 0xed, 0x21, 0x4d, 0xd3, 0x51, 0x85, 0x4c, 0x9a, 0x2a, 0xc1, 0x5d, 0x10,
	0x45, 0xc2, 0x16, 0x65, 0x83, 0xca, 0x2a, 0x48, 0x11, 0x8a, 0x54, 0x02, 0x32, 0x5d, 0xb1, 0xb1,
	0x26, 0xf6, 0x17, 0x7b, 0x54, 0xc7, 0x63, 0x3c, 0xe3, 0x42, 0x58, 0x72, 0x05, 0x76, 0x9c, 0x83,
	0x9b, 0x70, 0x05, 0x0e, 0xc0, 0x01, 0x58, 0x20, 0x7b, 0xc6, 0x21, 0x29, 0x85, 0x0d, 0x42, 0xac,
	0x12, 0xbf, 0xf7, 0xe6, 0xfb, 0xe6, 0xbd, 0x99, 0xf9, 0x50, 0x2f, 0x9d, 0x3b, 0x0b, 0xb2, 0xa4,
	0xf1, 0xca, 0x21, 0xb9, 0x88, 0x1c, 0xe2, 0xfb, 0x2c, 0x4f, 0x84, 0x9d, 0x66, 0x4c, 0x30, 0x7c,
	0x90, 0xce, 0x6d, 0xc9, 0xda, 0x8a, 0xe8, 0xf6, 0x42, 0xc6, 0xc2, 0x18, 0x1c, 0x92, 0x52, 0x87,
	0x24, 0x09, 0x13, 0x44, 0x50, 0x96, 0x70, 0xb9, 0xc0, 0xfa, 0xae, 0xa1, 0xdd, 0xb1, 0x54, 0xe2,
	0xc7, 0xc8, 0x28, 0x4a, 0x7a, 0x62, 0x95, 0x82, 0xa9, 0x0d, 0xb4, 0x61, 0xfb, 0xf4, 0xc8, 0xfe,
	0xa5, 0xa0, 0x3d, 0xce, 0x45, 0x74, 0xb1, 0x4a, 0xc1, 0xd5, 0x89, 0xfa, 0x87, 0x8f, 0x11, 0x52,
	0xac, 0x47, 0x03, 0xb3, 0x36, 0xd0, 0x86, 0x86, 0x6b, 0x28, 0x64, 0x1a, 0xe0, 0x43, 0xd4, 0x84,
	0x25, 0xa1, 0xb1, 0x59, 0x2f, 0x19, 0xf9, 0x81, 0xef, 0xa3, 0xfd, 0x88, 0xf0, 0x08, 0x02, 0x2f,
	0x25, 0x9c, 0xbf, 0x65, 0x59, 0x60, 0x36, 0x4a, 0xbe, 0x2d, 0xe1, 0x97, 0x0a, 0xc5, 0x47, 0xc8,
	0x58, 0xe4, 0x71, 0xec, 0x25, 0x64, 0x09, 0x66, 0xb3, 0x94, 0xe8, 0x05, 0x30, 0x23, 0xcb, 0xb2,
	0xb5, 0x9f, 0x01, 0x11, 0x10, 0x78, 0x44, 0x98, 0x3b, 0x03, 0x6d, 0x58, 0x77, 0x0d, 0x85, 0x8c,
	0x45, 0x41, 0xe7, 0x69, 0x50, 0xd1, 0xbb, 0x92, 0x56, 0xc8, 0x58, 0x58, 0x9f, 0x34, 0xb4, 0xef,
	0x42, 0x48, 0xb9, 0x80, 0xcc, 0x85, 0x37, 0x39, 0xf0, 0xbf, 0x89, 0x61, 0xed, 0xb3, 0xb6, 0xe9,
	0xb3, 0x8b, 0xf4, 0xb5, 0x41, 0x19, 0xc0, 0xfa, 0x7b, 0xdb, 0x5a, 0x63, 0xdb, 0x9a, 0x75, 0x89,
	0x3a, 0x3f, 0xf7, 0xc6, 0x53, 0x96, 0x70, 0xf8, 0x67, 0x67, 0x64, 0xbd, 0x47, 0xad, 0x73, 0x16,
	0x4e, 0x93, 0xff, 0x90, 0x82, 0xf5, 0x59, 0x43, 0x7b, 0xaa, 0xb9, 0xb2, 0x79, 0x0f, 0xb5, 0x88,
	0xef, 0x03, 0xe7, 0x9e, 0x60, 0x97, 0x90, 0x94, 0x1b, 0x30, 0xdc, 0xdb, 0x12, 0xbb, 0x28, 0x20,
	0x7c, 0x82, 0xf6, 0x32, 0x58, 0x64, 0xc0, 0x23, 0xa5, 0x91, 0xed, 0x5a, 0x0a, 0x94, 0xa2, 0x6d,
	0xd3, 0xf5, 0xeb, 0x17, 0xf3, 0x18, 0xa1, 0x72, 0xad, 0x74, 0x29, 0xf3, 0x37, 0x4a, 0xa4, 0x8a,
	0x0c, 0xde, 0xa5, 0x34, 0x03, 0xee, 0xd1, 0xa4, 0xbc, 0x79, 0x75, 0xd7, 0x50, 0xc8, 0x34, 0x19,
	0xf5, 0x91, 0x5e, 0xd9, 0xc7, 0x3a, 0x6a, 0xcc, 0x5e, 0xcc, 0x26, 0x9d, 0x5b, 0xd8, 0x40, 0xcd,
	0xc9, 0xf3, 0xf1, 0xf4, 0xbc, 0xa3, 0x9d, 0x7e, 0xd3, 0x50, 0x5b, 0x3d, 0xae, 0x57, 0x90, 0x5d,
	0x51, 0x1f, 0xb0, 0x40, 0x7a, 0x75, 0xa6, 0xd8, 0xba, 0x21, 0xcf, 0x6b, 0x97, 0xb1, 0x7b, 0xf2,
	0x47, 0x8d, 0x4c, 0xcb, 0xea, 0x7f, 0xf8, 0xf2, 0xf5, 0x63, 0xed, 0xae, 0x75, 0xe8, 0x5c, 0x3d,
	0xac, 0x06, 0x82, 0x93, 0x29, 0xd5, 0x99, 0x36, 0xc2, 0x11, 0x6a, 0x96, 0xf9, 0xe2, 0xfe, 0x0d,
	0xe5, 0x36, 0x8f, 0xbd, 0x3b, 0xf8, 0xbd, 0x40, 0x35, 0xeb, 0x95, 0xcd, 0xee, 0x58, 0x07, 0x9b,
	0xcd, 0xe2, 0x42, 0x72, 0xa6, 0x8d, 0x9e, 0x3e, 0x7b, 0x3d, 0x09, 0xa9, 0x88, 0xf2, 0xb9, 0xed,
	0xb3, 0xa5, 0x23, 0x08, 0x44, 0xec, 0x01, 0x65, 0xd5, 0xc4, 0xa2, 0x41, 0xec, 0x84, 0x90, 0x40,
	0x56, 0x3c, 0x3f, 0x27, 0x64, 0xce, 0xc6, 0x30, 0x93, 0x95, 0x9e, 0xa8, 0xdf, 0xf9, 0x4e, 0x39,
	0x9f, 0x1e, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x56, 0x4b, 0x74, 0xf0, 0x04, 0x00, 0x00,
}
