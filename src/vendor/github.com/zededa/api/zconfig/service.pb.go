// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ZSrvType int32

const (
	ZSrvType_ZsrvFirst      ZSrvType = 0
	ZSrvType_ZsrvStrongSwan ZSrvType = 1
	ZSrvType_ZsrvLISP       ZSrvType = 2
	ZSrvType_ZsrvBridge     ZSrvType = 3
	ZSrvType_ZsrvNAT        ZSrvType = 4
	ZSrvType_ZsrvLB         ZSrvType = 5
	ZSrvType_ZsrvLast       ZSrvType = 255
)

var ZSrvType_name = map[int32]string{
	0:   "ZsrvFirst",
	1:   "ZsrvStrongSwan",
	2:   "ZsrvLISP",
	3:   "ZsrvBridge",
	4:   "ZsrvNAT",
	5:   "ZsrvLB",
	255: "ZsrvLast",
}
var ZSrvType_value = map[string]int32{
	"ZsrvFirst":      0,
	"ZsrvStrongSwan": 1,
	"ZsrvLISP":       2,
	"ZsrvBridge":     3,
	"ZsrvNAT":        4,
	"ZsrvLB":         5,
	"ZsrvLast":       255,
}

func (x ZSrvType) String() string {
	return proto.EnumName(ZSrvType_name, int32(x))
}
func (ZSrvType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Service Opaque config. In future we might add more fields here
// but idea is here. This is service specific configuration.
type ServiceOpaqueConfig struct {
	Oconfig string `protobuf:"bytes,1,opt,name=oconfig" json:"oconfig,omitempty"`
}

func (m *ServiceOpaqueConfig) Reset()                    { *m = ServiceOpaqueConfig{} }
func (m *ServiceOpaqueConfig) String() string            { return proto.CompactTextString(m) }
func (*ServiceOpaqueConfig) ProtoMessage()               {}
func (*ServiceOpaqueConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ServiceOpaqueConfig) GetOconfig() string {
	if m != nil {
		return m.Oconfig
	}
	return ""
}

type ServiceInstanceConfig struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Displayname string   `protobuf:"bytes,2,opt,name=displayname" json:"displayname,omitempty"`
	Srvtype     ZSrvType `protobuf:"varint,3,opt,name=srvtype,enum=ZSrvType" json:"srvtype,omitempty"`
	// Optional in future we might need this
	// 	VmConfig fixedresources = 3;
	// 	repeated Drive drives = 4;
	Activate bool `protobuf:"varint,5,opt,name=activate" json:"activate,omitempty"`
	// towards application which networkUUID to use
	// FIXME: In future there might multiple application network assignment
	// so this will become repeated.
	Applink string `protobuf:"bytes,10,opt,name=applink" json:"applink,omitempty"`
	// towards devices which phyiscal or logical adapter to use
	Devlink *Adapter `protobuf:"bytes,20,opt,name=devlink" json:"devlink,omitempty"`
	// Opaque configuration
	Cfg *ServiceOpaqueConfig `protobuf:"bytes,30,opt,name=cfg" json:"cfg,omitempty"`
}

func (m *ServiceInstanceConfig) Reset()                    { *m = ServiceInstanceConfig{} }
func (m *ServiceInstanceConfig) String() string            { return proto.CompactTextString(m) }
func (*ServiceInstanceConfig) ProtoMessage()               {}
func (*ServiceInstanceConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ServiceInstanceConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceInstanceConfig) GetDisplayname() string {
	if m != nil {
		return m.Displayname
	}
	return ""
}

func (m *ServiceInstanceConfig) GetSrvtype() ZSrvType {
	if m != nil {
		return m.Srvtype
	}
	return ZSrvType_ZsrvFirst
}

func (m *ServiceInstanceConfig) GetActivate() bool {
	if m != nil {
		return m.Activate
	}
	return false
}

func (m *ServiceInstanceConfig) GetApplink() string {
	if m != nil {
		return m.Applink
	}
	return ""
}

func (m *ServiceInstanceConfig) GetDevlink() *Adapter {
	if m != nil {
		return m.Devlink
	}
	return nil
}

func (m *ServiceInstanceConfig) GetCfg() *ServiceOpaqueConfig {
	if m != nil {
		return m.Cfg
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceOpaqueConfig)(nil), "ServiceOpaqueConfig")
	proto.RegisterType((*ServiceInstanceConfig)(nil), "ServiceInstanceConfig")
	proto.RegisterEnum("ZSrvType", ZSrvType_name, ZSrvType_value)
}

