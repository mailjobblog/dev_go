package main

import "fmt"

type App struct {
	db *DB
}

func NewApp(db *DB) *App {
	return &App{
		db: db,
	}
}

func main() {
	app, err := InitApp()
	if err != nil {
		return
	}
	result := app.db.Find()
	fmt.Println(result)
}
