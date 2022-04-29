package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	// Calls the handler
	h, err := handler(ctx, req)

	// log.Println("Interceptor:::")
	// fmt.Println(metadata.FromIncomingContext(ctx))

	// Logging with grpclog (grpclog.LoggerV2)
	log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

func withUnaryServerInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}
