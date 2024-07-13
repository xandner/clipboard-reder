package server

import (
	"bufio"
	"fmt"
	"net"
)

func Main() {
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
		go handleConnection(conn) // handle each connection in a new goroutine
	}
}

func handleConnection(conn net.Conn) {
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
		conn.Write([]byte("ok\n"))
	}
}
