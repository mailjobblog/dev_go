package main

import (
	"fmt"
	"wire-example/internal/db"
)

type App struct {
	dao db.IDao
}

func NewApp(i db.IDao) *App {
	return &App{
		dao: i,
	}
}

func main() {
	app, _ := InitApp()
	version, _ := app.dao.Version()
	fmt.Println(version)
}
