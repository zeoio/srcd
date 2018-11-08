// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tx.proto

package transaction

import (
	"encoding/binary"
	"io"
	math "math"
	"fmt"

	"github.com/srchain/srcd/core/transaction/extend"
	proto "github.com/golang/protobuf/proto"
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

type Hash struct {
	V0 uint64 `protobuf:"fixed64,1,opt,name=v0" json:"v0,omitempty"`
	V1 uint64 `protobuf:"fixed64,2,opt,name=v1" json:"v1,omitempty"`
	V2 uint64 `protobuf:"fixed64,3,opt,name=v2" json:"v2,omitempty"`
	V3 uint64 `protobuf:"fixed64,4,opt,name=v3" json:"v3,omitempty"`
}

// NewHash convert the input byte array to hash
func NewHash(b32 [32]byte) (h Hash) {
	h.V0 = binary.BigEndian.Uint64(b32[0:8])
	h.V1 = binary.BigEndian.Uint64(b32[8:16])
	h.V2 = binary.BigEndian.Uint64(b32[16:24])
	h.V3 = binary.BigEndian.Uint64(b32[24:32])
	return h
}

// Bytes returns the byte representation
func (h Hash) Bytes() []byte {
	b32 := h.Byte32()
	return b32[:]
}
func (h *Hash) ReadFrom(r io.Reader) (int64, error) {
	var b32 [32]byte
	n, err := io.ReadFull(r, b32[:])
	if err != nil {
		return int64(n), err
	}
	*h = NewHash(b32)
	return int64(n), nil
}

// Byte32 return the byte array representation
func (h Hash) Byte32() (b32 [32]byte) {
	binary.BigEndian.PutUint64(b32[0:8], h.V0)
	binary.BigEndian.PutUint64(b32[8:16], h.V1)
	binary.BigEndian.PutUint64(b32[16:24], h.V2)
	binary.BigEndian.PutUint64(b32[24:32], h.V3)
	return b32
}
func (m *Hash) Reset()         { *m = Hash{} }
func (m *Hash) String() string { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()    {}
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_dcb76708e2c2a44f, []int{0}
}

// WriteTo satisfies the io.WriterTo interface.
func (h Hash) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(h.Bytes())
	return int64(n), err
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

func (m *Hash) GetV0() uint64 {
	if m != nil {
		return m.V0
	}
	return 0
}

func (m *Hash) GetV1() uint64 {
	if m != nil {
		return m.V1
	}
	return 0
}

func (m *Hash) GetV2() uint64 {
	if m != nil {
		return m.V2
	}
	return 0
}

func (m *Hash) GetV3() uint64 {
	if m != nil {
		return m.V3
	}
	return 0
}

type Program struct {
	VmVersion uint64 `protobuf:"varint,1,opt,name=vm_version,json=vmVersion" json:"vm_version,omitempty"`
	Code      []byte `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *Program) Reset()         { *m = Program{} }
func (m *Program) String() string { return proto.CompactTextString(m) }
func (*Program) ProtoMessage()    {}
func (*Program) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_dcb76708e2c2a44f, []int{1}
}

var xxx_messageInfo_Program proto.InternalMessageInfo

func (m *Program) GetVmVersion() uint64 {
	if m != nil {
		return m.VmVersion
	}
	return 0
}

func (m *Program) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

// This message type duplicates Hash, above. One alternative is to
// embed a Hash inside an AssetID. But it's useful for AssetID to be
// plain old data (without pointers). Another alternative is use Hash
// in any protobuf types where an AssetID is called for, but it's
// preferable to have type safety.
type AssetID struct {
	V0 uint64 `protobuf:"fixed64,1,opt,name=v0" json:"v0,omitempty"`
	V1 uint64 `protobuf:"fixed64,2,opt,name=v1" json:"v1,omitempty"`
	V2 uint64 `protobuf:"fixed64,3,opt,name=v2" json:"v2,omitempty"`
	V3 uint64 `protobuf:"fixed64,4,opt,name=v3" json:"v3,omitempty"`
}

// ReadFrom satisfies the io.ReaderFrom interface.
func (a *AssetID) ReadFrom(r io.Reader) (int64, error) { return (*Hash)(a).ReadFrom(r) }
func (a AssetID) WriteTo(w io.Writer) (int64, error)   { return Hash(a).WriteTo(w) }
func (m *AssetID) Reset()                              { *m = AssetID{} }
func (m *AssetID) String() string                      { return proto.CompactTextString(m) }
func (*AssetID) ProtoMessage()                         {}
func (*AssetID) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_dcb76708e2c2a44f, []int{2}
}

// Bytes returns the byte representation.
func (a AssetID) Bytes() []byte { return Hash(a).Bytes() }

var xxx_messageInfo_AssetID proto.InternalMessageInfo

func (m *AssetID) GetV0() uint64 {
	if m != nil {
		return m.V0
	}
	return 0
}

func (m *AssetID) GetV1() uint64 {
	if m != nil {
		return m.V1
	}
	return 0
}

func (m *AssetID) GetV2() uint64 {
	if m != nil {
		return m.V2
	}
	return 0
}

func (m *AssetID) GetV3() uint64 {
	if m != nil {
		return m.V3
	}
	return 0
}

type AssetAmount struct {
	AssetId *AssetID `protobuf:"bytes,1,opt,name=asset_id,json=assetId" json:"asset_id,omitempty"`
	Amount  uint64   `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

// ReadFrom read the AssetAmount from the bytes
func (a *AssetAmount) ReadFrom(r *extend.Reader) (err error) {
	var assetID AssetID
	if _, err = assetID.ReadFrom(r); err != nil {
		return err
	}
	a.AssetId = &assetID
	a.Amount, err = extend.ReadVarint63(r)
	return err
}

// WriteTo convert struct to byte and write to io
func (a AssetAmount) WriteTo(w io.Writer) (int64, error) {
	n, err := a.AssetId.WriteTo(w)
	if err != nil {
		return n, err
	}
	n2, err := extend.WriteVarint63(w, a.Amount)
	return n + int64(n2), err
}

func (m *AssetAmount) Reset()         { *m = AssetAmount{} }
func (m *AssetAmount) String() string { return proto.CompactTextString(m) }
func (*AssetAmount) ProtoMessage()    {}
func (*AssetAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_dcb76708e2c2a44f, []int{3}
}

var xxx_messageInfo_AssetAmount proto.InternalMessageInfo

func (m *AssetAmount) GetAssetId() *AssetID {
	if m != nil {
		return m.AssetId
	}
	return nil
}

func (m *AssetAmount) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ValueSource struct {
	Ref      *Hash        `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	Value    *AssetAmount `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Position uint64       `protobuf:"varint,3,opt,name=position" json:"position,omitempty"`
}

func (m *ValueSource) Reset()         { *m = ValueSource{} }
func (m *ValueSource) String() string { return proto.CompactTextString(m) }
func (*ValueSource) ProtoMessage()    {}
func (*ValueSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_dcb76708e2c2a44f, []int{4}
}

var xxx_messageInfo_ValueSource proto.InternalMessageInfo

func (m *ValueSource) GetRef() *Hash {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *ValueSource) GetValue() *AssetAmount {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ValueSource) GetPosition() uint64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type ValueDestination struct {
	Ref      *Hash        `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	Value    *AssetAmount `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Position uint64       `protobuf:"varint,3,opt,name=position" json:"position,omitempty"`
}

func (m *ValueDestination) Reset()         { *m = ValueDestination{} }
func (m *ValueDestination) String() string { return proto.CompactTextString(m) }
func (*ValueDestination) ProtoMessage()    {}
func (*ValueDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_dcb76708e2c2a44f, []int{5}
}

var xxx_messageInfo_ValueDestination proto.InternalMessageInfo

func (m *ValueDestination) GetRef() *Hash {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *ValueDestination) GetValue() *AssetAmount {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ValueDestination) GetPosition() uint64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type TxHeader struct {
	Version        uint64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	SerializedSize uint64  `protobuf:"varint,2,opt,name=serialized_size,json=serializedSize" json:"serialized_size,omitempty"`
	TimeRange      uint64  `protobuf:"varint,3,opt,name=time_range,json=timeRange" json:"time_range,omitempty"`
	ResultIds      []*Hash `protobuf:"bytes,4,rep,name=result_ids,json=resultIds" json:"result_ids,omitempty"`
}

// NewTxHeader creates an new TxHeader.
func NewTxHeader(version, serializedSize, timeRange uint64, resultIDs []*Hash) *TxHeader {
	return &TxHeader{
		Version:        version,
		SerializedSize: serializedSize,
		TimeRange:      timeRange,
		ResultIds:      resultIDs,
	}
}
func (TxHeader) typ() string { return "txheader" }
func (h *TxHeader) writeForHash(w io.Writer) {
	mustWriteForHash(w, h.Version)
	mustWriteForHash(w, h.TimeRange)
	mustWriteForHash(w, h.ResultIds)
}
func (m *TxHeader) Reset()         { *m = TxHeader{} }
func (m *TxHeader) String() string { return proto.CompactTextString(m) }
func (*TxHeader) ProtoMessage()    {}
func (*TxHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_dcb76708e2c2a44f, []int{6}
}

var xxx_messageInfo_TxHeader proto.InternalMessageInfo

func (m *TxHeader) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TxHeader) GetSerializedSize() uint64 {
	if m != nil {
		return m.SerializedSize
	}
	return 0
}

func (m *TxHeader) GetTimeRange() uint64 {
	if m != nil {
		return m.TimeRange
	}
	return 0
}

func (m *TxHeader) GetResultIds() []*Hash {
	if m != nil {
		return m.ResultIds
	}
	return nil
}

type Mux struct {
	Sources             []*ValueSource      `protobuf:"bytes,1,rep,name=sources" json:"sources,omitempty"`
	Program             *Program            `protobuf:"bytes,2,opt,name=program" json:"program,omitempty"`
	WitnessDestinations []*ValueDestination `protobuf:"bytes,3,rep,name=witness_destinations,json=witnessDestinations" json:"witness_destinations,omitempty"`
	WitnessArguments    [][]byte            `protobuf:"bytes,4,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
}

// NewMux creates a new Mux.
func NewMux(sources []*ValueSource, program *Program) *Mux {
	return &Mux{
		Sources: sources,
		Program: program,
	}
}
func (Mux) typ() string { return "mux1" }
func (m *Mux) writeForHash(w io.Writer) {
}
func (m *Mux) Reset()         { *m = Mux{} }
func (m *Mux) String() string { return proto.CompactTextString(m) }
func (*Mux) ProtoMessage()    {}
func (*Mux) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_dcb76708e2c2a44f, []int{7}
}

var xxx_messageInfo_Mux proto.InternalMessageInfo

func (m *Mux) GetSources() []*ValueSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Mux) GetProgram() *Program {
	if m != nil {
		return m.Program
	}
	return nil
}

func (m *Mux) GetWitnessDestinations() []*ValueDestination {
	if m != nil {
		return m.WitnessDestinations
	}
	return nil
}

func (m *Mux) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

type Output struct {
	Source         *ValueSource `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	ControlProgram *Program     `protobuf:"bytes,2,opt,name=control_program,json=controlProgram" json:"control_program,omitempty"`
	Ordinal        uint64       `protobuf:"varint,3,opt,name=ordinal" json:"ordinal,omitempty"`
}

// NewOutput creates a new Output.
func NewOutput(source *ValueSource, controlProgram *Program, ordinal uint64) *Output {
	return &Output{
		Source:         source,
		ControlProgram: controlProgram,
		Ordinal:        ordinal,
	}
}
func (Output) typ() string { return "output1" }
func (o *Output) writeForHash(w io.Writer) {
}
func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_dcb76708e2c2a44f, []int{8}
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetSource() *ValueSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Output) GetControlProgram() *Program {
	if m != nil {
		return m.ControlProgram
	}
	return nil
}

