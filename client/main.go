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

	// Get user input for operator sign (+, -, *, /)
	fmt.Print("Enter operator sign (+, -, *, /): ")
	_, err = fmt.Scan(&operator)
	if err != nil {
		log.Fatalf("Failed to read operator sign: %v", err)
	}
	// Get user input for num2
	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		log.Fatalf("Failed to read second number: %v", err)
	}

	switch operator {
	case "+":
		addResult, err := client.Add(context.Background(), &calculator.AddRequest{Num1: num1, Num2: num2})
		if err != nil {
			log.Fatalf("Add RPC failed: %v", err)
		}
		fmt.Printf("Add Result: %.2f\n", addResult.Response)
	case "-":
		subtractResult, err := client.Subtract(context.Background(), &calculator.SubtractRequest{Num1: num1, Num2: num2})
		if err != nil {
			log.Fatalf("Subtract RPC failed: %v", err)
		}
		fmt.Printf("Subtract Result: %.2f\n", subtractResult.Response)
	case "*":
		multiplyResult, err := client.Multiply(context.Background(), &calculator.MultiplyRequest{Num1: num1, Num2: num2})
		if err != nil {
			log.Fatalf("Multiply RPC failed: %v", err)
		}
		fmt.Printf("Multiply Result: %.2f\n", multiplyResult.Response)
	case "/":
		divideResult, err := client.Divide(context.Background(), &calculator.DivideRequest{Num1: num1, Num2: num2})
		if err != nil {
			log.Fatalf("Divide RPC failed: %v", err)
		}
		fmt.Printf("Divide Result: %.2f\n", divideResult.Response)
	default:
		fmt.Println("Invalid operation type. Please choose from add/subtract/multiply/divide.")
	}

}
