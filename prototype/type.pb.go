// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prototype/type.proto

package prototype // import "github.com/coschain/contentos-go/prototype"

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

type AccountName struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountName) Reset()         { *m = AccountName{} }
func (m *AccountName) String() string { return proto.CompactTextString(m) }
func (*AccountName) ProtoMessage()    {}
func (*AccountName) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{0}
}
func (m *AccountName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountName.Unmarshal(m, b)
}
func (m *AccountName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountName.Marshal(b, m, deterministic)
}
func (dst *AccountName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountName.Merge(dst, src)
}
func (m *AccountName) XXX_Size() int {
	return xxx_messageInfo_AccountName.Size(m)
}
func (m *AccountName) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountName.DiscardUnknown(m)
}

var xxx_messageInfo_AccountName proto.InternalMessageInfo

func (m *AccountName) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ChainId struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainId) Reset()         { *m = ChainId{} }
func (m *ChainId) String() string { return proto.CompactTextString(m) }
func (*ChainId) ProtoMessage()    {}
func (*ChainId) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{1}
}
func (m *ChainId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainId.Unmarshal(m, b)
}
func (m *ChainId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainId.Marshal(b, m, deterministic)
}
func (dst *ChainId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainId.Merge(dst, src)
}
func (m *ChainId) XXX_Size() int {
	return xxx_messageInfo_ChainId.Size(m)
}
func (m *ChainId) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainId.DiscardUnknown(m)
}

var xxx_messageInfo_ChainId proto.InternalMessageInfo

func (m *ChainId) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Coin struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coin) Reset()         { *m = Coin{} }
func (m *Coin) String() string { return proto.CompactTextString(m) }
func (*Coin) ProtoMessage()    {}
func (*Coin) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{2}
}
func (m *Coin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coin.Unmarshal(m, b)
}
func (m *Coin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coin.Marshal(b, m, deterministic)
}
func (dst *Coin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coin.Merge(dst, src)
}
func (m *Coin) XXX_Size() int {
	return xxx_messageInfo_Coin.Size(m)
}
func (m *Coin) XXX_DiscardUnknown() {
	xxx_messageInfo_Coin.DiscardUnknown(m)
}

var xxx_messageInfo_Coin proto.InternalMessageInfo

func (m *Coin) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Vest struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vest) Reset()         { *m = Vest{} }
func (m *Vest) String() string { return proto.CompactTextString(m) }
func (*Vest) ProtoMessage()    {}
func (*Vest) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{3}
}
func (m *Vest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vest.Unmarshal(m, b)
}
func (m *Vest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vest.Marshal(b, m, deterministic)
}
func (dst *Vest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vest.Merge(dst, src)
}
func (m *Vest) XXX_Size() int {
	return xxx_messageInfo_Vest.Size(m)
}
func (m *Vest) XXX_DiscardUnknown() {
	xxx_messageInfo_Vest.DiscardUnknown(m)
}

var xxx_messageInfo_Vest proto.InternalMessageInfo

func (m *Vest) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type PublicKeyType struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKeyType) Reset()         { *m = PublicKeyType{} }
func (m *PublicKeyType) String() string { return proto.CompactTextString(m) }
func (*PublicKeyType) ProtoMessage()    {}
func (*PublicKeyType) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{4}
}
func (m *PublicKeyType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKeyType.Unmarshal(m, b)
}
func (m *PublicKeyType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKeyType.Marshal(b, m, deterministic)
}
func (dst *PublicKeyType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKeyType.Merge(dst, src)
}
func (m *PublicKeyType) XXX_Size() int {
	return xxx_messageInfo_PublicKeyType.Size(m)
}
func (m *PublicKeyType) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKeyType.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKeyType proto.InternalMessageInfo

func (m *PublicKeyType) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PrivateKeyType struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivateKeyType) Reset()         { *m = PrivateKeyType{} }
func (m *PrivateKeyType) String() string { return proto.CompactTextString(m) }
func (*PrivateKeyType) ProtoMessage()    {}
func (*PrivateKeyType) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{5}
}
func (m *PrivateKeyType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateKeyType.Unmarshal(m, b)
}
func (m *PrivateKeyType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateKeyType.Marshal(b, m, deterministic)
}
func (dst *PrivateKeyType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateKeyType.Merge(dst, src)
}
func (m *PrivateKeyType) XXX_Size() int {
	return xxx_messageInfo_PrivateKeyType.Size(m)
}
func (m *PrivateKeyType) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateKeyType.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateKeyType proto.InternalMessageInfo

