package main

import (
	"fmt"
	"testproto/service"

	"google.golang.org/protobuf/proto"
)

func main() {
	user := &service.User{
		Username: "muqing",
		Age:      24,
	}

	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}

	fmt.Println(newUser.String())
}
