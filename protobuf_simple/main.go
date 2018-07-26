/*
1. install protoc
https://gist.github.com/sofyanhadia/37787e5ed098c97919b8c593f0ec44d8
2. write a proto file
/types/person.proto
3. generate go file
protoc --go_out=. ./types/*.proto
4. test
go run main.go
*/
package main

import (
	"log"

	"github.com/golang/protobuf/proto"

	"github.com/SUN-XIN/go-example/protobuf_simple/types"
)

func main() {
	loc := types.Location{
		Name: "paris",
	}
	p := types.Person{
		Name:      "toto",
		Age:       30,
		Locations: []*types.Location{&loc},
	}
	data, err := proto.Marshal(&p)
	if err != nil {
		log.Printf("Failed Marshal: %+v", err)
		return
	}

	log.Printf("encode get (%s)", data)

	var decodePers types.Person
	err = proto.Unmarshal(data, &decodePers)
	if err != nil {
		log.Printf("Failed Unmarshal: %+v", err)
		return
	}

	log.Printf("decode get (%+v)", decodePers)
}