func (m *PrivateKeyType) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Authority struct {
	Key                  *PublicKeyType `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Authority) Reset()         { *m = Authority{} }
func (m *Authority) String() string { return proto.CompactTextString(m) }
func (*Authority) ProtoMessage()    {}
func (*Authority) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{6}
}
func (m *Authority) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authority.Unmarshal(m, b)
}
func (m *Authority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authority.Marshal(b, m, deterministic)
}
func (dst *Authority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authority.Merge(dst, src)
}
func (m *Authority) XXX_Size() int {
	return xxx_messageInfo_Authority.Size(m)
}
func (m *Authority) XXX_DiscardUnknown() {
	xxx_messageInfo_Authority.DiscardUnknown(m)
}

var xxx_messageInfo_Authority proto.InternalMessageInfo

func (m *Authority) GetKey() *PublicKeyType {
	if m != nil {
		return m.Key
	}
	return nil
}

type TimePointSec struct {
	UtcSeconds           uint32   `protobuf:"varint,1,opt,name=utc_seconds,json=utcSeconds,proto3" json:"utc_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimePointSec) Reset()         { *m = TimePointSec{} }
func (m *TimePointSec) String() string { return proto.CompactTextString(m) }
func (*TimePointSec) ProtoMessage()    {}
func (*TimePointSec) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{7}
}
func (m *TimePointSec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimePointSec.Unmarshal(m, b)
}
func (m *TimePointSec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimePointSec.Marshal(b, m, deterministic)
}
func (dst *TimePointSec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimePointSec.Merge(dst, src)
}
func (m *TimePointSec) XXX_Size() int {
	return xxx_messageInfo_TimePointSec.Size(m)
}
func (m *TimePointSec) XXX_DiscardUnknown() {
	xxx_messageInfo_TimePointSec.DiscardUnknown(m)
}

var xxx_messageInfo_TimePointSec proto.InternalMessageInfo

func (m *TimePointSec) GetUtcSeconds() uint32 {
	if m != nil {
		return m.UtcSeconds
	}
	return 0
}

