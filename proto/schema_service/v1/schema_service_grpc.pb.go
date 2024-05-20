// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: schema_service/v1/schema_service.proto

package schemav1

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

const (
	SchemaService_Register_FullMethodName = "/schema.v1.SchemaService/Register"
	SchemaService_Fetch_FullMethodName    = "/schema.v1.SchemaService/Fetch"
)

// SchemaServiceClient is the client API for SchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchemaServiceClient interface {
	Register(ctx context.Context, in *RegisterSchemaRequest, opts ...grpc.CallOption) (*RegisterSchemaResponse, error)
	Fetch(ctx context.Context, in *FetchSchemaRequest, opts ...grpc.CallOption) (*FetchSchemaResponse, error)
}

type schemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemaServiceClient(cc grpc.ClientConnInterface) SchemaServiceClient {
	return &schemaServiceClient{cc}
}

func (c *schemaServiceClient) Register(ctx context.Context, in *RegisterSchemaRequest, opts ...grpc.CallOption) (*RegisterSchemaResponse, error) {
	out := new(RegisterSchemaResponse)
	err := c.cc.Invoke(ctx, SchemaService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) Fetch(ctx context.Context, in *FetchSchemaRequest, opts ...grpc.CallOption) (*FetchSchemaResponse, error) {
	out := new(FetchSchemaResponse)
	err := c.cc.Invoke(ctx, SchemaService_Fetch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServiceServer is the server API for SchemaService service.
// All implementations must embed UnimplementedSchemaServiceServer
// for forward compatibility
type SchemaServiceServer interface {
	Register(context.Context, *RegisterSchemaRequest) (*RegisterSchemaResponse, error)
	Fetch(context.Context, *FetchSchemaRequest) (*FetchSchemaResponse, error)
	mustEmbedUnimplementedSchemaServiceServer()
}

// UnimplementedSchemaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchemaServiceServer struct {
}

func (UnimplementedSchemaServiceServer) Register(context.Context, *RegisterSchemaRequest) (*RegisterSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedSchemaServiceServer) Fetch(context.Context, *FetchSchemaRequest) (*FetchSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedSchemaServiceServer) mustEmbedUnimplementedSchemaServiceServer() {}

// UnsafeSchemaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchemaServiceServer will
// result in compilation errors.
type UnsafeSchemaServiceServer interface {
	mustEmbedUnimplementedSchemaServiceServer()
}

func RegisterSchemaServiceServer(s grpc.ServiceRegistrar, srv SchemaServiceServer) {
	s.RegisterService(&SchemaService_ServiceDesc, srv)
}

func _SchemaService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).Register(ctx, req.(*RegisterSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_Fetch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).Fetch(ctx, req.(*FetchSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchemaService_ServiceDesc is the grpc.ServiceDesc for SchemaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchemaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schema.v1.SchemaService",
	HandlerType: (*SchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _SchemaService_Register_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _SchemaService_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema_service/v1/schema_service.proto",
}
