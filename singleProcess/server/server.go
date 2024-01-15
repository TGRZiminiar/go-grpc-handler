package main

import (
	"context"
	"fmt"
	"log"
	"net"
	grpcStreaming "tgrziminiar/grpcStreaming/singleProcess/proto"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	grpcStreaming.UnimplementedSingleMessageServer
}

func (g *grpcHandler) OneMessage(ctx context.Context, req *grpcStreaming.Request) (*grpcStreaming.Response, error) {

	fmt.Printf("Received Request: %+v\n", req)

	// Example: Send a response.
	response := &grpcStreaming.Response{
		Result: 42,
	}

	return response, nil
}

func main() {
	// create listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	grpcStreaming.RegisterSingleMessageServer(s, &grpcHandler{})

	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
