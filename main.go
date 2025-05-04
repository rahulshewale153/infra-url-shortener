package main

import "infra-url-shortener/server"

func main() {
	server := server.NewServer()
	server.Start()
}
