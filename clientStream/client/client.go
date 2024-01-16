package main

import (
	"context"
	"log"
	"time"

	grpcStreaming "tgrziminiar/grpcStreaming/clientStream/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcStreaming.NewClientStreamingClient(conn)

	// Create a stream
	stream, err := client.StreamMessage(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	// Send multiple requests
	for i := 0; i < 5; i++ {
		request := &grpcStreaming.Request{
			Key:     "test_key",
			Message: "hello",
		}

		if err := stream.Send(request); err != nil {
			log.Fatalf("Error sending request: %v", err)
		}

		time.Sleep(200 * time.Millisecond)
	}

	// Close the stream and receive the response
	response, err := stream.CloseAndRecv()
	if err != nil {
		if status.Code(err) == codes.Canceled {
			log.Println("Server closed the connection.")
		} else {
			log.Fatalf("Error receiving response: %v", err)
		}
		return
	}

	log.Printf("Response from server: %+v", response)
}
