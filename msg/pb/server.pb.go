// Code generated by protoc-gen-go.
// source: server.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	DomainMap
	MachineStatProto
	ServerStatProto
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type DomainMap struct {
	OperatorId       []int32  `protobuf:"varint,1,rep,name=operator_id" json:"operator_id,omitempty"`
	Domain           []string `protobuf:"bytes,2,rep,name=domain" json:"domain,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DomainMap) Reset()                    { *m = DomainMap{} }
func (m *DomainMap) String() string            { return proto.CompactTextString(m) }
func (*DomainMap) ProtoMessage()               {}
func (*DomainMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DomainMap) GetOperatorId() []int32 {
	if m != nil {
		return m.OperatorId
	}
	return nil
}

func (m *DomainMap) GetDomain() []string {
	if m != nil {
		return m.Domain
	}
	return nil
}

type MachineStatProto struct {
	Cpu              *string `protobuf:"bytes,1,opt,name=cpu" json:"cpu,omitempty"`
	Mem              *string `protobuf:"bytes,2,opt,name=mem" json:"mem,omitempty"`
	Disk             *string `protobuf:"bytes,3,opt,name=disk" json:"disk,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MachineStatProto) Reset()                    { *m = MachineStatProto{} }
func (m *MachineStatProto) String() string            { return proto.CompactTextString(m) }
func (*MachineStatProto) ProtoMessage()               {}
func (*MachineStatProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MachineStatProto) GetCpu() string {
	if m != nil && m.Cpu != nil {
		return *m.Cpu
	}
	return ""
}

func (m *MachineStatProto) GetMem() string {
	if m != nil && m.Mem != nil {
		return *m.Mem
	}
	return ""
}

func (m *MachineStatProto) GetDisk() string {
	if m != nil && m.Disk != nil {
		return *m.Disk
	}
	return ""
}

type ServerStatProto struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ResName          *string `protobuf:"bytes,2,opt,name=res_name" json:"res_name,omitempty"`
	GameCfg          []byte  `protobuf:"bytes,3,opt,name=game_cfg" json:"game_cfg,omitempty"`
	Stat             *int32  `protobuf:"varint,4,opt,name=stat" json:"stat,omitempty"`
	DoingStat        *int32  `protobuf:"varint,5,opt,name=doing_stat" json:"doing_stat,omitempty"`
	DoingStatErr     *string `protobuf:"bytes,6,opt,name=doing_stat_err" json:"doing_stat_err,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ServerStatProto) Reset()                    { *m = ServerStatProto{} }
func (m *ServerStatProto) String() string            { return proto.CompactTextString(m) }
func (*ServerStatProto) ProtoMessage()               {}
func (*ServerStatProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServerStatProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ServerStatProto) GetResName() string {
	if m != nil && m.ResName != nil {
		return *m.ResName
	}
	return ""
}

func (m *ServerStatProto) GetGameCfg() []byte {
	if m != nil {
		return m.GameCfg
	}
	return nil
}

func (m *ServerStatProto) GetStat() int32 {
	if m != nil && m.Stat != nil {
		return *m.Stat
	}
	return 0
}

func (m *ServerStatProto) GetDoingStat() int32 {
	if m != nil && m.DoingStat != nil {
		return *m.DoingStat
	}
	return 0
}

func (m *ServerStatProto) GetDoingStatErr() string {
	if m != nil && m.DoingStatErr != nil {
		return *m.DoingStatErr
	}
	return ""
}

func init() {
	proto.RegisterType((*DomainMap)(nil), "pb.DomainMap")
	proto.RegisterType((*MachineStatProto)(nil), "pb.MachineStatProto")
	proto.RegisterType((*ServerStatProto)(nil), "pb.ServerStatProto")
}

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x8e, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x46, 0xb9, 0xed, 0x6d, 0xb1, 0xd3, 0x52, 0x4b, 0x04, 0xc9, 0x52, 0xba, 0xea, 0x4a, 0x5c,
	0xbb, 0x76, 0x5b, 0x10, 0xfa, 0x00, 0x21, 0x36, 0x63, 0x0d, 0x92, 0x1f, 0x26, 0xd1, 0x9d, 0xef,
	0x6e, 0x12, 0x94, 0x2e, 0xcf, 0x61, 0xce, 0xc7, 0xc0, 0x10, 0x90, 0xbe, 0x91, 0x1e, 0x3d, 0xb9,
	0xe8, 0x58, 0xe5, 0xdf, 0xe6, 0x27, 0xe8, 0x5e, 0x9c, 0x91, 0xda, 0xae, 0xd2, 0xb3, 0x3b, 0xe8,
	0x9d, 0x47, 0x92, 0xd1, 0x91, 0xd0, 0x8a, 0x5f, 0x1e, 0xea, 0xa5, 0x61, 0x23, 0xb4, 0xaa, 0x5c,
	0xf0, 0x2a, 0x71, 0x37, 0x3f, 0xc3, 0xb4, 0xca, 0xfd, 0x43, 0x5b, 0xdc, 0xa2, 0x8c, 0xaf, 0x65,
	0xa9, 0x87, 0x7a, 0xf7, 0x5f, 0x29, 0xb8, 0x2c, 0x5d, 0x06, 0x83, 0x26, 0x5d, 0x67, 0x18, 0xe0,
	0xaa, 0x74, 0xf8, 0xe4, 0x75, 0xa6, 0xf9, 0x07, 0x6e, 0xb7, 0xf2, 0xc1, 0x99, 0xa6, 0x03, 0x2b,
	0x0d, 0xfe, 0xb5, 0x13, 0xdc, 0x10, 0x06, 0x51, 0x4c, 0xf5, 0x6f, 0x8e, 0x44, 0x62, 0x7f, 0x3f,
	0xca, 0xc8, 0x90, 0x8b, 0x90, 0x72, 0x7e, 0x4d, 0xd4, 0x30, 0x06, 0xa0, 0x9c, 0xb6, 0x87, 0x28,
	0xae, 0x29, 0xee, 0x1e, 0xc6, 0xd3, 0x09, 0x24, 0xe2, 0x6d, 0xde, 0xfa, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xcb, 0x58, 0x5c, 0x7d, 0xff, 0x00, 0x00, 0x00,
}