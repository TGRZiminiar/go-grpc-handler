// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: message.proto

package grpcStreaming

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

// ClientStreamingClient is the client API for ClientStreaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientStreamingClient interface {
	StreamMessage(ctx context.Context, opts ...grpc.CallOption) (ClientStreaming_StreamMessageClient, error)
}

type clientStreamingClient struct {
	cc grpc.ClientConnInterface
}

func NewClientStreamingClient(cc grpc.ClientConnInterface) ClientStreamingClient {
	return &clientStreamingClient{cc}
}

func (c *clientStreamingClient) StreamMessage(ctx context.Context, opts ...grpc.CallOption) (ClientStreaming_StreamMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientStreaming_ServiceDesc.Streams[0], "/protobuf.ClientStreaming/StreamMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStreamingStreamMessageClient{stream}
	return x, nil
}

type ClientStreaming_StreamMessageClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type clientStreamingStreamMessageClient struct {
	grpc.ClientStream
}

func (x *clientStreamingStreamMessageClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStreamingStreamMessageClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreamingServer is the server API for ClientStreaming service.
// All implementations must embed UnimplementedClientStreamingServer
// for forward compatibility
type ClientStreamingServer interface {
	StreamMessage(ClientStreaming_StreamMessageServer) error
	mustEmbedUnimplementedClientStreamingServer()
}

// UnimplementedClientStreamingServer must be embedded to have forward compatible implementations.
type UnimplementedClientStreamingServer struct {
}

func (UnimplementedClientStreamingServer) StreamMessage(ClientStreaming_StreamMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMessage not implemented")
}
func (UnimplementedClientStreamingServer) mustEmbedUnimplementedClientStreamingServer() {}

// UnsafeClientStreamingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientStreamingServer will
// result in compilation errors.
type UnsafeClientStreamingServer interface {
	mustEmbedUnimplementedClientStreamingServer()
}

func RegisterClientStreamingServer(s grpc.ServiceRegistrar, srv ClientStreamingServer) {
	s.RegisterService(&ClientStreaming_ServiceDesc, srv)
}

func _ClientStreaming_StreamMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStreamingServer).StreamMessage(&clientStreamingStreamMessageServer{stream})
}

type ClientStreaming_StreamMessageServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type clientStreamingStreamMessageServer struct {
	grpc.ServerStream
}

func (x *clientStreamingStreamMessageServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStreamingStreamMessageServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreaming_ServiceDesc is the grpc.ServiceDesc for ClientStreaming service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientStreaming_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.ClientStreaming",
	HandlerType: (*ClientStreamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMessage",
			Handler:       _ClientStreaming_StreamMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}
