#!/bin/bash
# Generate the source code for go
# This will [re]generate the `calculator/calculatorpb/*.pb.go` files

protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  calculator/calculatorpb/calculator.proto
