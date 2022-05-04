package main

import (
	"flag"
	"fmt"
	"go2022/hw3/pkg/crawler"
	"go2022/hw3/pkg/crawler/spider"
	"go2022/hw3/pkg/index/hash"
	"go2022/hw3/pkg/search"
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
	flag.IntVar(&depthVar, "d", 3, "scanning depth")
	flag.Parse()

	index := hash.New()
	scanner := spider.New()
	var data []crawler.Document
	for _, v := range sites {
		fmt.Println("scanning ", v)
		scan, err := scanner.Scan(v, depthVar)
		if err != nil {
			log.Print(err)
			continue
		}
		data = append(data, scan...)

	}
	// Add indexes and keywords in index
	index.Add(data)
	// slice of indexes where our keyword is
	idxArr := index.Search(flagVar)
	sort.Ints(idxArr)
	for i := range data {
		// binary search in index slice
		key := search.BinSearch(idxArr, data[i].ID)
		if key != -1 {
			fmt.Println(data[i].URL)
		}
	}
}
