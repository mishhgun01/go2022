package main

import (
	"flag"
	"fmt"
	"go2022/hw5/pkg/crawler"
	"go2022/hw5/pkg/crawler/spider"
	"go2022/hw5/pkg/fileworker"
	"log"
	"os"
	"strings"
)

var urls = []string{"https://go.dev/", "https://golang.org/"}

func main() {

	var flagVar string
	var depthVar int
	flag.StringVar(&flagVar, "s", "go", "string to scan")
	flag.IntVar(&depthVar, "d", 3, "scanning depth")
	flag.Parse()

	file, err := os.Open("output.txt")
	if err != nil {
		log.Print(err.Error())
	}
	if file == nil {
		file, err = os.Create("output.txt")
		defer file.Close()
		if err != nil {
			log.Fatal("unresolved error:", err)
		}
		scanner := spider.New()
		var data []crawler.Document
		for i := range urls {
			scan, err := scanner.Scan(urls[i], depthVar)
			fmt.Println("Scanning ", urls[i])
			if err != nil {
				log.Println(err.Error())
				continue
			}
			data = append(data, scan...)
		}

		b, err := fileworker.InputInFile(file, data)
		if b == 0 || err != nil {
			log.Fatal("error with writing: ", err)
		}
		return
	}
	file.Close()
	readFromFile, err := fileworker.ReadFromFile(file)
	if err != nil {
		log.Print(err.Error())
	}
	for i := range readFromFile {
		if strings.Contains(strings.ToLower(readFromFile[i]), strings.ToLower(flagVar)) {
			fmt.Println(readFromFile[i])
		}
	}
}
