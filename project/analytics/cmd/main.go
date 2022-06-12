package main

import (
	"go2022/project/queue/pkg"
	"log"
)

func main() {
	opts := pkg.Options{}
	r, err := pkg.NewConsumer(opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
