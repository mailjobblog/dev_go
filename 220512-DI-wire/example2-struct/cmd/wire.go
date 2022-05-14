//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitApp() (*FooMessage, error) {
	// 通过 wire.Struct 来指定那些字段要被注入到结构体中
	// 这里的 Msg、Ber 代表要导入的字段
	// 如果你要全部导入，可以这样写： wire.Struct(new(FooMessage), "*")
	wire.Build(ProvideMessage, ProvideBeer, wire.Struct(new(FooMessage), "Msg", "Ber"))
	return &FooMessage{}, nil
}

// 因为入参均是字符串类型, wire 无法区分具体的 string
// 所以这里我们使用自定义类型
type Message string
type Beer string

// 测试构造结构体
type FooMessage struct {
	Msg Message
	Ber Beer
}

func ProvideMessage() Message {
	return "info"
}

func ProvideBeer() Beer {
	return "xuehua666"
}
