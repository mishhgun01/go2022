package main

import (
	"go2022/hw18/links-service/pkg/api"
	"go2022/hw18/links-service/pkg/links"
)

var items = []links.Link{
	{
		Link:  "https://github.com/mishhgun01",
		Short: "github.com",
	},
	{
		Link:  "https://www.google.ru/",
		Short: "google.ru",
	},
}

func main() {
	db := links.New(items)
	api := api.New(db)
	api.Handle()
	api.ListenAndServe()
}
