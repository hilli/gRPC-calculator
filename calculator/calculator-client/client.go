package main

import (
	"context"
	"log"
	"os"
	"strconv"

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
	var first_number, second_number int32 = 2, 3
	if len(os.Args) > 2 {
		f64,_ := strconv.ParseInt(os.Args[1], 10, 32)
		s64,_ := strconv.ParseInt(os.Args[2], 10, 32)
		first_number = int32(f64)
		second_number = int32(s64)
	}
	req := &pb.SumRequest{
		FirstNumber: first_number,
		SecondNumber: int32(second_number),
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Panicf("Sum error: %v", err)
	}
	log.Printf("Sum result: %v", res.GetResult())
}