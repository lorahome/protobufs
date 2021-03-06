// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/join.proto

package openiot

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type JoinRequest struct {
	// Unique Device ID
	DeviceId uint64 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Device details
	DeviceName         string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	DeviceManufacturer string `protobuf:"bytes,3,opt,name=device_manufacturer,json=deviceManufacturer,proto3" json:"device_manufacturer,omitempty"`
	DeviceUrl          string `protobuf:"bytes,4,opt,name=device_url,json=deviceUrl,proto3" json:"device_url,omitempty"`
	// protobuf urls
	ProtoStatusMessageUrl string   `protobuf:"bytes,10,opt,name=proto_status_message_url,json=protoStatusMessageUrl,proto3" json:"proto_status_message_url,omitempty"`
	ProtoConfigMessageUrl string   `protobuf:"bytes,11,opt,name=proto_config_message_url,json=protoConfigMessageUrl,proto3" json:"proto_config_message_url,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85f778e2ac17fa52, []int{0}
}

func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetDeviceId() uint64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *JoinRequest) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *JoinRequest) GetDeviceManufacturer() string {
	if m != nil {
		return m.DeviceManufacturer
	}
	return ""
}

func (m *JoinRequest) GetDeviceUrl() string {
	if m != nil {
		return m.DeviceUrl
	}
	return ""
}

func (m *JoinRequest) GetProtoStatusMessageUrl() string {
	if m != nil {
		return m.ProtoStatusMessageUrl
	}
	return ""
}

func (m *JoinRequest) GetProtoConfigMessageUrl() string {
	if m != nil {
		return m.ProtoConfigMessageUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*JoinRequest)(nil), "openiot.JoinRequest")
}

func init() { proto.RegisterFile("proto/join.proto", fileDescriptor_85f778e2ac17fa52) }

var fileDescriptor_85f778e2ac17fa52 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x46, 0xe9, 0xf5, 0xa2, 0x76, 0xba, 0x91, 0x88, 0x10, 0x10, 0xb1, 0xb8, 0xea, 0xca, 0x2e,
	0x5c, 0xf8, 0x00, 0xae, 0x14, 0xea, 0xa2, 0xe2, 0x3a, 0xc4, 0x76, 0x5a, 0x22, 0x4d, 0x52, 0xf3,
	0xe3, 0x5b, 0xf8, 0xce, 0xc2, 0x24, 0x70, 0xdb, 0xdd, 0xf0, 0x9d, 0x73, 0x60, 0xe0, 0x6a, 0x75,
	0x36, 0xd8, 0xf6, 0xdb, 0x2a, 0xf3, 0x48, 0x27, 0xbb, 0xb0, 0x2b, 0x1a, 0x65, 0xc3, 0xc3, 0xdf,
	0x01, 0xaa, 0x37, 0xab, 0x4c, 0x8f, 0x3f, 0x11, 0x7d, 0x60, 0xb7, 0x50, 0x8e, 0xf8, 0xab, 0x06,
	0x14, 0x6a, 0xe4, 0x45, 0x5d, 0x34, 0xc7, 0xfe, 0x32, 0x0d, 0xaf, 0x23, 0xbb, 0x87, 0x2a, 0x43,
	0x23, 0x35, 0xf2, 0x43, 0x5d, 0x34, 0x65, 0x0f, 0x69, 0x7a, 0x97, 0x1a, 0x59, 0x0b, 0xd7, 0x59,
	0xd0, 0xd2, 0xc4, 0x49, 0x0e, 0x21, 0x3a, 0x74, 0xfc, 0x8c, 0x44, 0x96, 0x50, 0xb7, 0x21, 0xec,
	0x0e, 0x72, 0x2e, 0xa2, 0x5b, 0xf8, 0x91, 0xbc, 0xfc, 0xc0, 0xa7, 0x5b, 0xd8, 0x33, 0x70, 0xfa,
	0x57, 0xf8, 0x20, 0x43, 0xf4, 0x42, 0xa3, 0xf7, 0x72, 0x4e, 0x32, 0x90, 0x7c, 0x43, 0xfc, 0x83,
	0x70, 0x97, 0xe8, 0x2e, 0x1c, 0xac, 0x99, 0xd4, 0xbc, 0x0b, 0xab, 0x4d, 0xf8, 0x42, 0xf8, 0x14,
	0x7e, 0x9d, 0xd3, 0xfc, 0xf4, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x14, 0x2c, 0x90, 0xa4, 0x33, 0x01,
	0x00, 0x00,
}
