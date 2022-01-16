package main

import (
	"context"
	"demo1/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 创建一个 grpc server
	grpcServer := grpc.NewServer()
	// 注册 grpc
	pb.RegisterHelloServiceServer(grpcServer, new(HelloService))

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	// 监听端口上提供gRPC服务
	grpcServer.Serve(lis)
}

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (h *HelloService) Length(ctx context.Context, res *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := &pb.HelloResponse{Reply: int64(len(res.Res))}
	return reply, nil
}
