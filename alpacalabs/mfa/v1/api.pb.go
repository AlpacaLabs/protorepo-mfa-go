// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpacalabs/mfa/v1/api.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("alpacalabs/mfa/v1/api.proto", fileDescriptor_938fcc5f3151265b) }

var fileDescriptor_938fcc5f3151265b = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0x83, 0x30,
	0x18, 0xc7, 0x1d, 0x26, 0x66, 0xe9, 0x6d, 0x4d, 0xbc, 0xe0, 0xc9, 0x4d, 0x6f, 0xa3, 0x04, 0xbd,
	0x78, 0x65, 0x1a, 0xbd, 0x8c, 0xb8, 0xb8, 0xc8, 0x41, 0x97, 0x90, 0x02, 0x5f, 0xb1, 0x09, 0xd0,
	0x4a, 0x3b, 0x92, 0x3d, 0x83, 0x6f, 0xe1, 0xd1, 0xbb, 0x2f, 0xe1, 0x53, 0x99, 0xb1, 0x21, 0x4b,
	0xba, 0xa8, 0x57, 0xbe, 0xdf, 0xff, 0xf7, 0xa7, 0xdf, 0x87, 0x4e, 0x68, 0x2e, 0x69, 0x42, 0x73,
	0x1a, 0x2b, 0xb7, 0x60, 0xd4, 0xad, 0x3d, 0x97, 0x4a, 0x4e, 0x64, 0x25, 0xb4, 0xc0, 0x83, 0x6e,
	0x48, 0x0a, 0x46, 0x49, 0xed, 0xd9, 0x63, 0x93, 0xcf, 0x40, 0x47, 0x29, 0xe4, 0xbc, 0x86, 0x6a,
	0x15, 0x09, 0xa9, 0xb9, 0x28, 0xd5, 0x46, 0x60, 0x9f, 0x9a, 0xb4, 0x82, 0x32, 0x8d, 0x12, 0x91,
	0xc2, 0x16, 0x19, 0x99, 0x48, 0x0d, 0x15, 0x67, 0xab, 0x1d, 0xe8, 0xe2, 0xd3, 0x42, 0x28, 0xb8,
	0xf5, 0xe7, 0x50, 0xd5, 0x3c, 0x01, 0xac, 0x10, 0xbe, 0x03, 0x7d, 0xb3, 0xed, 0xbc, 0xdf, 0x54,
	0xe2, 0x31, 0x31, 0x7e, 0x97, 0x98, 0xd8, 0x03, 0xbc, 0x2e, 0x41, 0x69, 0xdb, 0xf9, 0x27, 0xad,
	0xa4, 0x28, 0x15, 0x0c, 0x0f, 0xf0, 0x23, 0xea, 0xcf, 0xa1, 0x4c, 0xaf, 0x45, 0x0a, 0x78, 0xb8,
	0x27, 0xdc, 0x0e, 0xdb, 0x82, 0xd1, 0xaf, 0xcc, 0x8f, 0xf6, 0x19, 0xa1, 0xb0, 0x79, 0x6f, 0x23,
	0x3e, 0xdb, 0x13, 0xea, 0xc6, 0xad, 0xfa, 0xfc, 0x0f, 0xaa, 0x95, 0x4f, 0xde, 0x7a, 0xe8, 0x38,
	0x11, 0x85, 0x89, 0x4f, 0xfa, 0x01, 0xa3, 0xb3, 0xf5, 0x6e, 0x67, 0xbd, 0xa7, 0xab, 0x8c, 0xeb,
	0x97, 0x65, 0x4c, 0x12, 0x51, 0xb8, 0x7e, 0x43, 0x4e, 0xd7, 0xd7, 0x68, 0x56, 0x5f, 0x81, 0x14,
	0x4e, 0xc1, 0xa8, 0x93, 0x09, 0xd7, 0xb8, 0xd4, 0xbb, 0x75, 0xe8, 0x4f, 0x83, 0x0f, 0x6b, 0xd0,
	0xa5, 0x48, 0xc0, 0x28, 0x09, 0xbd, 0xaf, 0xdd, 0x6f, 0x8b, 0x80, 0xd1, 0x45, 0xe8, 0xc5, 0x47,
	0x8d, 0xf1, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x18, 0x0b, 0x0d, 0x74, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MFAServiceClient is the client API for MFAService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MFAServiceClient interface {
	// initiate the MFA flow by providing an account ID.
	// the response will contain phone numbers and email addresses
	// that can receive the client's MFA code.
	GetDeliveryOptions(ctx context.Context, in *GetDeliveryOptionsRequest, opts ...grpc.CallOption) (*GetDeliveryOptionsResponse, error)
	// delivers the code to the specified device
	SendCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeResponse, error)
	// complete the MFA flow by sending the code and account ID.
	VerifyCode(ctx context.Context, in *VerifyCodeRequest, opts ...grpc.CallOption) (*VerifyCodeResponse, error)
}

