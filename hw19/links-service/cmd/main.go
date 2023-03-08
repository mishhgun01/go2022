package main

import (
	"go2022/hw19/links-service/pkg/api"
	"go2022/hw19/links-service/pkg/links"
	"log"
	"os"
	"strings"
)

var (
	brokers = strings.Split(os.Getenv("KFK_BROKERS"), ",")
	topic   = os.Getenv("KFK_TOPIC")
	address = os.Getenv("LINKS_ADDRESS")
)

func main() {
	db := links.New(make([]links.Link, 1))
	opts := api.KfkOptions{
		brokers,
		topic,
	}
	apiLinks, err := api.New(opts, db)

	if err != nil {
		log.Fatal(err.Error())
	}

	apiLinks.Handle()
	err = apiLinks.ListenAndServe(address)
	if err != nil {
		log.Fatal(err.Error())
	}
}
