package main

import (
	"fmt"
	"github.com/xavier-niu/grpc-demo/netrpc2"
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ netrpc2.HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		Client: client,
	}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Call(netrpc2.HelloServiceName+".Hello", request, reply)
}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

