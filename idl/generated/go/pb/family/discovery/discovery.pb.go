// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/family/discovery/discovery.proto

package discovery // import "github.com/taeho-io/family/idl/generated/go/pb/family/discovery"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Service int32

const (
	Service_AUTH       Service = 0
	Service_ACCOUNTS   Service = 1
	Service_TODOGROUPS Service = 2
	Service_TODOS      Service = 3
)

var Service_name = map[int32]string{
	0: "AUTH",
	1: "ACCOUNTS",
	2: "TODOGROUPS",
	3: "TODOS",
}
var Service_value = map[string]int32{
	"AUTH":       0,
	"ACCOUNTS":   1,
	"TODOGROUPS": 2,
	"TODOS":      3,
}

func (x Service) String() string {
	return proto.EnumName(Service_name, int32(x))
}
func (Service) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_discovery_fc8bbb6689165215, []int{0}
}

func init() {
	proto.RegisterEnum("pb.family.discovery.Service", Service_name, Service_value)
}

func init() {
	proto.RegisterFile("pb/family/discovery/discovery.proto", fileDescriptor_discovery_fc8bbb6689165215)
}

var fileDescriptor_discovery_fc8bbb6689165215 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x48, 0xd2, 0x4f,
	0x4b, 0xcc, 0xcd, 0xcc, 0xa9, 0xd4, 0x4f, 0xc9, 0x2c, 0x4e, 0xce, 0x2f, 0x4b, 0x2d, 0x42, 0x62,
	0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x17, 0x24, 0xe9, 0x41, 0x14, 0xe9, 0xc1, 0xa5,
	0xb4, 0x6c, 0xb8, 0xd8, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x38, 0xb8, 0x58, 0x1c,
	0x43, 0x43, 0x3c, 0x04, 0x18, 0x84, 0x78, 0xb8, 0x38, 0x1c, 0x9d, 0x9d, 0xfd, 0x43, 0xfd, 0x42,
	0x82, 0x05, 0x18, 0x85, 0xf8, 0xb8, 0xb8, 0x42, 0xfc, 0x5d, 0xfc, 0xdd, 0x83, 0xfc, 0x43, 0x03,
	0x82, 0x05, 0x98, 0x84, 0x38, 0xb9, 0x58, 0x41, 0xfc, 0x60, 0x01, 0x66, 0x27, 0xef, 0x28, 0xcf,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x92, 0xc4, 0xd4, 0x8c, 0x7c,
	0xdd, 0xcc, 0x7c, 0x98, 0x53, 0x32, 0x53, 0x72, 0xf4, 0xd3, 0x53, 0xf3, 0x52, 0x8b, 0x12, 0x4b,
	0x52, 0x53, 0xf4, 0xd3, 0xf3, 0xf5, 0xb1, 0xb8, 0xd2, 0x1a, 0xce, 0x4a, 0x62, 0x03, 0x3b, 0xd3,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x21, 0x2a, 0xa5, 0x1d, 0xcd, 0x00, 0x00, 0x00,
}
