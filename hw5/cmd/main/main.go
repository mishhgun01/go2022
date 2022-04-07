package main

import (
	"flag"
	"fmt"
	"go2022/hw5/cmd"
	"go2022/hw5/pkg/crawler"
	"go2022/hw5/pkg/crawler/spider"
	"log"
	"os"
)

const (
	urlGoDev     = "https://go.dev/"
	urlGolangOrg = "https://golang.org/"
)

func main() {

	var flagVar string
	var depthVar int
	flag.StringVar(&flagVar, "s", "", "string to scan")
	flag.IntVar(&depthVar, "d", 2, "scanning depth")
	flag.Parse()

	file, err := os.Open("output.txt")
	if file != nil {
		data, err := cmd.ReadFromFile(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)
		return
	}
	if file == nil {
		file, err = os.Create("output.txt")
		if err != nil {
			log.Fatal("unresolved error:", err)
		}
	}

	if err != nil {
		log.Fatal("unresolved error:", err)
	}
	scanner := spider.New()
	var data []crawler.Document
	scan, err := scanner.Scan(urlGoDev, depthVar)
	if err != nil {
		log.Fatal(err)
	}
	data = append(data, scan...)
	scan, err = scanner.Scan(urlGolangOrg, depthVar)
	if err != nil {
		log.Fatal(err)
	}
	data = append(data, scan...)

	b, err := cmd.InputInFile(file, data)
	if b == 0 || err != nil {
		log.Fatal("error with writing")
	}
}
