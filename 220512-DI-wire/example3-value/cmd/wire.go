//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-example/internal/config"
)

func InitApp() (*App, error) {
	// 绑定值用 wire.Value 进行绑定
	// 这里首先绑定了一个 string 类型的值
	// 然后绑定了 String2 类型的值，因为本例子需要绑定两个 string 类型的值，
	// 如果都用 string 那么注入的时候 wire 无法区分具体的 string, 所以另外一个 string 使用自定义string类型
	wire.Build(config.Provider, wire.Value("demo string1"), wire.Value(config.String2("demo string 2")), NewApp)
	return &App{}, nil
}
