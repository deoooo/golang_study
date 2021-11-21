package server

import (
	"fmt"
	"net"
)

type Server struct {
	ip   string
	port uint
}
k
func NewServer(ip string, port uint) *Server {
	server := &Server{
		ip,
		port,
	}
	return server
}

func (s *Server) handler(conn net.Conn) {
	fmt.Println("handler address:", conn.RemoteAddr().String())
}

func (s *Server) Start() {

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		fmt.Println("listen err:", err)
	}

	fmt.Printf("server listen on %s:%d \n", s.ip, s.port)

	defer listen.Close()

	for {
		// accept conn
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen err:", err)
		}
		go s.handler(conn)
	}
}
