package service

import "net/rpc"

// HelloServiceName 定义服务名字
// 客户端调用时可以根据此包名路径调用服务端的方法
const HelloServiceName = "rpc/service.HelloService"

// HelloServiceInterface 定义服务要实现的方法列表
type HelloServiceInterface = interface {
	Length(request string, reply *int) error
}

// HelloServiceClient 自定义内部结构体
type HelloServiceClient struct {
	*rpc.Client
}

// 强制要求该类型也必须满足 HelloServiceInterface 接口
// 这样客户端用户就可以直接通过接口对应的方法调用RPC函数
// 也就是说 HelloServiceInterface 接口要求实现的方法，HelloServiceClient 在此内部必须实现
// 写法说明：
// 1、创建一个HelloServiceInterface地址，但不会分配内存的,并且如果给字段赋值会报错。
// 2、nil断言HelloServiceClient类型的零值，如果是结构体则为指针类型，如果为切片和映射则会nil类型。此处nil转换为一个空指针赋值给 HelloServiceClient 结构体。
// 另一种写法：var _ HelloServiceInterface = &HelloServiceClient{}
var _ HelloServiceInterface = (*HelloServiceClient)(nil)

// Length 封装客户端请求的具体方法
func (p *HelloServiceClient) Length(request string, reply *int) error {
	return p.Client.Call(HelloServiceName+".Length", request, reply)
}

// RegisterHelloService 服务端注册服务
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

// DialHelloService 客户端拨号服务
// 用于客户端的拨号连接
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}
