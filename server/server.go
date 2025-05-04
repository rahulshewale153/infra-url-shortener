package server

import "log"

type server struct {
}

func NewServer() *server {
	return &server{}
}

func (s *server) Start() {
	log.Println("server started...")
}
