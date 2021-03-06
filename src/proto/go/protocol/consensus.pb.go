

package protocol

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

type PbftMessageType int32

const (
	PbftMessageType_PBFT_TYPE_PREPREPARE              PbftMessageType = 0
	PbftMessageType_PBFT_TYPE_PREPARE                 PbftMessageType = 1
	PbftMessageType_PBFT_TYPE_COMMIT                  PbftMessageType = 2
	PbftMessageType_PBFT_TYPE_VIEWCHANGE              PbftMessageType = 3
	PbftMessageType_PBFT_TYPE_NEWVIEW                 PbftMessageType = 4
	PbftMessageType_PBFT_TYPE_VIEWCHANG_WITH_RAWVALUE PbftMessageType = 5
)

var PbftMessageType_name = map[int32]string{
	0: "PBFT_TYPE_PREPREPARE",
	1: "PBFT_TYPE_PREPARE",
	2: "PBFT_TYPE_COMMIT",
	3: "PBFT_TYPE_VIEWCHANGE",
	4: "PBFT_TYPE_NEWVIEW",
	5: "PBFT_TYPE_VIEWCHANG_WITH_RAWVALUE",
}
var PbftMessageType_value = map[string]int32{
	"PBFT_TYPE_PREPREPARE":              0,
	"PBFT_TYPE_PREPARE":                 1,
	"PBFT_TYPE_COMMIT":                  2,
	"PBFT_TYPE_VIEWCHANGE":              3,
	"PBFT_TYPE_NEWVIEW":                 4,
	"PBFT_TYPE_VIEWCHANG_WITH_RAWVALUE": 5,
}

func (x PbftMessageType) String() string {
	return proto.EnumName(PbftMessageType_name, int32(x))
}
func (PbftMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{0}
}

type PbftValueType int32

const (
	PbftValueType_PBFT_VALUE_TX    PbftValueType = 0
	PbftValueType_PBFT_VALUE_TXSET PbftValueType = 1
)

var PbftValueType_name = map[int32]string{
	0: "PBFT_VALUE_TX",
	1: "PBFT_VALUE_TXSET",
}
var PbftValueType_value = map[string]int32{
	"PBFT_VALUE_TX":    0,
	"PBFT_VALUE_TXSET": 1,
}

func (x PbftValueType) String() string {
	return proto.EnumName(PbftValueType_name, int32(x))
}
func (PbftValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{1}
}

type FeeConfig_Type int32

const (
	FeeConfig_UNKNOWN      FeeConfig_Type = 0
	FeeConfig_GAS_PRICE    FeeConfig_Type = 1
	FeeConfig_BASE_RESERVE FeeConfig_Type = 2
)

var FeeConfig_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "GAS_PRICE",
	2: "BASE_RESERVE",
}
var FeeConfig_Type_value = map[string]int32{
	"UNKNOWN":      0,
	"GAS_PRICE":    1,
	"BASE_RESERVE": 2,
}

func (x FeeConfig_Type) String() string {
	return proto.EnumName(FeeConfig_Type_name, int32(x))
}
func (FeeConfig_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{12, 0}
}

// PBFT protocol
type PbftPrePrepare struct {
	ViewNumber           int64    `protobuf:"varint,1,opt,name=view_number,json=viewNumber" json:"view_number,omitempty"`
	Sequence             int64    `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	ReplicaId            int64    `protobuf:"varint,3,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	Value                []byte   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	ValueDigest          []byte   `protobuf:"bytes,5,opt,name=value_digest,json=valueDigest,proto3" json:"value_digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbftPrePrepare) Reset()         { *m = PbftPrePrepare{} }
func (m *PbftPrePrepare) String() string { return proto.CompactTextString(m) }
func (*PbftPrePrepare) ProtoMessage()    {}
func (*PbftPrePrepare) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{0}
}
func (m *PbftPrePrepare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftPrePrepare.Unmarshal(m, b)
}
func (m *PbftPrePrepare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftPrePrepare.Marshal(b, m, deterministic)
}
func (dst *PbftPrePrepare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftPrePrepare.Merge(dst, src)
}
func (m *PbftPrePrepare) XXX_Size() int {
	return xxx_messageInfo_PbftPrePrepare.Size(m)
}
func (m *PbftPrePrepare) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftPrePrepare.DiscardUnknown(m)
}

var xxx_messageInfo_PbftPrePrepare proto.InternalMessageInfo

func (m *PbftPrePrepare) GetViewNumber() int64 {
	if m != nil {
		return m.ViewNumber
	}
	return 0
}

