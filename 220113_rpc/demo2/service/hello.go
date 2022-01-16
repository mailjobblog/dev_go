package service

type HelloService struct{}

func (h *HelloService) Length(res string, reply *int) error {
	*reply = len(res)
	return nil
}
