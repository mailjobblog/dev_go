package main

import (
	"context"
	"demo1/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

// 测试GRPC客户端
// grpc.WithInsecure() 选项跳过了对服务器证书的验证
func TestClient(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)

	reply, err := client.Length(context.Background(), &pb.HelloRequest{Res: "test123456"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.Reply)
	fmt.Println(reply.GetReply())
}
