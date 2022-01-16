package main

import (
	"demo2/service"
	"log"
	"net"
	"net/rpc"
)

func main() {
	// 注册rpc服务
	// 这里使用了封装的方法注册了rpc的方法
	_ = service.RegisterHelloService(new(service.HelloService))

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
