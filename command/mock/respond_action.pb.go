// Code generated by protoc-gen-go. DO NOT EDIT.
// source: respond_action.proto

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
//  RespondAction
//
type RespondAction struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" firestore:"-"`
	XXX_unrecognized     []byte   `json:"-" firestore:"-"`
	XXX_sizecache        int32    `json:"-" firestore:"-"`
}

func (m *RespondAction) Reset()         { *m = RespondAction{} }
func (m *RespondAction) String() string { return proto.CompactTextString(m) }
func (*RespondAction) ProtoMessage()    {}
func (*RespondAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_3706df34e1021090, []int{0}
}

func (m *RespondAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespondAction.Unmarshal(m, b)
}
func (m *RespondAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespondAction.Marshal(b, m, deterministic)
}
func (m *RespondAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespondAction.Merge(m, src)
}
func (m *RespondAction) XXX_Size() int {
	return xxx_messageInfo_RespondAction.Size(m)
}
func (m *RespondAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RespondAction.DiscardUnknown(m)
}

var xxx_messageInfo_RespondAction proto.InternalMessageInfo

func (m *RespondAction) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*RespondAction)(nil), "RespondAction")
}

func init() { proto.RegisterFile("respond_action.proto", fileDescriptor_3706df34e1021090) }

var fileDescriptor_3706df34e1021090 = []byte{
	// 78 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x4b, 0x89, 0x4f, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x57, 0x52, 0xe6, 0xe2, 0x0d, 0x82, 0x88, 0x3b, 0x82, 0x85, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x52,
	0x2b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0xec, 0x24, 0x36, 0xb0, 0x5a, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x4e, 0x53, 0x89, 0x43, 0x00, 0x00, 0x00,
}


func (m *RespondAction) XXX_MapID() uint16 {
	return 1003
}

func (m *RespondAction) XXX_MapName() string {
	return "RespondAction"
}
