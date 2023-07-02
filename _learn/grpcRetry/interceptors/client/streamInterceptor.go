package client

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(parentCtx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		var clientStream grpc.ClientStream

		//your impl

		return clientStream, nil
	}
}
