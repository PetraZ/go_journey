package main

import (
	"fmt"
	"io/ioutil"
	"os"

	proto "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

//no matter how json looks like, it translate to byte array
func jsonMarshal() {
	// Open our jsonFile
	jsonFile, err := os.Open("person.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(byteValue)
	fmt.Println(string(byteValue))
}

// protocol is a format, similar to xml and json
// to store structured data
// protoc --go_out=. person.proto
func main() {
	yi := &Person{
		Name: "yi",
		Age:  10,
	}
	_ = protoMarshal(yi)
	/*
		ref: https://developers.google.com/protocol-buffers/docs/encoding
		basically different data type would have different encoding methods
		int:
		first bit used to indicate if having following bytes
		8 bits as 1 byte
		10 --> 0000 1010

		how many types exist there?
		0 varint int32 int64....
		1 64-bit fixed64, double..
		2 length-delimited string, bytes
		5 32-bit fixed32, float...

		first byte (least significant 3 bits) contains info of encoding type
		for this case
		0001 0000, 000-> variant 10 --> 2 (field)

		string type is a bit different
		(10) 0000 1010 -->  010->wire_type=2 length-delimited 1->field type 1
		(2) 0000 0010 length of 2
		(121) y
		(105) i
		0100 1000
	*/
	jsonMarshal()
}

func protoMarshal(inp protoiface.MessageV1) []byte {
	data, err := proto.Marshal(inp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("binary of proto string is %v \n", data)
	return data
}
