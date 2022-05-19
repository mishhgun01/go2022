package webapp

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go2022/hw12/pkg/crawler"
	"go2022/hw12/pkg/index/hash"
	"net/http"
	"sort"
)

type Docs struct {
	Docs []crawler.Document
}

func (d *Docs) docsHandler(w http.ResponseWriter, r *http.Request) {
	if len(d.Docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var ul string
	for _, doc := range d.Docs {
		ul += "<p><a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func (d *Docs) indexHandler(w http.ResponseWriter, r *http.Request) {
	if len(d.Docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	store := hash.New()
	store.Add(d.Docs)
	vars := mux.Vars(r)
	ids := store.Search(vars["query"])

	if len(ids) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var ul string
	for _, id := range ids {
		doc, err := search(id, d.Docs)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		ul += "<p><a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func search(id int, Docs []crawler.Document) (crawler.Document, error) {
	index := sort.Search(len(Docs), func(index int) bool { return Docs[index].ID >= id })
	if index >= len(Docs) || Docs[index].ID != id {
		doc := crawler.Document{}
		err := errors.New("No results")
		return doc, err
	}
	return Docs[index], nil
}
