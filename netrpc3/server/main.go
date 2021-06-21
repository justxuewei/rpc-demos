package main

import (
	"github.com/xavier-niu/grpc-demo/netrpc3"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func RegisterHelloService(svc netrpc3.HelloServiceInterface) error {
	return rpc.RegisterName(netrpc3.HelloServiceName, svc)
}

func main() {
	_ = RegisterHelloService(new(netrpc3.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// launch a goroutine to handle the RPC request
		// adopt the JsonRPC codec
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
