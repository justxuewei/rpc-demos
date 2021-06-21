package netrpc3

const HelloServiceName = "netrpc3.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}
