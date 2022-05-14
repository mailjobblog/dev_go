//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-example/internal/config"
	"wire-example/internal/db"
)

// 调用wire.Build方法传入所有的依赖对象以及构建最终对象的函数得到目标对象
func InitApp() (*App, error) {
	// 写法1（参考Kratos框架写法）
	// panic(wire.Build(config.Provider, db.Provider, NewApp))

	// 写法2（参考wire官方文档写法）
	wire.Build(config.Provider, db.Provider, NewApp)
	return &App{}, nil // 这里返回值没有实际意义，只需符合函数签名即可，生成的 wire_gen.go 会帮你包装该值
}
