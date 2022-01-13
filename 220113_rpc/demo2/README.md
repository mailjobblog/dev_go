# 原生rpc服务分割优化

作为开发人员一般至少有三种角色：首选是服务端实现RPC方法的开发人员，其次是客户端调用RPC方法的人员，最后也是最重要的是制定服务端和客户端RPC接口规范的设计人员。  

在前面的例子中我们为了简化将以上几种角色的工作全部放到了一起，虽然看似实现简单，但是不利于后期的维护和工作的切割。  

## 优化内容

- 对于调用的方法进行封装
- 对于调用的函数和服务进行抽离

## 优化步骤

### 封装与抽离

**定义服务名称**

定义服务的名字格式为包的路径前缀+服务名称
```go
const HelloServiceName = "rpc/service.HelloService"
```

**定义服务接口**

定义服务要实现的接口的详细方法列表
```go
type HelloServiceInterface = interface {
	Length(request string, reply *int) error
}
```

**定义客户端调用结构体**

此结构体用于在此封装的程序中，实现要调用的服务的结构体和方法。这里的方法相当于一个代理者的角色，客户端由原来通过TCP直接调用服务端的方式，改变成了，客户端先调用此封装程序的方法，然后封装程序的方法请求服务端的方法。  
```go
type HelloServiceClient struct {
	*rpc.Client
}
```
强制要求定义的内部结构体 HelloServiceClient 实现 HelloServiceInterface 的方法。  
在上面的说明中，我们说到，客户端先请求 HelloServiceClient 的方法，然后 HelloServiceClient 的方法作为代理者请求服务 HelloServiceInterface 的方法，所以这就存在一个问题，如果让我们新定义的结构体 HelloServiceClient 也实现服务的方法？  
我们可以用下面的定义实现此需求。
```go
var _ HelloServiceInterface = (*HelloServiceClient)(nil)
```
弄了这么多，就是为了实现以下的代理者的方法
```go
func (p *HelloServiceClient) Length(request string, reply *int) error {
	return p.Client.Call(HelloServiceName+".Length", request, reply)
}
```

**服务端注册服务**

此方法是对服务端注册服务的封装，你可以直接调用此方法实现对于服务的rpc注册
```go
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
```

**客户端拨号服务**

此方法是对客户端注册服务的封装，你可以直接调用此方法实现对server的拨号
```go
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}
```

### 客户端调用

经过一系列的优化，客户端再次调用服务端的时候，可以很方便的调用服务对象

```go
func main() {
	// 连接服务
	client, err := service.DialHelloService("tcp", "127.0.0.1:8888")

	if err != nil {
		log.Fatal("dialing:", err)
	}
	res := "hello world"
	var reply int
	err = client.Length(res, &reply) // 方法调用
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
```