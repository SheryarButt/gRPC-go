package main

import (
	"log"

	"google.golang.org/grpc"
)

var addr string = "localhost:50051"

func main() {

	conn, err := grpc.Dial(addr)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

}
