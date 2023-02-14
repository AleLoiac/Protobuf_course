package main

import (
	"Protobuf/src/simple"
	"fmt"
)

func doSimple() {
	sm := simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Sample message",
		SampleList: []int32{1, 3, 5, 7},
	}
	fmt.Println(sm)

	fmt.Println(sm.GetId())
}

func main() {
	doSimple()
}
