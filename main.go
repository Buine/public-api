package main

import "github.com/tesis/user-ms/cmd/httpserver"

var Server httpserver.EchoServer

func main() {
	Server = httpserver.EchoServer{}
	Server.RunServer()
}
