package tcp

import (
	"net"
	"testing"
	"time"
)

var server TcpServer

func init() {
	server = NewTcpServer("127.0.0.1", "8006")

	go server.Start()
}

func TestServerConnection(t *testing.T) {
	t.Log("Waiting for server to start...")

	var conn net.Conn
	var err error

	sleepTime := 1
	sleepTimeGrowth := 1

	// In practice its more like 1 or 2 seconds, but I want to do it more because why not.
	for i := 0; i < 5; i++ {
		t.Logf("Sleeping for %d seconds\n", sleepTime)

		conn, err = net.Dial("tcp", "127.0.0.1:8006")
		if err == nil {
			conn.Close()
			return
		}

		sleepTime += sleepTimeGrowth
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}

	t.Fatalf("Error connecting to server\n%s", err.Error())
}
