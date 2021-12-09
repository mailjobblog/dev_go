package main

import (
    "context"
    "github.com/tencentyun/scf-go-lib/cloudfunction"
    "strconv"
)

func main() {
    // Make the handler available for Remote Procedure Call by Cloud Function
    // 调用 serverless 函数
    cloudfunction.Start(hello)
}

// serverless 函数
// 入参1：go语言内置的context，用于上下文
// 入参2：腾讯serverless传递的参数会放在这个里面
func hello(ctx context.Context, event DefineEvent) (ResultDiy, error) {
    // 接收传递的参数
    x := event.QueryString.X
    y := event.QueryString.Y

    // 业务实现
    s := strJoin(x,y)

    result := ResultDiy{
        ErrorCode: 200,
        String:  s,
    }

    return result, nil
}

// DefineEvent 入参2定义
// 该参数的定义，需要遵循腾讯serverless给的规范进行定义
// api网关触发集成概述：https://cloud.tencent.com/document/product/583/12513
type DefineEvent struct {
    QueryString struct {
        X string `json:"x"`
        Y string `json:"y"`
    } `json:"queryString"`
    HttpMethod string `json:"httpMethod"`
}

// ResultDiy 定义返回的内容
type ResultDiy struct {
    ErrorCode int64 `json:"errorCode"`
    String string `json:"string"`
}

// 模拟实体业务
func strJoin(x,y string)  string {
    length := strconv.Itoa(len(x + y))
    return length + "_" + x + "_" + y
}