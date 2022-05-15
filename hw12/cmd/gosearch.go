package main

import (
	"go2022/hw12/pkg/webapp"
)

func main() {
	var s webapp.Server
	s.New()
	s.Handle()
	s.ListenAndServe()
}
