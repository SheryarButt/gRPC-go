package main

import (
	"context"
	"log"

	pb "github.com/SheryarButt/gRPC-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	// Create a context for the client.
	log.Println("doGreet function was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Sheryar",
	})

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
