# 初次尝试serverless

## 参考文档

- 代码开发：https://cloud.tencent.com/document/product/583/18032#.E6.9B.B4.E5.A4.9A.E6.8C.87.E5.BC.95
- Api网关触发：https://cloud.tencent.com/document/product/583/12513
- 腾讯serverless-github：https://github.com/tencentyun/scf-go-lib

## 如何使用

依赖安装

```bash
go mod tidy
```

编译打包

```bash
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```

最后将打包好的 `main.zip` 上传至 `函数代码` 中，进行 `部署`。

## 初始下载的代码备份

```go
package main

import (
    "context"
    "fmt"
    "github.com/tencentyun/scf-go-lib/cloudfunction"
)

type DefineEvent struct {
    // test event define
    Key1 string `json:"key1"`
    Key2 string `json:"key2"`
}

func hello(ctx context.Context, event DefineEvent) (string, error) {
    fmt.Println("key1:", event.Key1)
    fmt.Println("key2:", event.Key2)
    return fmt.Sprintf("Hello %s!", event.Key1), nil
}

func main() {
    // Make the handler available for Remote Procedure Call by Cloud Function
    cloudfunction.Start(hello)
}
```