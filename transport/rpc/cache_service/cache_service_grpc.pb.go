// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: transport/rpc/cache_service/cache_service.proto

package cacheService

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

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheServiceClient interface {
	GetValue(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	SetValue(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	SetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*SetUserResponse, error)
}

type cacheServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheServiceClient(cc grpc.ClientConnInterface) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) GetValue(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/CacheService/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) SetValue(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/CacheService/SetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/CacheService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) SetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*SetUserResponse, error) {
	out := new(SetUserResponse)
	err := c.cc.Invoke(ctx, "/CacheService/SetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
// All implementations must embed UnimplementedCacheServiceServer
// for forward compatibility
type CacheServiceServer interface {
	GetValue(context.Context, *GetRequest) (*GetResponse, error)
	SetValue(context.Context, *SetRequest) (*SetResponse, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	SetUser(context.Context, *User) (*SetUserResponse, error)
	mustEmbedUnimplementedCacheServiceServer()
}

// UnimplementedCacheServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCacheServiceServer struct {
}

func (UnimplementedCacheServiceServer) GetValue(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedCacheServiceServer) SetValue(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValue not implemented")
}
func (UnimplementedCacheServiceServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedCacheServiceServer) SetUser(context.Context, *User) (*SetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUser not implemented")
}
func (UnimplementedCacheServiceServer) mustEmbedUnimplementedCacheServiceServer() {}

// UnsafeCacheServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheServiceServer will
// result in compilation errors.
type UnsafeCacheServiceServer interface {
	mustEmbedUnimplementedCacheServiceServer()
}

func RegisterCacheServiceServer(s grpc.ServiceRegistrar, srv CacheServiceServer) {
	s.RegisterService(&CacheService_ServiceDesc, srv)
}

func _CacheService_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CacheService/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetValue(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CacheService/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).SetValue(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CacheService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_SetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).SetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CacheService/SetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).SetUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheService_ServiceDesc is the grpc.ServiceDesc for CacheService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValue",
			Handler:    _CacheService_GetValue_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _CacheService_SetValue_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _CacheService_GetUser_Handler,
		},
		{
			MethodName: "SetUser",
			Handler:    _CacheService_SetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport/rpc/cache_service/cache_service.proto",
}
