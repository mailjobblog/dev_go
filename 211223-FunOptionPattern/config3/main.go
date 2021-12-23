package main

import (
	"config3/server"
	"fmt"
)

func main() {
	s := server.New(
		server.WithHost("127.0.0.1"),
		server.WithPort(1234),
	)
	test, err := s.TestFunc()
	fmt.Println(test, err)
}
