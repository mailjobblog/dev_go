package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"protobuf/proto/api"
)

func main() {
	apiUser := &api.User{
		Name:   "kun",
		Status: true,
		Hobby:  []string{"唱歌", "跳舞", "rap"},
	}

	MarshalUser, err := proto.Marshal(apiUser)
	if err != nil {
		log.Fatal("Marshal error: ", err)
	}

	user := &api.User{}
	err = proto.Unmarshal(MarshalUser, user)
	if err != nil {
		log.Fatal("Unmarshal error: ", err)
	}

	if apiUser.GetName() != user.GetName() {
		log.Fatalf("data match error, %q != %q", apiUser.GetName(), user.GetName())
	}

	// success
	fmt.Println(apiUser.GetName(), user.GetName())
}
