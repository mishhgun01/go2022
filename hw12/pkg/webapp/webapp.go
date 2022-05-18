package webapp

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const addr = "0.0.0.0:8080"

type Server struct {
	r mux.Router
	d Docs
}

func New(d Docs) *Server {
	s := &Server{r: *(mux.NewRouter()), d: d}
	return s
}

func (s *Server) Handle() {
	s.r.HandleFunc("/docs", s.d.docsHandler).Methods(http.MethodGet)
	s.r.HandleFunc("/index/{query}", s.d.indexHandler).Methods(http.MethodGet)
}

func (s *Server) ListenAndServe() {
	log.Print("Listen on tcp://" + addr)
	http.ListenAndServe(addr, &s.r)
}
