package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/nats-io/nats"
	"net/http"
	"os"
	"sync"
	"time"
)

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

	m := mux.NewRouter()
	m.HandleFunc("/{id}", handleUserWithTime)

	http.ListenAndServe(":3000", m)
}

func handleUserWithTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	myUser := Transport.User{Id: vars["id"]}
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		data, err := proto.Marshal(&myUser)
		if err != nil || len(myUser.Id) == 0 {
			fmt.Println(err)
			w.WriteHeader(500)
			fmt.Println("Problem with parsing the user Id.")
			return
		}

		msg, err := nc.Request("UserNameById", data, 100*time.Millisecond)
		if err == nil && msg != nil {
			myUserWithName := Transport.Tweet{}
			err := proto.Unmarshal(msg.Data, &myUserWithName)
			if err == nil {
				myUser = myUserWithName
			}
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Fprintln(w, "Hello ", myUser.Name, " with id ", myUser.Id, ", the time is ")
}
