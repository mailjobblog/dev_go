package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc/service"
)

func main() {
	// 注册rpc服务
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
