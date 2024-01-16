package main

import (
	"context"
	"fmt"
	"io"
	"log"
	grpcStreaming "tgrziminiar/grpcStreaming/serverStream/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(":5000", opts...)
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	client := grpcStreaming.NewServerStreamingClient(conn)
	stream, err := client.StreamMessage(context.Background(), &grpcStreaming.Request{
		Key:     "testing key",
		Message: "hello",
	})
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan struct{})
	ctx := stream.Context()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("closing streaming from server")
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
				close(done)
				return
			}
			fmt.Printf("client receive message -> %+v\n", resp)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
}
