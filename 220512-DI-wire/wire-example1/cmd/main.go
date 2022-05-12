package main

import (
	"database/sql"
	"log"
)

type App struct { // 最终需要的对象
	db *sql.DB // db可以自定义命名，*sql.DB 需要和 internal/db/db.go 中 NewDb 方法返回的类型相同
}

func NewApp(db *sql.DB) *App {
	return &App{
		db: db,
	}
}

func main() {
	app, err := InitApp() // 使用 wire 生成的 injector 方法获取app对象
	if err != nil {
		log.Fatal(err)
	}

	// 测试数据库连接
	var version string
	row := app.db.QueryRow("SELECT VERSION()")
	if err := row.Scan(&version); err != nil {
		log.Fatal(err)
	}
	log.Println(version)
}
