package main

import (
	"fmt"

	pb "github.com/Harisalghifary/grpc-example/protofiles"
	"google.golang.org/protobuf/proto"
)

func main() {
	// create proto struct
	p := &pb.Person{
		Id:    1,
		Name:  "masamune date",
		Email: "masamune@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{
				Number: "072389",
				Type:   pb.Person_HOME,
			},
		},
	}

	// serialize the struct and attach to body
	body, _ := proto.Marshal(p)

	// deserialize the body and saving it to variable for testing
	p1 := &pb.Person{}
	_ = proto.Unmarshal(body, p1)

	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshalled proto data: ", body)
	fmt.Println("Unmarshalled struct: ", p1)

}
