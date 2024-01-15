package main

import (
	"context"
	"log"
	grpcStreaming "tgrziminiar/grpcStreaming/singleProcess/proto"

	"google.golang.org/grpc"
)

func main() {
	// Create a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client.
	client := grpcStreaming.NewSingleMessageClient(conn)

	// Prepare a request.
	request := &grpcStreaming.Request{
		Key:     "example_key",
		Errors:  "example_errors",
		Content: "example_content",
	}

	// Make an RPC call to the server.
	response, err := client.OneMessage(context.Background(), request)
	if err != nil {
		log.Fatalf("Error calling OneMessage: %v", err)
	}

	// Process the response.
	log.Printf("Response from server: %+v", response)
}
