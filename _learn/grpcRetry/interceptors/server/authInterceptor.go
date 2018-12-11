package interceptserver

import (
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"fmt"
)

/*func AuthUnary(optFuncs ...grpc.CallOption) grpc.UnaryServerInterceptor {
	return func(parentCtx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (err error) {
		fmt.Println("Inside StreamClientInterceptor with context: ", parentCtx)
		return err
	}
}*/

func LogUnary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println("Before LogUnary handler")
		h, err := handler(ctx, req)
		fmt.Println("After LogUnary handler")
		return h,err
	}

}

func AuthUnary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println("Before AuthUnary handler")
		h, err := handler(ctx, req)
		fmt.Println("After AuthUnary handler")
		return h,err
		return handler(ctx, req)
	}

}

/*
func defaultAuthentication() grpcauthentication.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		return ctx, nil
	}
}*/
