# 原生RPC和proto实现

由Go [原生RPC](https://github.com/mailjobblog/dev_go/tree/master/220113_rpc/demo1) 和 [protobuf](https://github.com/mailjobblog/dev_go/tree/master/220115_protobuf) 实现Go语言的RPC服务。 

## 如何测试

生成go代码
```bash
protoc --go_out=./ pb/*.proto
```

生成go代码
```bash
go run server.go
```

生成go代码
```bash
go run -run=TestClient -v
```