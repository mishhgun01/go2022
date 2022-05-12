package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go2022/hw12/pkg/crawler"
	"go2022/hw12/pkg/crawler/spider"
	"go2022/hw12/pkg/index/hash"
	"go2022/hw12/pkg/router"
	"log"
	"net/http"
	"sort"
	"strconv"
)

const depth, addr, file = 3, "0.0.0.0:8080", "storage.json"

var urls = []string{"https://go.dev", "https://golang.org"}

type docs struct {
	docs []crawler.Document
}

func main() {
	c := docs{}
	scanner := spider.New()
	log.Print("scanning")
	for _, v := range urls {
		fmt.Println("Scanning ", v)
		scan, err := scanner.Scan(v, depth)
		if err != nil {
			log.Print(err)
			continue
		}
		c.docs = append(c.docs, scan...)
	}
	fmt.Println("Finished scan!")
	r := router.Router()
	r.HandleFunc("/docs", c.docsHandler).Methods(http.MethodGet)
	r.HandleFunc("/index/{id}", c.indexHandler).Methods(http.MethodGet)
	r.HandleFunc("/search/{query}", c.searchHandler).Methods(http.MethodGet)
	router.ListenAndServe(addr, r)
}

func (d *docs) docsHandler(w http.ResponseWriter, r *http.Request) {
	if len(d.docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var ul string
	for _, doc := range d.docs {
		ul += "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func (d *docs) indexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	doc, err := search(id, d.docs)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	ul := "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"" + doc.URL + "\">" + doc.Title + "</a></p>"
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func (d *docs) searchHandler(w http.ResponseWriter, r *http.Request) {
	if len(d.docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	store := hash.New()
	store.Add(d.docs)
	vars := mux.Vars(r)
	ids := store.Search(vars["query"])

	if len(ids) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var ul string
	for _, id := range ids {
		doc, err := search(id, d.docs)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		ul += "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func search(id int, docs []crawler.Document) (crawler.Document, error) {
	index := sort.Search(len(docs), func(index int) bool { return docs[index].ID >= id })
	if index >= len(docs) || docs[index].ID != id {
		doc := crawler.Document{}
		err := errors.New("поиск не дал результатов")
		return doc, err
	}
	return docs[index], nil
}
