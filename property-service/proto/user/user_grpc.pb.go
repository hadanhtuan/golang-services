// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/user/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	sdk "property-service/proto/sdk"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Login(ctx context.Context, in *MsgUser, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	UpdateUser(ctx context.Context, in *MsgUser, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	GetUsers(ctx context.Context, in *MsgQueryUser, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	GetUsersByIds(ctx context.Context, in *MsgQueryUserByIds, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	Register(ctx context.Context, in *MsgUser, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	RefreshToken(ctx context.Context, in *MsgToken, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	Logout(ctx context.Context, in *MsgUser, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	GetProfile(ctx context.Context, in *MsgID, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
	VerifyToken(ctx context.Context, in *MsgToken, opts ...grpc.CallOption) (*sdk.BaseResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *MsgUser, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/userService.userService/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *MsgUser, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/userService.userService/updateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *MsgQueryUser, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/userService.userService/getUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsersByIds(ctx context.Context, in *MsgQueryUserByIds, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/userService.userService/getUsersByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Register(ctx context.Context, in *MsgUser, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/userService.userService/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RefreshToken(ctx context.Context, in *MsgToken, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/userService.userService/refreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Logout(ctx context.Context, in *MsgUser, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/userService.userService/logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetProfile(ctx context.Context, in *MsgID, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/userService.userService/getProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerifyToken(ctx context.Context, in *MsgToken, opts ...grpc.CallOption) (*sdk.BaseResponse, error) {
	out := new(sdk.BaseResponse)
	err := c.cc.Invoke(ctx, "/userService.userService/verifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Login(context.Context, *MsgUser) (*sdk.BaseResponse, error)
	UpdateUser(context.Context, *MsgUser) (*sdk.BaseResponse, error)
	GetUsers(context.Context, *MsgQueryUser) (*sdk.BaseResponse, error)
	GetUsersByIds(context.Context, *MsgQueryUserByIds) (*sdk.BaseResponse, error)
	Register(context.Context, *MsgUser) (*sdk.BaseResponse, error)
	RefreshToken(context.Context, *MsgToken) (*sdk.BaseResponse, error)
	Logout(context.Context, *MsgUser) (*sdk.BaseResponse, error)
	GetProfile(context.Context, *MsgID) (*sdk.BaseResponse, error)
	VerifyToken(context.Context, *MsgToken) (*sdk.BaseResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Login(context.Context, *MsgUser) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *MsgUser) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUsers(context.Context, *MsgQueryUser) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserServiceServer) GetUsersByIds(context.Context, *MsgQueryUserByIds) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByIds not implemented")
}
func (UnimplementedUserServiceServer) Register(context.Context, *MsgUser) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) RefreshToken(context.Context, *MsgToken) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedUserServiceServer) Logout(context.Context, *MsgUser) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServiceServer) GetProfile(context.Context, *MsgID) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedUserServiceServer) VerifyToken(context.Context, *MsgToken) (*sdk.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userService.userService/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*MsgUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userService.userService/updateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*MsgUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgQueryUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userService.userService/getUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*MsgQueryUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsersByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgQueryUserByIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsersByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userService.userService/getUsersByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsersByIds(ctx, req.(*MsgQueryUserByIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userService.userService/register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*MsgUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userService.userService/refreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RefreshToken(ctx, req.(*MsgToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userService.userService/logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Logout(ctx, req.(*MsgUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userService.userService/getProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetProfile(ctx, req.(*MsgID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userService.userService/verifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyToken(ctx, req.(*MsgToken))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userService.userService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "updateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "getUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
		{
			MethodName: "getUsersByIds",
			Handler:    _UserService_GetUsersByIds_Handler,
		},
		{
			MethodName: "register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "refreshToken",
			Handler:    _UserService_RefreshToken_Handler,
		},
		{
			MethodName: "logout",
			Handler:    _UserService_Logout_Handler,
		},
		{
			MethodName: "getProfile",
			Handler:    _UserService_GetProfile_Handler,
		},
		{
			MethodName: "verifyToken",
			Handler:    _UserService_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user/user.proto",
}
