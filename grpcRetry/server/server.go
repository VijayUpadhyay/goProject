package main

import (
	"_learn/grpcRetry/config"
	"_learn/grpcRetry/interceptors/server"
	"_learn/grpcRetry/pb"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
)

type server struct{}

var a = int32(1000)

func (s *server) CheckRetryOps(ctx context.Context, i *retrytest.Req) (*retrytest.Resp, error) {

	resp := new(retrytest.Resp)
	//context check
	timeVal, flag := ctx.Deadline()
	fmt.Println("Deadline details--> timeVal is: ", timeVal, " and flag is: ", flag)
	// to check deadline/timeout is over and client is still expecting the output or not
	if ctx.Err() == context.Canceled {
		fmt.Println("Inside ctx.Err()")
		resp.Status = codes.Canceled.String()
		return resp, errors.New("Client cancelled, abandoning.")
	}
	//fmt.Printf("Going to sleep for :%d sec", retryconfig.Sleep_Time)
	//time.Sleep(retryconfig.Sleep_Time * time.Second)
	//time.Sleep(retryconfig.Sleep_Time*time.Second)
	fmt.Println("\nFilling resp.....")
	a++
	resp.B, resp.Status = a, "SUCCESS"
	fmt.Println("Sent Response is:", a)
	fmt.Println("**************Done**********")
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", retryconfig.GrpcServerAdd)
	if err != nil {
		log.Println("Failed to listen to port ", retryconfig.GrpcServerAdd)
	} else {
		log.Print("Listening to port ", retryconfig.GrpcServerAdd)
	}
	fmt.Println("Going to sleep for a min")
	//time.Sleep(60*time.Second)
	//s := grpc.NewServer(withServerUnaryInterceptor()) // we can also go for chain interceptors

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptserver.LogUnary(),
			interceptserver.AuthUnary(),
		)),
	)
	retrytest.RegisterTestServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Println("Failed to serve")
	}
}

func withServerUnaryInterceptor() grpc.ServerOption {
	fmt.Println("Going to intercept server")
	//your logic for performing something which needs to be done before starting you actual rpc's task
	return grpc.UnaryInterceptor(interceptserver.ServerInterceptor)
}

func withServerUnaryInterceptor2() grpc.ServerOption {
	fmt.Println("Going to intercept server")
	//your logic for performing something which needs to be done before starting you actual rpc's task
	return grpc.UnaryInterceptor(interceptserver.ServerInterceptor)
}
