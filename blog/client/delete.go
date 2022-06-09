package main

import (
	"context"
	"log"

	pb "github.com/SheryarButt/gRPC-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("DeleteBlog was invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting blog: %v", err)
	}

	log.Printf("Blog was deleted: %v", id)
}
