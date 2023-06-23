package core_test

import (
	"RedisGoLang/core"
	"log"
	"testing"
)

func TestDecodeSimpleString(t *testing.T) {
	var str = "+OK\r\n"
	value, n, err := core.DecodeBytes([]byte(str))
	if err != nil || value != "OK" || n != 5 {
		t.Fail()
	}
}

func TestDecodeErrorString(t *testing.T) {
	var errString = "-ERR: ErrorMessage\r\n"
	value, _, err := core.DecodeBytes([]byte(errString))
	if err != nil || value != "ERR: ErrorMessage" {
		t.Fail()
	}
}

func TestDecodeInt64(t *testing.T) {
	var intbyte = ":1231232131232\r\n"
	value, _, err := core.DecodeBytes([]byte(intbyte))
	if err != nil || value != int64(1231232131232) {
		log.Println(value)
		t.Fail()
	}
}
