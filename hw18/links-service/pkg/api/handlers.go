package api

import (
	"encoding/json"
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
	http.Redirect(w, r, link, http.StatusSeeOther)
}

func (api *API) NewLink(w http.ResponseWriter, r *http.Request) {
	var s struct{ link string }
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	if len(s.link) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	short := api.db.NewLink(s.link)
	if len(short) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	ul := "<p><a href=\"/index/\">" + short + "</a></p>"
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}
