package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	_ = rpc.Register(new(HelloService))

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		// 用rpc.ServeCodec函数替代了rpc.ServeConn函数
		// 传入的参数是针对服务端的json编解码器
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

type HelloService struct{}

func (h *HelloService) Length(res string, reply *int) error {
	*reply = len(res)
	return nil
}
