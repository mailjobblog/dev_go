package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc_proto/pb"
)

func main() {
	rpc.Register(new(HelloService))
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

type HelloService struct{}

// Length 和原生相比，这里的接收参数和返回参数都用的是proto生成的代码
func (h *HelloService) Length(res pb.HelloRequest, reply *pb.HelloResponse) error {
	reply.Reply = int64(len(res.Res))
	return nil
}
