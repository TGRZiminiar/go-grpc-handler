package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	grpcStreaming "tgrziminiar/grpcStreaming/bidirectional/proto"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	grpcStreaming.UnimplementedBidirectionalMessageServer
	Count int
	Mu    sync.Mutex
}

func (g *grpcHandler) StreamMessage(s grpcStreaming.BidirectionalMessage_StreamMessageServer) error {
	ctx := s.Context()
	g.Count = 0
	for {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := s.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("End of stream from client closing stream...")
				// errorMsg = errors.New("end of stream from client closing stream")
				return err

			} else {
				log.Printf("receive message from client error %v", err)
				err := fmt.Errorf("receive message from client error %v", err)

				resp := &grpcStreaming.Response{
					Key:     fmt.Sprintf("test %d", g.Count),
					Errors:  err.Error(),
					Message: "hello",
				}

				if err := s.Send(resp); err != nil {
					log.Printf("sending message to client error %v", err)
					return err
				}
				return nil
			}
		}

		fmt.Printf("Receiving -> %+v\n", req)

		g.Count++

		resp := &grpcStreaming.Response{
			Key:     fmt.Sprintf("test %d", g.Count),
			Errors:  "",
			Message: "hello",
		}

		if err := s.Send(resp); err != nil {
			log.Printf("sending message to client error %v", err)
			return err
		}

	}
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcStreaming.RegisterBidirectionalMessageServer(s, &grpcHandler{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
