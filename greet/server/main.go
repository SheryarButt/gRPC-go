package main

import (
	"log"
	"net"

	pb "github.com/SheryarButt/gRPC-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", addr)

	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		certfile := "ssl/server.crt"
		keyfile := "ssl/server.pem"
		cred, err := credentials.NewServerTLSFromFile(certfile, keyfile)
		if err != nil {
			log.Fatalf("failed to load credentials: %v", err)
		}
		opts = append(opts, grpc.Creds(cred))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
