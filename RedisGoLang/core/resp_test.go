package core_test

import (
	"RedisGoLang/core"
	"log"
	"testing"
	"fmt"
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

func TestBulkString(t *testing.T) {
	var bulkString = "$4\r\nTEST\r\n"
	value, _, err := core.DecodeBytes([]byte(bulkString))
	if err != nil || value != "TEST" {
		log.Println(value)
		t.Fail()
	}
}

func TestArray(t *testing.T) {
	cases := map[string][]interface{}{
		"*0\r\n":                                                   {},
		"*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n":                     {"hello", "world"},
		"*3\r\n:1\r\n:2\r\n:3\r\n":                                 {int64(1), int64(2), int64(3)},
		"*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$5\r\nhello\r\n":            {int64(1), int64(2), int64(3), int64(4), "hello"},
		"*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Hello\r\n-World\r\n": {[]int64{int64(1), int64(2), int64(3)}, []interface{}{"Hello", "World"}},
	}
	for k, v := range cases {
		value, _ , err := core.DecodeArray([]byte(k))
		if err != nil {
			log.Println(err)
			t.Fail()
		}
		array := value.([]interface{})
		if len(array) != len(v) {
			t.Fail()
		}
		for i := range array {
			if fmt.Sprintf("%v", v[i]) != fmt.Sprintf("%v", array[i]) {
				t.Fail()
			}
		}
	}
}