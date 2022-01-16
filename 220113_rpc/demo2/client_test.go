package main

import (
	"demo2/service"
	"fmt"
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	client, err := service.DialHelloService("tcp", "127.0.0.1:8888")

	if err != nil {
		log.Fatal("dialing:", err)
	}
	res := "hello world"
	var reply int
	err = client.Length(res, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
