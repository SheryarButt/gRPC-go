package main

import (
	"context"
	"log"

	pb "github.com/SheryarButt/gRPC-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("UpdateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Sheryar",
		Title:    "My Updated Blog",
		Content:  "My updated content",
	}

	res, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error while updating blog: %v", err)
	}

	log.Printf("Blog was updated: %v", res)

}
