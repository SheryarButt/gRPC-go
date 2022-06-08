package main

import (
	"context"
	"log"
	"time"

	pb "github.com/SheryarButt/gRPC-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("Starting to do a LongGreet RPC...")
	reqs := []*pb.GreetRequest{
		{FirstName: "Muhammad"},
		{FirstName: "Sheryar"},
		{FirstName: "Butt"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet RPC: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet RPC: %v", err)
	}
	log.Printf("LongGreet response: %s", res.GetResult())
}
