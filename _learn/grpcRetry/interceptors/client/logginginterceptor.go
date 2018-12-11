package client

import (
	"google.golang.org/grpc"
	"time"
	"golang.org/x/net/context"
	"fmt"
)

// for chain interceptor
func LoggingUnary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Println("Printing details of context of 2nd interceptor:", ctx)
		fmt.Println("before invoking the invoker")
		start := time.Now()
		// Calls the invoker to execute RPC
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println("\nAfter invoking the invoker error is: ",err)
		fmt.Println("After invoking the invoker")
		fmt.Printf("\nInvoked RPC  details from 2nd interceptor method=%s; Duration=%s; Error=%v", method, time.Since(start), err)
		return err
	}
}

//This method can be used if we need single client interceptor
func WithClientUnaryInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

// called for intercepting gRPC communication
func clientInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption, ) error {
	fmt.Println("Printing details of context of 2nd interceptor:", ctx)
	// Logic before invoking the invoker
	fmt.Println("before invoking the invoker")
	start := time.Now()
	// Calls the invoker to execute RPC
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Println("After invoking the invoker error is: ",err)
	fmt.Println("After invoking the invoker")
	// Logic after invoking the invoker
	fmt.Printf("\nInvoked RPC method=%s; Duration=%s; Error=%v", method, time.Since(start), err)
	return err
}