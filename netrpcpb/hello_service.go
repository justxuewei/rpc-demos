package netrpcpb

type HelloService struct {}

func (p *HelloService) Hello(request, reply *String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}
