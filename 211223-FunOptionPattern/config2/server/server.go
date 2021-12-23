package server

import (
	"fmt"
)

type Server struct {
	cfg Config
}

type Config struct {
	Host string
	Port int
}

func New(cfg Config) *Server {
	return &Server{cfg}
}

// TestFunc 测试包内函数调用
func (s *Server) TestFunc() (string, error) {
	return fmt.Sprintf("Run Success, host:%s, port:%d \n", s.cfg.Host, s.cfg.Port), nil // test fun return
}
