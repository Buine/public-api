package tests

import (
	"os"
	"testing"

	"github.com/tesis/user-ms/cmd/httpserver"
)

var server httpserver.EchoServer

func TestMain(m *testing.M) {
	server = httpserver.EchoServer{}
	serverOn := make(chan bool)
	go server.RunServerAsync(serverOn)
	<-serverOn

	os.Exit(m.Run())
}
