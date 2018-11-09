// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_demo.proto

package table

import (
	_ "common/prototype"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoDemo struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	PostTime             uint32   `protobuf:"varint,2,opt,name=post_time,json=postTime,proto3" json:"post_time,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	LikeCount            int64    `protobuf:"varint,4,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	Idx                  int64    `protobuf:"varint,5,opt,name=idx,proto3" json:"idx,omitempty"`
	ReplayCount          int64    `protobuf:"varint,6,opt,name=replay_count,json=replayCount,proto3" json:"replay_count,omitempty"`
	Content              string   `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	Taglist              string   `protobuf:"bytes,8,opt,name=taglist,proto3" json:"taglist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoDemo) Reset()         { *m = SoDemo{} }
func (m *SoDemo) String() string { return proto.CompactTextString(m) }
func (*SoDemo) ProtoMessage()    {}
func (*SoDemo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9e5111d02436852, []int{0}
}

func (m *SoDemo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoDemo.Unmarshal(m, b)
}
func (m *SoDemo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoDemo.Marshal(b, m, deterministic)
}
func (m *SoDemo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoDemo.Merge(m, src)
}
func (m *SoDemo) XXX_Size() int {
	return xxx_messageInfo_SoDemo.Size(m)
}
func (m *SoDemo) XXX_DiscardUnknown() {
	xxx_messageInfo_SoDemo.DiscardUnknown(m)
}

var xxx_messageInfo_SoDemo proto.InternalMessageInfo

func (m *SoDemo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *SoDemo) GetPostTime() uint32 {
	if m != nil {
		return m.PostTime
	}
	return 0
}

func (m *SoDemo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SoDemo) GetLikeCount() int64 {
	if m != nil {
		return m.LikeCount
	}
	return 0
}

func (m *SoDemo) GetIdx() int64 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *SoDemo) GetReplayCount() int64 {
	if m != nil {
		return m.ReplayCount
	}
	return 0
}

func (m *SoDemo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SoDemo) GetTaglist() string {
	if m != nil {
		return m.Taglist
	}
	return ""
}

