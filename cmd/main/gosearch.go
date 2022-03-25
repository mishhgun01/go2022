package main

import (
	"flag"
	"fmt"
	"go2022/pkg/crawler/spider"
	"log"
	"strings"
)

const (
	urlGoDev     = "https://go.dev/"
	urlGolangOrg = "https://golang.org/"
)

func main() {
	var ans string
	flag.StringVar(&ans, "s", "pkg", "flag")
	flag.Parse()

	serv := spider.New()
	data, err := serv.Scan(urlGoDev, 2)
	data, err = serv.Scan(urlGolangOrg, 2)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(data); i++ {
		if strings.Contains(data[i].URL, ans) {
			fmt.Println(data[i].URL)
		}
	}
}
