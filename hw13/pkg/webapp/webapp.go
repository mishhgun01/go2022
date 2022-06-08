package webapp

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const addr = "0.0.0.0:8080"

func New(router *mux.Router, d Docs) *API {
	api := API{r: router, d: d}
	return &api
}

func (api *API) Handle() {
	api.r.HandleFunc("/api/v1/docs", api.docs).Methods(http.MethodGet)
	api.r.HandleFunc("/api/v1/index/{query}", api.index).Methods(http.MethodGet)
	api.r.HandleFunc("/api/v1/newDoc", api.newDocument).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/v1/deleteDoc/{id}", api.deleteDocument).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/v1/updateDoc/{id}", api.updateDocument).Methods(http.MethodPost, http.MethodOptions)
}

func (api *API) ListenAndServe() {
	log.Print("Listen on tcp://" + addr)
	http.ListenAndServe(addr, api.r)
}
