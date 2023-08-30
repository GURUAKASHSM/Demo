package main

import (
	"context"
	"fmt"
	
	"log"

	cus "Netxd_Customer/Netxd_Customer" 
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := cus.NewCustomerServiceClient(conn)
	

	response, err := client.CreateCustomer(context.Background(), &cus.Customer{FirstName: "GURU",LastName: "Akash",BankId: "ABCBANK",Balance: 55000.00,})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response.FirstName)
}
