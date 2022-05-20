package main

import (
	"github.com/gorilla/mux"
	"go2022/hw13/pkg/crawler/spider"
	"go2022/hw13/pkg/webapp"
	"log"
)

const depth = 3

var urls = []string{"https://go.dev", "https://golang.org"}

func main() {
	r := mux.NewRouter()
	docs := scan()
	a := webapp.New(r, docs)
	a.Handle()
	a.ListenAndServe()
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