func (m *Output) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type Spend struct {
	SpentOutputId      *Hash             `protobuf:"bytes,1,opt,name=spent_output_id,json=spentOutputId" json:"spent_output_id,omitempty"`
	WitnessDestination *ValueDestination `protobuf:"bytes,2,opt,name=witness_destination,json=witnessDestination" json:"witness_destination,omitempty"`
	WitnessArguments   [][]byte          `protobuf:"bytes,3,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
	Ordinal            uint64            `protobuf:"varint,4,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (Spend) typ() string { return "spend1" }
func (s *Spend) writeForHash(w io.Writer) {
}

// SetDestination will link the spend to the output
func (s *Spend) SetDestination(id *Hash, val *AssetAmount, pos uint64) {
	s.WitnessDestination = &ValueDestination{
		Ref:      id,
		Value:    val,
		Position: pos,
	}
}

// NewSpend creates a new Spend.
func NewSpend(spentOutputID *Hash, ordinal uint64) *Spend {
	return &Spend{
		SpentOutputId: spentOutputID,
		Ordinal:       ordinal,
	}
}
func (m *Spend) Reset()         { *m = Spend{} }
func (m *Spend) String() string { return proto.CompactTextString(m) }
func (*Spend) ProtoMessage()    {}
func (*Spend) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_dcb76708e2c2a44f, []int{9}
}

