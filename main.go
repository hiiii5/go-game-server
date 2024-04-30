package main

import "main/tcp"

func main() {
	tcp.NewTcpServer("127.0.0.1", "8006").Start()
}
