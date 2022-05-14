package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"wire-example/internal/biz"
	"wire-example/internal/server/config"
)

type App struct {
	conf      *config.Config
	db        *sql.DB
	OrderRepo biz.OrderRepo
}

func NewApp(cfg *config.Config, db *sql.DB, orderRepo biz.OrderRepo) *App {
	return &App{
		conf:      cfg,
		db:        db,
		OrderRepo: orderRepo,
	}
}

func main() {
	ctx := context.Background()
	app, cleanup, err := InitApp(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	// 测试 db 查询
	order, err := app.OrderRepo.Find(ctx, 1)
	if err != nil {
		fmt.Println("order find error:", err)
		return
	}
	fmt.Printf("查询成功， %+v", order)

	// 测试 db 插入
	var o = biz.Order{
		Name:  "测试上品",
		Type:  "1",
		Price: 9.9,
	}
	ID, err := app.OrderRepo.Create(ctx, &o)
	if err != nil {
		return
	}
	fmt.Printf("插入成功，id = %d", ID)
}
