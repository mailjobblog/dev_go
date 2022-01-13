package service

import "net/rpc"

const HelloServiceName = "rpc/service.HelloService"

type HelloServiceInterface = interface {
	Length(request string, reply *int) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Length(request string, reply *int) error {
	return p.Client.Call(HelloServiceName+".Length", request, reply)
}
