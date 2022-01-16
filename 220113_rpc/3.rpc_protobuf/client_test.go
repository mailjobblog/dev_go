package __rpc_protobuf

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc_proto/pb"
	"testing"
)

func TestClient(t *testing.T) {
	client, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 定义请求和接受参数
	// 接收参数和返回参数都用的是proto生成的代码
	res := pb.HelloRequest{Res: "test666"}
	var reply pb.HelloResponse

	err = client.Call("HelloService.Length", res, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
