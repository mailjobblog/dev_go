package main

import (
	"fmt"
	"log"
	"wire-example/internal/config"
)

type App struct {
	Config *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{
		Config: cfg,
	}
}

func main() {
	app, cleanup, err := InitApp()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup() // 处理需要关闭的资源

	fmt.Println("输出数据库配置：", app.Config.Database.Dsn)
}
