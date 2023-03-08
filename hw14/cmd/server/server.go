package main

import (
	"go2022/hw14/pkg/message"
	"log"
	"net"
	"net/rpc"
	"sync"
)

type Server struct {
	mu       sync.Mutex
	ID       int
	messages []message.Message
}

func (s *Server) Send(req []message.Message, _ *string) error {
	for _, item := range req {
		s.mu.Lock()
		item.ID = s.ID + 1

		s.messages = append(s.messages, item)
		s.ID += 1
		s.mu.Unlock()
	}
	return nil
}

func (s *Server) Messages(_ []message.Message, resp *[]message.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	*resp = s.messages
	return nil
}

func main() {
	srv := new(Server)
	err := rpc.Register(srv)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.Listen("tcp4", ":8080")
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
