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
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0xf2, 0x40,
	0x14, 0x85, 0x7f, 0x20, 0xf9, 0x63, 0x86, 0x15, 0x4d, 0xdc, 0xe0, 0x4e, 0xc1, 0xb8, 0x80, 0x69,
	0xaa, 0x1b, 0xb7, 0xa0, 0xd1, 0x0d, 0x8d, 0x44, 0x13, 0x16, 0x4a, 0x68, 0x2e, 0xe5, 0x0e, 0x4e,
	0x42, 0x99, 0x3a, 0x33, 0x34, 0xe1, 0x19, 0x7c, 0x07, 0x17, 0x2e, 0x7d, 0x14, 0x9f, 0xca, 0x74,
	0x1c, 0xa0, 0x66, 0x1a, 0x71, 0xd7, 0xf4, 0x7c, 0xe7, 0xdc, 0x39, 0x73, 0x87, 0x1c, 0xc1, 0x22,
	0x85, 0x18, 0x16, 0x30, 0x55, 0x7e, 0xc2, 0xc0, 0xcf, 0x02, 0x1f, 0x52, 0x4e, 0x53, 0x29, 0xb4,
	0xf0, 0x1a, 0x3b, 0x91, 0x26, 0x0c, 0x68, 0x16, 0x34, 0xcf, 0x4a, 0xf9, 0x48, 0xe2, 0xcb, 0x8a,
	0x4b, 0x54, 0x51, 0x8e, 0x19, 0x73, 0xb3, 0xe3, 0x92, 0x73, 0xd4, 0xd1, 0x0c, 0x17, 0x3c, 0x43,
	0xb9, 0x8e, 0x44, 0xaa, 0xb9, 0x58, 0x2a, 0x4b, 0xb7, 0x5c, 0xda, 0x92, 0x51, 0x2c, 0x66, 0x68,
	0xa9, 0x13, 0x97, 0xca, 0x50, 0x72, 0xb6, 0x2e, 0x40, 0xe7, 0x6f, 0x35, 0x42, 0xc2, 0x9b, 0xde,
	0x03, 0xca, 0x8c, 0xc7, 0xe8, 0x4d, 0x48, 0xfd, 0xde, 0x9e, 0x2e, 0x64, 0xe0, 0xb5, 0xa9, 0x53,
	0x8a, 0x16, 0xf4, 0xfc, 0x13, 0x95, 0x6e, 0x9e, 0xee, 0xc3, 0x54, 0x2a, 0x96, 0x0a, 0x8f, 0xff,
	0x79, 0x8a, 0x78, 0xb7, 0xa8, 0xaf, 0x6d, 0xad, 0xbb, 0xef, 0x56, 0x5e, 0xa7, 0xc4, 0xef, 0x62,
	0x9b, 0x69, 0xdd, 0x3f, 0xd2, 0xdb, 0xa1, 0x13, 0x52, 0xb7, 0xe2, 0x95, 0x98, 0x61, 0x69, 0xa9,
	0x82, 0xfe, 0x5b, 0xa9, 0x1f, 0xd8, 0x36, 0xff, 0x89, 0x90, 0x91, 0xb9, 0x58, 0x13, 0xdf, 0x2a,
	0xf1, 0xed, 0xe4, 0x4d, 0x7a, 0x7b, 0x0f, 0xb5, 0x09, 0xef, 0xbf, 0x56, 0xc8, 0x61, 0x2c, 0x12,
	0x17, 0xef, 0x1f, 0x84, 0x0c, 0x86, 0xf9, 0x12, 0x87, 0x95, 0xc7, 0xcb, 0x39, 0xd7, 0xcf, 0xab,
	0x29, 0x8d, 0x45, 0xe2, 0xf7, 0x0c, 0x39, 0xc8, 0xd7, 0x6e, 0x76, 0x2c, 0x31, 0x15, 0xdd, 0x84,
	0x41, 0x77, 0x2e, 0x7c, 0xe7, 0x49, 0xbc, 0x57, 0x6b, 0xbd, 0x41, 0xf8, 0x51, 0x6d, 0xec, 0x5c,
	0x34, 0x64, 0x40, 0x47, 0xc1, 0x67, 0xf1, 0xdf, 0x38, 0x64, 0x30, 0x1e, 0x05, 0xd3, 0xff, 0x26,
	0xf1, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xc6, 0x68, 0x31, 0x0a, 0x03, 0x00, 0x00,
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
	RequiresMfa(ctx context.Context, in *RequiresMfaRequest, opts ...grpc.CallOption) (*RequiresMfaResponse, error)
	// GetDeliveryOptions is Step 1 in the MFA flow.
	//
	// In exchange for some account identifier (an email, username, or phone number)
	// clients will receive their devices that can receive codes -- for example, any
	// of their confirmed email addresses and phone numbers -- so they can choose how
	// they want to receive their code.
	//
	// Behind the scenes, this will persist a code in the DB. The response will
	// contain that entity's primary key and delivery options for the client.
	GetDeliveryOptions(ctx context.Context, in *GetDeliveryOptionsRequest, opts ...grpc.CallOption) (*GetDeliveryOptionsResponse, error)
	// DeliverCode is Step 2 in the MFA flow.
	//
	// The client sends the primary key of their code, as well as the primary key of
	// an email address or phone number that will receive the code.
	DeliverCode(ctx context.Context, in *DeliverCodeRequest, opts ...grpc.CallOption) (*DeliverCodeResponse, error)
	// VerifyCode is Step 3 in the MFA flow.
	//
	// The client sends the random code they received and the primary key of their account.
	// If the code is valid, they're returned a JWT.
	VerifyCode(ctx context.Context, in *VerifyCodeRequest, opts ...grpc.CallOption) (*VerifyCodeResponse, error)
}

