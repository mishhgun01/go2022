package main

import (
	"go2022/project/cache/pkg/api"
	"go2022/project/cache/pkg/cache"
	"time"
)

func main() {
	c := cache.New("", 24*7*time.Hour)
	a := api.New(c)
	a.ListenAndServe("")
}
