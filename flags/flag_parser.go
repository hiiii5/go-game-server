package flags

import (
	"flag"
)

type Flags struct {
	Ip   string
	Port string
}

func ParseFlags() Flags {
	ip := flag.String("ip", "127.0.0.1", "IP address to listen on")
	port := flag.String("port", "8006", "Port to listen on")

	flag.Parse()

	return Flags{Ip: *ip, Port: *port}
}
