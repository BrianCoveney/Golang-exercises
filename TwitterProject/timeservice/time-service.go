package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats"
	"os"
	"time"

	"github.com/BrianCoveney/TwitterProject/transport"
)

// We use globals because it's a small application demonstrating NATS.

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

	nc.QueueSubscribe("TimeTeller", "TimeTellers", replyWithTime)
	select {} // Block forever
}

func replyWithTime(m *nats.Msg) {
	curTime := Transport.Time{time.Now().Format(time.RFC3339)}

	data, err := proto.Marshal(&curTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Replying to ", m.Reply)
	nc.Publish(m.Reply, data)

}
