package api

import (
	"github.com/gorilla/mux"
	"go2022/hw18/links-service/pkg/queue"
	"log"
	"net/http"
)

const addr = "0.0.0.0:8080"

type API struct {
	s queue.Storage
	r *mux.Router
}

func New(s queue.Storage) *API {
	return &API{s: s, r: mux.NewRouter()}
}

func (api *API) Handle() {
	api.r.HandleFunc("/newLink/{query}", api.NewLink).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/link/{query}", api.Link).Methods(http.MethodGet, http.MethodOptions)
}

func (api *API) ListenAndServe() {
	log.Print("Listen on tcp://" + addr)
	http.ListenAndServe(addr, api.r)
}
