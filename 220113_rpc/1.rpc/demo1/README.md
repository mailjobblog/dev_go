# Go原生方法实现RPC

## 如何测试

进入到 server 目录后，启动 server 服务

```bash
go run server.go
```

进入到 client 目录后，执行 client 程序，则会通过TCP协议请求 server

```bash
go run client.go
```

## 优化原生rpc

在 demo1 中，我们实现了一个简单的rpc服务，但是这样的服务存在一些新手经常遇到的问题，所以我们通过以下方式继续优化此rpc服务。

### 持续监听TCP

server.go是一个简易的程序，每次被 client 请求后就会被中断，所以可以用 for 进行持续的监听。

```go
func main() {
	_ = rpc.Register(new(HelloService))
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// for 持续监听tcp
	for {
		conn, err := listener.Accept()
		if err != nil {
		log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
```

### 采用http协议作为rpc载体

可以把函数方法注册到 http 协议上，方便调用者利用 http 的方式进行数据传递

```go
func main() {
	_ = rpc.Register(new(HelloService))

	// HTTP注册
	rpc.HandleHTTP()
	
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 阻塞监听服务
	http.Serve(listener, nil)
}
```

对于客户端 client.go 的调用也要改用 http 的方式

```go
rpc.Dial("tcp","127.0.0.1:8888")
改为
rpc.DialHTTP("tcp","127.0.0.1:8888")
```

### 客户端同步调用和异步调用

```go
func main() {
    client, err := rpc.Dial("tcp", "127.0.0.1:8888")
    if err != nil {
    log.Fatal("dialing:", err)
    }
    res := "test HelloService"
    var reply int
    
    // 同步调用
    err = client.Call("HelloService.Length", res, &reply)
    if err != nil {
    log.Fatal("同步调用失败，err：", err)
    }
    fmt.Println(reply)
    
    
    // 异步调用
    syncCall := client.Go("Arith.Opera", res, &reply, nil)
    // 阻塞，异步调用成功后解除阻塞
    replayDone := <-syncCall.Done
    fmt.Println(replayDone)
    fmt.Println(reply)
}
```