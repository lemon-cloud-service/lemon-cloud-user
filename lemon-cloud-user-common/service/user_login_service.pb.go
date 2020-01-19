// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/user_login_service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	dto "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common/dto"
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

func init() { proto.RegisterFile("service/user_login_service.proto", fileDescriptor_033eae90ef509fe2) }

var fileDescriptor_033eae90ef509fe2 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a, 0xcf, 0xc9, 0x4f, 0xcf, 0xcc, 0x8b, 0x87,
	0x0a, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xb1, 0x43, 0xb9, 0x52, 0x22, 0x29, 0x25, 0xf9,
	0xfa, 0xa9, 0x79, 0x45, 0xf9, 0x39, 0x39, 0xf1, 0x29, 0x25, 0xf9, 0x10, 0x69, 0x23, 0x6f, 0x2e,
	0x81, 0xd0, 0xe2, 0xd4, 0x22, 0x1f, 0x90, 0xce, 0x60, 0x88, 0x4a, 0x21, 0x73, 0x2e, 0x5e, 0x30,
	0xdf, 0xa9, 0xd2, 0xaf, 0x34, 0x37, 0x29, 0xb5, 0x48, 0x48, 0x54, 0x0f, 0xa4, 0x01, 0x45, 0x2c,
	0x28, 0xb5, 0x50, 0x8a, 0x17, 0x21, 0x1c, 0x54, 0x5c, 0xa0, 0xc4, 0xe0, 0x14, 0x18, 0xe5, 0x9f,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x93, 0x9a, 0x9b, 0x9f, 0xa7,
	0x9b, 0x9c, 0x93, 0x5f, 0x9a, 0xa2, 0x0b, 0x73, 0x26, 0xb2, 0x18, 0xc8, 0xc9, 0x18, 0x02, 0xba,
	0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0xfa, 0x50, 0x0d, 0x49, 0x6c, 0x60, 0x67, 0x1a, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x79, 0x07, 0xf0, 0xd5, 0xe9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserLoginServiceClient is the client API for UserLoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserLoginServiceClient interface {
	LoginByNumber(ctx context.Context, in *dto.LoginByNumberReq, opts ...grpc.CallOption) (*dto.LoginRsp, error)
}

type userLoginServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserLoginServiceClient(cc *grpc.ClientConn) UserLoginServiceClient {
	return &userLoginServiceClient{cc}
}

func (c *userLoginServiceClient) LoginByNumber(ctx context.Context, in *dto.LoginByNumberReq, opts ...grpc.CallOption) (*dto.LoginRsp, error) {
	out := new(dto.LoginRsp)
	err := c.cc.Invoke(ctx, "/service.UserLoginService/LoginByNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLoginServiceServer is the server API for UserLoginService service.
type UserLoginServiceServer interface {
	LoginByNumber(context.Context, *dto.LoginByNumberReq) (*dto.LoginRsp, error)
}

// UnimplementedUserLoginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserLoginServiceServer struct {
}

func (*UnimplementedUserLoginServiceServer) LoginByNumber(ctx context.Context, req *dto.LoginByNumberReq) (*dto.LoginRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByNumber not implemented")
}

func RegisterUserLoginServiceServer(s *grpc.Server, srv UserLoginServiceServer) {
	s.RegisterService(&_UserLoginService_serviceDesc, srv)
}

func _UserLoginService_LoginByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto.LoginByNumberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).LoginByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserLoginService/LoginByNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).LoginByNumber(ctx, req.(*dto.LoginByNumberReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserLoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.UserLoginService",
	HandlerType: (*UserLoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginByNumber",
			Handler:    _UserLoginService_LoginByNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/user_login_service.proto",
}
