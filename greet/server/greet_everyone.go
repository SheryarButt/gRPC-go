package main

import (
	"io"
	"log"

	pb "github.com/SheryarButt/gRPC-go/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		log.Printf("Received: %v", req.GetFirstName())
		res := req.GetFirstName() + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}
}
