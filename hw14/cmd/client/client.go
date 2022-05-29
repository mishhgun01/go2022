package main

import (
	"fmt"
	"go2022/hw14/pkg/message"
	"log"
	"net/rpc"
)

var messages = []string{
	"hello",
	"my name is Bob",
	"I'm from Houston",
	"I like Go",
}

func main() {
	client, err := rpc.Dial("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	var req []message.Message
	for _, text := range messages {
		item := message.Message{Text: text}
		req = append(req, item)
	}
	err = client.Call("Server.Send", req, nil)
	if err != nil {
		log.Fatal(err)
	}
	var resp []message.Message
	err = client.Call("Server.Messages", new([]message.Message), &resp)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range resp {
		fmt.Println(item.String())
	}
}
