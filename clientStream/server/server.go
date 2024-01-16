package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	grpcStreaming "tgrziminiar/grpcStreaming/clientStream/proto"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	grpcStreaming.UnimplementedClientStreamingServer
	Count chan int
	Mu    sync.Mutex
}

func (g *grpcHandler) StreamMessage(stream grpcStreaming.ClientStreaming_StreamMessageServer) error {
	var receivedValues []int

	for {
		req, err := stream.Recv()
		if err == io.EOF {

			// End of client stream; process received values and send response.
			g.Mu.Lock()
			defer g.Mu.Unlock()

			// Process received values (e.g., calculate the sum).
			var sum int
			for _, val := range receivedValues {
				sum += val
			}

			resp := &grpcStreaming.Response{
				Key:     fmt.Sprintf("sum: %d", sum),
				Errors:  "",
				Message: "hello",
			}

			// Send response to the client.
			if err := stream.SendAndClose(resp); err != nil {
				log.Printf("error sending response to client: %v", err)
				return err
			}

			return nil
		}

		if err != nil {
			log.Printf("error receiving message from client: %v", err)
			return err
		}

		fmt.Printf("Receiving -> %+v\n", req)

		// Lock and append the received value.
		g.Mu.Lock()
		receivedValues = append(receivedValues, len(req.Message))
		g.Mu.Unlock()
	}
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcStreaming.RegisterClientStreamingServer(s, &grpcHandler{
		Count: make(chan int),
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
