package main

import (
	"context"
	"log"

	pb "github.com/SheryarButt/gRPC-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*pb.Empty, error) {
	log.Printf("DeleteBlog was invoked")

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot delete object in MongoDB: %v", err,
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog in MongoDB",
		)
	}

	return &pb.Empty{}, nil
}
