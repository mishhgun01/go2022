package webapp

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router() *mux.Router {
	return mux.NewRouter()
}
func ListenAndServe(addr string, r http.Handler) {
	log.Print("LIsten on tcp://" + addr)
	http.ListenAndServe(addr, r)
}
