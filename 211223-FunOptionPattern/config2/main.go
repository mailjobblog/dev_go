package main

import (
	"config2/server"
	"fmt"
)

func main() {
	s := server.New(server.Config{
		Host: "127.0.0.1",
		Port: 1234,
	})
	test, err := s.TestFunc()
	fmt.Println(test, err)
}
