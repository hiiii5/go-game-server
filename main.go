package main

import (
	"main/flags"
	"main/tcp"
)

func main() {
	cmdFlags := flags.ParseFlags()
	tcp.NewTcpServer(cmdFlags.Ip, cmdFlags.Port).Start()
}
