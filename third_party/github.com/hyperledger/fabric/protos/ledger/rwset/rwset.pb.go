/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ledger/rwset/rwset.proto

/*
Package rwset is a generated protocol buffer package.

It is generated from these files:
	ledger/rwset/rwset.proto

It has these top-level messages:
	TxReadWriteSet
	NsReadWriteSet
	CollectionHashedReadWriteSet
	TxPvtReadWriteSet
	NsPvtReadWriteSet
	CollectionPvtReadWriteSet
*/
package rwset

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

type TxReadWriteSet_DataModel int32

const (
	TxReadWriteSet_KV TxReadWriteSet_DataModel = 0
)

var TxReadWriteSet_DataModel_name = map[int32]string{
	0: "KV",
}
var TxReadWriteSet_DataModel_value = map[string]int32{
	"KV": 0,
}

func (x TxReadWriteSet_DataModel) String() string {
	return proto.EnumName(TxReadWriteSet_DataModel_name, int32(x))
}
func (TxReadWriteSet_DataModel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// TxReadWriteSet encapsulates a read-write set for a transaction
// DataModel specifies the enum value of the data model
// ns_rwset field specifies a list of chaincode specific read-write set (one for each chaincode)
type TxReadWriteSet struct {
	DataModel TxReadWriteSet_DataModel `protobuf:"varint,1,opt,name=data_model,json=dataModel,enum=rwset.TxReadWriteSet_DataModel" json:"data_model,omitempty"`
	NsRwset   []*NsReadWriteSet        `protobuf:"bytes,2,rep,name=ns_rwset,json=nsRwset" json:"ns_rwset,omitempty"`
}

func (m *TxReadWriteSet) Reset()                    { *m = TxReadWriteSet{} }
func (m *TxReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*TxReadWriteSet) ProtoMessage()               {}
func (*TxReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TxReadWriteSet) GetDataModel() TxReadWriteSet_DataModel {
	if m != nil {
		return m.DataModel
	}
	return TxReadWriteSet_KV
}

func (m *TxReadWriteSet) GetNsRwset() []*NsReadWriteSet {
	if m != nil {
		return m.NsRwset
	}
	return nil
}

// NsReadWriteSet encapsulates the read-write set for a chaincode
type NsReadWriteSet struct {
	Namespace             string                          `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Rwset                 []byte                          `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"`
	CollectionHashedRwset []*CollectionHashedReadWriteSet `protobuf:"bytes,3,rep,name=collection_hashed_rwset,json=collectionHashedRwset" json:"collection_hashed_rwset,omitempty"`
}

func (m *NsReadWriteSet) Reset()                    { *m = NsReadWriteSet{} }
func (m *NsReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*NsReadWriteSet) ProtoMessage()               {}
func (*NsReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NsReadWriteSet) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NsReadWriteSet) GetRwset() []byte {
	if m != nil {
		return m.Rwset
	}
	return nil
}

func (m *NsReadWriteSet) GetCollectionHashedRwset() []*CollectionHashedReadWriteSet {
	if m != nil {
		return m.CollectionHashedRwset
	}
	return nil
}

// CollectionHashedReadWriteSet encapsulate the hashed representation for the private read-write set for a collection
type CollectionHashedReadWriteSet struct {
	CollectionName string `protobuf:"bytes,1,opt,name=collection_name,json=collectionName" json:"collection_name,omitempty"`
	HashedRwset    []byte `protobuf:"bytes,2,opt,name=hashed_rwset,json=hashedRwset,proto3" json:"hashed_rwset,omitempty"`
	PvtRwsetHash   []byte `protobuf:"bytes,3,opt,name=pvt_rwset_hash,json=pvtRwsetHash,proto3" json:"pvt_rwset_hash,omitempty"`
}

func (m *CollectionHashedReadWriteSet) Reset()                    { *m = CollectionHashedReadWriteSet{} }
func (m *CollectionHashedReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*CollectionHashedReadWriteSet) ProtoMessage()               {}
func (*CollectionHashedReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CollectionHashedReadWriteSet) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *CollectionHashedReadWriteSet) GetHashedRwset() []byte {
	if m != nil {
		return m.HashedRwset
	}
	return nil
}

func (m *CollectionHashedReadWriteSet) GetPvtRwsetHash() []byte {
	if m != nil {
		return m.PvtRwsetHash
	}
	return nil
}

