package main

import (
	"log"
	"net"

	pb "github.com/SheryarButt/gRPC-go/calculator/proto"
	"google.golang.org/grpc"
)

var addr = "localhost:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
