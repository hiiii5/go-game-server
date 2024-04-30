package tcp

import (
	"bufio"
	"io"
	"log"
	"main/command"
	"net"
)

type TcpServerError struct {
	Err error
}

func (e TcpServerError) Error() string {
	return e.Err.Error()
}

type TcpConnection struct {
	Id         int
	Connection net.Conn
	Messages   []string

	server *TcpServer
}

func (c TcpConnection) handleConnection() {
	for {
		netData, err := bufio.NewReader(c.Connection).ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("User disconnected...")
				return
			}

			log.Fatalf("Error reading bytes from connection\n%s", err.Error())
			return
		}

		msg := string(netData)
		c.Messages = append(c.Messages, msg)
		log.Printf(msg)

		log.Println("Attempting to parse a command")

		cmd, err := command.TryParse(netData)
		if err != nil {
			log.Printf("Error parsing command\n%s", err.Error())
			continue
		} else {
			c.server.Commander.ExecuteCommand(cmd)
		}
	}
}

type TcpServer struct {
	Ip, Port     string
	listenSocket net.Listener
	IsStarting   bool
	IsRunning    bool
	Connections  []TcpConnection
	Commander    command.Commander
}

func (s TcpServer) Start() {
	log.Println("Starting server...")
	s.IsStarting = true
	err := s.startListening()
	if err != nil {
		log.Fatalf("Error starting server\n%s", err.Error())
		s.IsStarting = false
		s.IsRunning = false
		return
	} else {
		s.IsRunning = true
		s.IsStarting = false
	}

	log.Printf("Server started on %s:%s\n", s.Ip, s.Port)

	for {
		if !s.IsRunning {
			break
		}

		s.acceptConnection()
	}
}

func (s *TcpServer) Stop() {
	log.Println("Stopping server...")
	s.IsRunning = false
	s.stopListening()
}

func (s *TcpServer) startListening() error {
	s.IsRunning = true

	log.Printf("Starting to listen on %s:%s TCP\n", s.Ip, s.Port)

	listenSocket, err := net.Listen("tcp", s.Ip+":"+s.Port)
	if err != nil {
		log.Fatalf(err.Error())
		return TcpServerError{Err: err}
	}

	log.Printf("Listening on %s:%s TCP\n", s.Ip, s.Port)
	s.listenSocket = listenSocket
	return nil
}

func (s *TcpServer) stopListening() {
	s.listenSocket.Close()
}

func (s *TcpServer) acceptConnection() {
	if !s.IsRunning {
		return
	}

	log.Println("Polling for connection...")
	c, err := s.listenSocket.Accept()
	if err != nil {
		log.Fatalf("Error accepting connection\n%s", err.Error())
		s.Stop()
		return
	}

	s.Connections = append(s.Connections, TcpConnection{Id: len(s.Connections), Connection: c, server: s})

	log.Printf("Accepted connection from %s", c.RemoteAddr().String())
	go s.Connections[len(s.Connections)-1].handleConnection()
}

func NewTcpServer(ip string, port string) TcpServer {
	s := TcpServer{Ip: ip, Port: port, Commander: command.Commander{}}
	return s
}
