package main

import (
	"log"

	pb "github.com/SheryarButt/gRPC-go/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {

	log.Printf("Primes function was invoked with %v", in)

	number := in.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			number = number / divisor
		} else {
			divisor++
		}
	}

	return nil
}