type SoListDemoByPostTime struct {
	PostTime             uint32   `protobuf:"varint,1,opt,name=post_time,json=postTime,proto3" json:"post_time,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoListDemoByPostTime) Reset()         { *m = SoListDemoByPostTime{} }
func (m *SoListDemoByPostTime) String() string { return proto.CompactTextString(m) }
func (*SoListDemoByPostTime) ProtoMessage()    {}
func (*SoListDemoByPostTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9e5111d02436852, []int{1}
}

func (m *SoListDemoByPostTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListDemoByPostTime.Unmarshal(m, b)
}
func (m *SoListDemoByPostTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListDemoByPostTime.Marshal(b, m, deterministic)
}
func (m *SoListDemoByPostTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListDemoByPostTime.Merge(m, src)
}
func (m *SoListDemoByPostTime) XXX_Size() int {
	return xxx_messageInfo_SoListDemoByPostTime.Size(m)
}
func (m *SoListDemoByPostTime) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListDemoByPostTime.DiscardUnknown(m)
}

var xxx_messageInfo_SoListDemoByPostTime proto.InternalMessageInfo

func (m *SoListDemoByPostTime) GetPostTime() uint32 {
	if m != nil {
		return m.PostTime
	}
	return 0
}

func (m *SoListDemoByPostTime) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type SoListDemoByReplayCount struct {
	ReplayCount          int64    `protobuf:"varint,1,opt,name=replay_count,json=replayCount,proto3" json:"replay_count,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoListDemoByReplayCount) Reset()         { *m = SoListDemoByReplayCount{} }
func (m *SoListDemoByReplayCount) String() string { return proto.CompactTextString(m) }
func (*SoListDemoByReplayCount) ProtoMessage()    {}
func (*SoListDemoByReplayCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9e5111d02436852, []int{2}
}

func (m *SoListDemoByReplayCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListDemoByReplayCount.Unmarshal(m, b)
}
func (m *SoListDemoByReplayCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListDemoByReplayCount.Marshal(b, m, deterministic)
}
func (m *SoListDemoByReplayCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListDemoByReplayCount.Merge(m, src)
}
func (m *SoListDemoByReplayCount) XXX_Size() int {
	return xxx_messageInfo_SoListDemoByReplayCount.Size(m)
}
func (m *SoListDemoByReplayCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListDemoByReplayCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoListDemoByReplayCount proto.InternalMessageInfo

func (m *SoListDemoByReplayCount) GetReplayCount() int64 {
	if m != nil {
		return m.ReplayCount
	}
	return 0
}

func (m *SoListDemoByReplayCount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type SoUniqueDemoByOwner struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniqueDemoByOwner) Reset()         { *m = SoUniqueDemoByOwner{} }
func (m *SoUniqueDemoByOwner) String() string { return proto.CompactTextString(m) }
func (*SoUniqueDemoByOwner) ProtoMessage()    {}
func (*SoUniqueDemoByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9e5111d02436852, []int{3}
}

func (m *SoUniqueDemoByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueDemoByOwner.Unmarshal(m, b)
}
func (m *SoUniqueDemoByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueDemoByOwner.Marshal(b, m, deterministic)
}
func (m *SoUniqueDemoByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueDemoByOwner.Merge(m, src)
}
func (m *SoUniqueDemoByOwner) XXX_Size() int {
	return xxx_messageInfo_SoUniqueDemoByOwner.Size(m)
}
func (m *SoUniqueDemoByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueDemoByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueDemoByOwner proto.InternalMessageInfo

func (m *SoUniqueDemoByOwner) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type SoUniqueDemoByLikeCount struct {
	LikeCount            int64    `protobuf:"varint,1,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniqueDemoByLikeCount) Reset()         { *m = SoUniqueDemoByLikeCount{} }
func (m *SoUniqueDemoByLikeCount) String() string { return proto.CompactTextString(m) }
func (*SoUniqueDemoByLikeCount) ProtoMessage()    {}
func (*SoUniqueDemoByLikeCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9e5111d02436852, []int{4}
}

func (m *SoUniqueDemoByLikeCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueDemoByLikeCount.Unmarshal(m, b)
}
func (m *SoUniqueDemoByLikeCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueDemoByLikeCount.Marshal(b, m, deterministic)
}
func (m *SoUniqueDemoByLikeCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueDemoByLikeCount.Merge(m, src)
}
func (m *SoUniqueDemoByLikeCount) XXX_Size() int {
	return xxx_messageInfo_SoUniqueDemoByLikeCount.Size(m)
}
func (m *SoUniqueDemoByLikeCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueDemoByLikeCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueDemoByLikeCount proto.InternalMessageInfo

func (m *SoUniqueDemoByLikeCount) GetLikeCount() int64 {
	if m != nil {
		return m.LikeCount
	}
	return 0
}

func (m *SoUniqueDemoByLikeCount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type SoUniqueDemoByIdx struct {
	Idx                  int64    `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniqueDemoByIdx) Reset()         { *m = SoUniqueDemoByIdx{} }
func (m *SoUniqueDemoByIdx) String() string { return proto.CompactTextString(m) }
func (*SoUniqueDemoByIdx) ProtoMessage()    {}
func (*SoUniqueDemoByIdx) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9e5111d02436852, []int{5}
}

func (m *SoUniqueDemoByIdx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueDemoByIdx.Unmarshal(m, b)
}
func (m *SoUniqueDemoByIdx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueDemoByIdx.Marshal(b, m, deterministic)
}
func (m *SoUniqueDemoByIdx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueDemoByIdx.Merge(m, src)
}
func (m *SoUniqueDemoByIdx) XXX_Size() int {
	return xxx_messageInfo_SoUniqueDemoByIdx.Size(m)
}
func (m *SoUniqueDemoByIdx) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueDemoByIdx.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueDemoByIdx proto.InternalMessageInfo

func (m *SoUniqueDemoByIdx) GetIdx() int64 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *SoUniqueDemoByIdx) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*SoDemo)(nil), "table.so_demo")
	proto.RegisterType((*SoListDemoByPostTime)(nil), "table.so_list_demo_by_post_time")
	proto.RegisterType((*SoListDemoByReplayCount)(nil), "table.so_list_demo_by_replay_count")
	proto.RegisterType((*SoUniqueDemoByOwner)(nil), "table.so_unique_demo_by_owner")
	proto.RegisterType((*SoUniqueDemoByLikeCount)(nil), "table.so_unique_demo_by_like_count")
	proto.RegisterType((*SoUniqueDemoByIdx)(nil), "table.so_unique_demo_by_idx")
}

