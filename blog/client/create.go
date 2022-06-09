package main

import (
	"context"
	"log"

	pb "github.com/SheryarButt/gRPC-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Creating the blog")
	blog := &pb.Blog{
		AuthorId: "Sheryar",
		Content:  "Hello World",
		Title:    "My first blog",
	}
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Error while creating blog: %v", err)
	}

	log.Printf("Blog has been created: %v", res)
	return res.Id
}
