package main

import (
	"github.com/xavier-niu/grpc-demo/grpcdm"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	grpcdm.RegisterHelloServiceServer(grpcServer, new(grpcdm.HelloService))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	_ = grpcServer.Serve(lis)
}
