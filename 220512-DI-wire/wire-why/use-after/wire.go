//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

// 调用wire.Build方法传入所有的依赖对象以及构建最终对象的函数得到目标对象
func InitApp() (*App, error) {
	wire.Build(NewConfig, NewDB, NewApp)
	return &App{}, nil // 这里返回值没有实际意义，只需符合函数签名即可，生成的 wire_gen.go 会帮你包装该值
}