func (m *PbftPrePrepare) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *PbftPrePrepare) GetReplicaId() int64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *PbftPrePrepare) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PbftPrePrepare) GetValueDigest() []byte {
	if m != nil {
		return m.ValueDigest
	}
	return nil
}

type PbftPrepare struct {
	ViewNumber           int64    `protobuf:"varint,1,opt,name=view_number,json=viewNumber" json:"view_number,omitempty"`
	Sequence             int64    `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	ReplicaId            int64    `protobuf:"varint,3,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	ValueDigest          []byte   `protobuf:"bytes,4,opt,name=value_digest,json=valueDigest,proto3" json:"value_digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbftPrepare) Reset()         { *m = PbftPrepare{} }
func (m *PbftPrepare) String() string { return proto.CompactTextString(m) }
func (*PbftPrepare) ProtoMessage()    {}
func (*PbftPrepare) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{1}
}
func (m *PbftPrepare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftPrepare.Unmarshal(m, b)
}
func (m *PbftPrepare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftPrepare.Marshal(b, m, deterministic)
}
func (dst *PbftPrepare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftPrepare.Merge(dst, src)
}
func (m *PbftPrepare) XXX_Size() int {
	return xxx_messageInfo_PbftPrepare.Size(m)
}
func (m *PbftPrepare) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftPrepare.DiscardUnknown(m)
}

var xxx_messageInfo_PbftPrepare proto.InternalMessageInfo

func (m *PbftPrepare) GetViewNumber() int64 {
	if m != nil {
		return m.ViewNumber
	}
	return 0
}

func (m *PbftPrepare) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *PbftPrepare) GetReplicaId() int64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *PbftPrepare) GetValueDigest() []byte {
	if m != nil {
		return m.ValueDigest
	}
	return nil
}

type PbftCommit struct {
	ViewNumber           int64    `protobuf:"varint,1,opt,name=view_number,json=viewNumber" json:"view_number,omitempty"`
	Sequence             int64    `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	ReplicaId            int64    `protobuf:"varint,3,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	ValueDigest          []byte   `protobuf:"bytes,4,opt,name=value_digest,json=valueDigest,proto3" json:"value_digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbftCommit) Reset()         { *m = PbftCommit{} }
func (m *PbftCommit) String() string { return proto.CompactTextString(m) }
func (*PbftCommit) ProtoMessage()    {}
func (*PbftCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{2}
}
func (m *PbftCommit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftCommit.Unmarshal(m, b)
}
func (m *PbftCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftCommit.Marshal(b, m, deterministic)
}
func (dst *PbftCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftCommit.Merge(dst, src)
}
func (m *PbftCommit) XXX_Size() int {
	return xxx_messageInfo_PbftCommit.Size(m)
}
func (m *PbftCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftCommit.DiscardUnknown(m)
}

var xxx_messageInfo_PbftCommit proto.InternalMessageInfo

func (m *PbftCommit) GetViewNumber() int64 {
	if m != nil {
		return m.ViewNumber
	}
	return 0
}

func (m *PbftCommit) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *PbftCommit) GetReplicaId() int64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *PbftCommit) GetValueDigest() []byte {
	if m != nil {
		return m.ValueDigest
	}
	return nil
}

