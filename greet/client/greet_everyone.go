package main

import (
	"context"
	"time"
	"log"
	"io"

	pb "github.com/SheryarButt/gRPC-go/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Printf("Starting to do a GreetEveryone RPC...")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetEveryone RPC: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Muhammad"},
		{FirstName: "Sheryar"},
		{FirstName: "Butt"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending req: %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reading client stream: %v", err)
			}

			log.Printf("Received: %v", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
