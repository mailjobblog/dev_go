package main

import (
	"fmt"
	"log"
	"rpc/service"
)

func main() {
	// 连接服务
	client, err := service.DialHelloService("tcp", "127.0.0.1:8888")

	if err != nil {
		log.Fatal("dialing:", err)
	}
	res := "hello world"
	var reply int
	err = client.Length(res, &reply) // 方法调用
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
