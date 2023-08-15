// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: completer.proto

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
	Completer_Complete_FullMethodName = "/completer.completer/Complete"
)

// CompleterClient is the client API for Completer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompleterClient interface {
	Complete(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (Completer_CompleteClient, error)
}

type completerClient struct {
	cc grpc.ClientConnInterface
}

func NewCompleterClient(cc grpc.ClientConnInterface) CompleterClient {
	return &completerClient{cc}
}

func (c *completerClient) Complete(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (Completer_CompleteClient, error) {
	stream, err := c.cc.NewStream(ctx, &Completer_ServiceDesc.Streams[0], Completer_Complete_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &completerCompleteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Completer_CompleteClient interface {
	Recv() (*Completion, error)
	grpc.ClientStream
}

type completerCompleteClient struct {
	grpc.ClientStream
}

func (x *completerCompleteClient)