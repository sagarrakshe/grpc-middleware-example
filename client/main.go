package main

import (
	"context"
	"log"

	ts "github.com/sagarrakshe/grpc-middleware-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial(":8001",
		grpc.WithInsecure(),
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
	})
	if err != nil {
		panic(err)
	}
	log.Println(res)

	// Second Request
	ctx = metadata.AppendToOutgoingContext(context.Background(), "idem_id", "1")
	resp, err := c.SayHello(ctx, &ts.HelloRequest{
		Name: "Hello!",
	})
	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
