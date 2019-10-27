// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/meoconbatu/cmsgrpc/view/proto-files/user.proto

package view

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

type User struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_b204e4bf47e49b86, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
}

func init() {
	proto.RegisterFile("github.com/meoconbatu/cmsgrpc/view/proto-files/user.proto", fileDescriptor_b204e4bf47e49b86)
}

var fileDescriptor_b204e4bf47e49b86 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x4d, 0xcd, 0x4f, 0xce, 0xcf, 0x4b, 0x4a, 0x2c,
	0x29, 0xd5, 0x4f, 0xce, 0x2d, 0x4e, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0xcb, 0x4c, 0x2d, 0xd7, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4d, 0xcb, 0xcc, 0x49, 0x2d, 0xd6, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2,
	0x03, 0x0b, 0x08, 0xb1, 0x80, 0xd8, 0x4a, 0x76, 0x5c, 0x2c, 0xa1, 0xc5, 0xa9, 0x45, 0x42, 0x52,
	0x5c, 0x1c, 0x20, 0xbe, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c,
	0x0f, 0x92, 0x2b, 0x48, 0x2c, 0x2e, 0x2e, 0xcf, 0x2f, 0x4a, 0x91, 0x60, 0x82, 0xc8, 0xc1, 0xf8,
	0x4e, 0x2a, 0x51, 0x4a, 0x84, 0x9d, 0x90, 0xc4, 0x06, 0xb6, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x76, 0x08, 0xbd, 0x36, 0xaf, 0x00, 0x00, 0x00,
}
