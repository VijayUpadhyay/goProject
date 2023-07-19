package client

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

func AuthUnary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Println("Printing ctx details of auth interceptor: ", ctx)
		// Logic before invoking the invoker
		start := time.Now()
		// Calls the invoker to execute RPC
		fmt.Println("before invoking the auth invoker")
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println("After invoking the auth invoker")
		// Logic after invoking the invoker
		fmt.Printf("Invoked RPC  details from auth interceptor method=%s; Duration=%s; error=%v", method, time.Since(start), err)
		return err
	}
}
