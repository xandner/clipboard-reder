package server

import (
	"bufio"
	"clip/logger"
	"clip/usecase"
	"encoding/json"
	"fmt"
	"net"
)

type server struct {
	l logger.Logger
	u usecase.Clipboard
}
type Server interface {
	Main()
}

func NewServer(l logger.Logger,u usecase.Clipboard) Server {
	return &server{
		l,
		u,
	}
}

func (s *server) Main() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is listening on port 9999...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go s.handleConnection(conn) // handle each connection in a new goroutine
	}
}

func (s *server) handleConnection(conn net.Conn) {
	defer func() {
		fmt.Println("Closing connection")
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		fmt.Printf("Received: %s", buf)
		if string(buf)=="get_10\n"{
			e,data:=s.u.GetLast10()
			if e!=nil{
				conn.Write([]byte("error\n"))
				s.l.Error(fmt.Sprintf("Error while getting last 10 data %v", e))
				return
			}
			jsonData, err := json.Marshal(data)
			if err != nil {
				conn.Write([]byte("error\n"))
				s.l.Error(fmt.Sprintf("Error while marshalling data %v", err))
				return
			}
			conn.Write([]byte (jsonData))
		}
		conn.Write([]byte("ok\n"))
	}
}
