package main

import (
	"github.com/xavier-niu/grpc-demo/grpcstm"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	grpcstm.RegisterHelloServiceServer(grpcServer, new(grpcstm.HelloService))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	_ = grpcServer.Serve(lis)
}
