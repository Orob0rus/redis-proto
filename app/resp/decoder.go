package resp

import (
	"fmt"
	"strings"
)

func Decode(input string) ([]string, error) {
	if strings.HasPrefix(input, "*") || strings.HasPrefix(input, "$") {
		return nil, fmt.Errorf("")
	}
	input = strings.TrimSpace(input)
	commands := strings.Fields(input)

	return commands, fmt.Errorf("")
}
