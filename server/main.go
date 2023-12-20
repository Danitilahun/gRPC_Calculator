package main

import (
	"context"
	"fmt"
	calculator "github.com/Danitilahun/gRPC_Calculator.git/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type calculatorServer struct {
	calculator.UnimplementedCalculatorServer
}

func (s *calculatorServer) Add(ctx context.Context, req *calculator.AddRequest) (*calculator.Response, error) {
	result := req.Num1 + req.Num2
	return &calculator.Response{Response: result}, nil
}

func (s *calculatorServer) Subtract(ctx context.Context, req *calculator.SubtractRequest) (*calculator.Response, error) {
	result := req.Num1 - req.Num2
	return &calculator.Response{Response: result}, nil
}

func (s *calculatorServer) Multiply(ctx context.Context, req *calculator.MultiplyRequest) (*calculator.Response, error) {
	result := req.Num1 * req.Num2
	return &calculator.Response{Response: result}, nil
}

func (s *calculatorServer) Divide(ctx context.Context, req *calculator.DivideRequest) (*calculator.Response, error) {
	result := req.Num1 / req.Num2
	return &calculator.Response{Response: result}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterCalculatorServer(s, &calculatorServer{})
	fmt.Println("Server is running on port 50051")
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
