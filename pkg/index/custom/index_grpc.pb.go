// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: index.proto

package custom

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
	Index_Query_FullMethodName = "/index.index/Query"
)

// IndexClient is the client API for Index service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndexClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*Results, error)
}

type indexClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexClient(cc grpc.ClientConnInterface) IndexClient {
	return &indexClient{cc}
}

func (c *indexClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*Results, error) {
	out := new(Results)
	err := c.cc.Invoke(ctx, Index_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServer is the server API for Index service.
// All implementations must embed UnimplementedIndexServer
// for forward compatibility
type IndexServer interface {
	Query(context.Context, *QueryRequest) (*Results, error)
	mustEmbedUnimplementedIndexServer()
}

// UnimplementedIndexServer must be embedded to have forward compatible implementations.
type UnimplementedIndexServer struct {
}

func (UnimplementedIndexServer) Query(context.Context, *QueryRequest) (*Results, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedIndexServer) mustEmbedUnimplementedIndexServer() {}

// UnsafeIndexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexServer will
// result in compilation errors.
type UnsafeIndexServer interface {
	mustEmbedUnimplementedIndexServer()
}

func RegisterIndexServer(s grpc.ServiceRegistrar, srv IndexServer) {
	s.RegisterService(&Index_ServiceDesc, srv)
}

func _Index_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
