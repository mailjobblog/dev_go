# GRPC流

RPC是远程函数调用，因此每次调用的函数参数和返回值不能太大，否则将严重影响每次调用的响应时间。  
因此传统的RPC方法调用对于上传和下载较大数据量场景并不适合。同时传统RPC模式也不适用于对时间不确定的订阅和发布模式。  
为此，gRPC框架针对服务器端和客户端分别提供了流特性。  

## 优化过程

在 proto 文件中，我们添加一个流方法 StreamLength 。关键字stream指定启用流特性，参数部分是接收客户端参数的流，返回值是返回给客户端的流。

```go
service HelloService {
    // ... ...
    
    // 计算字符串长度(grpc流)
    // stream 用于声明流处理
    rpc StreamLength (stream HelloRequest) returns (stream HelloResponse);

    // ... ...
}
```

重新生成代码可以看到接口中新增加的 StreamLength 方法的定义：  
在服务端的 StreamLength 方法参数是一个新的 HelloService_StreamLengthServer 类型的参数，可以用于和客户端双向通信。  
客户端的 StreamLength 方法返回一个 HelloService_StreamLengthClient 类型的返回值，可以用于和服务端进行双向通信。

```go
type HelloServiceServer interface {
	// 计算字符串长度
	Length(context.Context, *HelloRequest) (*HelloResponse, error)
	// 计算字符串长度(grpc流)
	// stream 用于声明流处理
	StreamLength(HelloService_StreamLengthServer) error
	mustEmbedUnimplementedHelloServiceServer()
}

type HelloServiceClient interface {
    // 计算字符串长度
    Length(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
    // 计算字符串长度(grpc流)
    // stream 用于声明流处理
    StreamLength(ctx context.Context, opts ...grpc.CallOption) (HelloService_StreamLengthClient, error)
}
```

HelloService_StreamLengthServer 和 HelloService_StreamLengthClient 都是接口类型

```go
type HelloService_StreamLengthServer interface {
    Send(*HelloResponse) error
    Recv() (*HelloRequest, error)
    grpc.ServerStream
}

type HelloService_StreamLengthClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}
```

由此可以发现：可以发现服务端和客户端的流辅助接口均定义了Send和Recv方法用于流数据的双向通信。

### 服务端流通信

服务端在循环中接收客户端发来的数据，如果遇到io.EOF表示客户端流被关闭，如果函数退出表示服务端流关闭。  
生成返回的数据通过流发送给客户端，双向流数据的发送和接收都是完全独立的行为。  
需要注意的是，发送和接收的操作并不需要一一对应，用户可以根据真实场景进行组织代码。

```go
func (h *HelloService) StreamLength(stream pb.HelloService_StreamLengthServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		reply :=  &pb.HelloResponse{Reply: int64(len(args.GetRes()))}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}
```

### 客户端流通信

客户端需要先调用 StreamLength 方法获取返回的流对象：

```go
stream, err := client.StreamLength(context.Background())
if err != nil {
    log.Fatal(err)
}
```

在客户端我们将发送和接收操作放到两个独立的 Goroutine。首先是向服务端发送数据；然后在循环中接收服务端返回的数据。

```go
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
```