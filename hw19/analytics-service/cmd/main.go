package main

import (
	"go2022/hw19/analytics-service/pkg/storage"
	"log"
	"os"
	"strings"
)

var (
	brokers = strings.Split(os.Getenv("KFK_BROKERS"), ",")
	groupId = os.Getenv("KFK_GROUP_ID")
	topic   = os.Getenv("KFK_TOPIC")
)

func main() {
	options := storage.ConsumerOptions{
		brokers,
		groupId,
		topic,
	}

	analytics, err := storage.New(options)

	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		analytics.Update()
	}
}
