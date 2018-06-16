package grpc_server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sagarrakshe/grpc-middleware-example/middleware"
	"google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	unaries := []grpc.UnaryServerInterceptor{
		middleware.UnaryServerInterceptor(),
	}

	return grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaries...)),
	)
}
