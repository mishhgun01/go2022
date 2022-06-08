package webapp

import (
	"github.com/gorilla/mux"
	"go2022/hw13/pkg/crawler"
)

type Docs struct {
	Docs []crawler.Document
}

type API struct {
	r *mux.Router
	d Docs
}
