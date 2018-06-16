package middleware

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getIdemId(ctx context.Context) string {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("Cannot get header metadata from context")
		panic(errors.New("cannot get header metadata from context"))
	}

	if headers["idem_id"] != nil {
		return headers["idem_id"][0]
	}
	return ""
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	store := map[string]interface{}{}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		id := getIdemId(ctx)

		if store[id] != nil {
			log.Printf("Cache hit")
			return store[id], nil
		} else {
			log.Printf("Cache miss")
		}

		resp, err := handler(ctx, req)
		if err == nil {
			store[id] = resp
		}
		return resp, err
	}
}
