package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"redis/app/commands"
	"redis/app/resp"
)

type Handler struct {
}

func (h *Handler) Handle(conn *net.Conn) {
	handleConn(conn)
}

func handleConn(conn *net.Conn) {
	defer func() {
		if err := (*conn).Close(); err != nil {
			fmt.Printf("[%s] occurred while closing the connection", err.Error())
		}
	}()
	log.Println("Handler interface invoked")

	reader := bufio.NewReader(*conn)

	// while there is something to read

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read stream finised on TCP connection, terminating reader")
			return
		}

		params, err := resp.Decode(msg)

		if err != nil {
			log.Fatalf("Error %v parsing the command: %s\n", err, msg)
			return
		}

		response, err := commands.ParseAndServeCommand(params)
		if err != nil {
			log.Fatalf("Error %v parsing the command: %s\n", err, msg)
			return
		}

		encodedResponse := resp.EncodeResp(response)

		_, err = (*conn).Write([]byte(encodedResponse))
		if err != nil {
			fmt.Println("Failed while responding to the request")
		}
	}
}
