package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	// 注册 rpc 服务
	_ = rpc.Register(new(HelloService))
	// _ = rpc.RegisterName("HelloService", new(HelloService)) // 自定义名称注册

	// 启动tcp
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 获取连接信息
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	// 阻塞监听服务
	rpc.ServeConn(conn)
}

// HelloService 测试被调用的结构体方法
type HelloService struct{}

func (h *HelloService) Length(res string, reply *int) error {
	*reply = len(res)
	return nil
}
