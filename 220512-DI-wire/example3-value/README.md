# Wire 值绑定

在此 demo 中，我们绑定了两个 string 类型的值，然后传递给 Config ，并在 config.New 方法中打印输出。

cmd/wire.go

```go
//go:build wireinject
// +build wireinject

package main

// ... ...

func InitApp() (*App, error) {
	// 绑定两个string
	wire.Build(config.Provider, wire.Value("demo string1"), wire.Value(config.String2("demo string 2")), NewApp)
	return &App{}, nil
}

```

internal/config/config.go

```go
package config

// ... ...

var Provider = wire.NewSet(New) // 将New方法声明为Provider，表示New方法可以创建一个被别人依赖的对象,也就是Config对象

type String2 string

func New(s string, s2 String2) (*Config, error) {

    // 输出绑定的两个 string
	fmt.Println("print wire value：" + s)
	fmt.Println("print wire value：" + s2)

	var cfg Config
	return &cfg, nil
}
```

**代码分析**

注意在 New 方法中，形参如果都写 string 那么 wire 无法区分具体的 string。所以第二个 string 采用自定义类型的 string

最后打印输出绑定的两个 string 验证是否正确。



