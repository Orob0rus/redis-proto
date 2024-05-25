package server

import (
	"fmt"
	"net"
	"os"
)

type Server struct {
	Handler Handler
}

func (s *Server) Init() {
	fmt.Println("Logs from your program will appear here!")

	lis, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleConn(&conn)
	}
}
