package main

import (
	"context"
	"log"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	ts "github.com/sagarrakshe/grpc-middleware-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func main() {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(10 * time.Second)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted, codes.Unavailable, codes.Unknown, codes.ResourceExhausted, codes.Unavailable),
	}

	conn, err := grpc.Dial(":8001",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		panic(err)
	}

	c := ts.NewGreeterClient(conn)

	// First Request
	ctx := metadata.AppendToOutgoingContext(context.Background(), "idem_id", "1")
	log.Printf("Context: %+v", ctx)

	res, err := c.SayHello(ctx, &ts.HelloRequest{
		Name: "World!",
	}, grpc_retry.WithMax(5), grpc_retry.WithPerRetryTimeout(10*time.Second))

	if err != nil {
		panic(err)
	}

	log.Println(res)

	// Second Request
	resp, err := c.SayHello(ctx, &ts.HelloRequest{
		Name: "New World!",
	}, grpc_retry.WithMax(3), grpc_retry.WithPerRetryTimeout(1*time.Second))

	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
