package main

import (
	"os"
	"fmt"
	"github.com/nats-io/go-nats"
	"github.com/golang/protobuf/proto"
	"github.com/BrianCoveney/nats2/transport"
)


// We use globals because it's a small application demonstrating NATS.
var users map[string]string
var nc *nats.Conn

func main() {

	uri := os.Getenv("NATS_URI")

	var err error

	nc, err = nats.Connect(uri)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to NATS server " + uri)

	users = make(map[string]string)
	users["1"] = "Brian"
	users["2"] = "Jones"
	users["3"] = "Mary"
	users["4"] = "Kate"

	nc.QueueSubscribe("UserNameById", "userNameByIdProviders", replyWithTwitterUserId)
	select {}
}

func replyWithTwitterUserId(m *nats.Msg) {

	myUser := Transport.User{}
	err := proto.Unmarshal(m.Data, &myUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	myUser.Name = users[myUser.Id]
	data, err := proto.Marshal(&myUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Replying to ", m.Reply)
	nc.Publish(m.Reply, data)
}