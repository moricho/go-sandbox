package main

import (
	"log"
	"net"

	"github.com/moricho/go-sandbox/microservice-sample/greeter"
	pb "github.com/moricho/go-sandbox/microservice-sample/greeter/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	greeterService := &greeter.GreeterService{}

	pb.RegisterGreeterServer(server, greeterService)
	server.Serve(lis)
}
