package webapp

import (
	"github.com/gorilla/mux"
	"go2022/hw12/pkg/crawler/spider"
	"log"
	"net/http"
)

const depth, addr = 3, "0.0.0.0:8080"

var urls = []string{"https://go.dev", "https://golang.org"}

type Server struct {
	r mux.Router
}

func (s *Server) New() {
	s = &Server{r: *(mux.NewRouter())}
}

func (s *Server) Handle() {
	d := scan()
	s.r.HandleFunc("/docs", d.docsHandler).Methods(http.MethodGet)
	s.r.HandleFunc("/index/{query}", d.indexHandler).Methods(http.MethodGet)
}

func (s *Server) ListenAndServe() {
	log.Print("Listen on tcp://" + addr)
	http.ListenAndServe(addr, &s.r)
}

func scan() Docs {
	scanner := spider.New()
	var d Docs
	for _, url := range urls {
		log.Println("Scanning ", url)
		scan, err := scanner.Scan(url, depth)
		if err != nil {
			log.Print(err)
			continue
		}
		d.Docs = append(d.Docs, scan...)
	}
	log.Println("Finished scan")
	return d
}
