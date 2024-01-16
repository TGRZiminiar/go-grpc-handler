package main

import (
	"fmt"
	"log"
	"net"
	grpcStreaming "tgrziminiar/grpcStreaming/serverStream/proto"
	"time"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	grpcStreaming.UnimplementedServerStreamingServer
}

func (g *grpcHandler) StreamMessage(req *grpcStreaming.Request, srv grpcStreaming.ServerStreaming_StreamMessageServer) error {

	fmt.Printf("Received Request: %+v\n", req)

	for i := 1; i <= 5; i++ {
		response := &grpcStreaming.Response{
			Key:     "testing",
			Errors:  "",
			Message: fmt.Sprintf("hello from server %d", i),
		}

		if err := srv.Send(response); err != nil {
			log.Printf("Error sending response: %v", err)
			return err
		}

		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcStreaming.RegisterServerStreamingServer(s, &grpcHandler{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
