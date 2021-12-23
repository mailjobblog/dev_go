package server

import "fmt"

type Server struct {
	host string
	port int
}

func New(host string, port int) *Server {
	return &Server{host, port}
}

// TestFunc 测试包内函数调用
func (s *Server) TestFunc() (string, error) {
	return fmt.Sprintf("Run Success, host:%s, port:%d \n", s.host, s.port), nil // test fun return
}
