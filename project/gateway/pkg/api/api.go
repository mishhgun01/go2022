package api

import (
	"github.com/gorilla/mux"
	"go2022/project/cache/pkg/cache"
	"go2022/project/storage"
	"log"
	"net/http"
)

type API struct {
	r     *mux.Router
	db    *storage.Storage
	cache *cache.Storage
}

func New(db *storage.Storage) *API {
	return &API{
		r:  mux.NewRouter(),
		db: db,
	}
}

func (api *API) Handle() {
	api.r.HandleFunc("/api/v1/newLink/{url}", api.newUrl).Methods(http.MethodPost)
	api.r.HandleFunc("/api/v1/url/{short}", api.url).Methods(http.MethodGet)
}

func (api *API) ListenAndServe(addr string) {
	log.Print("Listen on tcp://" + addr)
	http.ListenAndServe(addr, api.r)
}