func init() { proto.RegisterFile("app/table/so_demo.proto", fileDescriptor_c9e5111d02436852) }

var fileDescriptor_c9e5111d02436852 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0x94, 0xdb, 0xaf, 0x7f, 0xfb, 0x81, 0x84, 0x2c, 0x50, 0x0d, 0x05, 0xa9, 0xe4, 0xd4, 0x13,
	0x39, 0xf0, 0x00, 0x1c, 0xb8, 0x73, 0x28, 0x48, 0x1c, 0xad, 0x24, 0xb5, 0x90, 0x45, 0xe2, 0x35,
	0xc9, 0x46, 0x90, 0x57, 0xe5, 0x69, 0x90, 0xed, 0xd0, 0x96, 0x24, 0x97, 0x68, 0x77, 0x36, 0x33,
	0xd9, 0xd9, 0x09, 0x2c, 0x13, 0x6b, 0x63, 0x4a, 0xd2, 0x5c, 0xc5, 0x15, 0xca, 0x9d, 0x2a, 0xf0,
	0xce, 0x96, 0x48, 0xc8, 0x27, 0x1e, 0xbc, 0x5a, 0x65, 0x58, 0x14, 0x68, 0x62, 0x0f, 0x52, 0x63,
	0x55, 0xec, 0x1e, 0xe1, 0x9d, 0xe8, 0x9b, 0xc1, 0xac, 0x65, 0xf1, 0x73, 0x98, 0xe0, 0xa7, 0x51,
	0xa5, 0x60, 0x6b, 0xb6, 0x59, 0x6c, 0x43, 0xc3, 0x57, 0xb0, 0xb0, 0x58, 0x91, 0x24, 0x5d, 0x28,
	0x31, 0x5a, 0xb3, 0xcd, 0xe9, 0x76, 0xee, 0x80, 0x17, 0x5d, 0x28, 0x47, 0x21, 0x4d, 0xb9, 0x12,
	0xe3, 0x40, 0xf1, 0x0d, 0xbf, 0x01, 0xc8, 0xf5, 0xbb, 0x92, 0x19, 0xd6, 0x86, 0xc4, 0xbf, 0x35,
	0xdb, 0x8c, 0xb7, 0x0b, 0x87, 0x3c, 0x3a, 0x80, 0x9f, 0xc1, 0x58, 0xef, 0xbe, 0xc4, 0xc4, 0xe3,
	0xae, 0xe4, 0xb7, 0x70, 0x52, 0x2a, 0x9b, 0x27, 0x4d, 0x4b, 0x99, 0xfa, 0xd1, 0xff, 0x80, 0x05,
	0x92, 0x80, 0x59, 0x86, 0x86, 0x94, 0x21, 0x31, 0xf3, 0xdf, 0xfa, 0x6d, 0xdd, 0x84, 0x92, 0xb7,
	0x5c, 0x57, 0x24, 0xe6, 0x61, 0xd2, 0xb6, 0xd1, 0x13, 0x5c, 0x56, 0x28, 0x5d, 0xe9, 0x0d, 0xca,
	0xb4, 0x91, 0x7b, 0x2b, 0x7f, 0x7d, 0xb1, 0xbe, 0xaf, 0x70, 0x8a, 0xd1, 0xd1, 0x29, 0xa2, 0x57,
	0xb8, 0xee, 0xea, 0x1d, 0xaf, 0xdd, 0xb3, 0xc1, 0xfa, 0x36, 0x86, 0x85, 0x63, 0x58, 0x56, 0x28,
	0x6b, 0xa3, 0x3f, 0x6a, 0xb5, 0x97, 0x0e, 0xe7, 0x1f, 0x0c, 0x25, 0x7a, 0xf6, 0x9b, 0x74, 0x08,
	0x87, 0x9b, 0x77, 0x12, 0x60, 0xdd, 0x04, 0x86, 0xb7, 0x78, 0x80, 0x8b, 0xbe, 0xa8, 0x8b, 0xa7,
	0x0d, 0x8c, 0x1d, 0x02, 0x1b, 0x14, 0x48, 0xa7, 0xfe, 0x9f, 0xba, 0xff, 0x09, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0xcb, 0x12, 0x4c, 0x92, 0x02, 0x00, 0x00,
}