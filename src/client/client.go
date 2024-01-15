package main

import (
	"context"
	"io"
	"log"
	"math/rand"

	pb "tgrziminiar/grpcStreaming/src/proto"

	"time"

	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().Unix())

	// dial server
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := pb.NewMathClient(conn)
	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	var max int32
	ctx := stream.Context()
	done := make(chan bool)
	responseCh := make(chan *pb.Response)

	// first goroutine sends random increasing numbers to stream
	// and closes it after 10 iterations
	go func() {
		defer close(done)
		for i := 1; i <= 10; i++ {
			req := pb.Request{Num: int32(i)}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			log.Printf("%d sent", req.Num)
			// time.Sleep(time.Millisecond * 200)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// second goroutine receives data from stream
	// and saves result in max variable
	go func() {
		defer close(responseCh)
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
				return
			}
			responseCh <- resp
		}
	}()

	// third goroutine waits for responses and updates max
	// closes done channel when server is done
	go func() {
		for {
			select {
			case resp, ok := <-responseCh:
				if !ok {
					return
				}
				max = resp.Result
				log.Printf("new max %d received", max)
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					log.Println(err)
				}
				return
			}
		}
	}()

	<-done
	log.Printf("finished with max=%d", max)
}
