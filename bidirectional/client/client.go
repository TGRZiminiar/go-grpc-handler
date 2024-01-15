package main

import (
	"context"
	"fmt"
	"io"
	"log"
	grpcStreaming "tgrziminiar/grpcStreaming/bidirectional/proto"
	"time"

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

	client := grpcStreaming.NewBidirectionalMessageClient(conn)
	stream, err := client.StreamMessage(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan struct{})
	ctx := stream.Context()

	go func() {
		defer close(done)
		for i := 1; i <= 10; i++ {
			req := grpcStreaming.Request{
				Key:     fmt.Sprintf("testing key %d", i),
				Message: "hello",
			}
			fmt.Printf("client sending msg -> %d\n", i)
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %+v", err)
			}
			time.Sleep(10 * time.Millisecond)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
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
