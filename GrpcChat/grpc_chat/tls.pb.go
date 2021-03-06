// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tls.proto

package grpc_chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TlsService service

type TlsServiceClient interface {
	SendMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Msg, error)
}

type tlsServiceClient struct {
	cc *grpc.ClientConn
}

func NewTlsServiceClient(cc *grpc.ClientConn) TlsServiceClient {
	return &tlsServiceClient{cc}
}

func (c *tlsServiceClient) SendMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := grpc.Invoke(ctx, "/grpc_chat.TlsService/SendMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TlsService service

type TlsServiceServer interface {
	SendMsg(context.Context, *Msg) (*Msg, error)
}

func RegisterTlsServiceServer(s *grpc.Server, srv TlsServiceServer) {
	s.RegisterService(&_TlsService_serviceDesc, srv)
}

func _TlsService_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TlsServiceServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_chat.TlsService/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TlsServiceServer).SendMsg(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

var _TlsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_chat.TlsService",
	HandlerType: (*TlsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _TlsService_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tls.proto",
}

func init() { proto.RegisterFile("tls.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xc9, 0x29, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x2f, 0x2a, 0x48, 0x8e, 0x4f, 0xce, 0x48, 0x2c,
	0x91, 0xe2, 0x87, 0x33, 0x21, 0x72, 0x46, 0xe6, 0x5c, 0x5c, 0x21, 0x39, 0xc5, 0xc1, 0xa9, 0x45,
	0x65, 0x99, 0xc9, 0xa9, 0x42, 0x9a, 0x5c, 0xec, 0xc1, 0xa9, 0x79, 0x29, 0xbe, 0xc5, 0xe9, 0x42,
	0x7c, 0x7a, 0x08, 0xa5, 0xbe, 0xc5, 0xe9, 0x52, 0x68, 0x7c, 0x27, 0x0d, 0x2e, 0xa9, 0xcc, 0x7c,
	0xb0, 0x98, 0x5e, 0x6a, 0x45, 0x62, 0x6e, 0x41, 0x4e, 0x6a, 0x31, 0x42, 0x85, 0x13, 0x87, 0x7b,
	0x51, 0x41, 0xb2, 0x73, 0x46, 0x62, 0x49, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0x26, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xaa, 0x44, 0xbc, 0x05, 0x92, 0x00, 0x00, 0x00,
}
