package main

import (
	"context"
	"log"
	"time"

	grpcStreaming "tgrziminiar/grpcStreaming/singleProcess/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:5000", opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcStreaming.NewSingleMessageClient(conn)

	responseChan := make(chan *grpcStreaming.Response)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {

		request := &grpcStreaming.Request{
			Key:     "testing key",
			Message: "hello request",
		}

		response, err := client.OneMessage(context.Background(), request)
		if err != nil {
			log.Fatalf("Error calling OneMessage: %v", err)
			return
		}

		responseChan <- response
		close(responseChan)
	}()

	select {
	case response := <-responseChan:
		log.Printf("Response from server: %+v", response)
	case <-ctx.Done():
		log.Println("Call timed out.")
	}
}
