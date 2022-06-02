package webapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go2022/hw13/pkg/crawler"
	"go2022/hw13/pkg/index/hash"
	"net/http"
	"sort"
	"strconv"
)

func (api *API) docs(w http.ResponseWriter, r *http.Request) {
	if len(api.d.Docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var ul string
	for _, doc := range api.d.Docs {
		ul += "<p><a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func (api *API) index(w http.ResponseWriter, r *http.Request) {
	if len(api.d.Docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	store := hash.New()
	store.Add(api.d.Docs)
	vars := mux.Vars(r)
	ids := store.Search(vars["query"])

	if len(ids) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var ul string
	for _, id := range ids {
		doc, err := search(id, api.d.Docs)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		ul += "<p><a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func (api *API) newDocument(w http.ResponseWriter, r *http.Request) {
	var doc crawler.Document
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	api.d.Docs = append(api.d.Docs, doc)
}

func (api *API) deleteDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for i := range api.d.Docs {
		if id == api.d.Docs[i].ID {
			api.d.Docs[i] = api.d.Docs[len(api.d.Docs)-1]
			api.d.Docs[len(api.d.Docs)-1] = crawler.Document{}
			api.d.Docs = api.d.Docs[:len(api.d.Docs)-1]
		}
	}
}

func (api *API) updateDocument(w http.ResponseWriter, r *http.Request) {
	var doc crawler.Document
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, v := range api.d.Docs {
		if doc.ID == v.ID {
			v = doc
		}
	}
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
