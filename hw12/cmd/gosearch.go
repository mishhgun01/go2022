package main

import (
	"fmt"
	"go2022/hw12/pkg/crawler/spider"
	"go2022/hw12/pkg/handlers"
	"go2022/hw12/pkg/webapp"
	"log"
	"net/http"
)

const depth, addr = 3, "0.0.0.0:8080"

var urls = []string{"https://go.dev", "https://golang.org"}

func main() {
	c := handlers.Docs{}
	scanner := spider.New()
	log.Print("scanning")
	for _, v := range urls {
		fmt.Println("Scanning ", v)
		scan, err := scanner.Scan(v, depth)
		if err != nil {
			log.Print(err)
			continue
		}
		c.Docs = append(c.Docs, scan...)
	}
	fmt.Println("Finished scan!")
	r := webapp.Router()
	r.HandleFunc("/docs", c.DocsHandler).Methods(http.MethodGet)
	r.HandleFunc("/index/{id}", c.IndexHandler).Methods(http.MethodGet)
	r.HandleFunc("/search/{query}", c.SearchHandler).Methods(http.MethodGet)
	webapp.ListenAndServe(addr, r)
}
