package interceptserver

import (
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"fmt"
)

func LogUnary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println("Before LogUnary handler")
		//ur logic
		h, err := handler(ctx, req)
		fmt.Println("After LogUnary handler")
		//ur logic
		return h,err
	}

}

func AuthUnary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println("Before AuthUnary handler")
		//ur logic
		h, err := handler(ctx, req)
		fmt.Println("After AuthUnary handler")
		//ur logic
		return h,err
	}

}