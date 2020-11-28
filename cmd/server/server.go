package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/mukuro690528/chatroom/pkg/user"
)

var m = map[string]*user.Info{
	"001": {
		Name:  "Shan",
		Email: "shan@gmail.com",
		Age:   20,
		Phone: "0912345678",
	},
}

type HelloService struct{}

func (s *HelloService) Hello(req string, reply *string) error {
	*reply = "hello! " + req
	return nil
}

func (s *HelloService) UserUnfo(id string, info *user.Info) error {
	*info = *m[id]
	return nil
}

func main() {
	if err := rpc.RegisterName("HelloService", new(HelloService)); err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", "0.0.0.0:5288")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go rpc.ServeConn(conn)
	}
}
