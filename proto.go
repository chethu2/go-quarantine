package main

import (
	"fmt"

	"github.com/golang/protobuf/jsonpb"
)

type tags struct {
	data string
	beta string
}

func main() {
	de := tags{}
	data := "{\"data\": \"a\", \"beta\": \"b\"}"
	err := jsonpb.Unmarshal([]byte(data), &de)
	fmt.Println("err de", err, de)
}
