// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: cyclical.proto

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

// CyclicalClient is the client API for Cyclical service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CyclicalClient interface {
	Test(ctx context.Context, in *Foo, opts ...grpc.CallOption) (*Bar, error)
}

type cyclicalClient struct {
	cc grpc.ClientConnInterface
}

func NewCyclicalClient(cc grpc.ClientConnInterface) CyclicalClient {
	return &cyclicalClient{cc}
}

func (c *cyclicalClient) Test(ctx context.Context, in *Foo, opts ...grpc.CallOption) (*Bar, error) {
	out := new(Bar)
	err := c.cc.Invoke(ctx, "/example.Cyclical/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CyclicalServer is the server API for Cyclical service.
// All implementations must embed UnimplementedCyclicalServer
// for forward compatibility
type CyclicalServer interface {
	Test(context.Context, *Foo) (*Bar, error)
	mustEmbedUnimplementedCyclicalServer()
}

// UnimplementedCyclicalServer must be embedded to have forward compatible implementations.
type UnimplementedCyclicalServer struct {
}

func (UnimplementedCyclicalServer) Test(context.Context, *Foo) (*Bar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedCyclicalServer) mustEmbedUnimplementedCyclicalServer() {}

// UnsafeCyclicalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CyclicalServer will
// result in compilation errors.
type UnsafeCyclicalServer interface {
	mustEmbedUnimplementedCyclicalServer()
}

func RegisterCyclicalServer(s grpc.ServiceRegistrar, srv CyclicalServer) {
	s.RegisterService(&Cyclical_ServiceDesc, srv)
}

func _Cyclical_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Foo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CyclicalServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Cyclical/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CyclicalServer).Test(ctx, req.(*Foo))
	}
	return interceptor(ctx, in, info, handler)
}

// Cyclical_ServiceDesc is the grpc.ServiceDesc for Cyclical service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cyclical_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.Cyclical",
	HandlerType: (*CyclicalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _Cyclical_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cyclical.proto",
}
