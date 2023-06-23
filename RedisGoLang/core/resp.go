package core

import (
	"errors"
	"strconv"
)

func decodeSimpleString(data []byte) (string, int, error) {
	if len(data) == 0 {
		return "", 0, nil
	}

	if data[0] != '+' {
		panic("Not a simple string")
	}

	pos := 1
	for ; data[pos] != '\r'; pos++ {
	}

	return string(data[1:pos]), pos + 2, nil
}

func decodeErrorString(data []byte) (string, int, error) {
	if len(data) == 0 {
		return "", 0, nil
	}

	if data[0] != '-' {
		panic("Not a error message")
	}

	pos := 1
	for ; data[pos] != '\r'; pos++ {
	}

	return string(data[1:pos]), pos + 2, nil
}

func decodeInt64(data []byte) (int64, int, error) {
	if len(data) == 0 {
		return 0, 0, nil
	}

	if data[0] != ':' {
		panic("Not a integer byte")
	}

	pos := 1
	for ; data[pos] != '\r'; pos++ {
	}

	value, err := strconv.ParseInt(string(data[1:pos]), 10, 64)
	if err != nil {
		panic(err)
	}

	return value, pos, nil
}

func decodeBulkString(data []byte) (string, int, error) {
	if len(data) == 0 {
		return "", 0, nil
	}

	if data[0] != '$' {
		panic("not a bulk string")
	}

	pos := 1
	var lenBulkString = 0
	for ; data[pos] != '\r'; pos++ {
		lenBulkString = lenBulkString*10 + int(data[pos]-'0')
	}

	//pos
}

// DecodeBytes : Decode bytes using the RESP protocol. Returns the data,
// position of the next byte and error if any
func DecodeBytes(data []byte) (interface{}, int, error) {

	if len(data) == 0 {
		return nil, 0, errors.New("no data")
	}

	switch data[0] {
	case '+':
		return decodeSimpleString(data)
	case '-':
		return decodeErrorString(data)
	case ':':
		return decodeInt64(data)

	}

	return nil, 0, nil
}
