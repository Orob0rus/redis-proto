package commands

import (
	"fmt"
	"strings"
)

func ParseAndServeCommand(commands []string) ([]string, error) {
	if len(commands) < 1 {
		return nil, fmt.Errorf("no valid message found")
	}
	fmt.Println(commands)
	switch strings.ToUpper(commands[0]) {
	case "PING":
		return generatePingResponse(commands), nil
	case "ECHO":
		return generateEchoResponse(commands), nil
	}
	return nil, fmt.Errorf("unsupported command found")
}

func generatePingResponse(commands []string) []string {
	if len(commands) == 1 {
		return []string{"+PONG\r\n"}
	}
	return []string{strings.Join(commands[1:], " ")}
}

func generateEchoResponse(commands []string) []string {
	return nil
}
