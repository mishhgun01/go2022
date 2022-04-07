package main

import (
	"flag"
	"fmt"
	"go2022/hw2/pkg/crawler/spider"
	"log"
	"strings"
)

const (
	urlGoDev     = "https://go.dev/"
	urlGolangOrg = "https://golang.org/"
)

func main() {
	var flagVar string
	flag.StringVar(&flagVar, "s", "", "string to scan")
	flag.Parse()

	scanner := spider.New()
	data, err := scanner.Scan(urlGoDev, 2)
	data, err = scanner.Scan(urlGolangOrg, 2)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(data); i++ {
		if strings.Contains(strings.ToLower(data[i].URL), strings.ToLower(flagVar)) {
			fmt.Println(data[i].URL)
		}
	}
}
