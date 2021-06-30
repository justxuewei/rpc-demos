package grpcdm

import "context"

type HelloService struct {}

func (p *HelloService) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "hello: " + args.GetValue()}
	return reply, nil
}
