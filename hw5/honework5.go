package main

import (
	"flag"
	"fmt"
	"go2022/pkg/crawler"
	"go2022/pkg/crawler/spider"
	"io"
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
	if err != nil {
		file, err = os.Create("output.txt")
		if err != nil {
			log.Fatal("unresolved error:", err)
		}
		scanner := spider.New()
		data, err := scanner.Scan(urlGoDev, depthVar)
		if err != nil {
			log.Fatal("unresolved error:", err)
		}
		data, err = scanner.Scan(urlGolangOrg, depthVar)
		if err != nil {
			log.Fatal("unresolved error:", err)
		}
		b, err := inputInFile(file, data)
		if b == 0 || err != nil {
			log.Fatal("error with writing")
		}
	}
	data, err := readFromFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func inputInFile(writer io.Writer, data []crawler.Document) (int, error) {
	bytesRead := 0
	if data == nil {
		return 0, nil
	}
	var err error
	for _, v := range data {
		bytesRead, err = writer.Write([]byte(v.URL + "\n"))
		if err != nil {
			return 0, err
		}
	}
	return bytesRead, nil
}

func readFromFile(reader io.Reader) ([]string, error) {
	output := make([]byte, 64)
	var URLS []string
	for {
		_, err := reader.Read(output)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		URLS = append(URLS, string(output))
	}
	return URLS, nil
}
