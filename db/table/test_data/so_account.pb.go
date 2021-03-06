// Code generated by protoc-gen-go. DO NOT EDIT.
// source: so_account.proto

package test_data

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
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

type SoAccount struct {
	Name                 *prototype.AccountName   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreatedTime          *prototype.TimePointSec  `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Creator              *prototype.AccountName   `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	PubKey               *prototype.PublicKeyType `protobuf:"bytes,4,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Balance              *prototype.Coin          `protobuf:"bytes,5,opt,name=balance,proto3" json:"balance,omitempty"`
	VestingShares        *prototype.Vest          `protobuf:"bytes,6,opt,name=vesting_shares,json=vestingShares,proto3" json:"vesting_shares,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SoAccount) Reset()         { *m = SoAccount{} }
func (m *SoAccount) String() string { return proto.CompactTextString(m) }
func (*SoAccount) ProtoMessage()    {}
func (*SoAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac55b55358b6f556, []int{0}
}

func (m *SoAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoAccount.Unmarshal(m, b)
}
func (m *SoAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoAccount.Marshal(b, m, deterministic)
}
func (m *SoAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoAccount.Merge(m, src)
}
func (m *SoAccount) XXX_Size() int {
	return xxx_messageInfo_SoAccount.Size(m)
}
func (m *SoAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SoAccount proto.InternalMessageInfo

func (m *SoAccount) GetName() *prototype.AccountName {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SoAccount) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *SoAccount) GetCreator() *prototype.AccountName {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *SoAccount) GetPubKey() *prototype.PublicKeyType {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *SoAccount) GetBalance() *prototype.Coin {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *SoAccount) GetVestingShares() *prototype.Vest {
	if m != nil {
		return m.VestingShares
	}
	return nil
}

func init() {
	proto.RegisterType((*SoAccount)(nil), "test_data.so_account")
}

func init() { proto.RegisterFile("so_account.proto", fileDescriptor_ac55b55358b6f556) }

var fileDescriptor_ac55b55358b6f556 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8d, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0x47, 0x69, 0xbf, 0x7e, 0x33, 0x78, 0xeb, 0x3f, 0x82, 0x60, 0xec, 0x4a, 0x5c, 0x29, 0xc2,
	0x88, 0x16, 0x5c, 0xf9, 0x06, 0xee, 0xa2, 0xfb, 0x4b, 0x92, 0x5e, 0x34, 0xd8, 0x49, 0xc2, 0xe4,
	0x8e, 0x30, 0xaf, 0xe7, 0x93, 0x49, 0xe2, 0x54, 0x07, 0x04, 0x37, 0x81, 0xfb, 0x3b, 0xe7, 0x10,
	0x38, 0x4e, 0x01, 0xb5, 0xb5, 0xa1, 0xf7, 0xdc, 0xc4, 0x2e, 0x70, 0x10, 0x7b, 0x4c, 0x89, 0x71,
	0xa3, 0x59, 0xaf, 0x4e, 0xca, 0xc2, 0x43, 0xa4, 0x9b, 0xfc, 0x7c, 0x09, 0x17, 0x1f, 0x73, 0x80,
	0x9f, 0x4a, 0x5c, 0xc3, 0xc2, 0xeb, 0x96, 0xe4, 0xec, 0x7c, 0x76, 0xb9, 0xbc, 0x3b, 0x6d, 0xbe,
	0x9b, 0x66, 0x34, 0x30, 0x63, 0x55, 0x24, 0xf1, 0x00, 0xfb, 0xb6, 0x23, 0xcd, 0xb4, 0x41, 0x76,
	0x2d, 0xc9, 0x79, 0x89, 0xce, 0x26, 0x51, 0x9e, 0x31, 0x06, 0xe7, 0x19, 0x13, 0x59, 0xb5, 0x1c,
	0xf5, 0x67, 0xd7, 0x92, 0xb8, 0x85, 0xba, 0x9c, 0xa1, 0x93, 0xff, 0xfe, 0xfe, 0x6d, 0xe7, 0x89,
	0x35, 0xd4, 0xb1, 0x37, 0xf8, 0x46, 0x83, 0x5c, 0x94, 0x64, 0x35, 0x49, 0x62, 0x6f, 0xb6, 0xce,
	0x66, 0x88, 0xf9, 0x56, 0x55, 0xec, 0xcd, 0x23, 0x0d, 0xe2, 0x0a, 0x6a, 0xa3, 0xb7, 0xda, 0x5b,
	0x92, 0xff, 0x4b, 0x74, 0x34, 0x89, 0x6c, 0x70, 0x5e, 0xed, 0xb8, 0xb8, 0x87, 0xc3, 0x77, 0x4a,
	0xec, 0xfc, 0x0b, 0xa6, 0x57, 0xdd, 0x51, 0x92, 0xd5, 0xaf, 0x22, 0x0b, 0xea, 0x60, 0xd4, 0x9e,
	0x8a, 0x65, 0xaa, 0x82, 0xd7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xb3, 0x54, 0xe8, 0x80,
	0x01, 0x00, 0x00,
}
