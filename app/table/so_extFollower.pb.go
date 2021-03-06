// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extFollower.proto

package table // import "github.com/coschain/contentos-go/app/table"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import prototype "github.com/coschain/contentos-go/prototype"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoExtFollower struct {
	FollowerInfo         *prototype.FollowerRelation     `protobuf:"bytes,1,opt,name=follower_info,json=followerInfo,proto3" json:"follower_info,omitempty"`
	FollowerCreatedOrder *prototype.FollowerCreatedOrder `protobuf:"bytes,2,opt,name=follower_created_order,json=followerCreatedOrder,proto3" json:"follower_created_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *SoExtFollower) Reset()         { *m = SoExtFollower{} }
func (m *SoExtFollower) String() string { return proto.CompactTextString(m) }
func (*SoExtFollower) ProtoMessage()    {}
func (*SoExtFollower) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extFollower_52edb3ec288f3f73, []int{0}
}
func (m *SoExtFollower) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtFollower.Unmarshal(m, b)
}
func (m *SoExtFollower) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtFollower.Marshal(b, m, deterministic)
}
func (dst *SoExtFollower) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtFollower.Merge(dst, src)
}
func (m *SoExtFollower) XXX_Size() int {
	return xxx_messageInfo_SoExtFollower.Size(m)
}
func (m *SoExtFollower) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtFollower.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtFollower proto.InternalMessageInfo

func (m *SoExtFollower) GetFollowerInfo() *prototype.FollowerRelation {
	if m != nil {
		return m.FollowerInfo
	}
	return nil
}

func (m *SoExtFollower) GetFollowerCreatedOrder() *prototype.FollowerCreatedOrder {
	if m != nil {
		return m.FollowerCreatedOrder
	}
	return nil
}

type SoMemExtFollowerByFollowerInfo struct {
	FollowerInfo         *prototype.FollowerRelation `protobuf:"bytes,1,opt,name=follower_info,json=followerInfo,proto3" json:"follower_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SoMemExtFollowerByFollowerInfo) Reset()         { *m = SoMemExtFollowerByFollowerInfo{} }
