package main

import (
	"github.com/xavier-niu/grpc-demo/rpcctx"
	"log"
	"net"
	"net/rpc"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go func() {
			//goland:noinspection GoUnhandledErrorResult
			defer conn.Close()

			p := rpc.NewServer()
			_ = p.Register(rpcctx.NewHelloService(conn))
			p.ServeConn(conn)
		}()
	}
}
