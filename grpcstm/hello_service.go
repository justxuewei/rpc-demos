package grpcstm

import (
	"context"
	"io"
)

type HelloService struct {}

func (p *HelloService) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "hello: " + args.GetValue()}
	return reply, nil
}

func (p *HelloService) Channel(stream HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &String{
			Value: "hello: " + args.GetValue(),
		}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}
