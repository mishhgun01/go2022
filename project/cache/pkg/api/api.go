package api

import (
	"github.com/gorilla/mux"
	"go2022/project/cache/pkg/cache"
	"log"
	"net/http"
)

type API struct {
	r     *mux.Router
	cache *cache.Storage
}

func New(cache *cache.Storage) *API {
	return &API{
		r:     mux.NewRouter(),
		cache: cache,
	}
}

func (api *API) Handle() {
	api.r.HandleFunc("/api/v1/url/{short}", api.url).Methods(http.MethodGet)
}

func (api *API) ListenAndServe(addr string) {
	log.Print("Listen on tcp://" + addr)
	http.ListenAndServe(addr, api.r)
}
