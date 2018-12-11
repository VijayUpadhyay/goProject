package main

import (
	"google.golang.org/grpc"
	"log"
	pb "demoMicroservice/vijayU_ms_interfaces/simRegistration"
	"testing"
	"golang.org/x/net/context"
	"strings"
)

const (
	simRegPort = "localhost:9000"
)

func RegisterCustomerTest(t *testing.T) {
	conn, err := grpc.Dial(simRegPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: ", err)
	} else {
		log.Println("Connected to server")
	}
	defer conn.Close()

	client := pb.NewSimRegistrationClient(conn)

	customer := &pb.CustomerDetails{
		CustId:   101,
		CustName: "Vijay Upadhyay",
		Email:    "vijay.upadhyay@nokia.com",
		AlternateContactNumber: "900000000",
		AadharNumber:           221044448888,
		Addresses: []*pb.CustomerDetails_Addrress{
			&pb.CustomerDetails_Addrress{
				HouseNumber: "34/C",
				Street:      "Gorakhnath",
				City:        "Gorakhpur",
				State:       "UP",
				Zip:         274402,
			},
			&pb.CustomerDetails_Addrress{
				HouseNumber: "36/C",
				Street:      "Friends Layout",
				City:        "Bangalore",
				State:       "Karnatak",
				Zip:         560037,
			},
		},
	}
	actualResponse, err:=client.RegisterCustomer(context.Background(),customer)
	if err != nil {
		log.Fatalf("Error while converting from Decimal to Binary: %v", err)
	}
	if strings.Compare(actualResponse.SuccessMsg,"Success") !=0 {
		t.Errorf("Test failed, expected: '%v', got:  '%t', for input: '%d'", "Success", actualResponse.SuccessMsg,customer)
	}
}
