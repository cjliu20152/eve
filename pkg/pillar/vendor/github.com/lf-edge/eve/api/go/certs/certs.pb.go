// Code generated by protoc-gen-go. DO NOT EDIT.
// source: certs/certs.proto

package certs

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

//This is same as hashAlgorithm in auth/auth.proto
//Keep these two in sync
//XXX: import auth/auth.proto and avoid this duplication
type CertHashAlgorithm int32

const (
	CertHashAlgorithm_HASH_NONE           CertHashAlgorithm = 0
	CertHashAlgorithm_HASH_SHA256_16bytes CertHashAlgorithm = 1
	CertHashAlgorithm_HASH_SHA256_32bytes CertHashAlgorithm = 2
)

var CertHashAlgorithm_name = map[int32]string{
	0: "HASH_NONE",
	1: "HASH_SHA256_16bytes",
	2: "HASH_SHA256_32bytes",
}

var CertHashAlgorithm_value = map[string]int32{
	"HASH_NONE":           0,
	"HASH_SHA256_16bytes": 1,
	"HASH_SHA256_32bytes": 2,
}

func (x CertHashAlgorithm) String() string {
	return proto.EnumName(CertHashAlgorithm_name, int32(x))
}

func (CertHashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9bf61e17e095098d, []int{0}
}

type ZCertType int32

const (
	ZCertType_CERT_TYPE_CONTROLLER_NONE          ZCertType = 0
	ZCertType_CERT_TYPE_CONTROLLER_SIGNING       ZCertType = 1
	ZCertType_CERT_TYPE_CONTROLLER_INTERMEDIATE  ZCertType = 2
	ZCertType_CERT_TYPE_CONTROLLER_ECDH_EXCHANGE ZCertType = 3
)

var ZCertType_name = map[int32]string{
	0: "CERT_TYPE_CONTROLLER_NONE",
	1: "CERT_TYPE_CONTROLLER_SIGNING",
	2: "CERT_TYPE_CONTROLLER_INTERMEDIATE",
	3: "CERT_TYPE_CONTROLLER_ECDH_EXCHANGE",
}

var ZCertType_value = map[string]int32{
	"CERT_TYPE_CONTROLLER_NONE":          0,
	"CERT_TYPE_CONTROLLER_SIGNING":       1,
	"CERT_TYPE_CONTROLLER_INTERMEDIATE":  2,
	"CERT_TYPE_CONTROLLER_ECDH_EXCHANGE": 3,
}

func (x ZCertType) String() string {
	return proto.EnumName(ZCertType_name, int32(x))
}

func (ZCertType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9bf61e17e095098d, []int{1}
}

// This is the response payload for GET /api/v1/edgeDevice/certs
// or /api/v2/edgeDevice/certs
// ZControllerCert carries a set of X.509 certificate and their properties
// from Controller to EVE.
type ZControllerCert struct {
	Certs                []*ZCert `protobuf:"bytes,1,rep,name=certs,proto3" json:"certs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZControllerCert) Reset()         { *m = ZControllerCert{} }
func (m *ZControllerCert) String() string { return proto.CompactTextString(m) }
func (*ZControllerCert) ProtoMessage()    {}
func (*ZControllerCert) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bf61e17e095098d, []int{0}
}

func (m *ZControllerCert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZControllerCert.Unmarshal(m, b)
}
func (m *ZControllerCert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZControllerCert.Marshal(b, m, deterministic)
}
func (m *ZControllerCert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZControllerCert.Merge(m, src)
}
func (m *ZControllerCert) XXX_Size() int {
	return xxx_messageInfo_ZControllerCert.Size(m)
}
func (m *ZControllerCert) XXX_DiscardUnknown() {
	xxx_messageInfo_ZControllerCert.DiscardUnknown(m)
}

var xxx_messageInfo_ZControllerCert proto.InternalMessageInfo

func (m *ZControllerCert) GetCerts() []*ZCert {
	if m != nil {
		return m.Certs
	}
	return nil
}

type ZCert struct {
	HashAlgo             CertHashAlgorithm `protobuf:"varint,1,opt,name=hashAlgo,proto3,enum=CertHashAlgorithm" json:"hashAlgo,omitempty"`
	CertHash             []byte            `protobuf:"bytes,2,opt,name=certHash,proto3" json:"certHash,omitempty"`
	Type                 ZCertType         `protobuf:"varint,3,opt,name=type,proto3,enum=ZCertType" json:"type,omitempty"`
	Cert                 []byte            `protobuf:"bytes,4,opt,name=cert,proto3" json:"cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ZCert) Reset()         { *m = ZCert{} }
func (m *ZCert) String() string { return proto.CompactTextString(m) }
func (*ZCert) ProtoMessage()    {}
func (*ZCert) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bf61e17e095098d, []int{1}
}

func (m *ZCert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZCert.Unmarshal(m, b)
}
func (m *ZCert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZCert.Marshal(b, m, deterministic)
}
func (m *ZCert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZCert.Merge(m, src)
}
func (m *ZCert) XXX_Size() int {
	return xxx_messageInfo_ZCert.Size(m)
}
func (m *ZCert) XXX_DiscardUnknown() {
	xxx_messageInfo_ZCert.DiscardUnknown(m)
}

var xxx_messageInfo_ZCert proto.InternalMessageInfo

func (m *ZCert) GetHashAlgo() CertHashAlgorithm {
	if m != nil {
		return m.HashAlgo
	}
	return CertHashAlgorithm_HASH_NONE
}

func (m *ZCert) GetCertHash() []byte {
	if m != nil {
		return m.CertHash
	}
	return nil
}

