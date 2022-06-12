package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (api *API) url(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["short"]
	url := api.cache.Url(key)
	if url == "" {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}
}
