package handlers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go2022/hw12/pkg/crawler"
	"go2022/hw12/pkg/index/hash"
	"net/http"
	"sort"
	"strconv"
)

type Docs struct {
	Docs []crawler.Document
}

func (d *Docs) DocsHandler(w http.ResponseWriter, r *http.Request) {
	if len(d.Docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var ul string
	for _, doc := range d.Docs {
		ul += "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func (d *Docs) IndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	doc, err := search(id, d.Docs)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	ul := "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"" + doc.URL + "\">" + doc.Title + "</a></p>"
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func (d *Docs) SearchHandler(w http.ResponseWriter, r *http.Request) {
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
		ul += "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func search(id int, Docs []crawler.Document) (crawler.Document, error) {
	index := sort.Search(len(Docs), func(index int) bool { return Docs[index].ID >= id })
	if index >= len(Docs) || Docs[index].ID != id {
		doc := crawler.Document{}
		err := errors.New("поиск не дал результатов")
		return doc, err
	}
	return Docs[index], nil
}
