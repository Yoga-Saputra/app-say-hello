package main

import (
	"fmt"

	go_say_hello "github.com/Yoga-Saputra/go-say-hello/v2"
)

// add dependency : go get github.com/Yoga-Saputra/go-say-hello
// update dependency : only update tag version in go.mod and run go get in terminal

func main() {
	fmt.Println(go_say_hello.SayHello("Yoga Saputra"))
}