// TxPvtReadWriteSet encapsulate the private read-write set for a transaction
type TxPvtReadWriteSet struct {
	DataModel  TxReadWriteSet_DataModel `protobuf:"varint,1,opt,name=data_model,json=dataModel,enum=rwset.TxReadWriteSet_DataModel" json:"data_model,omitempty"`
	NsPvtRwset []*NsPvtReadWriteSet     `protobuf:"bytes,2,rep,name=ns_pvt_rwset,json=nsPvtRwset" json:"ns_pvt_rwset,omitempty"`
}

func (m *TxPvtReadWriteSet) Reset()                    { *m = TxPvtReadWriteSet{} }
func (m *TxPvtReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*TxPvtReadWriteSet) ProtoMessage()               {}
func (*TxPvtReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TxPvtReadWriteSet) GetDataModel() TxReadWriteSet_DataModel {
	if m != nil {
		return m.DataModel
	}
	return TxReadWriteSet_KV
}

func (m *TxPvtReadWriteSet) GetNsPvtRwset() []*NsPvtReadWriteSet {
	if m != nil {
		return m.NsPvtRwset
	}
	return nil
}

// NsPvtReadWriteSet encapsulates the private read-write set for a chaincode
type NsPvtReadWriteSet struct {
	Namespace          string                       `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	CollectionPvtRwset []*CollectionPvtReadWriteSet `protobuf:"bytes,2,rep,name=collection_pvt_rwset,json=collectionPvtRwset" json:"collection_pvt_rwset,omitempty"`
}

func (m *NsPvtReadWriteSet) Reset()                    { *m = NsPvtReadWriteSet{} }
func (m *NsPvtReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*NsPvtReadWriteSet) ProtoMessage()               {}
func (*NsPvtReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NsPvtReadWriteSet) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NsPvtReadWriteSet) GetCollectionPvtRwset() []*CollectionPvtReadWriteSet {
	if m != nil {
		return m.CollectionPvtRwset
	}
	return nil
}

// CollectionPvtReadWriteSet encapsulates the private read-write set for a collection
type CollectionPvtReadWriteSet struct {
	CollectionName string `protobuf:"bytes,1,opt,name=collection_name,json=collectionName" json:"collection_name,omitempty"`
	Rwset          []byte `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"`
}

func (m *CollectionPvtReadWriteSet) Reset()                    { *m = CollectionPvtReadWriteSet{} }
func (m *CollectionPvtReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*CollectionPvtReadWriteSet) ProtoMessage()               {}
func (*CollectionPvtReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CollectionPvtReadWriteSet) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *CollectionPvtReadWriteSet) GetRwset() []byte {
	if m != nil {
		return m.Rwset
	}
	return nil
}

func init() {
	proto.RegisterType((*TxReadWriteSet)(nil), "rwset.TxReadWriteSet")
	proto.RegisterType((*NsReadWriteSet)(nil), "rwset.NsReadWriteSet")
	proto.RegisterType((*CollectionHashedReadWriteSet)(nil), "rwset.CollectionHashedReadWriteSet")
	proto.RegisterType((*TxPvtReadWriteSet)(nil), "rwset.TxPvtReadWriteSet")
	proto.RegisterType((*NsPvtReadWriteSet)(nil), "rwset.NsPvtReadWriteSet")
	proto.RegisterType((*CollectionPvtReadWriteSet)(nil), "rwset.CollectionPvtReadWriteSet")
	proto.RegisterEnum("rwset.TxReadWriteSet_DataModel", TxReadWriteSet_DataModel_name, TxReadWriteSet_DataModel_value)
}