var xxx_messageInfo_Spend proto.InternalMessageInfo

func (m *Spend) GetSpentOutputId() *Hash {
	if m != nil {
		return m.SpentOutputId
	}
	return nil
}

func (m *Spend) GetWitnessDestination() *ValueDestination {
	if m != nil {
		return m.WitnessDestination
	}
	return nil
}

func (m *Spend) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

func (m *Spend) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func init() {
	proto.RegisterType((*Hash)(nil), "transaction.Hash")
	proto.RegisterType((*Program)(nil), "transaction.Program")
	proto.RegisterType((*AssetID)(nil), "transaction.AssetID")
	proto.RegisterType((*AssetAmount)(nil), "transaction.AssetAmount")
	proto.RegisterType((*ValueSource)(nil), "transaction.ValueSource")
	proto.RegisterType((*ValueDestination)(nil), "transaction.ValueDestination")
	proto.RegisterType((*TxHeader)(nil), "transaction.TxHeader")
	proto.RegisterType((*Mux)(nil), "transaction.Mux")
	proto.RegisterType((*Output)(nil), "transaction.Output")
	proto.RegisterType((*Spend)(nil), "transaction.Spend")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor_tx_dcb76708e2c2a44f) }

var fileDescriptor_tx_dcb76708e2c2a44f = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4f, 0x8b, 0xd3, 0x4e,
	0x18, 0x26, 0x9b, 0x6c, 0xd3, 0x7d, 0xbb, 0xbf, 0x76, 0xf7, 0xdd, 0xe5, 0x47, 0x10, 0x16, 0x96,
	0x78, 0x50, 0x10, 0x6a, 0xb7, 0x3d, 0x09, 0x7a, 0x28, 0x2c, 0xb2, 0x3d, 0xa8, 0xcb, 0x54, 0x7a,
	0x0d, 0x63, 0x33, 0xd6, 0x81, 0x66, 0x26, 0xcc, 0x4c, 0x62, 0xe9, 0xc1, 0x8b, 0x1f, 0xc1, 0xbb,
	0xdf, 0xcb, 0x6f, 0xe0, 0xc7, 0x90, 0x4c, 0x26, 0xb6, 0xa5, 0x15, 0x05, 0xc1, 0x5b, 0x9f, 0x77,
	0xde, 0x7f, 0xcf, 0xf3, 0xbc, 0x29, 0xb4, 0xcd, 0xaa, 0x9f, 0x2b, 0x69, 0x24, 0x76, 0x8c, 0xa2,
	0x42, 0xd3, 0xb9, 0xe1, 0x52, 0xc4, 0x2f, 0x21, 0xb8, 0xa3, 0xfa, 0x03, 0x76, 0xe1, 0xa8, 0x1c,
	0x44, 0xde, 0xb5, 0xf7, 0xb8, 0x45, 0x8e, 0xca, 0x81, 0xc5, 0x37, 0xd1, 0x91, 0xc3, 0x37, 0x16,
	0x0f, 0x23, 0xdf, 0xe1, 0xa1, 0xc5, 0xa3, 0x28, 0x70, 0x78, 0x14, 0x3f, 0x87, 0xf0, 0x5e, 0xc9,
	0x85, 0xa2, 0x19, 0x5e, 0x01, 0x94, 0x59, 0x52, 0x32, 0xa5, 0xb9, 0x14, 0xb6, 0x65, 0x40, 0x4e,
	0xca, 0x6c, 0x56, 0x07, 0x10, 0x21, 0x98, 0xcb, 0x94, 0xd9, 0xde, 0xa7, 0xc4, 0xfe, 0x8e, 0x27,
	0x10, 0x8e, 0xb5, 0x66, 0x66, 0x72, 0xfb, 0xd7, 0x8b, 0xcc, 0xa0, 0x63, 0x5b, 0x8d, 0x33, 0x59,
	0x08, 0x83, 0x4f, 0xa1, 0x4d, 0x2b, 0x98, 0xf0, 0xd4, 0x36, 0xed, 0x0c, 0x2f, 0xfb, 0x5b, 0xfc,
	0xfb, 0x6e, 0x2c, 0x09, 0x6d, 0xd6, 0x24, 0xc5, 0xff, 0xa1, 0x45, 0x6d, 0xa9, 0x9d, 0x19, 0x10,
	0x87, 0xe2, 0x4f, 0xd0, 0x99, 0xd1, 0x65, 0xc1, 0xa6, 0xb2, 0x50, 0x73, 0x86, 0x0f, 0xc1, 0x57,
	0xec, 0xbd, 0x6b, 0x79, 0xbe, 0xd3, 0xb2, 0xd2, 0x93, 0x54, 0xaf, 0xd8, 0x87, 0xe3, 0xb2, 0xaa,
	0xb1, 0xad, 0x3a, 0xc3, 0x68, 0x7f, 0x72, 0xbd, 0x25, 0xa9, 0xd3, 0xf0, 0x01, 0xb4, 0x73, 0xa9,
	0x79, 0xf5, 0x6c, 0x19, 0x06, 0xe4, 0x27, 0x8e, 0x3f, 0x7b, 0x70, 0x66, 0x17, 0xb8, 0x65, 0xda,
	0x70, 0x41, 0xab, 0xe0, 0xbf, 0xdf, 0xe2, 0xab, 0x07, 0xed, 0xb7, 0xab, 0x3b, 0x46, 0x53, 0xa6,
	0x30, 0x82, 0x70, 0xd7, 0xe5, 0x06, 0xe2, 0x23, 0xe8, 0x69, 0xa6, 0x38, 0x5d, 0xf2, 0x35, 0x4b,
	0x13, 0xcd, 0xd7, 0xcc, 0xa9, 0xd9, 0xdd, 0x84, 0xa7, 0x7c, 0xcd, 0xaa, 0x5b, 0x31, 0x3c, 0x63,
	0x89, 0xa2, 0x62, 0xc1, 0xdc, 0xb4, 0x93, 0x2a, 0x42, 0xaa, 0x00, 0x0e, 0x00, 0x14, 0xd3, 0xc5,
	0xb2, 0xb2, 0x4f, 0x47, 0xc1, 0xb5, 0x7f, 0x98, 0xe6, 0x49, 0x9d, 0x34, 0x49, 0x75, 0xfc, 0xdd,
	0x03, 0xff, 0x55, 0xb1, 0xc2, 0x21, 0x84, 0xda, 0x3a, 0xa5, 0x23, 0xcf, 0x96, 0xed, 0xd2, 0xde,
	0xb2, 0x92, 0x34, 0x89, 0xd8, 0x87, 0x30, 0xaf, 0x6f, 0xd8, 0x49, 0xb5, 0x7b, 0x2a, 0xee, 0xbe,
	0x49, 0x93, 0x84, 0xf7, 0x70, 0xf9, 0x91, 0x1b, 0xc1, 0xb4, 0x4e, 0xd2, 0x8d, 0x29, 0x3a, 0xf2,
	0xed, 0xc0, 0xab, 0xfd, 0x81, 0x5b, 0xd6, 0x91, 0x0b, 0x57, 0xba, 0x15, 0xd3, 0xf8, 0x04, 0xce,
	0x9b, 0x8e, 0x54, 0x2d, 0x8a, 0x8c, 0x09, 0x53, 0xd3, 0x3e, 0x25, 0x67, 0xee, 0x61, 0xdc, 0xc4,
	0xe3, 0x2f, 0x1e, 0xb4, 0xde, 0x14, 0x26, 0x2f, 0x0c, 0x0e, 0xa0, 0x55, 0x93, 0x70, 0xa7, 0xf0,
	0x6b, 0xb2, 0x2e, 0x0f, 0x5f, 0x40, 0x6f, 0x2e, 0x85, 0x51, 0x72, 0x99, 0xfc, 0x09, 0xe7, 0xae,
	0x4b, 0x6e, 0xbe, 0xf1, 0x08, 0x42, 0xa9, 0x52, 0x2e, 0xe8, 0xd2, 0x99, 0xd6, 0xc0, 0xf8, 0x9b,
	0x07, 0xc7, 0xd3, 0x9c, 0x89, 0x14, 0x9f, 0x41, 0x4f, 0xe7, 0x4c, 0x98, 0x44, 0xda, 0x25, 0x37,
	0x5f, 0xe0, 0x01, 0x07, 0xff, 0xb3, 0x99, 0x35, 0x9b, 0x49, 0x8a, 0xaf, 0xe1, 0xe2, 0x80, 0xb2,
	0x6e, 0xc3, 0xdf, 0x08, 0x8b, 0xfb, 0xc2, 0x1e, 0xd6, 0xd5, 0x3f, 0xac, 0xeb, 0x36, 0xb7, 0x60,
	0x87, 0xdb, 0xbb, 0x96, 0xfd, 0x03, 0x1d, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0x46, 0x7d, 0x65,
	0xae, 0x4c, 0x05, 0x00, 0x00,
}
