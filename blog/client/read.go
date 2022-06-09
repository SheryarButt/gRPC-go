package main

import (
	"context"
	"log"

	pb "github.com/SheryarButt/gRPC-go/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("ReadBlog was invoked")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while reading blog: %v", err)
	}

	log.Printf("Blog was read: %v", res)
	return res
}