type mFAServiceClient struct {
	cc *grpc.ClientConn
}

func NewMFAServiceClient(cc *grpc.ClientConn) MFAServiceClient {
	return &mFAServiceClient{cc}
}

func (c *mFAServiceClient) RequiresMfa(ctx context.Context, in *RequiresMfaRequest, opts ...grpc.CallOption) (*RequiresMfaResponse, error) {
	out := new(RequiresMfaResponse)
	err := c.cc.Invoke(ctx, "/alpacalabs.mfa.v1.MFAService/RequiresMfa", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mFAServiceClient) GetDeliveryOptions(ctx context.Context, in *GetDeliveryOptionsRequest, opts ...grpc.CallOption) (*GetDeliveryOptionsResponse, error) {
	out := new(GetDeliveryOptionsResponse)
	err := c.cc.Invoke(ctx, "/alpacalabs.mfa.v1.MFAService/GetDeliveryOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mFAServiceClient) DeliverCode(ctx context.Context, in *DeliverCodeRequest, opts ...grpc.CallOption) (*DeliverCodeResponse, error) {
	out := new(DeliverCodeResponse)
	err := c.cc.Invoke(ctx, "/alpacalabs.mfa.v1.MFAService/DeliverCode", in, out, opts...)
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
	RequiresMfa(context.Context, *RequiresMfaRequest) (*RequiresMfaResponse, error)
	// GetDeliveryOptions is Step 1 in the MFA flow.
	//
	// In exchange for some account identifier (an email, username, or phone number)
	// clients will receive their devices that can receive codes -- for example, any
	// of their confirmed email addresses and phone numbers -- so they can choose how
	// they want to receive their code.
	//
	// Behind the scenes, this will persist a code in the DB. The response will
	// contain that entity's primary key and delivery options for the client.
	GetDeliveryOptions(context.Context, *GetDeliveryOptionsRequest) (*GetDeliveryOptionsResponse, error)
	// DeliverCode is Step 2 in the MFA flow.
	//
	// The client sends the primary key of their code, as well as the primary key of
	// an email address or phone number that will receive the code.
	DeliverCode(context.Context, *DeliverCodeRequest) (*DeliverCodeResponse, error)
	// VerifyCode is Step 3 in the MFA flow.
	//
	// The client sends the random code they received and the primary key of their account.
	// If the code is valid, they're returned a JWT.
	VerifyCode(context.Context, *VerifyCodeRequest) (*VerifyCodeResponse, error)
}

func RegisterMFAServiceServer(s *grpc.Server, srv MFAServiceServer) {
	s.RegisterService(&_MFAService_serviceDesc, srv)
}

func _MFAService_RequiresMfa_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequiresMfaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MFAServiceServer).RequiresMfa(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpacalabs.mfa.v1.MFAService/RequiresMfa",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MFAServiceServer).RequiresMfa(ctx, req.(*RequiresMfaRequest))
	}
	return interceptor(ctx, in, info, handler)
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

func _MFAService_DeliverCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MFAServiceServer).DeliverCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpacalabs.mfa.v1.MFAService/DeliverCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MFAServiceServer).DeliverCode(ctx, req.(*DeliverCodeRequest))
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
			MethodName: "RequiresMfa",
			Handler:    _MFAService_RequiresMfa_Handler,
		},
		{
			MethodName: "GetDeliveryOptions",
			Handler:    _MFAService_GetDeliveryOptions_Handler,
		},
		{
			MethodName: "DeliverCode",
			Handler:    _MFAService_DeliverCode_Handler,
		},
		{
			MethodName: "VerifyCode",
			Handler:    _MFAService_VerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alpacalabs/mfa/v1/api.proto",
}
