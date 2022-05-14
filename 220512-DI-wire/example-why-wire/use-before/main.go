package main

import "fmt"

func main() {
	conf := NewConfig()
	db := NewDB(conf) // DB 依赖 Config
	result := db.Find()
	fmt.Println(result)
}
