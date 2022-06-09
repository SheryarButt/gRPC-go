package main

import (
	"context"
	"io"
	"log"

	pb "github.com/SheryarButt/gRPC-go/blog/proto"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("ListBlog was invoked")

	steam, err := c.ListBlogs(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling ListBlogs RPC: %v", err)
	}

	for {
		res, err := steam.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Blog was received: %v", res)
	}
}
