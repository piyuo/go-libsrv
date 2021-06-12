// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd-respond.proto

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
//  CmdRespond
//
type CmdRespond struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" firestore:"-"`
	XXX_unrecognized     []byte   `json:"-" firestore:"-"`
	XXX_sizecache        int32    `json:"-" firestore:"-"`
}

func (m *CmdRespond) Reset()         { *m = CmdRespond{} }
func (m *CmdRespond) String() string { return proto.CompactTextString(m) }
func (*CmdRespond) ProtoMessage()    {}
func (*CmdRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_d423697d8819ed3b, []int{0}
}

func (m *CmdRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdRespond.Unmarshal(m, b)
}
func (m *CmdRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdRespond.Marshal(b, m, deterministic)
}
func (m *CmdRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdRespond.Merge(m, src)
}
func (m *CmdRespond) XXX_Size() int {
	return xxx_messageInfo_CmdRespond.Size(m)
}
func (m *CmdRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdRespond.DiscardUnknown(m)
}

var xxx_messageInfo_CmdRespond proto.InternalMessageInfo

func (m *CmdRespond) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*CmdRespond)(nil), "CmdRespond")
}

func init() { proto.RegisterFile("cmd-respond.proto", fileDescriptor_d423697d8819ed3b) }

var fileDescriptor_d423697d8819ed3b = []byte{
	// 75 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xce, 0x4d, 0xd1,
	0x2d, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe0,
	0xe2, 0x72, 0xce, 0x4d, 0x09, 0x82, 0x88, 0x09, 0x09, 0x71, 0xb1, 0x94, 0xa4, 0x56, 0x94, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x49, 0x6c, 0x60, 0x85, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd5, 0xa2, 0xdb, 0x68, 0x3d, 0x00, 0x00, 0x00,
}


func (m *CmdRespond) XXX_MapID() uint16 {
	return 1004
}

func (m *CmdRespond) XXX_MapName() string {
	return "CmdRespond"
}
