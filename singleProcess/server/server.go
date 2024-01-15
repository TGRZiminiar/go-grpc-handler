package main

import (
	"context"
	"fmt"
	"log"
	"net"
	grpcStreaming "tgrziminiar/grpcStreaming/singleProcess/proto"
	"time"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	grpcStreaming.UnimplementedSingleMessageServer
}

func (g *grpcHandler) OneMessage(ctx context.Context, req *grpcStreaming.Request) (*grpcStreaming.Response, error) {

	fmt.Printf("Received Request: %+v\n", req)

	response := &grpcStreaming.Response{
		Key:     "testing",
		Errors:  "",
		Message: "hello from server",
	}
	time.Sleep(1 * time.Second)
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcStreaming.RegisterSingleMessageServer(s, &grpcHandler{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
