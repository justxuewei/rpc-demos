package main

import (
	"github.com/xavier-niu/grpc-demo/netrpc2"
	"log"
	"net"
	"net/rpc"
)

func RegisterHelloService(svc netrpc2.HelloServiceInterface) error {
	return rpc.RegisterName(netrpc2.HelloServiceName, svc)
}

func main() {
	_ = RegisterHelloService(new(netrpc2.HelloService))

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
		go rpc.ServeConn(conn)
	}
}
