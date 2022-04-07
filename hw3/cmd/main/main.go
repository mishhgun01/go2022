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

	index := hash.New()
	scanner := spider.New()
	var data []crawler.Document
	scanDev, err := scanner.Scan(urlGoDev, depthVar)
	if err != nil {
		log.Fatal(err)
	}
	scanOrg, err := scanner.Scan(urlGolangOrg, depthVar)
	data = append(data, scanOrg...)
	data = append(data, scanDev...)
	index.Add(data)
	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})
	idxArr := index.Search(flagVar)
	sort.Ints(idxArr)
	for _, v := range data {
		key := cmd.BinSearch(idxArr, v.ID) //бинарный поиск
		if key != -1 {
			fmt.Println(v.URL)
		}
	}
}
