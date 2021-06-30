package main

import (
	"context"
	"fmt"
	"github.com/xavier-niu/grpc-demo/grpcdm"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer conn.Close()

	client := grpcdm.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &grpcdm.String{Value: "Hello"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())
}
