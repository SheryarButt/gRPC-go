package main

import (
	"fmt"
	"log"

	pb "github.com/SheryarButt/gRPC-go/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {

	log.Printf("GreetManyTimes function was invoked with %v", in)

	for i := 0; i < 10; i++ {
		result := "Hello " + in.GetFirstName() + " number " + fmt.Sprint(i)
		response := &pb.GreetResponse{
			Result: result,
		}
		stream.Send(response)
		log.Printf("Sent: %v", response)
	}
	return nil
}
