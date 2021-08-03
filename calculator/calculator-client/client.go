package main

import (
	"context"
	"log"

	pb "github.com/hilli/gRPC-calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
)

func main() {
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := pb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c pb.CalculatorServiceClient) {
	log.Printf("Starting Sum request")
	req := &pb.SumRequest{
		FirstNumber: 2,
		SecondNumber: 3,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Panicf("Sum error: %v", err)
	}
	log.Printf("Sum result: %v", res.GetResult())
}