func (m *SoMemExtFollowerByFollowerInfo) String() string { return proto.CompactTextString(m) }
func (*SoMemExtFollowerByFollowerInfo) ProtoMessage()    {}
func (*SoMemExtFollowerByFollowerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extFollower_52edb3ec288f3f73, []int{1}
}
func (m *SoMemExtFollowerByFollowerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtFollowerByFollowerInfo.Unmarshal(m, b)
}
func (m *SoMemExtFollowerByFollowerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtFollowerByFollowerInfo.Marshal(b, m, deterministic)
}
func (dst *SoMemExtFollowerByFollowerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtFollowerByFollowerInfo.Merge(dst, src)
}
func (m *SoMemExtFollowerByFollowerInfo) XXX_Size() int {
	return xxx_messageInfo_SoMemExtFollowerByFollowerInfo.Size(m)
}
func (m *SoMemExtFollowerByFollowerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtFollowerByFollowerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtFollowerByFollowerInfo proto.InternalMessageInfo

func (m *SoMemExtFollowerByFollowerInfo) GetFollowerInfo() *prototype.FollowerRelation {
	if m != nil {
		return m.FollowerInfo
	}
	return nil
}

type SoMemExtFollowerByFollowerCreatedOrder struct {
	FollowerCreatedOrder *prototype.FollowerCreatedOrder `protobuf:"bytes,1,opt,name=follower_created_order,json=followerCreatedOrder,proto3" json:"follower_created_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *SoMemExtFollowerByFollowerCreatedOrder) Reset() {
	*m = SoMemExtFollowerByFollowerCreatedOrder{}
}
func (m *SoMemExtFollowerByFollowerCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*SoMemExtFollowerByFollowerCreatedOrder) ProtoMessage()    {}
func (*SoMemExtFollowerByFollowerCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extFollower_52edb3ec288f3f73, []int{2}
}
func (m *SoMemExtFollowerByFollowerCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtFollowerByFollowerCreatedOrder.Unmarshal(m, b)
}
func (m *SoMemExtFollowerByFollowerCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtFollowerByFollowerCreatedOrder.Marshal(b, m, deterministic)
}
func (dst *SoMemExtFollowerByFollowerCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtFollowerByFollowerCreatedOrder.Merge(dst, src)
}
func (m *SoMemExtFollowerByFollowerCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_SoMemExtFollowerByFollowerCreatedOrder.Size(m)
}
func (m *SoMemExtFollowerByFollowerCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtFollowerByFollowerCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtFollowerByFollowerCreatedOrder proto.InternalMessageInfo

func (m *SoMemExtFollowerByFollowerCreatedOrder) GetFollowerCreatedOrder() *prototype.FollowerCreatedOrder {
	if m != nil {
		return m.FollowerCreatedOrder
	}
	return nil
}

type SoListExtFollowerByFollowerCreatedOrder struct {
	FollowerCreatedOrder *prototype.FollowerCreatedOrder `protobuf:"bytes,1,opt,name=follower_created_order,json=followerCreatedOrder,proto3" json:"follower_created_order,omitempty"`
	FollowerInfo         *prototype.FollowerRelation     `protobuf:"bytes,2,opt,name=follower_info,json=followerInfo,proto3" json:"follower_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *SoListExtFollowerByFollowerCreatedOrder) Reset() {
	*m = SoListExtFollowerByFollowerCreatedOrder{}
}
func (m *SoListExtFollowerByFollowerCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*SoListExtFollowerByFollowerCreatedOrder) ProtoMessage()    {}
func (*SoListExtFollowerByFollowerCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extFollower_52edb3ec288f3f73, []int{3}
}
func (m *SoListExtFollowerByFollowerCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtFollowerByFollowerCreatedOrder.Unmarshal(m, b)
}
func (m *SoListExtFollowerByFollowerCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtFollowerByFollowerCreatedOrder.Marshal(b, m, deterministic)
}
func (dst *SoListExtFollowerByFollowerCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtFollowerByFollowerCreatedOrder.Merge(dst, src)
}
func (m *SoListExtFollowerByFollowerCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_SoListExtFollowerByFollowerCreatedOrder.Size(m)
}
func (m *SoListExtFollowerByFollowerCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtFollowerByFollowerCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtFollowerByFollowerCreatedOrder proto.InternalMessageInfo

func (m *SoListExtFollowerByFollowerCreatedOrder) GetFollowerCreatedOrder() *prototype.FollowerCreatedOrder {
	if m != nil {
		return m.FollowerCreatedOrder
	}
	return nil
}

func (m *SoListExtFollowerByFollowerCreatedOrder) GetFollowerInfo() *prototype.FollowerRelation {
	if m != nil {
		return m.FollowerInfo
	}
	return nil
}

type SoUniqueExtFollowerByFollowerInfo struct {
	FollowerInfo         *prototype.FollowerRelation `protobuf:"bytes,1,opt,name=follower_info,json=followerInfo,proto3" json:"follower_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SoUniqueExtFollowerByFollowerInfo) Reset()         { *m = SoUniqueExtFollowerByFollowerInfo{} }
func (m *SoUniqueExtFollowerByFollowerInfo) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtFollowerByFollowerInfo) ProtoMessage()    {}
func (*SoUniqueExtFollowerByFollowerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extFollower_52edb3ec288f3f73, []int{4}
}
func (m *SoUniqueExtFollowerByFollowerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtFollowerByFollowerInfo.Unmarshal(m, b)
}
func (m *SoUniqueExtFollowerByFollowerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtFollowerByFollowerInfo.Marshal(b, m, deterministic)
}
func (dst *SoUniqueExtFollowerByFollowerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtFollowerByFollowerInfo.Merge(dst, src)
}
func (m *SoUniqueExtFollowerByFollowerInfo) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtFollowerByFollowerInfo.Size(m)
}
func (m *SoUniqueExtFollowerByFollowerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtFollowerByFollowerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtFollowerByFollowerInfo proto.InternalMessageInfo

func (m *SoUniqueExtFollowerByFollowerInfo) GetFollowerInfo() *prototype.FollowerRelation {
	if m != nil {
		return m.FollowerInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*SoExtFollower)(nil), "table.so_extFollower")
	proto.RegisterType((*SoMemExtFollowerByFollowerInfo)(nil), "table.so_mem_extFollower_by_follower_info")
	proto.RegisterType((*SoMemExtFollowerByFollowerCreatedOrder)(nil), "table.so_mem_extFollower_by_follower_created_order")
	proto.RegisterType((*SoListExtFollowerByFollowerCreatedOrder)(nil), "table.so_list_extFollower_by_follower_created_order")
	proto.RegisterType((*SoUniqueExtFollowerByFollowerInfo)(nil), "table.so_unique_extFollower_by_follower_info")
}

func init() {
	proto.RegisterFile("app/table/so_extFollower.proto", fileDescriptor_so_extFollower_52edb3ec288f3f73)
}

var fileDescriptor_so_extFollower_52edb3ec288f3f73 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xcd, 0x4a, 0xc4, 0x30,
	0x18, 0x24, 0x0b, 0x7a, 0x88, 0x3f, 0x87, 0x22, 0x52, 0x44, 0x44, 0x2b, 0x88, 0xc8, 0x6e, 0x03,
	0xfa, 0x04, 0x2a, 0x08, 0x9e, 0x84, 0xbd, 0x08, 0x5e, 0x42, 0x9a, 0x7e, 0xdd, 0x06, 0xd3, 0x7c,
	0x35, 0xf9, 0x8a, 0xee, 0x13, 0xf8, 0x32, 0xbe, 0x89, 0x2f, 0x25, 0x76, 0xdd, 0x62, 0x51, 0x14,
	0x51, 0xd9, 0x5b, 0x32, 0xdf, 0x64, 0x32, 0x33, 0x21, 0x7c, 0x47, 0xd5, 0xb5, 0x20, 0x95, 0x59,
	0x10, 0x01, 0x25, 0x3c, 0xd0, 0x05, 0x5a, 0x8b, 0xf7, 0xe0, 0xd3, 0xda, 0x23, 0x61, 0xb4, 0xd4,
	0xce, 0xb6, 0xe2, 0x76, 0x47, 0xd3, 0x1a, 0x44, 0xd5, 0x58, 0x32, 0xd2, 0xe4, 0x33, 0x42, 0xf2,
	0xc4, 0xf8, 0x7a, 0xff, 0x64, 0x74, 0xca, 0xd7, 0x8a, 0xb7, 0xb5, 0x34, 0xae, 0xc0, 0x98, 0xed,
	0xb2, 0xc3, 0x95, 0xe3, 0xed, 0xb4, 0x13, 0x49, 0xbb, 0xb9, 0x07, 0xab, 0xc8, 0xa0, 0x1b, 0xaf,
	0xce, 0xa1, 0x4b, 0x57, 0x60, 0x74, 0xcd, 0x37, 0x3b, 0x8a, 0xf6, 0xa0, 0x08, 0x72, 0x89, 0x3e,
	0x07, 0x1f, 0x0f, 0x5a, 0xad, 0xbd, 0xcf, 0xb4, 0x7a, 0xc4, 0xf1, 0xc6, 0x1c, 0x3f, 0x9f, 0xc1,
	0x57, 0xaf, 0x68, 0x52, 0xf2, 0xfd, 0x80, 0xb2, 0x82, 0xea, 0xbd, 0x63, 0x99, 0x4d, 0x65, 0xcf,
	0xf1, 0x1f, 0x44, 0x48, 0x1e, 0x19, 0x1f, 0x7e, 0x73, 0x55, 0xcf, 0xf0, 0x17, 0x99, 0xd9, 0xef,
	0x32, 0x3f, 0x33, 0x3e, 0x0a, 0x28, 0xad, 0x09, 0xb4, 0x60, 0x2b, 0x1f, 0x7b, 0x1d, 0xfc, 0xb8,
	0xd7, 0x5b, 0x7e, 0x10, 0x50, 0x36, 0xce, 0xdc, 0x35, 0xf0, 0xdf, 0x8f, 0x78, 0x36, 0xbc, 0x39,
	0x9a, 0x18, 0x2a, 0x9b, 0x2c, 0xd5, 0x58, 0x09, 0x8d, 0x41, 0x97, 0xca, 0x38, 0xa1, 0xd1, 0x11,
	0x38, 0xc2, 0x30, 0x9a, 0xa0, 0xe8, 0x7e, 0x50, 0xb6, 0xdc, 0x0a, 0x9f, 0xbc, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x84, 0x6c, 0x82, 0xe7, 0x55, 0x03, 0x00, 0x00,
}
