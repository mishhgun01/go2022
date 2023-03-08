package main

import (
	"go2022/hw18/links-service/pkg/api"
	"go2022/hw18/links-service/pkg/links"
)

var urls = []links.Link{
	{
		Short: "https://test.ru/googleKfk",
		Link:  "https://www.google.com/search?q=kafka+%D0%BA%D0%B0%D0%BA+%D0%BF%D0%BE%D0%B4%D0%BA%D0%BB%D1%8E%D1%87%D0%B8%D1%82%D1%8C%D1%81%D1%8F&oq=kafka+%D0%BF%D0%BE%D0%B4%D0%BA%D0%BB%D0%B1%D1%87%D0%B8%D1%82%D1%8C&aqs=chrome.1.69i57j0i8i13i30.13457j0j7&sourceid=chrome&ie=UTF-8",
	},
	{
		Short: "https://test.ru/mgRepoFiles",
		Link:  "https://github.com/mishhgun01/go2022/pull/28/files",
	},
}

func main() {
	db := links.New(urls)
	a := api.New(db)
	a.Handle()
	a.ListenAndServe("0.0.0.0:8080")
}
