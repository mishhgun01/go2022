package main

import (
	"go2022/hw18/links-service/pkg/api"
	"go2022/hw18/links-service/pkg/links"
)

func main() {
	db := links.DB{}
	api := api.New(&db)
	api.Handle()
	api.ListenAndServe("0.0.0.0:8080")
}
