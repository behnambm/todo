// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protobuf/auth.proto

package authservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	IsTokenValid(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error)
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenReply, error)
	ValidateTokenWithClaims(ctx context.Context, in *ValidateTokenWithClaimsRequest, opts ...grpc.CallOption) (*ValidateTokenWithClaimsReply, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) IsTokenValid(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error) {
	out := new(TokenReply)
	err := c.cc.Invoke(ctx, "/Auth/IsTokenValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenReply, error) {
	out := new(GetTokenReply)
	err := c.cc.Invoke(ctx, "/Auth/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidateTokenWithClaims(ctx context.Context, in *ValidateTokenWithClaimsRequest, opts ...grpc.CallOption) (*ValidateTokenWithClaimsReply, error) {
	out := new(ValidateTokenWithClaimsReply)
	err := c.cc.Invoke(ctx, "/Auth/ValidateTokenWithClaims", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	IsTokenValid(context.Context, *TokenRequest) (*TokenReply, error)
	GetToken(context.Context, *GetTokenRequest) (*GetTokenReply, error)
	ValidateTokenWithClaims(context.Context, *ValidateTokenWithClaimsRequest) (*ValidateTokenWithClaimsReply, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) IsTokenValid(context.Context, *TokenRequest) (*TokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsTokenValid not implemented")
}
func (UnimplementedAuthServer) GetToken(context.Context, *GetTokenRequest) (*GetTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedAuthServer) ValidateTokenWithClaims(context.Context, *ValidateTokenWithClaimsRequest) (*ValidateTokenWithClaimsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateTokenWithClaims not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_IsTokenValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).IsTokenValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/IsTokenValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).IsTokenValid(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ValidateTokenWithClaims_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenWithClaimsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ValidateTokenWithClaims(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/ValidateTokenWithClaims",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ValidateTokenWithClaims(ctx, req.(*ValidateTokenWithClaimsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsTokenValid",
			Handler:    _Auth_IsTokenValid_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _Auth_GetToken_Handler,
		},
		{
			MethodName: "ValidateTokenWithClaims",
			Handler:    _Auth_ValidateTokenWithClaims_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/auth.proto",
}
