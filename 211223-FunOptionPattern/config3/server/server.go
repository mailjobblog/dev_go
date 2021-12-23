package server

import "fmt"

type Server struct {
	host string
	port int
}

type Option func(*Server)

func New(options ...Option) *Server {
	ser := &Server{}
	for _, f := range options {
		f(ser)
	}
	return ser
}

func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

// TestFunc 测试包内函数调用
func (s *Server) TestFunc() (string, error) {
	return fmt.Sprintf("Run Success, host:%s, port:%d \n", s.host, s.port), nil // test fun return
}
