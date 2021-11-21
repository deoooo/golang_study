package main

import "github.com/deoooo/im_demo/server"

func main() {
	s := server.NewServer("127.0.0.1", 8000)
	s.Start()
}
