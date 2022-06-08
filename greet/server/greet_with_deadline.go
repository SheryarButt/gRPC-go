package main

import (
	"context"
	"log"
	"time"

	pb "github.com/SheryarButt/gRPC-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("GreetWithDeadline was invoked with %v", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Deadline exceeded! Client cancelled the request")
			return nil, status.Error(codes.DeadlineExceeded, "The deadline has been exceeded")
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.GetFirstName(),
	}, nil
}
