package main

import (
	"context"
	"log"
	"time"

	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	ts "github.com/sagarrakshe/grpc-middleware-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func main() {
	opts := []grpcRetry.CallOption{
		grpcRetry.WithBackoff(grpcRetry.BackoffLinear(10 * time.Second)),
		grpcRetry.WithCodes(codes.NotFound, codes.Aborted, codes.Unavailable, codes.Unknown, codes.ResourceExhausted, codes.Unavailable),
	}

	conn, err := grpc.Dial(":8001",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpcRetry.UnaryClientInterceptor(opts...)),
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
	}, grpcRetry.WithMax(5), grpcRetry.WithPerRetryTimeout(10*time.Second))

	if err != nil {
		panic(err)
	}

	log.Println(res)

	// Second Request
	resp, err := c.SayHello(ctx, &ts.HelloRequest{
		Name: "New World!",
	}, grpcRetry.WithMax(3), grpcRetry.WithPerRetryTimeout(1*time.Second))

	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
