package main

import (
	"github.com/xavier-niu/grpc-demo/netrpc1"
	"log"
	"net"
	"net/rpc"
)

func main() {
	_ = rpc.RegisterName("HelloService", new(netrpc1.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
