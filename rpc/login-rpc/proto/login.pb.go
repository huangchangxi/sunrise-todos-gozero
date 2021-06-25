// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	Code                 int64      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *UserToken `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetData() *UserToken {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpiresIn            int32    `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserToken) Reset()         { *m = UserToken{} }
func (m *UserToken) String() string { return proto.CompactTextString(m) }
func (*UserToken) ProtoMessage()    {}
func (*UserToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{2}
}

func (m *UserToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserToken.Unmarshal(m, b)
}
func (m *UserToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserToken.Marshal(b, m, deterministic)
}
func (m *UserToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserToken.Merge(m, src)
}
func (m *UserToken) XXX_Size() int {
	return xxx_messageInfo_UserToken.Size(m)
}
func (m *UserToken) XXX_DiscardUnknown() {
	xxx_messageInfo_UserToken.DiscardUnknown(m)
}

var xxx_messageInfo_UserToken proto.InternalMessageInfo

func (m *UserToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserToken) GetExpiresIn() int32 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "login.Request")
	proto.RegisterType((*Response)(nil), "login.Response")
	proto.RegisterType((*UserToken)(nil), "login.UserToken")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0x69, 0xda, 0x66, 0x0a, 0x5a, 0x06, 0x0f, 0xa1, 0x20, 0x94, 0xe0, 0xa1, 0x5e,
	0x22, 0xc4, 0x2f, 0xa0, 0xde, 0x04, 0x4f, 0xeb, 0x9f, 0x83, 0x17, 0x49, 0xcd, 0x10, 0x16, 0xcd,
	0xee, 0xba, 0xb3, 0x41, 0x3f, 0xbe, 0xec, 0x1f, 0x73, 0xda, 0xf7, 0x9b, 0xc7, 0xec, 0xcc, 0x1b,
	0xd8, 0x7c, 0xe9, 0x41, 0xaa, 0xc6, 0x58, 0xed, 0x34, 0x16, 0x01, 0xea, 0x3b, 0x58, 0x09, 0xfa,
	0x9e, 0x88, 0x1d, 0xee, 0x60, 0x3d, 0x31, 0x59, 0xd5, 0x8d, 0x54, 0x65, 0xfb, 0xec, 0x50, 0x8a,
	0x99, 0xbd, 0x67, 0x3a, 0xe6, 0x1f, 0x6d, 0xfb, 0xea, 0x24, 0x7a, 0xff, 0x5c, 0xbf, 0xc2, 0x5a,
	0x10, 0x1b, 0xad, 0x98, 0x10, 0x61, 0xf1, 0xa1, 0xfb, 0xd8, 0x9f, 0x8b, 0xa0, 0x71, 0x0b, 0xf9,
	0xc8, 0x43, 0x6a, 0xf3, 0x12, 0x2f, 0x61, 0xd1, 0x77, 0xae, 0xab, 0xf2, 0x7d, 0x76, 0xd8, 0xb4,
	0xdb, 0x26, 0xee, 0xf5, 0xc2, 0x64, 0x9f, 0xf5, 0x27, 0x29, 0x11, 0xdc, 0xfa, 0x16, 0xca, 0xb9,
	0x84, 0xe7, 0x50, 0x38, 0x2f, 0xd2, 0x66, 0x11, 0xf0, 0x02, 0x80, 0x7e, 0x8d, 0xb4, 0xc4, 0xef,
	0x52, 0x85, 0x09, 0x85, 0x28, 0x53, 0xe5, 0x41, 0xb5, 0x2d, 0x14, 0x8f, 0xfe, 0x6b, 0xbc, 0x82,
	0xe5, 0xd3, 0x74, 0x1c, 0xa5, 0xc3, 0xd3, 0x34, 0x2c, 0x85, 0xde, 0x9d, 0xcd, 0x1c, 0x13, 0xdc,
	0x97, 0x6f, 0xab, 0xe6, 0x3a, 0x9c, 0xe8, 0xb8, 0x0c, 0xcf, 0xcd, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x17, 0xd1, 0x17, 0xb7, 0x38, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginClient interface {
	Submit(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type loginClient struct {
	cc *grpc.ClientConn
}

func NewLoginClient(cc *grpc.ClientConn) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Submit(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/login.Login/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
type LoginServer interface {
	Submit(context.Context, *Request) (*Response, error)
}

// UnimplementedLoginServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (*UnimplementedLoginServer) Submit(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.Login/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Submit(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "login.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _Login_Submit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login.proto",
}
