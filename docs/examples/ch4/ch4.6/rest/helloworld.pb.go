// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StringMessage struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMessage) Reset()         { *m = StringMessage{} }
func (m *StringMessage) String() string { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()    {}
func (*StringMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_b65656966bdaa93a, []int{0}
}
func (m *StringMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMessage.Unmarshal(m, b)
}
func (m *StringMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMessage.Marshal(b, m, deterministic)
}
func (dst *StringMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMessage.Merge(dst, src)
}
func (m *StringMessage) XXX_Size() int {
	return xxx_messageInfo_StringMessage.Size(m)
}
func (m *StringMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StringMessage proto.InternalMessageInfo

func (m *StringMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "main.StringMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RestServiceClient is the client API for RestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RestServiceClient interface {
	Get(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
	Post(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type restServiceClient struct {
	cc *grpc.ClientConn
}

func NewRestServiceClient(cc *grpc.ClientConn) RestServiceClient {
	return &restServiceClient{cc}
}

func (c *restServiceClient) Get(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/main.RestService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restServiceClient) Post(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/main.RestService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestServiceServer is the server API for RestService service.
type RestServiceServer interface {
	Get(context.Context, *StringMessage) (*StringMessage, error)
	Post(context.Context, *StringMessage) (*StringMessage, error)
}

func RegisterRestServiceServer(s *grpc.Server, srv RestServiceServer) {
	s.RegisterService(&_RestService_serviceDesc, srv)
}

func _RestService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.RestService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestServiceServer).Get(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.RestService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestServiceServer).Post(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _RestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.RestService",
	HandlerType: (*RestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RestService_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _RestService_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_helloworld_b65656966bdaa93a) }

var fileDescriptor_helloworld_b65656966bdaa93a = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d,
	0xcc, 0xcc, 0x93, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f,
	0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x51, 0x52, 0xe5, 0xe2,
	0x0d, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe1,
	0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x8c,
	0x66, 0x30, 0x72, 0x71, 0x07, 0xa5, 0x16, 0x97, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a,
	0xb9, 0x72, 0x31, 0xbb, 0xa7, 0x96, 0x08, 0x09, 0xeb, 0x81, 0xac, 0xd0, 0x43, 0x31, 0x41, 0x0a,
	0x9b, 0xa0, 0x92, 0x48, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xf8, 0x84, 0x78, 0xf4, 0xd3, 0x53, 0x4b,
	0xf4, 0xab, 0xc1, 0xa6, 0xd6, 0x0a, 0x39, 0x71, 0xb1, 0x04, 0xe4, 0x17, 0x93, 0x62, 0x8e, 0x00,
	0xd8, 0x1c, 0x2e, 0x25, 0x56, 0xfd, 0x82, 0xfc, 0xe2, 0x12, 0x2b, 0x46, 0xad, 0x24, 0x36, 0xb0,
	0x47, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x75, 0x88, 0x64, 0x8b, 0x00, 0x01, 0x00, 0x00,
}