// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package userservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	GetUserTokenByCredentials(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*UserToken, error)
	GetUserTokenByEmail(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserToken, error)
	CreteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserState, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) GetUserTokenByCredentials(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*UserToken, error) {
	out := new(UserToken)
	err := c.cc.Invoke(ctx, "/users.Users/GetUserTokenByCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserTokenByEmail(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserToken, error) {
	out := new(UserToken)
	err := c.cc.Invoke(ctx, "/users.Users/GetUserTokenByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CreteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserState, error) {
	out := new(UserState)
	err := c.cc.Invoke(ctx, "/users.Users/CreteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	GetUserTokenByCredentials(context.Context, *Credentials) (*UserToken, error)
	GetUserTokenByEmail(context.Context, *UserInfo) (*UserToken, error)
	CreteUser(context.Context, *User) (*UserState, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) GetUserTokenByCredentials(context.Context, *Credentials) (*UserToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTokenByCredentials not implemented")
}
func (*UnimplementedUsersServer) GetUserTokenByEmail(context.Context, *UserInfo) (*UserToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTokenByEmail not implemented")
}
func (*UnimplementedUsersServer) CreteUser(context.Context, *User) (*UserState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreteUser not implemented")
}
func (*UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_GetUserTokenByCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserTokenByCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/GetUserTokenByCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserTokenByCredentials(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserTokenByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserTokenByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/GetUserTokenByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserTokenByEmail(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_CreteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/CreteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserTokenByCredentials",
			Handler:    _Users_GetUserTokenByCredentials_Handler,
		},
		{
			MethodName: "GetUserTokenByEmail",
			Handler:    _Users_GetUserTokenByEmail_Handler,
		},
		{
			MethodName: "CreteUser",
			Handler:    _Users_CreteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}