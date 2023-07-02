package interceptserver

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

func ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	fmt.Println("Request came at: ", start)
	fmt.Println("Server Interceptor before handler")
	h, err := handler(ctx, req)
	fmt.Println("Server Interceptor after handler")
	fmt.Printf("\nRequest - Method:%s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)
	return h, err
}
