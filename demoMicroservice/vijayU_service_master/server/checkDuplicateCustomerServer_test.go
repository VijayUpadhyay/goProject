package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "demoMicroservice/vijayU_ms_interfaces/checkDuplicateCustomer"
	"testing"
)

const (
	chkCustPortClient = "localhost:50058"
)



func CheckDuplicateCustomerBeforeCreateTest(t *testing.T) {
	conn, err := grpc.Dial(chkCustPortClient, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: ", err)
	} else {
		log.Println("Connected to server")
	}
	defer conn.Close()

	tables := []struct {
		input int64
		expected bool
	}{
		{222233336448, false},
		{222266666444, false},
		{222233334444, true},
		{222233334464, false},
	}

	c := pb.NewCheckDuplicateCustomerClient(conn)

	for _, table := range tables {
		actualResponse, err := c.CheckDuplicateCustomerBeforeCreate(context.Background(), &pb.RequestMsg{AadharNumber: table.input})
		if err != nil {
			log.Fatalf("Error while converting from Decimal to Binary: %v", err)
		}
		if actualResponse.CustExist != table.expected {
			t.Errorf("Test failed, expected: '%t', got:  '%t', for input: '%d'", table.expected, actualResponse.CustExist,table.input)
		}
	}

}
