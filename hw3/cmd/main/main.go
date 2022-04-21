package main

import (
	"flag"
	"fmt"
	"go2022/hw3/cmd"
	"go2022/hw3/pkg/crawler"
	"go2022/hw3/pkg/crawler/spider"
	"go2022/hw3/pkg/index/hash"
	"log"
	"sort"
)

var sites = []string{
	"https://go.dev/", "https://golang.org/",
}

func main() {
	var flagVar string
	var depthVar int
	flag.StringVar(&flagVar, "s", "", "string to scan")
	flag.IntVar(&depthVar, "d", 2, "scanning depth")
	flag.Parse()

	index := hash.New()
	scanner := spider.New()
	var data []crawler.Document
	for _, v := range sites {
		fmt.Println("scanning ", v)
		scan, err := scanner.Scan(v, depthVar)
		data = append(data, scan...)
		if err != nil {
			log.Fatal(err)
		}
	}
	index.Add(data)                 // Add indexes and keywords in index
	idxArr := index.Search(flagVar) // slice of indexes where our keyword is
	sort.Ints(idxArr)               //sorting
	for i := range data {
		key := cmd.BinSearch(idxArr, data[i].ID) // binary search in index slice
		if key != -1 {
			fmt.Println(data[i].URL)
		}
	}
}
