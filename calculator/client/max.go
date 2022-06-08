package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/SheryarButt/gRPC-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {

	log.Printf("Max function was invoked with a streaming request")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Max RPC: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{1, 5, 3, 6, 2, 20}

		for _, number := range numbers {
			log.Printf("Sending number: %v", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
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
				log.Printf("Error while reading stream: %v", err)
				break
			}
			log.Printf("Recieved a new Max value: %v", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
