// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/proto/service.proto

package proto

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

// service_map defines the mapping between a service ID and its tuple network_address
// e.g., map[autopeering:&{tcp, 198.51.100.1:80}]
type ServiceMap struct {
	Map                  map[string]*NetworkAddress `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ServiceMap) Reset()         { *m = ServiceMap{} }
func (m *ServiceMap) String() string { return proto.CompactTextString(m) }
func (*ServiceMap) ProtoMessage()    {}
func (*ServiceMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_7271203e88dabbb0, []int{0}
}

func (m *ServiceMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceMap.Unmarshal(m, b)
}
func (m *ServiceMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceMap.Marshal(b, m, deterministic)
}
func (m *ServiceMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceMap.Merge(m, src)
}
func (m *ServiceMap) XXX_Size() int {
	return xxx_messageInfo_ServiceMap.Size(m)
}
func (m *ServiceMap) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceMap.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceMap proto.InternalMessageInfo

func (m *ServiceMap) GetMap() map[string]*NetworkAddress {
	if m != nil {
		return m.Map
	}
	return nil
}

// network_address defines the service type (e.g., tcp, upd) and the address (e.g., 198.51.100.1:80)
type NetworkAddress struct {
	Network              string   `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkAddress) Reset()         { *m = NetworkAddress{} }
func (m *NetworkAddress) String() string { return proto.CompactTextString(m) }
func (*NetworkAddress) ProtoMessage()    {}
func (*NetworkAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_7271203e88dabbb0, []int{1}
}

func (m *NetworkAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkAddress.Unmarshal(m, b)
}
func (m *NetworkAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkAddress.Marshal(b, m, deterministic)
}
func (m *NetworkAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkAddress.Merge(m, src)
}
func (m *NetworkAddress) XXX_Size() int {
	return xxx_messageInfo_NetworkAddress.Size(m)
}
func (m *NetworkAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkAddress proto.InternalMessageInfo

func (m *NetworkAddress) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *NetworkAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceMap)(nil), "proto.ServiceMap")
	proto.RegisterMapType((map[string]*NetworkAddress)(nil), "proto.ServiceMap.MapEntry")
	proto.RegisterType((*NetworkAddress)(nil), "proto.NetworkAddress")
}

func init() { proto.RegisterFile("peer/proto/service.proto", fileDescriptor_7271203e88dabbb0) }

var fileDescriptor_7271203e88dabbb0 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x48, 0x4d, 0x2d,
	0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03,
	0xf3, 0x84, 0x58, 0xc1, 0x94, 0x52, 0x27, 0x23, 0x17, 0x57, 0x30, 0x44, 0xc2, 0x37, 0xb1, 0x40,
	0x48, 0x87, 0x8b, 0x39, 0x37, 0xb1, 0x40, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x0a, 0xa2,
	0x54, 0x0f, 0x21, 0xaf, 0x07, 0xc4, 0xae, 0x79, 0x25, 0x45, 0x95, 0x41, 0x20, 0x65, 0x52, 0xbe,
	0x5c, 0x1c, 0x30, 0x01, 0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0xa0, 0x4e, 0x46, 0x0d, 0xce,
	0x20, 0x10, 0x53, 0x48, 0x9b, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x09, 0x28, 0xc6,
	0x6d, 0x24, 0x0a, 0x35, 0xcd, 0x2f, 0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0xdb, 0x31, 0x25, 0xa5, 0x28,
	0xb5, 0xb8, 0x38, 0x08, 0xa2, 0xc6, 0x8a, 0xc9, 0x82, 0x51, 0xc9, 0x85, 0x8b, 0x0f, 0x55, 0x52,
	0x48, 0x82, 0x8b, 0x3d, 0x0f, 0x22, 0x02, 0x35, 0x18, 0xc6, 0x05, 0xc9, 0x24, 0x42, 0x14, 0x81,
	0x8d, 0x07, 0xca, 0x40, 0xb9, 0x4e, 0x46, 0x51, 0x06, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a,
	0xc9, 0xf9, 0xb9, 0xfa, 0x99, 0xf9, 0x25, 0x89, 0x39, 0xa9, 0x29, 0xe9, 0xc0, 0x50, 0x48, 0x2c,
	0x2d, 0xc9, 0x07, 0x05, 0x47, 0x66, 0x5e, 0xba, 0x6e, 0x71, 0x66, 0xae, 0x3e, 0x22, 0x68, 0x92,
	0xd8, 0xc0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xf7, 0x68, 0x31, 0x2f, 0x01, 0x00,
	0x00,
}
