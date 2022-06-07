package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (api *API) Link(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["query"]
	if len(short) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	link := api.db.Link(short)
	if len(link) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	ul := "<p><a href=\"/index/\">" + link + "</a></p>"
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func (api *API) NewLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	link := vars["query"]
	if len(link) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	short := api.db.NewLink(link)
	if len(short) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	ul := "<p><a href=\"/index/\">" + short + "</a></p>"
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}
