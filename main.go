package main

import (
	"Protobuf/src/simple"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"log"
	"os"
)

func doSimple() *simple.SimpleMessage {
	sm := simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Sample message",
		SampleList: []int32{1, 3, 5, 7},
	}

	fmt.Println(sm.GetId())

	return &sm
}

func writeToFile(fileName string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalf("Marshaling error: %v", err)
	}

	if err := os.WriteFile(fileName, out, 0644); err != nil {
		log.Fatalf("Can't write to file: %v", err)
	}

	fmt.Println("Data has been written")
}

func readFromFile(fileName string, pb proto.Message) {
	in, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Reading error: %v", err)
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalf("Unmarshalling problem: %v", err2)
	}
}

// encrypts the message to a JSON format
func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalf("Can't converto to JSON: %v", err)
	}
	return out
}

func main() {

	sm := doSimple()

	writeToFile("simple.bin", sm)

	sm2 := &simple.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)

	smAsString := toJSON(sm)
	fmt.Println(smAsString)
}
