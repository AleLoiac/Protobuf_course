package main

import (
	"Protobuf/src/complex"
	"Protobuf/src/enum"
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

	//fmt.Println(sm.GetId())

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

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalf("Can't converto to JSON: %v", err)
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalf("Counldn't unmarshal the JSON: %v", err)
	}
}

func doEnum() {
	em := enum.EnumMessage{
		Id:           42,
		DayOfTheWeek: enum.DayOfTheWeek_FRIDAY,
	}

	em.DayOfTheWeek = enum.DayOfTheWeek_WEDNESDAY

	//fmt.Println(em)

	fmt.Println(em.GetDayOfTheWeek())
}

func doComplex() {
	cm := complex.ComplexMessage{
		One_Message: &complex.BasicMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleMessage: []*complex.BasicMessage{
			&complex.BasicMessage{
				Id:   2,
				Name: "Second message",
			},
			&complex.BasicMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	fmt.Println(cm)
}

func main() {

	// creates a simple message
	sm := doSimple()

	// writes the message to a binary file
	writeToFile("simple.bin", sm)

	// reads the binary file and unmarshals the content to a protobuf message
	sm2 := &simple.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)

	// encrypts the message to a JSON format
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	// takes a string in a JSON format and unmarshals it to a protobuf message
	sm3 := &simple.SimpleMessage{}
	fromJSON(smAsString, sm3)
	fmt.Println(sm3)

	// creates an enum
	doEnum()

	// creates nested messages
	doComplex()
}
