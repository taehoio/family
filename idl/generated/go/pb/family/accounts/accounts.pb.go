// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/family/accounts/accounts.proto

package accounts // import "github.com/taeho-io/family/idl/generated/go/pb/family/accounts"

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
	return fileDescriptor_accounts_3ef18cf231cdb3d2, []int{0}
}

type Account struct {
	AuthType             AuthType `protobuf:"varint,1,opt,name=auth_type,json=authType,proto3,enum=pb.family.accounts.AuthType" json:"auth_type,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	HashedPassword       string   `protobuf:"bytes,4,opt,name=hashed_password,json=hashedPassword,proto3" json:"hashed_password,omitempty"`
	FullName             string   `protobuf:"bytes,5,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	CreatedAt            int64    `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_accounts_3ef18cf231cdb3d2, []int{0}
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
	AuthType             AuthType `protobuf:"varint,1,opt,name=auth_type,json=authType,proto3,enum=pb.family.accounts.AuthType" json:"auth_type,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	FullName             string   `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_accounts_3ef18cf231cdb3d2, []int{1}
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
	AuthType             AuthType `protobuf:"varint,1,opt,name=auth_type,json=authType,proto3,enum=pb.family.accounts.AuthType" json:"auth_type,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_accounts_3ef18cf231cdb3d2, []int{2}
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
	AuthType             AuthType `protobuf:"varint,1,opt,name=auth_type,json=authType,proto3,enum=pb.family.accounts.AuthType" json:"auth_type,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInRequest) Reset()         { *m = LogInRequest{} }
func (m *LogInRequest) String() string { return proto.CompactTextString(m) }
func (*LogInRequest) ProtoMessage()    {}
func (*LogInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_accounts_3ef18cf231cdb3d2, []int{3}
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
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	AccountId            string   `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	TokenType            string   `protobuf:"bytes,4,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	ExpiresIn            int64    `protobuf:"varint,5,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInResponse) Reset()         { *m = LogInResponse{} }
func (m *LogInResponse) String() string { return proto.CompactTextString(m) }
func (*LogInResponse) ProtoMessage()    {}
func (*LogInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_accounts_3ef18cf231cdb3d2, []int{4}
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
	proto.RegisterType((*Account)(nil), "pb.family.accounts.Account")
	proto.RegisterType((*RegisterRequest)(nil), "pb.family.accounts.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "pb.family.accounts.RegisterResponse")
	proto.RegisterType((*LogInRequest)(nil), "pb.family.accounts.LogInRequest")
	proto.RegisterType((*LogInResponse)(nil), "pb.family.accounts.LogInResponse")
	proto.RegisterEnum("pb.family.accounts.AuthType", AuthType_name, AuthType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountsServiceClient is the client API for AccountsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
}

type accountsServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountsServiceClient(cc *grpc.ClientConn) AccountsServiceClient {
	return &accountsServiceClient{cc}
}

func (c *accountsServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/pb.family.accounts.AccountsService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	out := new(LogInResponse)
	err := c.cc.Invoke(ctx, "/pb.family.accounts.AccountsService/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServiceServer is the server API for AccountsService service.
type AccountsServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	LogIn(context.Context, *LogInRequest) (*LogInResponse, error)
}

func RegisterAccountsServiceServer(s *grpc.Server, srv AccountsServiceServer) {
	s.RegisterService(&_AccountsService_serviceDesc, srv)
}

func _AccountsService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.accounts.AccountsService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.family.accounts.AccountsService/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.family.accounts.AccountsService",
	HandlerType: (*AccountsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AccountsService_Register_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _AccountsService_LogIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/family/accounts/accounts.proto",
}

func init() {
	proto.RegisterFile("pb/family/accounts/accounts.proto", fileDescriptor_accounts_3ef18cf231cdb3d2)
}

var fileDescriptor_accounts_3ef18cf231cdb3d2 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xc5, 0x4d, 0xd3, 0xda, 0x97, 0x3e, 0xa2, 0x11, 0x08, 0x2b, 0x34, 0x22, 0x71, 0x91, 0x88,
	0x22, 0x61, 0x8b, 0xb2, 0xa2, 0xac, 0x82, 0x54, 0xa1, 0x48, 0x25, 0x20, 0xd3, 0x15, 0x1b, 0x6b,
	0x62, 0xdf, 0xd8, 0xa3, 0x3a, 0x1e, 0xe3, 0x19, 0xb7, 0x44, 0xec, 0xf8, 0x05, 0x96, 0xfc, 0x06,
	0x7f, 0xc2, 0x2f, 0xf0, 0x0f, 0xac, 0x90, 0x50, 0x3c, 0xe3, 0xb4, 0x69, 0xab, 0x6e, 0x10, 0x62,
	0x37, 0x3e, 0xe7, 0x78, 0xee, 0x9c, 0x73, 0x67, 0x2e, 0xf4, 0xf2, 0x89, 0x37, 0xa5, 0x33, 0x96,
	0xce, 0x3d, 0x1a, 0x86, 0xbc, 0xcc, 0xa4, 0x58, 0x2e, 0xdc, 0xbc, 0xe0, 0x92, 0x13, 0x92, 0x4f,
	0x5c, 0x25, 0x71, 0x6b, 0xa6, 0xbd, 0x17, 0x73, 0x1e, 0xa7, 0xe8, 0xd1, 0x9c, 0x79, 0x34, 0xcb,
	0xb8, 0xa4, 0x92, 0xf1, 0x4c, 0xff, 0xe1, 0xfc, 0x36, 0x60, 0x73, 0xa8, 0xa4, 0xe4, 0x05, 0x58,
	0xb4, 0x94, 0x49, 0x20, 0xe7, 0x39, 0xda, 0x46, 0xd7, 0xe8, 0xef, 0x1c, 0xec, 0xb9, 0xd7, 0x77,
	0x74, 0x87, 0xa5, 0x4c, 0x4e, 0xe6, 0x39, 0xfa, 0x26, 0xd5, 0x2b, 0xd2, 0x01, 0xd0, 0x74, 0xc0,
	0x22, 0x7b, 0xad, 0x6b, 0xf4, 0x2d, 0xdf, 0xd2, 0xc8, 0x28, 0x22, 0xf7, 0xa0, 0x89, 0x33, 0xca,
	0x52, 0xbb, 0x51, 0x31, 0xea, 0x83, 0x3c, 0x81, 0xdd, 0x84, 0x8a, 0x04, 0xa3, 0x20, 0xa7, 0x42,
	0x9c, 0xf3, 0x22, 0xb2, 0xd7, 0x2b, 0x7e, 0x47, 0xc1, 0xef, 0x34, 0x4a, 0x1e, 0x82, 0x35, 0x2d,
	0xd3, 0x34, 0xc8, 0xe8, 0x0c, 0xed, 0x66, 0x25, 0x31, 0x17, 0xc0, 0x98, 0xce, 0xaa, 0xd2, 0x61,
	0x81, 0x54, 0x62, 0x14, 0x50, 0x69, 0x6f, 0x74, 0x8d, 0x7e, 0xc3, 0xb7, 0x34, 0x32, 0x94, 0x0b,
	0xba, 0xcc, 0xa3, 0x9a, 0xde, 0x54, 0xb4, 0x46, 0x86, 0xd2, 0xf9, 0x66, 0xc0, 0xae, 0x8f, 0x31,
	0x13, 0x12, 0x0b, 0x1f, 0x3f, 0x96, 0x28, 0xfe, 0x2a, 0x87, 0xa5, 0xd1, 0xb5, 0xcb, 0x46, 0xdb,
	0x60, 0x2e, 0x1d, 0xaa, 0x04, 0x96, 0xdf, 0xab, 0xde, 0xd6, 0x57, 0xbd, 0x39, 0x29, 0xb4, 0x2e,
	0x0e, 0x27, 0x72, 0x9e, 0x09, 0xfc, 0x77, 0x5d, 0x72, 0x3e, 0xc3, 0xd6, 0x31, 0x8f, 0x47, 0xd9,
	0xff, 0xc8, 0xc1, 0xf9, 0x6e, 0xc0, 0xb6, 0xae, 0xae, 0x8d, 0xf6, 0x60, 0x8b, 0x86, 0x21, 0x0a,
	0x11, 0x48, 0x7e, 0x8a, 0x59, 0x75, 0x02, 0xcb, 0xbf, 0xab, 0xb0, 0x93, 0x05, 0x44, 0xf6, 0x61,
	0xbb, 0xc0, 0x69, 0x81, 0x22, 0xd1, 0x1a, 0x55, 0x6e, 0x4b, 0x83, 0x4a, 0xb4, 0xea, 0xba, 0x71,
	0xf5, 0x6e, 0x76, 0x00, 0xaa, 0x7f, 0x95, 0x4d, 0xd5, 0x01, 0xab, 0x42, 0xea, 0xcc, 0xf0, 0x53,
	0xce, 0x0a, 0x14, 0x01, 0xcb, 0xaa, 0xcb, 0xd7, 0xf0, 0x2d, 0x8d, 0x8c, 0xb2, 0xc1, 0x23, 0x30,
	0x6b, 0xfb, 0xc4, 0x84, 0xf5, 0xf1, 0xdb, 0xf1, 0x51, 0xeb, 0x0e, 0xb1, 0xa0, 0x79, 0xf4, 0x66,
	0x38, 0x3a, 0x6e, 0x19, 0x07, 0xbf, 0x0c, 0xd8, 0xd5, 0x0f, 0x4c, 0xbc, 0xc7, 0xe2, 0x8c, 0x85,
	0x48, 0xce, 0xc1, 0xac, 0xdb, 0x4a, 0xf6, 0x6f, 0x4a, 0xf4, 0xca, 0x8d, 0x6c, 0x3f, 0xbe, 0x5d,
	0xa4, 0x02, 0x73, 0xba, 0x5f, 0x7e, 0xfc, 0xfc, 0xba, 0xd6, 0x76, 0xee, 0x7b, 0x67, 0xcf, 0x2e,
	0x46, 0x44, 0xa1, 0x65, 0x87, 0xc6, 0x80, 0x9c, 0x42, 0xb3, 0xca, 0x98, 0x74, 0x6f, 0xda, 0xf0,
	0x72, 0xf3, 0xdb, 0xbd, 0x5b, 0x14, 0xba, 0x5e, 0xa7, 0xaa, 0xf7, 0xc0, 0x21, 0x2b, 0xf5, 0xd2,
	0x85, 0xe6, 0xd0, 0x18, 0xbc, 0x1a, 0x7d, 0x78, 0x1d, 0x33, 0x99, 0x94, 0x13, 0x37, 0xe4, 0x33,
	0x4f, 0x52, 0x4c, 0xf8, 0x53, 0xc6, 0xeb, 0x11, 0xc6, 0xa2, 0xd4, 0x8b, 0x31, 0xc3, 0x62, 0xf1,
	0x10, 0xbd, 0x98, 0x7b, 0xd7, 0xa7, 0xdb, 0xcb, 0x7a, 0x31, 0xd9, 0xa8, 0x86, 0xd5, 0xf3, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x4f, 0x96, 0x0a, 0x03, 0x05, 0x00, 0x00,
}
