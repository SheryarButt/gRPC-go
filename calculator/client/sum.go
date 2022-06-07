package main

import (
	"context"
	"log"

	pb "github.com/SheryarButt/gRPC-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.Sum)
}