func init() { proto.RegisterFile("service.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xdf, 0x8a, 0xd4, 0x30,
	0x14, 0xc6, 0x4d, 0xc7, 0xdd, 0x76, 0xce, 0x38, 0xb5, 0xc4, 0x15, 0xca, 0x82, 0x5a, 0x46, 0x90,
	0xe2, 0x45, 0x0a, 0xeb, 0x0b, 0xb8, 0x23, 0x08, 0x0b, 0xa2, 0xd2, 0xee, 0xd5, 0xdc, 0x65, 0x92,
	0x4c, 0x0d, 0x4e, 0x93, 0x98, 0xa4, 0x95, 0x99, 0xa7, 0xf6, 0x0d, 0x94, 0xfe, 0x13, 0x85, 0xbd,
	0xcb, 0xef, 0x7c, 0xbf, 0x03, 0x1f, 0x39, 0xb0, 0x76, 0xc2, 0x76, 0x92, 0x09, 0x62, 0xac, 0xf6,
	0xfa, 0xfa, 0x29, 0x17, 0x1d, 0xd3, 0x4d, 0xa3, 0xd5, 0x38, 0xd8, 0x14, 0xf0, 0xac, 0x1a, 0x8d,
	0x2f, 0x86, 0xfe, 0x68, 0xc5, 0x07, 0xad, 0x0e, 0xb2, 0xc6, 0x29, 0x84, 0x9a, 0x0d, 0xcf, 0x14,
	0x65, 0x28, 0x5f, 0x96, 0x33, 0x6e, 0x7e, 0x21, 0x78, 0x3e, 0x6d, 0xdc, 0x29, 0xe7, 0xa9, 0x62,
	0xf3, 0x4e, 0x0c, 0x81, 0xe4, 0x93, 0x1e, 0x48, 0x8e, 0x33, 0x58, 0x71, 0xe9, 0xcc, 0x91, 0x9e,
	0x14, 0x6d, 0x44, 0x1a, 0x0c, 0xc1, 0xbf, 0x23, 0xfc, 0x1a, 0x42, 0x67, 0x3b, 0x7f, 0x32, 0x22,
	0x5d, 0x64, 0x28, 0x8f, 0x6f, 0x96, 0x64, 0x57, 0xd9, 0xee, 0xfe, 0x64, 0x44, 0x39, 0x27, 0xf8,
	0x1a, 0x22, 0xca, 0xbc, 0xec, 0xa8, 0x17, 0xe9, 0x45, 0x86, 0xf2, 0xa8, 0xfc, 0xcb, 0x7d, 0x4d,
	0x6a, 0xcc, 0x51, 0xaa, 0xef, 0x29, 0x8c, 0x35, 0x27, 0xc4, 0x1b, 0x08, 0xb9, 0xe8, 0x86, 0xe4,
	0x2a, 0x43, 0xf9, 0xea, 0x26, 0x22, 0xb7, 0x9c, 0x1a, 0x2f, 0x6c, 0x39, 0x07, 0xf8, 0x0d, 0x2c,
	0xd8, 0xa1, 0x4e, 0x5f, 0x0e, 0xf9, 0x15, 0x79, 0xe0, 0x1f, 0xca, 0x5e, 0x78, 0xeb, 0x20, 0x9a,
	0x6b, 0xe1, 0x35, 0x2c, 0x77, 0xce, 0x76, 0x1f, 0xa5, 0x75, 0x3e, 0x79, 0x84, 0x31, 0xc4, 0x3d,
	0x56, 0xde, 0x6a, 0x55, 0x57, 0x3f, 0xa9, 0x4a, 0x10, 0x7e, 0x02, 0x51, 0x3f, 0xfb, 0x74, 0x57,
	0x7d, 0x4d, 0x02, 0x1c, 0x03, 0xf4, 0xb4, 0xb5, 0x92, 0xd7, 0x22, 0x59, 0xe0, 0x15, 0x84, 0x3d,
	0x7f, 0xbe, 0xbd, 0x4f, 0x1e, 0x63, 0x80, 0xcb, 0x41, 0xdd, 0x26, 0x17, 0x78, 0x3d, 0xad, 0x51,
	0xe7, 0x93, 0xdf, 0x68, 0xfb, 0x1e, 0x5e, 0x31, 0xdd, 0x90, 0xb3, 0xe0, 0x82, 0x53, 0xc2, 0x8e,
	0xba, 0xe5, 0xa4, 0xfd, 0xef, 0x98, 0xbb, 0x17, 0xb5, 0xf4, 0xdf, 0xda, 0x3d, 0x61, 0xba, 0x29,
	0x46, 0xaf, 0xa0, 0x46, 0x16, 0xe7, 0xf1, 0x52, 0xfb, 0xcb, 0xc1, 0x7a, 0xf7, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x96, 0x01, 0xda, 0x03, 0x03, 0x02, 0x00, 0x00,
}