// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth.proto

package auth

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

type Request struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_84c28b9eff86c7f4, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

type Result struct {
	Customer             *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_84c28b9eff86c7f4, []int{1}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type Customer struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthToken            string   `protobuf:"bytes,2,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_84c28b9eff86c7f4, []int{2}
}
func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (dst *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(dst, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "auth.Request")
	proto.RegisterType((*Result)(nil), "auth.Result")
	proto.RegisterType((*Customer)(nil), "auth.Customer")
}

func init() { proto.RegisterFile("proto/auth.proto", fileDescriptor_auth_84c28b9eff86c7f4) }

var fileDescriptor_auth_84c28b9eff86c7f4 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x03, 0x33, 0x85, 0x58, 0x40, 0x6c, 0x25, 0x75, 0x2e,
	0xf6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x19, 0x2e, 0x4e, 0x90, 0x50, 0x48, 0x7e,
	0x76, 0x6a, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x42, 0x40, 0xc9, 0x84, 0x8b, 0x2d,
	0x28, 0xb5, 0xb8, 0x34, 0xa7, 0x44, 0x48, 0x8b, 0x8b, 0x23, 0xb9, 0xb4, 0xb8, 0x24, 0x3f, 0x37,
	0xb5, 0x08, 0xac, 0x8c, 0xdb, 0x88, 0x4f, 0x0f, 0x6c, 0xae, 0x33, 0x54, 0x34, 0x08, 0x2e, 0xaf,
	0x64, 0xc1, 0xc5, 0x01, 0x13, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x01, 0xeb, 0x60, 0x0d, 0x62,
	0xca, 0x4c, 0x41, 0xb5, 0x8f, 0x09, 0xcd, 0x3e, 0x23, 0x13, 0x2e, 0x16, 0xc7, 0xd2, 0x92, 0x0c,
	0x21, 0x1d, 0x2e, 0xee, 0xb0, 0xd4, 0xa2, 0xcc, 0xb4, 0x4a, 0xb0, 0xb0, 0x10, 0x2f, 0xc4, 0x2a,
	0xa8, 0x9b, 0xa5, 0x78, 0x60, 0x5c, 0x90, 0xcb, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x7e, 0x33, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x0d, 0xfb, 0xf8, 0xef, 0x00, 0x00, 0x00,
}
