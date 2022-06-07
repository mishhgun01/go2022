package api

import (
	"github.com/gorilla/mux"
	"go2022/hw18/links-service/pkg/links"
	"log"
	"net/http"
)

type API struct {
	db *links.DB
	r  *mux.Router
}

func New(db *links.DB) *API {
	return &API{db: db, r: mux.NewRouter()}
}

func (api *API) Handle() {
	api.r.HandleFunc("/newLink/{query}", api.NewLink).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/link/{query}", api.Link).Methods(http.MethodGet, http.MethodOptions)
}

func (api *API) ListenAndServe(addr string) {
	log.Print("Listen on tcp://" + addr)
	http.ListenAndServe(addr, api.r)
}