type PbftPreparedSet struct {
	PrePrepare           *PbftEnv   `protobuf:"bytes,1,opt,name=pre_prepare,json=prePrepare" json:"pre_prepare,omitempty"`
	Prepare              []*PbftEnv `protobuf:"bytes,2,rep,name=prepare" json:"prepare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PbftPreparedSet) Reset()         { *m = PbftPreparedSet{} }
func (m *PbftPreparedSet) String() string { return proto.CompactTextString(m) }
func (*PbftPreparedSet) ProtoMessage()    {}
func (*PbftPreparedSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{3}
}
func (m *PbftPreparedSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftPreparedSet.Unmarshal(m, b)
}
func (m *PbftPreparedSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftPreparedSet.Marshal(b, m, deterministic)
}
func (dst *PbftPreparedSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftPreparedSet.Merge(dst, src)
}
func (m *PbftPreparedSet) XXX_Size() int {
	return xxx_messageInfo_PbftPreparedSet.Size(m)
}
func (m *PbftPreparedSet) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftPreparedSet.DiscardUnknown(m)
}

var xxx_messageInfo_PbftPreparedSet proto.InternalMessageInfo

func (m *PbftPreparedSet) GetPrePrepare() *PbftEnv {
	if m != nil {
		return m.PrePrepare
	}
	return nil
}

func (m *PbftPreparedSet) GetPrepare() []*PbftEnv {
	if m != nil {
		return m.Prepare
	}
	return nil
}

type PbftViewChange struct {
	ViewNumber           int64    `protobuf:"varint,1,opt,name=view_number,json=viewNumber" json:"view_number,omitempty"`
	Sequence             int64    `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	PrepredValueDigest   []byte   `protobuf:"bytes,3,opt,name=prepred_value_digest,json=prepredValueDigest,proto3" json:"prepred_value_digest,omitempty"`
	ReplicaId            int64    `protobuf:"varint,4,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbftViewChange) Reset()         { *m = PbftViewChange{} }
func (m *PbftViewChange) String() string { return proto.CompactTextString(m) }
func (*PbftViewChange) ProtoMessage()    {}
func (*PbftViewChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{4}
}
func (m *PbftViewChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftViewChange.Unmarshal(m, b)
}
func (m *PbftViewChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftViewChange.Marshal(b, m, deterministic)
}
func (dst *PbftViewChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftViewChange.Merge(dst, src)
}
func (m *PbftViewChange) XXX_Size() int {
	return xxx_messageInfo_PbftViewChange.Size(m)
}
func (m *PbftViewChange) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftViewChange.DiscardUnknown(m)
}

var xxx_messageInfo_PbftViewChange proto.InternalMessageInfo

func (m *PbftViewChange) GetViewNumber() int64 {
	if m != nil {
		return m.ViewNumber
	}
	return 0
}

func (m *PbftViewChange) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *PbftViewChange) GetPrepredValueDigest() []byte {
	if m != nil {
		return m.PrepredValueDigest
	}
	return nil
}

func (m *PbftViewChange) GetReplicaId() int64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

type PbftViewChangeWithRawValue struct {
	ViewChangeEnv        *PbftEnv         `protobuf:"bytes,1,opt,name=view_change_env,json=viewChangeEnv" json:"view_change_env,omitempty"`
	PreparedSet          *PbftPreparedSet `protobuf:"bytes,2,opt,name=prepared_set,json=preparedSet" json:"prepared_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PbftViewChangeWithRawValue) Reset()         { *m = PbftViewChangeWithRawValue{} }
func (m *PbftViewChangeWithRawValue) String() string { return proto.CompactTextString(m) }
func (*PbftViewChangeWithRawValue) ProtoMessage()    {}
func (*PbftViewChangeWithRawValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{5}
}
func (m *PbftViewChangeWithRawValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftViewChangeWithRawValue.Unmarshal(m, b)
}
func (m *PbftViewChangeWithRawValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftViewChangeWithRawValue.Marshal(b, m, deterministic)
}
func (dst *PbftViewChangeWithRawValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftViewChangeWithRawValue.Merge(dst, src)
}
func (m *PbftViewChangeWithRawValue) XXX_Size() int {
	return xxx_messageInfo_PbftViewChangeWithRawValue.Size(m)
}
func (m *PbftViewChangeWithRawValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftViewChangeWithRawValue.DiscardUnknown(m)
}

var xxx_messageInfo_PbftViewChangeWithRawValue proto.InternalMessageInfo

func (m *PbftViewChangeWithRawValue) GetViewChangeEnv() *PbftEnv {
	if m != nil {
		return m.ViewChangeEnv
	}
	return nil
}

func (m *PbftViewChangeWithRawValue) GetPreparedSet() *PbftPreparedSet {
	if m != nil {
		return m.PreparedSet
	}
	return nil
}

type PbftNewView struct {
	ViewNumber           int64      `protobuf:"varint,1,opt,name=view_number,json=viewNumber" json:"view_number,omitempty"`
	Sequence             int64      `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	ReplicaId            int64      `protobuf:"varint,3,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	ViewChanges          []*PbftEnv `protobuf:"bytes,4,rep,name=view_changes,json=viewChanges" json:"view_changes,omitempty"`
	PrePrepare           *PbftEnv   `protobuf:"bytes,5,opt,name=pre_prepare,json=prePrepare" json:"pre_prepare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PbftNewView) Reset()         { *m = PbftNewView{} }
func (m *PbftNewView) String() string { return proto.CompactTextString(m) }
func (*PbftNewView) ProtoMessage()    {}
func (*PbftNewView) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{6}
}
func (m *PbftNewView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftNewView.Unmarshal(m, b)
}
func (m *PbftNewView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftNewView.Marshal(b, m, deterministic)
}
func (dst *PbftNewView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftNewView.Merge(dst, src)
}
func (m *PbftNewView) XXX_Size() int {
	return xxx_messageInfo_PbftNewView.Size(m)
}
func (m *PbftNewView) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftNewView.DiscardUnknown(m)
}

var xxx_messageInfo_PbftNewView proto.InternalMessageInfo

func (m *PbftNewView) GetViewNumber() int64 {
	if m != nil {
		return m.ViewNumber
	}
	return 0
}

func (m *PbftNewView) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *PbftNewView) GetReplicaId() int64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *PbftNewView) GetViewChanges() []*PbftEnv {
	if m != nil {
		return m.ViewChanges
	}
	return nil
}

func (m *PbftNewView) GetPrePrepare() *PbftEnv {
	if m != nil {
		return m.PrePrepare
	}
	return nil
}

type Pbft struct {
	RoundNumber            int64                       `protobuf:"varint,1,opt,name=round_number,json=roundNumber" json:"round_number,omitempty"`
	Type                   PbftMessageType             `protobuf:"varint,2,opt,name=type,enum=protocol.PbftMessageType" json:"type,omitempty"`
	PrePrepare             *PbftPrePrepare             `protobuf:"bytes,3,opt,name=pre_prepare,json=prePrepare" json:"pre_prepare,omitempty"`
	Prepare                *PbftPrepare                `protobuf:"bytes,4,opt,name=prepare" json:"prepare,omitempty"`
	Commit                 *PbftCommit                 `protobuf:"bytes,5,opt,name=commit" json:"commit,omitempty"`
	ViewChange             *PbftViewChange             `protobuf:"bytes,6,opt,name=view_change,json=viewChange" json:"view_change,omitempty"`
	NewView                *PbftNewView                `protobuf:"bytes,7,opt,name=new_view,json=newView" json:"new_view,omitempty"`
	ViewChangeWithRawvalue *PbftViewChangeWithRawValue `protobuf:"bytes,8,opt,name=view_change_with_rawvalue,json=viewChangeWithRawvalue" json:"view_change_with_rawvalue,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                    `json:"-"`
	XXX_unrecognized       []byte                      `json:"-"`
	XXX_sizecache          int32                       `json:"-"`
}

func (m *Pbft) Reset()         { *m = Pbft{} }
func (m *Pbft) String() string { return proto.CompactTextString(m) }
func (*Pbft) ProtoMessage()    {}
func (*Pbft) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{7}
}
func (m *Pbft) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pbft.Unmarshal(m, b)
}
func (m *Pbft) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pbft.Marshal(b, m, deterministic)
}
func (dst *Pbft) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pbft.Merge(dst, src)
}
func (m *Pbft) XXX_Size() int {
	return xxx_messageInfo_Pbft.Size(m)
}
func (m *Pbft) XXX_DiscardUnknown() {
	xxx_messageInfo_Pbft.DiscardUnknown(m)
}