func init() { proto.RegisterFile("ledger/rwset/rwset.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x71, 0xab, 0x16, 0xf2, 0x36, 0x0a, 0xd4, 0xb4, 0x22, 0x48, 0x95, 0x28, 0x05, 0x89,
	0x8a, 0x43, 0x02, 0xe5, 0xc6, 0x81, 0x03, 0x70, 0x40, 0x42, 0x54, 0xc8, 0x54, 0x4c, 0xea, 0x0e,
	0x91, 0x9b, 0x78, 0x4d, 0xa4, 0xfc, 0x53, 0xec, 0x75, 0xdd, 0x07, 0xd8, 0x79, 0xbb, 0xed, 0xbc,
	0x6f, 0x3a, 0xd5, 0x4e, 0xd3, 0x24, 0xdd, 0xbf, 0xc3, 0x2e, 0x51, 0xfc, 0xfa, 0x79, 0xf2, 0xfc,
	0xec, 0x37, 0x2f, 0x98, 0x21, 0xf3, 0x96, 0x2c, 0xb3, 0xb3, 0x13, 0xce, 0x84, 0x7a, 0x5a, 0x69,
	0x96, 0x88, 0x04, 0xb7, 0xe4, 0x62, 0x74, 0x89, 0xc0, 0x98, 0xad, 0x09, 0xa3, 0xde, 0x41, 0x16,
	0x08, 0xf6, 0x8f, 0x09, 0xfc, 0x0d, 0xc0, 0xa3, 0x82, 0x3a, 0x51, 0xe2, 0xb1, 0xd0, 0x44, 0x43,
	0x34, 0x36, 0x26, 0x6f, 0x2c, 0xe5, 0xad, 0x4a, 0xad, 0x9f, 0x54, 0xd0, 0x3f, 0x1b, 0x19, 0xd1,
	0xbc, 0xed, 0x2b, 0xfe, 0x04, 0xcf, 0x62, 0xee, 0x48, 0xbd, 0xd9, 0x18, 0x36, 0xc7, 0x9d, 0x49,
	0x3f, 0x77, 0x4f, 0x79, 0xd9, 0x4d, 0x9e, 0xc6, 0x9c, 0x48, 0x88, 0x97, 0xa0, 0x15, 0x5f, 0xc2,
	0x6d, 0x68, 0xfc, 0xfe, 0xff, 0xe2, 0xc9, 0xe8, 0x0a, 0x81, 0x51, 0x35, 0xe0, 0x01, 0x68, 0x31,
	0x8d, 0x18, 0x4f, 0xa9, 0xcb, 0x24, 0x98, 0x46, 0x76, 0x05, 0xdc, 0x83, 0xd6, 0x36, 0x14, 0x8d,
	0x75, 0xa2, 0x16, 0xf8, 0x10, 0x5e, 0xb9, 0x49, 0x18, 0x32, 0x57, 0x04, 0x49, 0xec, 0xf8, 0x94,
	0xfb, 0xcc, 0xcb, 0xe1, 0x9a, 0x12, 0xee, 0x5d, 0x0e, 0xf7, 0xa3, 0x50, 0xfd, 0x92, 0xa2, 0x0a,
	0x6a, 0xdf, 0xad, 0xef, 0x4a, 0xf0, 0x0b, 0x04, 0x83, 0xbb, 0x7c, 0xf8, 0x03, 0x3c, 0x2f, 0xa5,
	0x6f, 0x58, 0x73, 0x6e, 0x63, 0x57, 0x9e, 0xd2, 0x88, 0xe1, 0xb7, 0xa0, 0x57, 0xd8, 0xd4, 0x19,
	0x3a, 0xfe, 0x2e, 0x0c, 0xbf, 0x07, 0x23, 0x5d, 0x09, 0xb5, 0x2f, 0x0f, 0x62, 0x36, 0xa5, 0x48,
	0x4f, 0x57, 0x42, 0x2a, 0x36, 0xf9, 0xa3, 0x73, 0x04, 0xdd, 0xd9, 0xfa, 0xef, 0x4a, 0x3c, 0x6a,
	0x4f, 0xbf, 0x82, 0x1e, 0x73, 0xa7, 0x88, 0xcf, 0xfb, 0x6a, 0x16, 0x7d, 0xad, 0xe5, 0x11, 0x88,
	0x65, 0x49, 0x5e, 0xd2, 0x19, 0x82, 0xee, 0x9e, 0xe2, 0x9e, 0x5e, 0x12, 0xe8, 0x95, 0xee, 0xad,
	0x9e, 0x3b, 0xdc, 0x6b, 0x59, 0x3d, 0x1f, 0xbb, 0x95, 0x2d, 0xc9, 0x31, 0x87, 0xd7, 0xb7, 0x1a,
	0x1e, 0xde, 0xa8, 0x1b, 0xff, 0xb2, 0xef, 0x0e, 0x7c, 0x4c, 0xb2, 0xa5, 0xe5, 0x9f, 0xa6, 0x2c,
	0x53, 0x23, 0x67, 0x1d, 0xd1, 0x45, 0x16, 0xb8, 0x6a, 0xda, 0xb8, 0x95, 0x17, 0xa5, 0x7a, 0xfe,
	0x79, 0x19, 0x08, 0xff, 0x78, 0x61, 0xb9, 0x49, 0x64, 0x97, 0x2c, 0xb6, 0xb2, 0xd8, 0xca, 0x62,
	0x97, 0x47, 0x77, 0xd1, 0x96, 0xc5, 0x2f, 0xd7, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x8a, 0xfa,
	0x65, 0xd1, 0x03, 0x00, 0x00,
}