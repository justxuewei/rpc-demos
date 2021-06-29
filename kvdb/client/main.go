package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	go func() {
		var keyChanged string
		err := client.Call("StoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch: " + keyChanged)
	}()

	err = client.Call("StoreService.Set", [2]string{"key1", "value1"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}

}
