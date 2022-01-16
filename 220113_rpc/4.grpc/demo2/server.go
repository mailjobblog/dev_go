package main

import (
	"context"
	"demo2/pb"
	"google.golang.org/grpc"
	"io"
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
	grpcServer.Serve(lis)
}

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

// Length 普通处理
func (h *HelloService) Length(ctx context.Context, res *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := &pb.HelloResponse{Reply: int64(len(res.Res))}
	return reply, nil
}

// StreamLength 流处理
func (h *HelloService) StreamLength(stream pb.HelloService_StreamLengthServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		reply := &pb.HelloResponse{Reply: int64(len(args.GetRes()))}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}