type mFAServiceClient struct {
	cc *grpc.ClientConn
}

func NewMFAServiceClient(cc *grpc.ClientConn) MFAServiceClient {
	return &mFAServiceClient{cc}
}

func (c *mFAServiceClient) GetDeliveryOptions(ctx context.Context, in *GetDeliveryOptionsRequest, opts ...grpc.CallOption) (*GetDeliveryOptionsResponse, error) {
	out := new(GetDeliveryOptionsResponse)
	err := c.cc.Invoke(ctx, "/alpacalabs.mfa.v1.MFAService/GetDeliveryOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mFAServiceClient) SendCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeResponse, error) {
	out := new(SendCodeResponse)
	err := c.cc.Invoke(ctx, "/alpacalabs.mfa.v1.MFAService/SendCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mFAServiceClient) VerifyCode(ctx context.Context, in *VerifyCodeRequest, opts ...grpc.CallOption) (*VerifyCodeResponse, error) {
	out := new(VerifyCodeResponse)
	err := c.cc.Invoke(ctx, "/alpacalabs.mfa.v1.MFAService/VerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MFAServiceServer is the server API for MFAService service.
type MFAServiceServer interface {
	// initiate the MFA flow by providing an account ID.
	// the response will contain phone numbers and email addresses
	// that can receive the client's MFA code.
	GetDeliveryOptions(context.Context, *GetDeliveryOptionsRequest) (*GetDeliveryOptionsResponse, error)
	// delivers the code to the specified device
	SendCode(context.Context, *SendCodeRequest) (*SendCodeResponse, error)
	// complete the MFA flow by sending the code and account ID.
	VerifyCode(context.Context, *VerifyCodeRequest) (*VerifyCodeResponse, error)
}

func RegisterMFAServiceServer(s *grpc.Server, srv MFAServiceServer) {
	s.RegisterService(&_MFAService_serviceDesc, srv)
}

func _MFAService_GetDeliveryOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeliveryOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MFAServiceServer).GetDeliveryOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpacalabs.mfa.v1.MFAService/GetDeliveryOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MFAServiceServer).GetDeliveryOptions(ctx, req.(*GetDeliveryOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MFAService_SendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MFAServiceServer).SendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpacalabs.mfa.v1.MFAService/SendCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MFAServiceServer).SendCode(ctx, req.(*SendCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MFAService_VerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MFAServiceServer).VerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpacalabs.mfa.v1.MFAService/VerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MFAServiceServer).VerifyCode(ctx, req.(*VerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MFAService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alpacalabs.mfa.v1.MFAService",
	HandlerType: (*MFAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeliveryOptions",
			Handler:    _MFAService_GetDeliveryOptions_Handler,
		},
		{
			MethodName: "SendCode",
			Handler:    _MFAService_SendCode_Handler,
		},
		{
			MethodName: "VerifyCode",
			Handler:    _MFAService_VerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alpacalabs/mfa/v1/api.proto",
}
