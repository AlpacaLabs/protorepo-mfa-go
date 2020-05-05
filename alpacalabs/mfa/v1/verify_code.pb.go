// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpacalabs/mfa/v1/verify_code.proto

package v1

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

type VerifyCodeRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyCodeRequest) Reset()         { *m = VerifyCodeRequest{} }
func (m *VerifyCodeRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyCodeRequest) ProtoMessage()    {}
func (*VerifyCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d997ced59e054de, []int{0}
}

func (m *VerifyCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyCodeRequest.Unmarshal(m, b)
}
func (m *VerifyCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyCodeRequest.Marshal(b, m, deterministic)
}
func (m *VerifyCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyCodeRequest.Merge(m, src)
}
func (m *VerifyCodeRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyCodeRequest.Size(m)
}
func (m *VerifyCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyCodeRequest proto.InternalMessageInfo

func (m *VerifyCodeRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *VerifyCodeRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type VerifyCodeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyCodeResponse) Reset()         { *m = VerifyCodeResponse{} }
func (m *VerifyCodeResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyCodeResponse) ProtoMessage()    {}
func (*VerifyCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d997ced59e054de, []int{1}
}

func (m *VerifyCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyCodeResponse.Unmarshal(m, b)
}
func (m *VerifyCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyCodeResponse.Marshal(b, m, deterministic)
}
func (m *VerifyCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyCodeResponse.Merge(m, src)
}
func (m *VerifyCodeResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyCodeResponse.Size(m)
}
func (m *VerifyCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyCodeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VerifyCodeRequest)(nil), "alpacalabs.mfa.v1.VerifyCodeRequest")
	proto.RegisterType((*VerifyCodeResponse)(nil), "alpacalabs.mfa.v1.VerifyCodeResponse")
}

func init() {
	proto.RegisterFile("alpacalabs/mfa/v1/verify_code.proto", fileDescriptor_5d997ced59e054de)
}

var fileDescriptor_5d997ced59e054de = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xcc, 0x29, 0x48,
	0x4c, 0x4e, 0xcc, 0x49, 0x4c, 0x2a, 0xd6, 0xcf, 0x4d, 0x4b, 0xd4, 0x2f, 0x33, 0xd4, 0x2f, 0x4b,
	0x2d, 0xca, 0x4c, 0xab, 0x8c, 0x4f, 0xce, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x44, 0x28, 0xd2, 0xcb, 0x4d, 0x4b, 0xd4, 0x2b, 0x33, 0x54, 0x72, 0xe3, 0x12, 0x0c, 0x03,
	0xab, 0x73, 0xce, 0x4f, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62,
	0x01, 0xe9, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x64, 0xb9, 0xb8, 0x12,
	0x93, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0xe2, 0x33, 0x53, 0x24, 0x98, 0xc0, 0x32, 0x9c, 0x50, 0x11,
	0xcf, 0x14, 0x25, 0x11, 0x2e, 0x21, 0x64, 0x73, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x9d, 0xba,
	0x19, 0xb9, 0x44, 0x93, 0xf3, 0x73, 0xf5, 0x30, 0xec, 0x75, 0xe2, 0xf0, 0x4d, 0x4b, 0x0c, 0x00,
	0x39, 0x2a, 0x80, 0x31, 0xca, 0x22, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0xdf, 0x11, 0xac, 0xd2, 0x07, 0xe4, 0x0d, 0xb0, 0x9b, 0x8b, 0x52, 0x0b, 0xf2, 0x75, 0x73, 0xd3,
	0x12, 0x75, 0xd3, 0xf3, 0xf5, 0x31, 0xbc, 0xb8, 0x88, 0x89, 0xd9, 0xd1, 0xc7, 0x77, 0x15, 0x93,
	0x20, 0x42, 0x97, 0x9e, 0x6f, 0x5a, 0xa2, 0x5e, 0x98, 0xe1, 0x29, 0x64, 0xb1, 0x18, 0xdf, 0xb4,
	0xc4, 0x98, 0x30, 0xc3, 0x24, 0x36, 0xb0, 0x89, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9,
	0xa9, 0x0d, 0xcf, 0x2c, 0x01, 0x00, 0x00,
}
