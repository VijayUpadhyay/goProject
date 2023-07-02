package client

import (
	cfg "_learn/grpcRetry/config"
	"_learn/grpcRetry/pb"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/backoffutils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

// for chain interceptor
func RetryUnary(clientreq *retrytest.Req) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		hr, min, sec := time.Now().Clock()
		fmt.Printf("Received request at: %d::%d::%d at RetryUnary \n", hr, min, sec)
		// Calls the invoker to execute RPC
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			fmt.Println("Going to retry using Backoff algo")
			sleepAndRetry(clientreq, start)
		}
		fmt.Println("\nAfter invoking the invoker error is: ", err)
		fmt.Println("After invoking the invoker")
		fmt.Printf("\nInvoked RPC  details from 2nd interceptor method=%s; Duration=%s; Error=%v", method, time.Since(start), err)
		return err
	}
}

func sleepAndRetry(req *retrytest.Req, start time.Time) {
	var attempt uint8
	for attempt = 1; (attempt <= cfg.MAX_RETRY) && (int(start.Sub(time.Now()).Seconds()) < cfg.MAX_BACKOFF); attempt++ {
		fmt.Println(" int(start.Sub(time.Now()).Seconds()) < cfg.MAX_BACKOFF ", int(start.Sub(time.Now()).Seconds()) < cfg.MAX_BACKOFF)
		fmt.Println("Attemp number: ", attempt)
		sleeptime := GetSleepTimeInterval(attempt, float64(cfg.INITIAL_BACKOFF))
		fmt.Println("Going to sleep for: ", sleeptime)
		time.Sleep(sleeptime)
		err := ClientRetry(req)
		if err == nil && attempt == 3 {
			fmt.Println("Going to break for loop")
			break
		}
	}
}

func GetSleepTimeInterval(attempt uint8, backOff float64) time.Duration {
	var sleeptime time.Duration
	if attempt == 1 {
		sleeptime = cfg.INITIAL_BACKOFF * time.Second
	} else {
		backOff = backOff * cfg.MULTIPLIER
		current_backofftime := time.Duration(backOff+float64(30)) * time.Second
		sleeptime = backoffutils.JitterUp(current_backofftime, float64(cfg.JITTER))
	}
	return sleeptime
}

func ClientRetry(req *retrytest.Req) error {
	ctx, _ := context.WithTimeout(context.Background(), cfg.GrpcTimeout*time.Second)
	conn, err := grpc.DialContext(ctx, cfg.GrpcServerAdd, grpc.WithInsecure())
	fmt.Printf("ClientRetry: conn is: %v and err is: %v\n", conn, err)
	if err != nil {
		log.Println("Error at client retry: %v", err)
	}
	defer conn.Close()
	client := retrytest.NewTestServiceClient(conn)
	resp, err := client.CheckRetryOps(context.Background(), req)
	if err != nil {
		fmt.Println("Error while sending client request: ", err)
	}
	fmt.Println("\nResp from server at ClientRetry is: ", resp)
	fmt.Println("*****************Done***********************")
	return err
}
