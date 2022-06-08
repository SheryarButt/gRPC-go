package main

import (
	"io"
	"log"

	pb "github.com/SheryarButt/gRPC-go/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {

	log.Printf("Avg function was invoked with a streaming request")

	var sum int32
	var count int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum / count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		log.Printf("Received number %v", req.GetNumber())
		sum += req.GetNumber()
		count++
	}
}
