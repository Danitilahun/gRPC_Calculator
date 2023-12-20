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
