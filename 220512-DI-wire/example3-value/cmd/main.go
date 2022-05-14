package main

import (
	"fmt"
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
	app, _ := InitApp()
	fmt.Println("输出数据库配置：", app.Config.Database.Dsn)
}
