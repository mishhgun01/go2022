package webapp

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const addr = "0.0.0.0:8080"

type Server struct {
	r mux.Router
}

func New() *Server {
	s := &Server{r: *(mux.NewRouter())}
	return s
}

func (s *Server) Handle(d Docs) {
	s.r.HandleFunc("/docs", d.docsHandler).Methods(http.MethodGet)
	s.r.HandleFunc("/index/{query}", d.indexHandler).Methods(http.MethodGet)
}

func (s *Server) ListenAndServe() {
	log.Print("Listen on tcp://" + addr)
	http.ListenAndServe(addr, &s.r)
}
