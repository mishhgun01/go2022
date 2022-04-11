package main

import (
	"flag"
	"fmt"
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
	var URLs []string
	URLs = append(URLs, urlGolangOrg, urlGoDev)
	var data []crawler.Document
	for _, v := range URLs {
		scan, err := scanner.Scan(v, depthVar)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, scan...)
	}
	index.Add(data)
	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})
	idxArr := index.Search(flagVar)
	sort.Ints(idxArr)
	for _, v := range data {
		key := binSearch(idxArr, v.ID) //бинарный поиск
		if key != -1 {
			fmt.Println(v.URL)
		}
	}
}

//бинарный поиск
func binSearch(arr []int, need int) int {
	lowKey := 0
	highKey := len(arr) - 1
	var index int
	if arr[lowKey] > need || arr[highKey] < need {
		index = -1
	}
	for lowKey <= highKey {
		mid := (lowKey + highKey) / 2
		if arr[mid] == need {
			return mid
		}
		if arr[mid] < need {
			lowKey = mid + 1
			continue
		}
		if arr[mid] > need {
			highKey = mid - 1
		}
	}
	return index
}