var xxx_messageInfo_Pbft proto.InternalMessageInfo

func (m *Pbft) GetRoundNumber() int64 {
	if m != nil {
		return m.RoundNumber
	}
	return 0
}

func (m *Pbft) GetType() PbftMessageType {
	if m != nil {
		return m.Type
	}
	return PbftMessageType_PBFT_TYPE_PREPREPARE
}

func (m *Pbft) GetPrePrepare() *PbftPrePrepare {
	if m != nil {
		return m.PrePrepare
	}
	return nil
}

func (m *Pbft) GetPrepare() *PbftPrepare {
	if m != nil {
		return m.Prepare
	}
	return nil
}

func (m *Pbft) GetCommit() *PbftCommit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *Pbft) GetViewChange() *PbftViewChange {
	if m != nil {
		return m.ViewChange
	}
	return nil
}

func (m *Pbft) GetNewView() *PbftNewView {
	if m != nil {
		return m.NewView
	}
	return nil
}

func (m *Pbft) GetViewChangeWithRawvalue() *PbftViewChangeWithRawValue {
	if m != nil {
		return m.ViewChangeWithRawvalue
	}
	return nil
}

type PbftEnv struct {
	Pbft                 *Pbft      `protobuf:"bytes,1,opt,name=pbft" json:"pbft,omitempty"`
	Signature            *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PbftEnv) Reset()         { *m = PbftEnv{} }
func (m *PbftEnv) String() string { return proto.CompactTextString(m) }
func (*PbftEnv) ProtoMessage()    {}
func (*PbftEnv) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{8}
}
func (m *PbftEnv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftEnv.Unmarshal(m, b)
}
func (m *PbftEnv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftEnv.Marshal(b, m, deterministic)
}
func (dst *PbftEnv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftEnv.Merge(dst, src)
}
func (m *PbftEnv) XXX_Size() int {
	return xxx_messageInfo_PbftEnv.Size(m)
}
func (m *PbftEnv) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftEnv.DiscardUnknown(m)
}

