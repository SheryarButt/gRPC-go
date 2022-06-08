package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/SheryarButt/gRPC-go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked with a streaming request")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		log.Printf("Received a greet request: %v", req)
		res += fmt.Sprintf("Hello %s! ", req.GetFirstName())
	}
}
