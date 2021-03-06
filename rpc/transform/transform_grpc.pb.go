// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: transform.proto

package transform

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

// TransformServiceClient is the client API for TransformService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransformServiceClient interface {
	Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*ExpandResponse, error)
	Shorten(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error)
}

type transformServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransformServiceClient(cc grpc.ClientConnInterface) TransformServiceClient {
	return &transformServiceClient{cc}
}

func (c *transformServiceClient) Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*ExpandResponse, error) {
	out := new(ExpandResponse)
	err := c.cc.Invoke(ctx, "/transform.TransformService/Expand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transformServiceClient) Shorten(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error) {
	out := new(ShortenResponse)
	err := c.cc.Invoke(ctx, "/transform.TransformService/Shorten", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransformServiceServer is the server API for TransformService service.
// All implementations must embed UnimplementedTransformServiceServer
// for forward compatibility
type TransformServiceServer interface {
	Expand(context.Context, *ExpandRequest) (*ExpandResponse, error)
	Shorten(context.Context, *ShortenRequest) (*ShortenResponse, error)
	mustEmbedUnimplementedTransformServiceServer()
}

// UnimplementedTransformServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransformServiceServer struct {
}

func (UnimplementedTransformServiceServer) Expand(context.Context, *ExpandRequest) (*ExpandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expand not implemented")
}
func (UnimplementedTransformServiceServer) Shorten(context.Context, *ShortenRequest) (*ShortenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shorten not implemented")
}
func (UnimplementedTransformServiceServer) mustEmbedUnimplementedTransformServiceServer() {}

// UnsafeTransformServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransformServiceServer will
// result in compilation errors.
type UnsafeTransformServiceServer interface {
	mustEmbedUnimplementedTransformServiceServer()
}

func RegisterTransformServiceServer(s grpc.ServiceRegistrar, srv TransformServiceServer) {
	s.RegisterService(&TransformService_ServiceDesc, srv)
}

func _TransformService_Expand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformServiceServer).Expand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.TransformService/Expand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformServiceServer).Expand(ctx, req.(*ExpandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransformService_Shorten_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformServiceServer).Shorten(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.TransformService/Shorten",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformServiceServer).Shorten(ctx, req.(*ShortenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransformService_ServiceDesc is the grpc.ServiceDesc for TransformService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransformService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transform.TransformService",
	HandlerType: (*TransformServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Expand",
			Handler:    _TransformService_Expand_Handler,
		},
		{
			MethodName: "Shorten",
			Handler:    _TransformService_Shorten_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transform.proto",
}
