package main

import (
	"context"
	"io"
	"log"

	pb "github.com/SheryarButt/gRPC-go/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("Starting to do a GreetManyTimes RPC...")
	req := &pb.GreetRequest{
		FirstName: "Sheryar",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}
