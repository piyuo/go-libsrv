// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd-big-data.proto

package mock

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

//*
//  CmdBigData
//
type CmdBigData struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" firestore:"-"`
	XXX_unrecognized     []byte   `json:"-" firestore:"-"`
	XXX_sizecache        int32    `json:"-" firestore:"-"`
}

func (m *CmdBigData) Reset()         { *m = CmdBigData{} }
func (m *CmdBigData) String() string { return proto.CompactTextString(m) }
func (*CmdBigData) ProtoMessage()    {}
func (*CmdBigData) Descriptor() ([]byte, []int) {
	return fileDescriptor_74ba3abe540ef2a1, []int{0}
}

func (m *CmdBigData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdBigData.Unmarshal(m, b)
}
func (m *CmdBigData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdBigData.Marshal(b, m, deterministic)
}
func (m *CmdBigData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdBigData.Merge(m, src)
}
func (m *CmdBigData) XXX_Size() int {
	return xxx_messageInfo_CmdBigData.Size(m)
}
func (m *CmdBigData) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdBigData.DiscardUnknown(m)
}

var xxx_messageInfo_CmdBigData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CmdBigData)(nil), "CmdBigData")
}

func init() { proto.RegisterFile("cmd-big-data.proto", fileDescriptor_74ba3abe540ef2a1) }

var fileDescriptor_74ba3abe540ef2a1 = []byte{
	// 63 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0x4d, 0xd1,
	0x4d, 0xca, 0x4c, 0xd7, 0x4d, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xe2,
	0xe1, 0xe2, 0x72, 0xce, 0x4d, 0x71, 0xca, 0x4c, 0x77, 0x49, 0x2c, 0x49, 0x4c, 0x62, 0x03, 0x0b,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x9c, 0x8d, 0x86, 0x2a, 0x00, 0x00, 0x00,
}


func (m *CmdBigData) XXX_MapID() uint16 {
	return 1001
}

func (m *CmdBigData) XXX_MapName() string {
	return "CmdBigData"
}
