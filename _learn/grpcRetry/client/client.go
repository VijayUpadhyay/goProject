package main

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"_learn/grpcRetry/pb"
	"golang.org/x/net/context"
	"fmt"
	"_learn/grpcRetry/config"
	"time"
	"_learn/grpcRetry/interceptors/client"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc/codes"
)

func main() {
	resp,_:=gRPCReq()
	//fmt.Println("Resp from server is: ", resp.B, " and status is: ", resp.Status)
	fmt.Println("\nResp from server is: ", resp)
}

func gRPCReq() (*retrytest.Resp, error) {
	//ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(retryconfig.GrpcTimeout) * time.Second))
	ctx, _ := context.WithTimeout(context.Background(), retryconfig.GrpcTimeout * time.Second)
	TLSOpt := grpc.WithInsecure()
	// MaxDelay is the upper bound of backoff delay.
	/*timeconfig := grpc.BackoffConfig{MaxDelay: 120 * time.Second,}*/
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponentialWithJitter(20 * time.Second, 0.2)),
		grpc_retry.WithCodes(codes.Aborted, codes.Canceled, codes.DeadlineExceeded, codes.NotFound,codes.Unavailable),
		//grpc_retry.WithMax(3),
	}
	//fmt.Println("In client opts is: ",opts[0])
	//including multiple interceptors for client with retry ==> but no implementation of Initial_BackOff
	/*conn, err := grpc.Dial(retryconfig.GrpcServerAdd, TLSOpt,
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(client.RetryUnary(),grpc_retry.UnaryClientInterceptor(opts...), client.LoggingUnary(),client.AuthUnary())),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(grpc_retry.StreamClientInterceptor(opts...))),
	)*/
	req := new(retrytest.Req)
	req.A = 100
	conn, err := grpc.DialContext(ctx,retryconfig.GrpcServerAdd, TLSOpt,
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(client.RetryUnary(req), client.LoggingUnary(),client.AuthUnary())),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(grpc_retry.StreamClientInterceptor(opts...))),
	)

	/*conn, err := grpc.DialContext(ctx, retryconfig.GrpcServerAdd, TLSOpt, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(3)), client.UnaryClientInterceptorForLogging())),
	)*/

	//grpc_retry.WithMax(3) ---->  can be used if we need an interceptor with retry option, instead of opts..., use this method

	//grpc_retry.WithPerRetryTimeout()  -->  to set retry timeout for each call
	//conn, err := grpc.DialContext(ctx,retryconfig.GrpcServerAdd,TLSOpt, grpc.WithInsecure(), client.WithClientUnaryInterceptor()) // --->  for single interceptor
	//fmt.Println("conn is: ",conn)
	/*conn, err := grpc.DialContext(ctx, retryconfig.GrpcServerAdd, TLSOpt,
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),  --->  default client interceptor from go-grpc-middleware/retry pkg
	)*/

	/*conn, err := grpc.DialContext(ctx, retryconfig.GrpcServerAdd, TLSOpt,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(grpc_retry.UnaryClientInterceptor(opts...))),   ---->   2 interceptors for both simple and streaming rpc calls
	)*/

	if err != nil {
		log.Println("did not connect: %v", err)
	}
	defer conn.Close()
	client := retrytest.NewTestServiceClient(conn)
	resp, err := client.CheckRetryOps(context.Background(), req)
	if err != nil {
		fmt.Println("\nError while sending client request: ", err)
	}
	return resp,err
}