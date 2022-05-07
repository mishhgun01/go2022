package main

import (
	"fmt"
	"go2022/hw11/pkg/crawler"
	"go2022/hw11/pkg/crawler/spider"
	"go2022/hw11/pkg/netsrv"
	"log"
	"net"
)

var sites = []string{
	"https://go.dev/", "https://golang.org/",
}

const network, addr, depth = "tcp4", "0.0.0.0:8000", 3

func main() {
	scanner := spider.New()
	var data []crawler.Document
	for _, v := range sites {
		fmt.Println("Scanning ", v)
		scan, err := scanner.Scan(v, depth)
		if err != nil {
			log.Print(err)
			continue
		}
		data = append(data, scan...)
	}
	fmt.Println("Finished scan!")
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go netsrv.Handler(conn, data)
	}
}
