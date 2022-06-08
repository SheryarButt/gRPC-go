package main

import (
	"context"
	"log"
	"math"

	pb "github.com/SheryarButt/gRPC-go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {

	log.Printf("Received a Sqrt request")

	number := in.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Received a negative number: %f", number,
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
