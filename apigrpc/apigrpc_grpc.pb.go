// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.1
// source: apigrpc.proto

package apigrpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SocialClient is the client API for Social service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocialClient interface {
	// Add friends by ID or username to a user's account.
	CheckGoogleToken(ctx context.Context, in *CheckGoogleTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CheckAppleToken(ctx context.Context, in *CheckAppleTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type socialClient struct {
	cc grpc.ClientConnInterface
}

func NewSocialClient(cc grpc.ClientConnInterface) SocialClient {
	return &socialClient{cc}
}

func (c *socialClient) CheckGoogleToken(ctx context.Context, in *CheckGoogleTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/socialservice.api.Social/CheckGoogleToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) CheckAppleToken(ctx context.Context, in *CheckAppleTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/socialservice.api.Social/CheckAppleToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocialServer is the server API for Social service.
// All implementations must embed UnimplementedSocialServer
// for forward compatibility
type SocialServer interface {
	// Add friends by ID or username to a user's account.
	CheckGoogleToken(context.Context, *CheckGoogleTokenRequest) (*empty.Empty, error)
	CheckAppleToken(context.Context, *CheckAppleTokenRequest) (*empty.Empty, error)
	mustEmbedUnimplementedSocialServer()
}

// UnimplementedSocialServer must be embedded to have forward compatible implementations.
type UnimplementedSocialServer struct {
}

func (UnimplementedSocialServer) CheckGoogleToken(context.Context, *CheckGoogleTokenRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckGoogleToken not implemented")
}
func (UnimplementedSocialServer) CheckAppleToken(context.Context, *CheckAppleTokenRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAppleToken not implemented")
}
func (UnimplementedSocialServer) mustEmbedUnimplementedSocialServer() {}

// UnsafeSocialServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocialServer will
// result in compilation errors.
type UnsafeSocialServer interface {
	mustEmbedUnimplementedSocialServer()
}

func RegisterSocialServer(s grpc.ServiceRegistrar, srv SocialServer) {
	s.RegisterService(&Social_ServiceDesc, srv)
}

func _Social_CheckGoogleToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckGoogleTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).CheckGoogleToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialservice.api.Social/CheckGoogleToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).CheckGoogleToken(ctx, req.(*CheckGoogleTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_CheckAppleToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAppleTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).CheckAppleToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialservice.api.Social/CheckAppleToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).CheckAppleToken(ctx, req.(*CheckAppleTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Social_ServiceDesc is the grpc.ServiceDesc for Social service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Social_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "socialservice.api.Social",
	HandlerType: (*SocialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckGoogleToken",
			Handler:    _Social_CheckGoogleToken_Handler,
		},
		{
			MethodName: "CheckAppleToken",
			Handler:    _Social_CheckAppleToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apigrpc.proto",
}
