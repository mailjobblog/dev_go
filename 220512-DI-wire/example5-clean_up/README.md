# Wire 清理函数

在创建依赖资源时，如果由某个资源创建失败，那么其他资源需要关闭的情况下，可以使用cleanup函数来关闭资源。比如咱们给`config.New`方法返回一个`cleanup`函数来关闭文件句柄资源，相关代码修改如下：

cmd/wire.go

```go
package main

// ... ...

func InitApp() (*App, func(), error) {
	wire.Build(config.Provider, NewApp)
	return &App{}, func() {
	}, nil
}
```

internal/config/config.go

```go
package config

// ... ...

func New() (*Config, func(), error) {
	fp, err := os.Open("../config/app.json")
	if err != nil {
		return nil, func() {}, err
	}
	// defer fp.Close() // 这里注释了，交由 cleanup 处理
	var cfg Config
	if err := json.NewDecoder(fp).Decode(&cfg); err != nil {
		return nil, func() {
			fp.Close()
		}, err
	}
	return &cfg, func() {
		fp.Close()
		fmt.Println("app.json 资源句柄关闭成功")
	}, nil
}
```

cmd/main.go

```go
package main

// ... ...

func main() {
	app, cleanup, err := InitApp()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup() // 处理需要关闭的资源

	fmt.Println("输出数据库配置：", app.Config.Database.Dsn)
}
```



