package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	calculator "github.com/Danitilahun/gRPC_Calculator.git/proto"
)

func main() {
	serverAddress := "localhost:50051"
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
}
