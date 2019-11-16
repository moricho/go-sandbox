package main

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/moricho/go-sandbox/microservice-sample/greeter"
	pb "github.com/moricho/go-sandbox/microservice-sample/greeter/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// Prepare zap logger
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// Generate gRPC server and set interceptor
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zapLogger),
			// grpc_auth.UnaryServerInterceptor(auth),
		)),
	)

	// Register service on gRPC server
	greeterService := &greeter.GreeterService{}
	pb.RegisterGreeterServer(server, greeterService)

	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("gRPC server started on :%d\n", port)
	server.Serve(lis)
}
