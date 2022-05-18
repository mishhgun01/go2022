package main

import (
	"go2022/hw12/pkg/crawler/spider"
	"go2022/hw12/pkg/webapp"
	"log"
)

const depth = 3

var urls = []string{"https://go.dev", "https://golang.org"}

func main() {
	docs := scan()
	s := webapp.New()
	s.Handle(docs)
	s.ListenAndServe()
}

func scan() webapp.Docs {
	scanner := spider.New()
	var d webapp.Docs
	for _, url := range urls {
		log.Println("Scanning ", url)
		scan, err := scanner.Scan(url, depth)
		if err != nil {
			log.Print(err)
			continue
		}
		d.Docs = append(d.Docs, scan...)
	}
	log.Println("Finished scan")
	return d
}
