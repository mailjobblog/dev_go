package main

import (
	"context"
	"demo2/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)

	// 返回 stream 流对象
	stream, err := client.StreamLength(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 向服务器发送消息的流
	go func() {
		for {
			if err := stream.Send(&pb.HelloRequest{Res: "test123456"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second) // 每秒调用一次
		}
	}()

	// 接收服务端返回的数据
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetReply())
	}
}
