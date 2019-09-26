// Code generated by protoc-gen-go. DO NOT EDIT.

package command

import (
	fmt "fmt"
	math "math"

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TestAction struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" firestore:"-"`
	XXX_unrecognized     []byte   `json:"-" firestore:"-"`
	XXX_sizecache        int32    `json:"-" firestore:"-"`
}

func (m *TestAction) Reset()         { *m = TestAction{} }
func (m *TestAction) String() string { return proto.CompactTextString(m) }
func (*TestAction) ProtoMessage()    {}
func (*TestAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83791fb8244f8c0, []int{0}
}

func (m *TestAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestAction.Unmarshal(m, b)
}
func (m *TestAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestAction.Marshal(b, m, deterministic)
}
func (m *TestAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAction.Merge(m, src)
}
func (m *TestAction) XXX_Size() int {
	return xxx_messageInfo_TestAction.Size(m)
}
func (m *TestAction) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAction.DiscardUnknown(m)
}

var xxx_messageInfo_TestAction proto.InternalMessageInfo

func (m *TestAction) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*TestAction)(nil), "TestAction")
}

func init() { proto.RegisterFile("echo_action.proto", fileDescriptor_d83791fb8244f8c0) }

var fileDescriptor_d83791fb8244f8c0 = []byte{
	// 76 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4d, 0xce, 0xc8,
	0x8f, 0x4f, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe0,
	0xe2, 0x72, 0x4d, 0xce, 0xc8, 0x77, 0x04, 0x8b, 0x09, 0x09, 0x71, 0xb1, 0x94, 0xa4, 0x56, 0x94,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x49, 0x6c, 0x60, 0x85, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xee, 0xb7, 0x79, 0xda, 0x3d, 0x00, 0x00, 0x00,
}

func (m *TestAction) XXX_MapID() uint16 {
	return 1
}

func (m *TestAction) XXX_MapName() string {
	return "TestAction"
}