var xxx_messageInfo_PbftEnv proto.InternalMessageInfo

func (m *PbftEnv) GetPbft() *Pbft {
	if m != nil {
		return m.Pbft
	}
	return nil
}

func (m *PbftEnv) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Validator struct {
	Address              string   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	PledgeCoinAmount     int64    `protobuf:"varint,2,opt,name=pledge_coin_amount,json=pledgeCoinAmount" json:"pledge_coin_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{9}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validator.Unmarshal(m, b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
}
func (dst *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(dst, src)
}
func (m *Validator) XXX_Size() int {
	return xxx_messageInfo_Validator.Size(m)
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Validator) GetPledgeCoinAmount() int64 {
	if m != nil {
		return m.PledgeCoinAmount
	}
	return 0
}

type ValidatorSet struct {
	Validators           []*Validator `protobuf:"bytes,1,rep,name=validators" json:"validators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ValidatorSet) Reset()         { *m = ValidatorSet{} }
func (m *ValidatorSet) String() string { return proto.CompactTextString(m) }
func (*ValidatorSet) ProtoMessage()    {}
func (*ValidatorSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{10}
}
func (m *ValidatorSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorSet.Unmarshal(m, b)
}
func (m *ValidatorSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorSet.Marshal(b, m, deterministic)
}
func (dst *ValidatorSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSet.Merge(dst, src)
}
func (m *ValidatorSet) XXX_Size() int {
	return xxx_messageInfo_ValidatorSet.Size(m)
}
func (m *ValidatorSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSet.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSet proto.InternalMessageInfo

func (m *ValidatorSet) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

type PbftProof struct {
	Commits              []*PbftEnv `protobuf:"bytes,1,rep,name=commits" json:"commits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PbftProof) Reset()         { *m = PbftProof{} }
func (m *PbftProof) String() string { return proto.CompactTextString(m) }
func (*PbftProof) ProtoMessage()    {}
func (*PbftProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{11}
}
func (m *PbftProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftProof.Unmarshal(m, b)
}
func (m *PbftProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftProof.Marshal(b, m, deterministic)
}
func (dst *PbftProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftProof.Merge(dst, src)
}
func (m *PbftProof) XXX_Size() int {
	return xxx_messageInfo_PbftProof.Size(m)
}
func (m *PbftProof) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftProof.DiscardUnknown(m)
}

var xxx_messageInfo_PbftProof proto.InternalMessageInfo

func (m *PbftProof) GetCommits() []*PbftEnv {
	if m != nil {
		return m.Commits
	}
	return nil
}

type FeeConfig struct {
	GasPrice             int64    `protobuf:"varint,1,opt,name=gas_price,json=gasPrice" json:"gas_price,omitempty"`
	BaseReserve          int64    `protobuf:"varint,2,opt,name=base_reserve,json=baseReserve" json:"base_reserve,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeeConfig) Reset()         { *m = FeeConfig{} }
func (m *FeeConfig) String() string { return proto.CompactTextString(m) }
func (*FeeConfig) ProtoMessage()    {}
func (*FeeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_711869b4bd553306, []int{12}
}
func (m *FeeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeeConfig.Unmarshal(m, b)
}
func (m *FeeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeeConfig.Marshal(b, m, deterministic)
}
func (dst *FeeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeConfig.Merge(dst, src)
}
func (m *FeeConfig) XXX_Size() int {
	return xxx_messageInfo_FeeConfig.Size(m)
}
func (m *FeeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FeeConfig proto.InternalMessageInfo

func (m *FeeConfig) GetGasPrice() int64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *FeeConfig) GetBaseReserve() int64 {
	if m != nil {
		return m.BaseReserve
	}
	return 0
}

func init() {
	proto.RegisterType((*PbftPrePrepare)(nil), "protocol.PbftPrePrepare")
	proto.RegisterType((*PbftPrepare)(nil), "protocol.PbftPrepare")
	proto.RegisterType((*PbftCommit)(nil), "protocol.PbftCommit")
	proto.RegisterType((*PbftPreparedSet)(nil), "protocol.PbftPreparedSet")
	proto.RegisterType((*PbftViewChange)(nil), "protocol.PbftViewChange")
	proto.RegisterType((*PbftViewChangeWithRawValue)(nil), "protocol.PbftViewChangeWithRawValue")
	proto.RegisterType((*PbftNewView)(nil), "protocol.PbftNewView")
	proto.RegisterType((*Pbft)(nil), "protocol.Pbft")
	proto.RegisterType((*PbftEnv)(nil), "protocol.PbftEnv")
	proto.RegisterType((*Validator)(nil), "protocol.Validator")
	proto.RegisterType((*ValidatorSet)(nil), "protocol.ValidatorSet")
	proto.RegisterType((*PbftProof)(nil), "protocol.PbftProof")
	proto.RegisterType((*FeeConfig)(nil), "protocol.FeeConfig")
	proto.RegisterEnum("protocol.PbftMessageType", PbftMessageType_name, PbftMessageType_value)
	proto.RegisterEnum("protocol.PbftValueType", PbftValueType_name, PbftValueType_value)
	proto.RegisterEnum("protocol.FeeConfig_Type", FeeConfig_Type_name, FeeConfig_Type_value)
}

func init() { proto.RegisterFile("consensus.proto", fileDescriptor_consensus_711869b4bd553306) }

var fileDescriptor_consensus_711869b4bd553306 = []byte{
	// 921 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xce, 0xc6, 0x9b, 0x38, 0x7e, 0xeb, 0x24, 0x9b, 0xc1, 0x45, 0xdb, 0x20, 0x44, 0xba, 0x02,
	0xa9, 0x2a, 0xc5, 0x94, 0xb4, 0x87, 0x56, 0xe2, 0xe2, 0x98, 0x6d, 0x6b, 0x41, 0x5c, 0x6b, 0xec,
	0xda, 0x70, 0x1a, 0xd6, 0xbb, 0xcf, 0xce, 0x0a, 0x7b, 0x77, 0x99, 0x59, 0xdb, 0xf4, 0xc8, 0x19,
	0xae, 0x5c, 0xb9, 0x71, 0xe7, 0x4f, 0xe0, 0x7f, 0xe0, 0x1f, 0x42, 0x33, 0x3b, 0xfe, 0xb1, 0xae,
	0x23, 0x21, 0xa1, 0x0a, 0x29, 0x52, 0x3c, 0xdf, 0xfb, 0xe6, 0xf9, 0x7b, 0xef, 0x7d, 0x6f, 0x0c,
	0xa7, 0x41, 0x12, 0x0b, 0x8c, 0xc5, 0x4c, 0xd4, 0x53, 0x9e, 0x64, 0x09, 0x39, 0x52, 0xff, 0x82,
	0x64, 0x72, 0x5e, 0x0d, 0x92, 0xe9, 0x34, 0x89, 0x73, 0xdc, 0xfd, 0xc3, 0x80, 0x93, 0xce, 0x70,
	0x94, 0x75, 0x38, 0x76, 0x38, 0xa6, 0x3e, 0x47, 0xf2, 0x11, 0x58, 0xf3, 0x08, 0x17, 0x2c, 0x9e,
	0x4d, 0x87, 0xc8, 0x1d, 0xe3, 0xc2, 0xb8, 0x5f, 0xa2, 0x20, 0xa1, 0xb6, 0x42, 0xc8, 0x39, 0x1c,
	0x09, 0xfc, 0x71, 0x86, 0x71, 0x80, 0xce, 0xbe, 0x8a, 0xae, 0xce, 0xe4, 0x43, 0x00, 0x8e, 0xe9,
	0x24, 0x0a, 0x7c, 0x16, 0x85, 0x4e, 0x49, 0x45, 0x2b, 0x1a, 0x69, 0x85, 0xa4, 0x06, 0x07, 0x73,
	0x7f, 0x32, 0x43, 0xc7, 0xbc, 0x30, 0xee, 0x57, 0x69, 0x7e, 0x20, 0xf7, 0xa0, 0xaa, 0x3e, 0xb0,
	0x30, 0x1a, 0xa3, 0xc8, 0x9c, 0x03, 0x15, 0xb4, 0x14, 0xf6, 0x95, 0x82, 0xdc, 0x5f, 0x0d, 0xb0,
	0xb4, 0xce, 0x77, 0x2e, 0x72, 0x5b, 0x8e, 0xf9, 0xb6, 0x9c, 0x5f, 0x0c, 0x00, 0x29, 0xa7, 0x99,
	0x4c, 0xa7, 0x51, 0xf6, 0x7f, 0xab, 0xe1, 0x70, 0xba, 0xd1, 0x9b, 0xb0, 0x8b, 0x19, 0xb9, 0x04,
	0x2b, 0xe5, 0xc8, 0xd2, 0x1c, 0x52, 0x8a, 0xac, 0xcb, 0xb3, 0xfa, 0xd2, 0x05, 0x75, 0xc9, 0xf7,
	0xe2, 0x39, 0x85, 0x74, 0x3d, 0xf8, 0x4f, 0xa1, 0xbc, 0xe4, 0xef, 0x5f, 0x94, 0x76, 0xf3, 0x97,
	0x0c, 0xf7, 0x77, 0x6d, 0x9c, 0x7e, 0x84, 0x8b, 0xe6, 0x8d, 0x1f, 0x8f, 0xff, 0xe3, 0x4c, 0x1e,
	0x41, 0x4d, 0xa6, 0xe6, 0x18, 0xb2, 0x42, 0xb9, 0x25, 0x55, 0x2e, 0xd1, 0xb1, 0xfe, 0xba, 0xea,
	0xad, 0xbe, 0x99, 0x5b, 0x7d, 0x73, 0x7f, 0x33, 0xe0, 0xbc, 0x28, 0x70, 0x10, 0x65, 0x37, 0xd4,
	0x5f, 0xa8, 0x1c, 0xe4, 0x19, 0x9c, 0x2a, 0xb1, 0x81, 0x0a, 0x31, 0x8c, 0xe7, 0xb7, 0x37, 0xe9,
	0x78, 0xbe, 0xca, 0xe1, 0xc5, 0x73, 0xf2, 0x25, 0x54, 0x75, 0x17, 0x42, 0x26, 0x30, 0x53, 0xa5,
	0x58, 0x97, 0x77, 0x8b, 0xf7, 0x36, 0x86, 0x41, 0xad, 0x74, 0x7d, 0x70, 0xff, 0xd6, 0x4e, 0x6e,
	0xe3, 0x42, 0x4a, 0x7b, 0xa7, 0xde, 0x79, 0x02, 0xd5, 0x8d, 0x22, 0x85, 0x63, 0xde, 0x36, 0x56,
	0x6b, 0x5d, 0xa1, 0xd8, 0xf6, 0xce, 0xc1, 0xbf, 0xf0, 0x8e, 0xfb, 0x57, 0x09, 0x4c, 0x89, 0x4b,
	0xbb, 0xf2, 0x64, 0x16, 0x87, 0xc5, 0x7a, 0x2c, 0x85, 0xe9, 0x82, 0x3e, 0x03, 0x33, 0x7b, 0x93,
	0xe6, 0xc5, 0x9c, 0x6c, 0xf7, 0xed, 0x1a, 0x85, 0xf0, 0xc7, 0xd8, 0x7b, 0x93, 0x22, 0x55, 0x34,
	0xf2, 0xac, 0x28, 0xa7, 0xa4, 0xe4, 0x38, 0x6f, 0x75, 0x5b, 0x2b, 0x29, 0x38, 0xfa, 0xf3, 0xb5,
	0xa3, 0x4d, 0x75, 0xed, 0xce, 0xce, 0x21, 0xad, 0x5c, 0x4d, 0x1e, 0xc2, 0x61, 0xa0, 0x56, 0x5a,
	0x57, 0x5d, 0x2b, 0xf2, 0xf3, 0x75, 0xa7, 0x9a, 0x23, 0x95, 0x6d, 0xb4, 0xd7, 0x39, 0xdc, 0xa5,
	0x6c, 0x6d, 0xbf, 0x7c, 0xa8, 0x7a, 0x57, 0x1e, 0xc1, 0x51, 0x8c, 0x0b, 0x26, 0x11, 0xa7, 0xbc,
	0x4b, 0x9a, 0xb6, 0x07, 0x2d, 0xc7, 0xda, 0x27, 0x0c, 0xee, 0x6e, 0x1a, 0x76, 0x11, 0x65, 0x37,
	0x8c, 0xfb, 0x8b, 0xfc, 0x39, 0x3d, 0x52, 0x29, 0x3e, 0xbe, 0xed, 0xab, 0x37, 0x9d, 0x4f, 0xdf,
	0x9f, 0x6f, 0xe3, 0x2a, 0x87, 0xfb, 0x3d, 0x94, 0xf5, 0x64, 0x89, 0x0b, 0x66, 0x3a, 0x1c, 0x65,
	0x7a, 0x23, 0x4e, 0x8a, 0x69, 0xa9, 0x8a, 0x91, 0x2f, 0xa0, 0x22, 0xa2, 0x71, 0xec, 0x67, 0x33,
	0x8e, 0x7a, 0x05, 0xde, 0x5b, 0x13, 0xbb, 0xcb, 0x10, 0x5d, 0xb3, 0xdc, 0x2e, 0x54, 0xfa, 0xfe,
	0x24, 0x0a, 0xfd, 0x2c, 0xe1, 0xc4, 0x81, 0xb2, 0x1f, 0x86, 0x1c, 0x85, 0x50, 0x5f, 0x53, 0xa1,
	0xcb, 0x23, 0x79, 0x08, 0x24, 0x9d, 0x60, 0x38, 0x46, 0x16, 0x24, 0x51, 0xcc, 0xfc, 0x69, 0x32,
	0x8b, 0x33, 0x6d, 0x7d, 0x3b, 0x8f, 0x34, 0x93, 0x28, 0x6e, 0x28, 0xdc, 0x6d, 0x42, 0x75, 0x95,
	0x54, 0xbe, 0x7c, 0x8f, 0x01, 0xe6, 0xcb, 0xb3, 0x4c, 0x5d, 0x2a, 0x0a, 0x5b, 0x71, 0xe9, 0x06,
	0xcd, 0x7d, 0x0a, 0x95, 0xdc, 0x0f, 0x49, 0x32, 0x92, 0xef, 0x60, 0x3e, 0xe0, 0xe5, 0xf5, 0x5d,
	0xef, 0xa0, 0x66, 0xb8, 0x3f, 0x1b, 0x50, 0x79, 0x8e, 0xd8, 0x4c, 0xe2, 0x51, 0x34, 0x26, 0x1f,
	0x40, 0x65, 0xec, 0x0b, 0x96, 0xf2, 0x28, 0x40, 0x6d, 0xfd, 0xa3, 0xb1, 0x2f, 0x3a, 0xf2, 0x2c,
	0x57, 0x63, 0xe8, 0x0b, 0x64, 0x1c, 0x05, 0xf2, 0xf9, 0x72, 0x99, 0x2d, 0x89, 0xd1, 0x1c, 0x72,
	0x9f, 0x80, 0x29, 0x9d, 0x4f, 0x2c, 0x28, 0xbf, 0x6e, 0x7f, 0xdd, 0x7e, 0x35, 0x68, 0xdb, 0x7b,
	0xe4, 0x18, 0x2a, 0x2f, 0x1a, 0x5d, 0xd6, 0xa1, 0xad, 0xa6, 0x67, 0x1b, 0xc4, 0x86, 0xea, 0x55,
	0xa3, 0xeb, 0x31, 0xea, 0x75, 0x3d, 0xda, 0xf7, 0xec, 0xfd, 0x07, 0x7f, 0x1a, 0xf9, 0x0f, 0xc0,
	0xc6, 0xee, 0x10, 0x07, 0x6a, 0x9d, 0xab, 0xe7, 0x3d, 0xd6, 0xfb, 0xae, 0xe3, 0xb1, 0x0e, 0xf5,
	0xe4, 0x5f, 0x83, 0x7a, 0xf6, 0x1e, 0xb9, 0x03, 0x67, 0xc5, 0x88, 0x84, 0x0d, 0x52, 0x03, 0x7b,
	0x0d, 0x37, 0x5f, 0x5d, 0x5f, 0xb7, 0x7a, 0xf6, 0x7e, 0x31, 0x4d, 0xbf, 0xe5, 0x0d, 0x9a, 0x2f,
	0x1b, 0xed, 0x17, 0x9e, 0x5d, 0x2a, 0xa6, 0x69, 0x7b, 0x03, 0x19, 0xb4, 0x4d, 0xf2, 0x09, 0xdc,
	0xdb, 0x71, 0x81, 0x0d, 0x5a, 0xbd, 0x97, 0x8c, 0x36, 0x06, 0xfd, 0xc6, 0x37, 0xaf, 0x3d, 0xfb,
	0xe0, 0xc1, 0x53, 0x38, 0x56, 0x16, 0x95, 0xce, 0x53, 0x7a, 0xcf, 0xe0, 0x58, 0xdd, 0x53, 0x04,
	0xd6, 0xfb, 0xd6, 0xde, 0x5b, 0x29, 0x5a, 0x42, 0x5d, 0xaf, 0x67, 0x1b, 0x57, 0x2e, 0x5c, 0x44,
	0x49, 0x7d, 0x38, 0x9b, 0x26, 0x75, 0x11, 0xfe, 0x50, 0x0f, 0x12, 0x8e, 0x75, 0xfc, 0x29, 0xc3,
	0x38, 0xcc, 0xa7, 0x34, 0x9c, 0x8d, 0x86, 0x87, 0xea, 0xd3, 0xe3, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x69, 0x47, 0x83, 0x71, 0x07, 0x09, 0x00, 0x00,
}
