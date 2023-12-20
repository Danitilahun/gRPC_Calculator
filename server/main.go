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
