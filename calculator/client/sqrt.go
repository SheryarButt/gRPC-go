package main

import (
	"context"
	"log"

	pb "github.com/SheryarButt/gRPC-go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("Starting to do a Sqrt RPC...")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Fatalf("Error message from server: %v", e.Message())
			log.Fatalf("Error code from server: %v", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Fatalf("Received a negative number")
			}
		} else {
			log.Fatalf("A non gRPC error : %v", err)
		}
	}
	log.Printf("Response from Sum: %v", res.Result)
}
