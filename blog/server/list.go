package main

import (
	"context"
	"log"

	pb "github.com/SheryarButt/gRPC-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBlogs(in *pb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("ListBlogs was invoked")

	cursor, err := collection.Find(context.Background(), primitive.D{})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Unknown internal error",
		)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		data := &BlogItem{}
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				"Error while decoding data from MongoDB",
			)
		}

		stream.Send(documentToBlog(data))
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			"Error while iterating cursor",
		)
	}

	return nil
}
