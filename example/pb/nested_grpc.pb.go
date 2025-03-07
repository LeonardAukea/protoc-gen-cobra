// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: nested.proto

package pb

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

// NestedClient is the client API for Nested service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NestedClient interface {
	Get(ctx context.Context, in *NestedRequest, opts ...grpc.CallOption) (*NestedResponse, error)
	GetDeep(ctx context.Context, in *DeepRequest, opts ...grpc.CallOption) (*NestedResponse, error)
	GetOneOf(ctx context.Context, in *OneOfRequest, opts ...grpc.CallOption) (*NestedResponse, error)
	GetOneOfDeep(ctx context.Context, in *OneOfDeepRequest, opts ...grpc.CallOption) (*NestedResponse, error)
}

type nestedClient struct {
	cc grpc.ClientConnInterface
}

func NewNestedClient(cc grpc.ClientConnInterface) NestedClient {
	return &nestedClient{cc}
}

func (c *nestedClient) Get(ctx context.Context, in *NestedRequest, opts ...grpc.CallOption) (*NestedResponse, error) {
	out := new(NestedResponse)
	err := c.cc.Invoke(ctx, "/example.Nested/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nestedClient) GetDeep(ctx context.Context, in *DeepRequest, opts ...grpc.CallOption) (*NestedResponse, error) {
	out := new(NestedResponse)
	err := c.cc.Invoke(ctx, "/example.Nested/GetDeep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nestedClient) GetOneOf(ctx context.Context, in *OneOfRequest, opts ...grpc.CallOption) (*NestedResponse, error) {
	out := new(NestedResponse)
	err := c.cc.Invoke(ctx, "/example.Nested/GetOneOf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nestedClient) GetOneOfDeep(ctx context.Context, in *OneOfDeepRequest, opts ...grpc.CallOption) (*NestedResponse, error) {
	out := new(NestedResponse)
	err := c.cc.Invoke(ctx, "/example.Nested/GetOneOfDeep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NestedServer is the server API for Nested service.
// All implementations must embed UnimplementedNestedServer
// for forward compatibility
type NestedServer interface {
	Get(context.Context, *NestedRequest) (*NestedResponse, error)
	GetDeep(context.Context, *DeepRequest) (*NestedResponse, error)
	GetOneOf(context.Context, *OneOfRequest) (*NestedResponse, error)
	GetOneOfDeep(context.Context, *OneOfDeepRequest) (*NestedResponse, error)
	mustEmbedUnimplementedNestedServer()
}

// UnimplementedNestedServer must be embedded to have forward compatible implementations.
type UnimplementedNestedServer struct {
}

func (UnimplementedNestedServer) Get(context.Context, *NestedRequest) (*NestedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNestedServer) GetDeep(context.Context, *DeepRequest) (*NestedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeep not implemented")
}
func (UnimplementedNestedServer) GetOneOf(context.Context, *OneOfRequest) (*NestedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneOf not implemented")
}
func (UnimplementedNestedServer) GetOneOfDeep(context.Context, *OneOfDeepRequest) (*NestedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneOfDeep not implemented")
}
func (UnimplementedNestedServer) mustEmbedUnimplementedNestedServer() {}

// UnsafeNestedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NestedServer will
// result in compilation errors.
type UnsafeNestedServer interface {
	mustEmbedUnimplementedNestedServer()
}

func RegisterNestedServer(s grpc.ServiceRegistrar, srv NestedServer) {
	s.RegisterService(&Nested_ServiceDesc, srv)
}

func _Nested_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NestedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NestedServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Nested/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NestedServer).Get(ctx, req.(*NestedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nested_GetDeep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NestedServer).GetDeep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Nested/GetDeep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NestedServer).GetDeep(ctx, req.(*DeepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nested_GetOneOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneOfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NestedServer).GetOneOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Nested/GetOneOf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NestedServer).GetOneOf(ctx, req.(*OneOfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nested_GetOneOfDeep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneOfDeepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NestedServer).GetOneOfDeep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Nested/GetOneOfDeep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NestedServer).GetOneOfDeep(ctx, req.(*OneOfDeepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Nested_ServiceDesc is the grpc.ServiceDesc for Nested service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nested_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.Nested",
	HandlerType: (*NestedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Nested_Get_Handler,
		},
		{
			MethodName: "GetDeep",
			Handler:    _Nested_GetDeep_Handler,
		},
		{
			MethodName: "GetOneOf",
			Handler:    _Nested_GetOneOf_Handler,
		},
		{
			MethodName: "GetOneOfDeep",
			Handler:    _Nested_GetOneOfDeep_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nested.proto",
}
