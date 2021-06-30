package rpcctx

import (
	"fmt"
	"log"
	"net"
)

type HelloService struct {
	conn    net.Conn
	isLogin bool
}

func NewHelloService(conn net.Conn) *HelloService {
	return &HelloService{
		conn: conn,
		isLogin: false,
	}
}

func (p *HelloService) Login(request string, reply *string) error {
	if request != "usr:pwd" {
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	p.isLogin = true
	return nil
}

func (p *HelloService) Hello(request string, reply *string) error {
	if !p.isLogin {
		return fmt.Errorf("please login")
	}
	*reply = "hello: " + request + ", from " + p.conn.RemoteAddr().String()
	return nil
}
