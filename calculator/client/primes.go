package main

import (
	"context"
	"io"
	"log"

	pb "github.com/SheryarButt/gRPC-go/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from Primes: %v", res.Result)
	}
}
