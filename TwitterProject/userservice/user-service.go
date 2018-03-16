package main

import (
	"fmt"
	"os"

	"github.com/BrianCoveney/TwitterProject/transport"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats"
)

var twitterUsers map[string]string
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

	twitterUsers = make(map[string]string)
	twitterUsers["1"] = "Bob"
	twitterUsers["2"] = "John"
	twitterUsers["3"] = "Dan"
	twitterUsers["4"] = "Kate"

	nc.QueueSubscribe("UserNameById", "userNameByIdProviders", replyWithUserId)
	select {}
}

func replyWithUserId(m *nats.Msg) {

	myUser := Transport.User{}
	err := proto.Unmarshal(m.Data, &myUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	myUser.Name = twitterUsers[myUser.Id]
	data, err := proto.Marshal(&myUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Replying to ", m.Reply)
	nc.Publish(m.Reply, data)
}
