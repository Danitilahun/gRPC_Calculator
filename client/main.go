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

	defer conn.Close()

	client := calculator.NewCalculatorClient(conn)

	var num1, num2 float64
	var operator string

	// Get user input for num1
	fmt.Print("Enter first number: ")
	_, err = fmt.Scan(&num1)
	if err != nil {
		log.Fatalf("Failed to read first number: %v", err)
	}

	// Get user input for num2
	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		log.Fatalf("Failed to read second number: %v", err)
	}

	// Get user input for operator sign (+, -, *, /)
	fmt.Print("Enter operator sign (+, -, *, /): ")
	_, err = fmt.Scan(&operator)
	if err != nil {
		log.Fatalf("Failed to read operator sign: %v", err)
	}

}
