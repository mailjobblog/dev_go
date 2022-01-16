# GRPC

## 如何测试

生成grpc代码
```bash
protoc -I. --go_out=. --go-grpc_out=. pb/*.proto
```
启动server
```bash
go run server.go
```
测试client
```bash
go test -v
```

### GRPC拓展学习

- gRPC证书认证：https://www.bookstack.cn/read/advanced-go-programming-book/ch4-rpc-ch4-05-grpc-hack.md
- Protobuf扩展：https://www.bookstack.cn/read/advanced-go-programming-book/ch4-rpc-ch4-06-grpc-ext.md

## 常见问题

### UnimplementedHelloServiceServer 的作用

之前我们说写业务逻辑的时候，写一个 XxxServerImpl 通过实现 XxxServer Interface 的方法来实现这个接口，从而可以将 XxxServerImpl 注册到 rpc 服务中去。像这样：
```go
// 第一种方法
 
// 通过下面代码将这个 service 注册到 grpc 服务中去
// pb.RegisterHelloWorldServer(s, &HelloWorldService{})
 
type HelloWorldService struct {
}
 
func (s *HelloWorldService) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Msg: fmt.Sprintf("Hello, %s!", req.GetName())}, nil
}
```
但是还有一种方法：  
因为在 protoc 帮我们生成的 .pb.go 文件中定义了 UnimplementedXxxServer 结构体，并且 *UnimplementedXxxServer 实现了 XxxServer 这个接口。  
所以我们写一个 XxxServerImpl，嵌入 *UnimplementedXxxServer 类型，也就实现了 XxxServer 这个接口。  
```go
// 第二种方法
 
// 通过下面代码将这个 service 注册到 grpc 服务中去
// pb.RegisterHelloWorldServer(s, &HelloWorldService{})
 
type HelloWorldService struct {
	*pb.UnimplementedHelloWorldServer
}
 
func (s *HelloWorldService) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Msg: fmt.Sprintf("Hello, %s!", req.GetName())}, nil
}
```
如果以后 .proto 协议中的 service 有变更函数，重新生成 .pb.go 文件后，XxxServer 这个接口也相应的变化了。  
这时候第一种方法会在编译时报错，因为 XxxServerImpl 写好的接口跟 XxxServer 不一样了，所以XxxServerImpl 没有实现 XxxServer，这样 Register 的时候类型不符合。
  
但是用第二种方法就不会报错。因为重新生成 .pb.go 文件的时候，protoc 帮我们重新生成了 UnimplementedXxxServer，它一定是实现了 XxxServer 这个接口的，所以我们的 XxxServerImpl 也是实现了 XxxServer 这个接口的。  
只不过在调用我们重写的 rpc 方法（例如这个SayHello）时，调用的可能就是嵌入类型 *pb.UnimplementedHelloWorldServer 它的 SayHello 了。  
（这涉及到 golang 的嵌入和组合：现有一个 Struct ，嵌入了一个其他类型，用外部类型调用某方法，如果外部类型包含了符合要求的接口实现，它的方法将会被使用。否则，通过方法提升，内部类型的接口实现可以直接被外部类型使用。）
