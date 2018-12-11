package main

import (
	pb "demoMicroservice/vijayU_ms_interfaces/checkDuplicateCustomer"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const chkCustport = "0.0.0.0:50058"

type chkDupCustServer struct{}
type custDtls struct {
	custName string
	email    string
	phone    int64
}

func (s *chkDupCustServer) CheckDuplicateCustomerBeforeCreate(ctx context.Context, i *pb.RequestMsg) (*pb.ResponseMsg, error) {
	log.Printf("Received Aadhar Number is %d", i.AadharNumber)
	fmt.Println("Available Customers are")
	// Use DB to get data.Here, using map to keep data.
	var m = map[int64]custDtls{
		222233334444: custDtls{"vijay", "vijay@nokia.com", 978222222},
		222244448888: custDtls{"ram", "ram@nokia.com", 7983334444,},
	}
	fmt.Println(m)
	if m[i.AadharNumber] != (custDtls{}) {
		return &pb.ResponseMsg{CustExist: true}, nil
	}
	return &pb.ResponseMsg{CustExist: false}, nil
}

func main() {
	lis, err := net.Listen("tcp", chkCustport)
	if err != nil {
		log.Fatalf("Failed to listen to port ", chkCustport)
	} else {
		log.Print("Listening to port ", chkCustport)
	}

	s := grpc.NewServer()

	pb.RegisterCheckDuplicateCustomerServer(s, &chkDupCustServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve")
	}
}
