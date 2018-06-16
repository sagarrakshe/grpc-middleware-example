package main

import (
	"context"
	"log"
	"net"

	"github.com/sagarrakshe/grpc-middleware-example/grpc_server"
	ts "github.com/sagarrakshe/grpc-middleware-example/proto"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *ts.HelloRequest) (*ts.HelloReply, error) {
	return &ts.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {

	lis, err := net.Listen("tcp", "[::]:8001")
	if err != nil {
		panic(err)
	}

	s := grpc_server.NewServer()
	ts.RegisterGreeterServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("Running server on port 8081")
	s.Serve(lis)
}
