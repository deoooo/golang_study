package server

import (
	"encoding/json"
	"fmt"
	"github.com/deoooo/im_demo/server/route"
	"github.com/deoooo/im_demo/server/route/v1"
	"io"
	"net"
)

type Server struct {
	ip         string
	port       uint
	orderRoute *route.OrderRoute
}

func NewServer(ip string, port uint) *Server {
	server := &Server{
		ip,
		port,
		&route.OrderRoute{},
	}
	return server
}

func (s *Server) connHandler(conn net.Conn) {
	fmt.Println("Handler address:", conn.RemoteAddr().String())
	defer conn.Close()
	msg := 2048
	for {
		msg := make([]byte, msg)
		l, err := conn.Read(msg)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client close:", conn.RemoteAddr())
				return
			}
			fmt.Println("read message err:", err)
		}
		// TODO 判断是否超出限制大小
		msg = msg[:l]
		// 非信息格式返回错误
		if !json.Valid(msg) {
			conn.Write([]byte("invalid message\n"))
			continue
		}
		s.orderRoute.MsgHandle(conn, msg)
	}
}

func (s *Server) Start() {

	s.orderRoute.Register("hello", v1.Hello)

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
		_, err = conn.Write([]byte("连接服务器成功...\n"))
		if err != nil {
			fmt.Println("返回连接成功信息失败:", err)
		}
		go s.connHandler(conn)
	}
}
