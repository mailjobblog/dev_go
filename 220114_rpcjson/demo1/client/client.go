package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 手工调用net.Dial函数建立TCP链接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	// 针对客户端的json编解码器
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	res := "test HelloService"
	var reply int
	err = client.Call("HelloService.Length", res, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