func (m *ZCert) GetType() ZCertType {
	if m != nil {
		return m.Type
	}
	return ZCertType_CERT_TYPE_CONTROLLER_NONE
}

func (m *ZCert) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func init() {
	proto.RegisterEnum("CertHashAlgorithm", CertHashAlgorithm_name, CertHashAlgorithm_value)
	proto.RegisterEnum("ZCertType", ZCertType_name, ZCertType_value)
	proto.RegisterType((*ZControllerCert)(nil), "ZControllerCert")
	proto.RegisterType((*ZCert)(nil), "ZCert")
}

func init() {
	proto.RegisterFile("certs/certs.proto", fileDescriptor_9bf61e17e095098d)
}

var fileDescriptor_9bf61e17e095098d = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xdd, 0x8a, 0xda, 0x40,
	0x14, 0xc7, 0x3b, 0x7e, 0xa1, 0xa7, 0x5f, 0xe9, 0xf4, 0xa2, 0x69, 0xb1, 0x6d, 0x6a, 0x69, 0x11,
	0xa1, 0x13, 0x1a, 0xa9, 0xf7, 0x69, 0x9c, 0x9a, 0x80, 0x8d, 0x65, 0xcc, 0x45, 0xeb, 0x4d, 0xd0,
	0x64, 0x9a, 0x04, 0xe2, 0x4e, 0x48, 0x46, 0xc1, 0x7d, 0x81, 0x7d, 0x88, 0x7d, 0xd9, 0xc5, 0xc9,
	0x2a, 0xec, 0xe2, 0xcd, 0x30, 0xe7, 0xf7, 0xff, 0xcd, 0x39, 0x0c, 0x07, 0x5e, 0x45, 0xbc, 0x94,
	0x95, 0xa9, 0x4e, 0x52, 0x94, 0x42, 0x8a, 0x81, 0x09, 0x2f, 0x57, 0x8e, 0xb8, 0x92, 0xa5, 0xc8,
	0x73, 0x5e, 0x3a, 0xbc, 0x94, 0xb8, 0x0f, 0x6d, 0x65, 0xe8, 0xc8, 0x68, 0x0e, 0x9f, 0x5a, 0x1d,
	0xb2, 0x3a, 0x62, 0x56, 0xc3, 0xc1, 0x0d, 0x82, 0xb6, 0x02, 0x98, 0x40, 0x37, 0x5d, 0x57, 0xa9,
	0x9d, 0x27, 0x42, 0x47, 0x06, 0x1a, 0xbe, 0xb0, 0x30, 0x39, 0x3a, 0xee, 0x3d, 0x2c, 0x33, 0x99,
	0x6e, 0xd9, 0xd9, 0xc1, 0xef, 0xa0, 0x7b, 0x8a, 0xf5, 0x86, 0x81, 0x86, 0xcf, 0xd8, 0xb9, 0xc6,
	0x1f, 0xa0, 0x25, 0x0f, 0x05, 0xd7, 0x9b, 0xaa, 0x0f, 0xd4, 0x23, 0x83, 0x43, 0xc1, 0x99, 0xe2,
	0x18, 0x43, 0xeb, 0xe8, 0xea, 0x2d, 0xf5, 0x4e, 0xdd, 0x47, 0x41, 0xfd, 0x9f, 0x07, 0xe3, 0xf0,
	0x73, 0xe8, 0xb9, 0xf6, 0xd2, 0x0d, 0xfd, 0x85, 0x4f, 0xb5, 0x27, 0xf8, 0x0d, 0xbc, 0x56, 0xe5,
	0xd2, 0xb5, 0xad, 0x1f, 0x93, 0xf0, 0xfb, 0x64, 0x73, 0x90, 0xbc, 0xd2, 0xd0, 0xe3, 0x60, 0x6c,
	0xd5, 0x41, 0x63, 0x74, 0x8b, 0xa0, 0x77, 0x9e, 0x8e, 0xdf, 0xc3, 0x5b, 0x87, 0xb2, 0x20, 0x0c,
	0xfe, 0xfd, 0xa1, 0xa1, 0xb3, 0xf0, 0x03, 0xb6, 0x98, 0xcf, 0x29, 0x3b, 0xb5, 0x37, 0xa0, 0x7f,
	0x31, 0x5e, 0x7a, 0x33, 0xdf, 0xf3, 0x67, 0x1a, 0xc2, 0x5f, 0xe0, 0xd3, 0x45, 0xc3, 0xf3, 0x03,
	0xca, 0x7e, 0xd3, 0xa9, 0x67, 0x07, 0x54, 0x6b, 0xe0, 0xaf, 0x30, 0xb8, 0xa8, 0x51, 0x67, 0xea,
	0x86, 0xf4, 0xaf, 0xe3, 0xda, 0xfe, 0x8c, 0x6a, 0xcd, 0x9f, 0xbf, 0xe0, 0x63, 0x24, 0xb6, 0xe4,
	0x9a, 0xc7, 0x3c, 0x5e, 0x93, 0x28, 0x17, 0xbb, 0x98, 0xec, 0x2a, 0x5e, 0xee, 0xb3, 0x88, 0xd7,
	0x1b, 0x5d, 0x7d, 0x4e, 0x32, 0x99, 0xee, 0x36, 0x24, 0x12, 0x5b, 0x33, 0xff, 0xff, 0x8d, 0xc7,
	0x09, 0x37, 0xf9, 0x9e, 0x9b, 0xeb, 0x22, 0x33, 0x13, 0x51, 0x2f, 0x7f, 0xd3, 0x51, 0xee, 0xf8,
	0x2e, 0x00, 0x00, 0xff, 0xff, 0x89, 0xf4, 0x8b, 0xae, 0x12, 0x02, 0x00, 0x00,
}