type SignatureType struct {
	Sig                  []byte   `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureType) Reset()         { *m = SignatureType{} }
func (m *SignatureType) String() string { return proto.CompactTextString(m) }
func (*SignatureType) ProtoMessage()    {}
func (*SignatureType) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{8}
}
func (m *SignatureType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureType.Unmarshal(m, b)
}
func (m *SignatureType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureType.Marshal(b, m, deterministic)
}
func (dst *SignatureType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureType.Merge(dst, src)
}
func (m *SignatureType) XXX_Size() int {
	return xxx_messageInfo_SignatureType.Size(m)
}
func (m *SignatureType) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureType.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureType proto.InternalMessageInfo

func (m *SignatureType) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

type Sha256 struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sha256) Reset()         { *m = Sha256{} }
func (m *Sha256) String() string { return proto.CompactTextString(m) }
func (*Sha256) ProtoMessage()    {}
func (*Sha256) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{9}
}
func (m *Sha256) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sha256.Unmarshal(m, b)
}
func (m *Sha256) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sha256.Marshal(b, m, deterministic)
}
func (dst *Sha256) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sha256.Merge(dst, src)
}
func (m *Sha256) XXX_Size() int {
	return xxx_messageInfo_Sha256.Size(m)
}
func (m *Sha256) XXX_DiscardUnknown() {
	xxx_messageInfo_Sha256.DiscardUnknown(m)
}

var xxx_messageInfo_Sha256 proto.InternalMessageInfo

func (m *Sha256) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type ChainProperties struct {
	AccountCreationFee   *Coin    `protobuf:"bytes,1,opt,name=account_creation_fee,json=accountCreationFee,proto3" json:"account_creation_fee,omitempty"`
	MaximumBlockSize     uint32   `protobuf:"varint,2,opt,name=maximum_block_size,json=maximumBlockSize,proto3" json:"maximum_block_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainProperties) Reset()         { *m = ChainProperties{} }
func (m *ChainProperties) String() string { return proto.CompactTextString(m) }
func (*ChainProperties) ProtoMessage()    {}
func (*ChainProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{10}
}
func (m *ChainProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainProperties.Unmarshal(m, b)
}
func (m *ChainProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainProperties.Marshal(b, m, deterministic)
}
func (dst *ChainProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainProperties.Merge(dst, src)
}
func (m *ChainProperties) XXX_Size() int {
	return xxx_messageInfo_ChainProperties.Size(m)
}
func (m *ChainProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainProperties.DiscardUnknown(m)
}

var xxx_messageInfo_ChainProperties proto.InternalMessageInfo

func (m *ChainProperties) GetAccountCreationFee() *Coin {
	if m != nil {
		return m.AccountCreationFee
	}
	return nil
}

func (m *ChainProperties) GetMaximumBlockSize() uint32 {
	if m != nil {
		return m.MaximumBlockSize
	}
	return 0
}

type DynamicProperties struct {
	Id                   int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	HeadBlockId          *Sha256       `protobuf:"bytes,2,opt,name=head_block_id,json=headBlockId,proto3" json:"head_block_id,omitempty"`
	HeadBlockNumber      uint64        `protobuf:"varint,3,opt,name=head_block_number,json=headBlockNumber,proto3" json:"head_block_number,omitempty"`
	MaximumBlockSize     uint32        `protobuf:"varint,4,opt,name=maximum_block_size,json=maximumBlockSize,proto3" json:"maximum_block_size,omitempty"`
	TotalCos             *Coin         `protobuf:"bytes,5,opt,name=total_cos,json=totalCos,proto3" json:"total_cos,omitempty"`
	Time                 *TimePointSec `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
	CurrentWitness       *AccountName  `protobuf:"bytes,7,opt,name=current_witness,json=currentWitness,proto3" json:"current_witness,omitempty"`
	IrreversibleBlockNum uint64        `protobuf:"varint,8,opt,name=irreversible_block_num,json=irreversibleBlockNum,proto3" json:"irreversible_block_num,omitempty"`
	Tps                  uint32        `protobuf:"varint,9,opt,name=tps,proto3" json:"tps,omitempty"`
	TotalVestingShares   *Vest         `protobuf:"bytes,10,opt,name=total_vesting_shares,json=totalVestingShares,proto3" json:"total_vesting_shares,omitempty"`
	CurrentSupply        *Coin         `protobuf:"bytes,11,opt,name=current_supply,json=currentSupply,proto3" json:"current_supply,omitempty"`
	CurrentAslot         uint32        `protobuf:"varint,12,opt,name=current_aslot,json=currentAslot,proto3" json:"current_aslot,omitempty"`
	WeightedVps          uint64        `protobuf:"varint,13,opt,name=weighted_vps,json=weightedVps,proto3" json:"weighted_vps,omitempty"`
	PostRewards          *Vest         `protobuf:"bytes,14,opt,name=post_rewards,json=postRewards,proto3" json:"post_rewards,omitempty"`
	ReplyRewards         *Vest         `protobuf:"bytes,15,opt,name=reply_rewards,json=replyRewards,proto3" json:"reply_rewards,omitempty"`
	TotalTrxCnt          uint64        `protobuf:"varint,16,opt,name=total_trx_cnt,json=totalTrxCnt,proto3" json:"total_trx_cnt,omitempty"`
	TotalPostCnt         uint64        `protobuf:"varint,17,opt,name=total_post_cnt,json=totalPostCnt,proto3" json:"total_post_cnt,omitempty"`
	TotalUserCnt         uint64        `protobuf:"varint,18,opt,name=total_user_cnt,json=totalUserCnt,proto3" json:"total_user_cnt,omitempty"`
	MaxTps               uint32        `protobuf:"varint,19,opt,name=max_tps,json=maxTps,proto3" json:"max_tps,omitempty"`
	MaxTpsBlockNum       uint64        `protobuf:"varint,20,opt,name=max_tps_block_num,json=maxTpsBlockNum,proto3" json:"max_tps_block_num,omitempty"`
	HeadBlockPrefix      uint32        `protobuf:"varint,21,opt,name=head_block_prefix,json=headBlockPrefix,proto3" json:"head_block_prefix,omitempty"`
	ReportRewards        *Vest         `protobuf:"bytes,22,opt,name=report_rewards,json=reportRewards,proto3" json:"report_rewards,omitempty"`
	IthYear              uint32        `protobuf:"varint,23,opt,name=ith_year,json=ithYear,proto3" json:"ith_year,omitempty"`
	AnnualBudget         *Vest         `protobuf:"bytes,24,opt,name=annual_budget,json=annualBudget,proto3" json:"annual_budget,omitempty"`
	AnnualMinted         *Vest         `protobuf:"bytes,25,opt,name=annual_minted,json=annualMinted,proto3" json:"annual_minted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DynamicProperties) Reset()         { *m = DynamicProperties{} }
func (m *DynamicProperties) String() string { return proto.CompactTextString(m) }
func (*DynamicProperties) ProtoMessage()    {}
func (*DynamicProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{11}
}
func (m *DynamicProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicProperties.Unmarshal(m, b)
}
func (m *DynamicProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicProperties.Marshal(b, m, deterministic)
}
func (dst *DynamicProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicProperties.Merge(dst, src)
}
func (m *DynamicProperties) XXX_Size() int {
	return xxx_messageInfo_DynamicProperties.Size(m)
}
func (m *DynamicProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicProperties.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicProperties proto.InternalMessageInfo

func (m *DynamicProperties) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DynamicProperties) GetHeadBlockId() *Sha256 {
	if m != nil {
		return m.HeadBlockId
	}
	return nil
}

func (m *DynamicProperties) GetHeadBlockNumber() uint64 {
	if m != nil {
		return m.HeadBlockNumber
	}
	return 0
}

func (m *DynamicProperties) GetMaximumBlockSize() uint32 {
	if m != nil {
		return m.MaximumBlockSize
	}
	return 0
}

func (m *DynamicProperties) GetTotalCos() *Coin {
	if m != nil {
		return m.TotalCos
	}
	return nil
}

func (m *DynamicProperties) GetTime() *TimePointSec {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *DynamicProperties) GetCurrentWitness() *AccountName {
	if m != nil {
		return m.CurrentWitness
	}
	return nil
}

func (m *DynamicProperties) GetIrreversibleBlockNum() uint64 {
	if m != nil {
		return m.IrreversibleBlockNum
	}
	return 0
}

func (m *DynamicProperties) GetTps() uint32 {
	if m != nil {
		return m.Tps
	}
	return 0
}

func (m *DynamicProperties) GetTotalVestingShares() *Vest {
	if m != nil {
		return m.TotalVestingShares
	}
	return nil
}

func (m *DynamicProperties) GetCurrentSupply() *Coin {
	if m != nil {
		return m.CurrentSupply
	}
	return nil
}

func (m *DynamicProperties) GetCurrentAslot() uint32 {
	if m != nil {
		return m.CurrentAslot
	}
	return 0
}

func (m *DynamicProperties) GetWeightedVps() uint64 {
	if m != nil {
		return m.WeightedVps
	}
	return 0
}

func (m *DynamicProperties) GetPostRewards() *Vest {
	if m != nil {
		return m.PostRewards
	}
	return nil
}

func (m *DynamicProperties) GetReplyRewards() *Vest {
	if m != nil {
		return m.ReplyRewards
	}
	return nil
}

func (m *DynamicProperties) GetTotalTrxCnt() uint64 {
	if m != nil {
		return m.TotalTrxCnt
	}
	return 0
}

func (m *DynamicProperties) GetTotalPostCnt() uint64 {
	if m != nil {
		return m.TotalPostCnt
	}
	return 0
}

func (m *DynamicProperties) GetTotalUserCnt() uint64 {
	if m != nil {
		return m.TotalUserCnt
	}
	return 0
}

func (m *DynamicProperties) GetMaxTps() uint32 {
	if m != nil {
		return m.MaxTps
	}
	return 0
}

func (m *DynamicProperties) GetMaxTpsBlockNum() uint64 {
	if m != nil {
		return m.MaxTpsBlockNum
	}
	return 0
}

func (m *DynamicProperties) GetHeadBlockPrefix() uint32 {
	if m != nil {
		return m.HeadBlockPrefix
	}
	return 0
}

func (m *DynamicProperties) GetReportRewards() *Vest {
	if m != nil {
		return m.ReportRewards
	}
	return nil
}

func (m *DynamicProperties) GetIthYear() uint32 {
	if m != nil {
		return m.IthYear
	}
	return 0
}

func (m *DynamicProperties) GetAnnualBudget() *Vest {
	if m != nil {
		return m.AnnualBudget
	}
	return nil
}

func (m *DynamicProperties) GetAnnualMinted() *Vest {
	if m != nil {
		return m.AnnualMinted
	}
	return nil
}

type InternalRewardsKeeper struct {
	Id                   int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Rewards              map[string]*Vest `protobuf:"bytes,2,rep,name=rewards,proto3" json:"rewards,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *InternalRewardsKeeper) Reset()         { *m = InternalRewardsKeeper{} }
func (m *InternalRewardsKeeper) String() string { return proto.CompactTextString(m) }
func (*InternalRewardsKeeper) ProtoMessage()    {}
func (*InternalRewardsKeeper) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{12}
}
func (m *InternalRewardsKeeper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalRewardsKeeper.Unmarshal(m, b)
}
func (m *InternalRewardsKeeper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalRewardsKeeper.Marshal(b, m, deterministic)
}
func (dst *InternalRewardsKeeper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalRewardsKeeper.Merge(dst, src)
}
func (m *InternalRewardsKeeper) XXX_Size() int {
	return xxx_messageInfo_InternalRewardsKeeper.Size(m)
}
func (m *InternalRewardsKeeper) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalRewardsKeeper.DiscardUnknown(m)
}

var xxx_messageInfo_InternalRewardsKeeper proto.InternalMessageInfo

func (m *InternalRewardsKeeper) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InternalRewardsKeeper) GetRewards() map[string]*Vest {
	if m != nil {
		return m.Rewards
	}
	return nil
}

type BeneficiaryRouteType struct {
	Name                 *AccountName `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight               uint32       `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BeneficiaryRouteType) Reset()         { *m = BeneficiaryRouteType{} }
func (m *BeneficiaryRouteType) String() string { return proto.CompactTextString(m) }
func (*BeneficiaryRouteType) ProtoMessage()    {}
func (*BeneficiaryRouteType) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_6ec313e39d9a16ac, []int{13}
}
func (m *BeneficiaryRouteType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeneficiaryRouteType.Unmarshal(m, b)
}
func (m *BeneficiaryRouteType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeneficiaryRouteType.Marshal(b, m, deterministic)
}
func (dst *BeneficiaryRouteType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeneficiaryRouteType.Merge(dst, src)
}
func (m *BeneficiaryRouteType) XXX_Size() int {
	return xxx_messageInfo_BeneficiaryRouteType.Size(m)
}
func (m *BeneficiaryRouteType) XXX_DiscardUnknown() {
	xxx_messageInfo_BeneficiaryRouteType.DiscardUnknown(m)
}

var xxx_messageInfo_BeneficiaryRouteType proto.InternalMessageInfo

func (m *BeneficiaryRouteType) GetName() *AccountName {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *BeneficiaryRouteType) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountName)(nil), "prototype.account_name")
	proto.RegisterType((*ChainId)(nil), "prototype.chain_id")
	proto.RegisterType((*Coin)(nil), "prototype.coin")
	proto.RegisterType((*Vest)(nil), "prototype.vest")
	proto.RegisterType((*PublicKeyType)(nil), "prototype.public_key_type")
	proto.RegisterType((*PrivateKeyType)(nil), "prototype.private_key_type")
	proto.RegisterType((*Authority)(nil), "prototype.authority")
	proto.RegisterType((*TimePointSec)(nil), "prototype.time_point_sec")
	proto.RegisterType((*SignatureType)(nil), "prototype.signature_type")
	proto.RegisterType((*Sha256)(nil), "prototype.sha256")
	proto.RegisterType((*ChainProperties)(nil), "prototype.chain_properties")
	proto.RegisterType((*DynamicProperties)(nil), "prototype.dynamic_properties")
	proto.RegisterType((*InternalRewardsKeeper)(nil), "prototype.internal_rewards_keeper")
	proto.RegisterMapType((map[string]*Vest)(nil), "prototype.internal_rewards_keeper.RewardsEntry")
	proto.RegisterType((*BeneficiaryRouteType)(nil), "prototype.beneficiary_route_type")
}

func init() { proto.RegisterFile("prototype/type.proto", fileDescriptor_type_6ec313e39d9a16ac) }

var fileDescriptor_type_6ec313e39d9a16ac = []byte{
	// 940 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0x6d, 0x6f, 0x1b, 0x45,
	0x10, 0xc7, 0x65, 0xc7, 0x4d, 0xe2, 0xf1, 0x43, 0x92, 0x25, 0x24, 0x9b, 0x0a, 0x89, 0x70, 0xb4,
	0xa8, 0x2a, 0x69, 0x0c, 0xa6, 0xad, 0x80, 0x57, 0x34, 0x11, 0x48, 0x15, 0x02, 0x55, 0x97, 0x52,
	0x04, 0x12, 0x5a, 0xad, 0xef, 0x26, 0xbe, 0x55, 0xee, 0x76, 0x4f, 0xbb, 0x7b, 0x8e, 0xaf, 0x6f,
	0xf9, 0x64, 0x7c, 0x1f, 0x3e, 0x04, 0xda, 0xbd, 0x3b, 0xe7, 0x42, 0x9d, 0xbc, 0xb1, 0xf6, 0x66,
	0x7e, 0x33, 0x9e, 0xf9, 0xcf, 0x3e, 0xc0, 0x7e, 0xae, 0x95, 0x55, 0xb6, 0xcc, 0x71, 0xe2, 0x7e,
	0x4e, 0xfd, 0x27, 0xe9, 0xaf, 0xac, 0xc1, 0x23, 0x18, 0xf2, 0x28, 0x52, 0x85, 0xb4, 0x4c, 0xf2,
	0x0c, 0xc9, 0x3e, 0x3c, 0x58, 0xf0, 0xb4, 0x40, 0xda, 0x39, 0xee, 0x3c, 0xe9, 0x87, 0xd5, 0x47,
	0x70, 0x0c, 0xdb, 0x51, 0xc2, 0x85, 0x64, 0x22, 0xbe, 0x4d, 0x8c, 0x6e, 0x88, 0x5e, 0xa4, 0x84,
	0x24, 0xb4, 0xed, 0xed, 0x9d, 0x75, 0xbf, 0xea, 0xb4, 0x88, 0x05, 0x1a, 0x7b, 0x0f, 0xf1, 0x18,
	0x76, 0xf2, 0x62, 0x96, 0x8a, 0x88, 0x5d, 0x61, 0xc9, 0x5c, 0x79, 0x84, 0x40, 0x2f, 0xe6, 0x96,
	0x7b, 0x76, 0x18, 0xfa, 0x75, 0xf0, 0x05, 0xec, 0xe6, 0x5a, 0x2c, 0xb8, 0xc5, 0xfb, 0xb9, 0xef,
	0xa0, 0xcf, 0x0b, 0x9b, 0x28, 0x2d, 0x6c, 0x49, 0x4e, 0x60, 0xe3, 0x0a, 0x4b, 0xef, 0x1f, 0x4c,
	0x1f, 0x9e, 0xae, 0x04, 0x38, 0xfd, 0xdf, 0x3f, 0x86, 0x0e, 0x0b, 0xbe, 0x86, 0xb1, 0x15, 0x19,
	0xb2, 0x5c, 0x09, 0x69, 0x99, 0xc1, 0x88, 0x7c, 0x0a, 0x83, 0xc2, 0x46, 0x6e, 0xa9, 0x64, 0x6c,
	0xea, 0xde, 0xa1, 0xb0, 0xd1, 0x45, 0x65, 0x09, 0x02, 0x18, 0x1b, 0x31, 0x97, 0xdc, 0x16, 0x1a,
	0xab, 0x9a, 0x76, 0x61, 0xc3, 0x88, 0x79, 0x5d, 0x92, 0x5b, 0x06, 0x9f, 0xc0, 0xa6, 0x49, 0xf8,
	0xf4, 0xc5, 0x4b, 0x57, 0x6f, 0xc2, 0x4d, 0xd2, 0xd4, 0xeb, 0xd6, 0xc1, 0xdf, 0x1d, 0xd8, 0xad,
	0x54, 0xce, 0xb5, 0xca, 0x51, 0x5b, 0x81, 0x86, 0xbc, 0x82, 0xfd, 0x66, 0x3e, 0x91, 0x46, 0x6e,
	0x85, 0x92, 0xec, 0x12, 0xb1, 0x6e, 0x64, 0xa7, 0xd5, 0x88, 0x93, 0x3f, 0x24, 0x35, 0x7c, 0x5e,
	0xb3, 0x3f, 0x21, 0x92, 0x13, 0x20, 0x19, 0x5f, 0x8a, 0xac, 0xc8, 0xd8, 0x2c, 0x55, 0xd1, 0x15,
	0x33, 0xe2, 0x3d, 0xd2, 0xae, 0xef, 0x60, 0xb7, 0xf6, 0x9c, 0x39, 0xc7, 0x85, 0x78, 0x8f, 0xc1,
	0xbf, 0xdb, 0x40, 0xe2, 0x52, 0xf2, 0x4c, 0x44, 0xed, 0x3a, 0xc6, 0xd0, 0x15, 0xb1, 0xff, 0xd7,
	0x07, 0x61, 0x57, 0xc4, 0xe4, 0x05, 0x8c, 0x12, 0xe4, 0x71, 0x9d, 0x51, 0xc4, 0x3e, 0xdf, 0x60,
	0xba, 0xd7, 0x2a, 0xa8, 0x6a, 0x35, 0x1c, 0x38, 0xce, 0xe7, 0x7f, 0x1d, 0x93, 0xa7, 0xb0, 0xd7,
	0x0a, 0x93, 0x45, 0x36, 0x43, 0x4d, 0x37, 0xdc, 0x46, 0x08, 0x77, 0x56, 0xdc, 0xaf, 0xde, 0x7c,
	0x47, 0xdd, 0xbd, 0xf5, 0x75, 0x93, 0x13, 0xe8, 0x5b, 0x65, 0x79, 0xca, 0x22, 0x65, 0xe8, 0x83,
	0xf5, 0xea, 0x6c, 0x7b, 0xe2, 0x5c, 0x19, 0xf2, 0x0c, 0x7a, 0x6e, 0xc0, 0x74, 0xd3, 0x83, 0x47,
	0x2d, 0xf0, 0xf6, 0xdc, 0x43, 0x8f, 0x91, 0x1f, 0x60, 0x27, 0x2a, 0xb4, 0x46, 0x69, 0xd9, 0xb5,
	0xb0, 0x12, 0x8d, 0xa1, 0x5b, 0x3e, 0xf2, 0xb0, 0x15, 0xd9, 0x3e, 0x47, 0xe1, 0xb8, 0xe6, 0x7f,
	0xaf, 0x70, 0xf2, 0x1c, 0x0e, 0x84, 0xd6, 0xb8, 0x40, 0x6d, 0xc4, 0x2c, 0xc5, 0x1b, 0x01, 0xe8,
	0xb6, 0xef, 0x7e, 0xbf, 0xed, 0x6d, 0x54, 0x70, 0x5b, 0xc8, 0xe6, 0x86, 0xf6, 0x7d, 0xcf, 0x6e,
	0xe9, 0xf6, 0x43, 0xd5, 0xa6, 0x3b, 0x4b, 0x42, 0xce, 0x99, 0x49, 0xb8, 0x46, 0x43, 0xe1, 0x83,
	0x8e, 0x1d, 0x10, 0x12, 0x0f, 0xbf, 0xab, 0xd8, 0x0b, 0x8f, 0x92, 0x97, 0xd0, 0x14, 0xc7, 0x4c,
	0x91, 0xe7, 0x69, 0x49, 0x07, 0xeb, 0xe5, 0x1a, 0xd5, 0xd8, 0x85, 0xa7, 0xc8, 0xe7, 0xd0, 0x18,
	0x18, 0x37, 0xa9, 0xb2, 0x74, 0xe8, 0xcb, 0x1a, 0xd6, 0xc6, 0x57, 0xce, 0x46, 0x3e, 0x83, 0xe1,
	0x35, 0x8a, 0x79, 0x62, 0x31, 0x66, 0x8b, 0xdc, 0xd0, 0x91, 0xef, 0x6e, 0xd0, 0xd8, 0xde, 0xe5,
	0x86, 0x4c, 0x61, 0x98, 0x2b, 0x63, 0x99, 0xc6, 0x6b, 0xae, 0x63, 0x43, 0xc7, 0xeb, 0x4b, 0x1f,
	0x38, 0x28, 0xac, 0x18, 0xf2, 0x1c, 0x46, 0x1a, 0xf3, 0xb4, 0x5c, 0x05, 0xed, 0xac, 0x0f, 0x1a,
	0x7a, 0xaa, 0x89, 0x0a, 0x60, 0x54, 0x89, 0x65, 0xf5, 0x92, 0x45, 0xd2, 0xd2, 0xdd, 0xaa, 0x1a,
	0x6f, 0x7c, 0xab, 0x97, 0xe7, 0xd2, 0x92, 0x47, 0x30, 0xae, 0x18, 0x5f, 0x93, 0x83, 0xf6, 0x3c,
	0x34, 0xf4, 0xd6, 0x37, 0xca, 0xd8, 0x5b, 0x54, 0x61, 0x50, 0x7b, 0x8a, 0xb4, 0xa8, 0xdf, 0x0c,
	0x6a, 0x47, 0x1d, 0xc2, 0x56, 0xc6, 0x97, 0xcc, 0x8d, 0xec, 0x23, 0xaf, 0xcd, 0x66, 0xc6, 0x97,
	0x6f, 0x73, 0xb7, 0xdd, 0xf6, 0x6a, 0x47, 0x6b, 0xf0, 0xfb, 0xab, 0xfb, 0x6f, 0x5c, 0x61, 0xab,
	0xb1, 0xdf, 0x3e, 0x25, 0xb9, 0xc6, 0x4b, 0xb1, 0xa4, 0x1f, 0xfb, 0x8c, 0x37, 0xa7, 0xe4, 0x8d,
	0x37, 0xbb, 0x69, 0x6a, 0xcc, 0x95, 0xbe, 0xd1, 0xf3, 0x60, 0xbd, 0x34, 0xa3, 0x0a, 0x6b, 0xb4,
	0x39, 0x82, 0x6d, 0x61, 0x13, 0x56, 0x22, 0xd7, 0xf4, 0xd0, 0xa7, 0xde, 0x12, 0x36, 0xf9, 0x03,
	0xb9, 0x76, 0x62, 0x73, 0x29, 0x0b, 0x9e, 0xb2, 0x59, 0x11, 0xcf, 0xd1, 0x52, 0x7a, 0x87, 0xd8,
	0x15, 0x75, 0xe6, 0xa1, 0x56, 0x54, 0x26, 0xa4, 0xc5, 0x98, 0x1e, 0xdd, 0x1b, 0xf5, 0x8b, 0x87,
	0x82, 0x7f, 0x3a, 0x70, 0xe8, 0x56, 0x5a, 0xf2, 0xb4, 0xe9, 0x80, 0x5d, 0x21, 0xe6, 0xa8, 0x3f,
	0xb8, 0x73, 0x5e, 0xc3, 0x56, 0xd3, 0x63, 0xf7, 0x78, 0xe3, 0xc9, 0x60, 0x3a, 0x69, 0xe5, 0xbe,
	0x23, 0xc9, 0x69, 0xdd, 0xef, 0x8f, 0xd2, 0xea, 0x32, 0x6c, 0xe2, 0x1f, 0xfe, 0x0c, 0xc3, 0xb6,
	0xc3, 0x1d, 0xb4, 0xe6, 0x79, 0xe8, 0xfb, 0x27, 0x80, 0x3c, 0x6e, 0x9e, 0xa9, 0xee, 0xfa, 0x36,
	0x2a, 0xef, 0xf7, 0xdd, 0x6f, 0x3b, 0xc1, 0x5f, 0x70, 0x30, 0x43, 0x89, 0x97, 0x22, 0x12, 0x5c,
	0x97, 0x4c, 0xab, 0xc2, 0xd6, 0x4f, 0xc0, 0x97, 0xd0, 0x73, 0xb7, 0x41, 0x7d, 0x5b, 0xdf, 0x79,
	0x59, 0x78, 0x88, 0x1c, 0xc0, 0x66, 0x75, 0x4c, 0xea, 0xbb, 0xb9, 0xfe, 0x3a, 0x3b, 0xf9, 0xf3,
	0xe9, 0x5c, 0xd8, 0xa4, 0x98, 0x9d, 0x46, 0x2a, 0x9b, 0x44, 0xca, 0xf8, 0x47, 0x62, 0x12, 0x29,
	0x69, 0x51, 0x5a, 0x65, 0x9e, 0xcd, 0xd5, 0x64, 0x95, 0x78, 0xb6, 0xe9, 0x97, 0xdf, 0xfc, 0x17,
	0x00, 0x00, 0xff, 0xff, 0xf1, 0x0e, 0x44, 0x43, 0xfa, 0x07, 0x00, 0x00,
}
