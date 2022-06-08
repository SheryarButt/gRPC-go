package main

import (
	"context"
	"log"
	"time"

	pb "github.com/SheryarButt/gRPC-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	// Create a context for the client.
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancle := context.WithTimeout(context.Background(), timeout)
	defer cancle()

	req := &pb.GreetRequest{
		FirstName: "Sheryar",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		if statusErr, ok := status.FromError(err); ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("Timeout was hit! Deadline was exceeded")
				return
			} else {
				log.Fatalf("Error while calling GreetWithDeadline RPC: %v", statusErr)
			}
		} else {
			log.Fatalf("Error while calling GreetWithDeadline RPC: %v", err)
		}
	}

	log.Printf("Response from GreetWithDeadline: %v", res.GetResult())
}
