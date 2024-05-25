package resp

import "fmt"

const (
	NULL_STRING = "$-1\r\n"
)

type RespParser struct {
}

type BulkString string

func toBulkString(str *string) BulkString {
	if str == nil {
		return BulkString(NULL_STRING)
	}
	return BulkString(fmt.Sprintf("$%d\r\n%s\r\n", len(*str), *str))
}

func (r *RespParser) toString(msg string) BulkString {
	return toBulkString(&msg)
}

func (r *RespParser) fromBulkString(msg BulkString) *string {
	if msg == NULL_STRING {
		return nil
	}
	return nil
}
