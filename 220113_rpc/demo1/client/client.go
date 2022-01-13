package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 拨号连接
	client, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 定义请求和接受参数
	res := "test HelloService"
	var reply int

	// 调用rpc方法
	err = client.Call("HelloService.Length", res, &reply)
	if err != nil {
		log.Fatal(err)
	}

	// 输出结果
	fmt.Println(reply)
}
