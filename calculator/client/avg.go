package main

import (
	"context"
	"log"

	pb "github.com/SheryarButt/gRPC-go/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("Starting to do a Avg RPC...")

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg RPC: %v", err)
	}

	numbers := []int32{3, 5, 9, 27, 81}

	for _, n := range numbers {
		log.Printf("Sending number: %v", n)
		stream.Send(&pb.AvgRequest{
			Number: n,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg RPC: %v", err)
	}

	log.Printf("Avg response: %v", res.GetResult())
}
