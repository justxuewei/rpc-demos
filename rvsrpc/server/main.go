package main

import (
	"github.com/xavier-niu/grpc-demo/rvsrpc"
	"net"
	"net/rpc"
	"time"
)

func main() {
	// a service from local network
	_ = rpc.Register(new(rvsrpc.HelloService))

	for {
		// establish a tcp connection with a client from internet
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}

		rpc.ServeConn(conn)
		_ = conn.Close()
	}
}
