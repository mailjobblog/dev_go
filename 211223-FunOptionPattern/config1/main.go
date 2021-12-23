package main

import (
	"config1/server"
	"fmt"
)

func main() {
	s := server.New("127.0.0.1", 1234)
	test, err := s.TestFunc()
	fmt.Println(test, err)
}
