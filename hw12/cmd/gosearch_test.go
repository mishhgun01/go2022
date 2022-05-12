package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go2022/hw12/pkg/crawler"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var documents docs
var r *mux.Router

func TestMain(m *testing.M) {
	r = mux.NewRouter()
	r.HandleFunc("/docs", documents.docsHandler).Methods(http.MethodGet)
	r.HandleFunc("/index/{id}", documents.indexHandler).Methods(http.MethodGet)
	r.HandleFunc("/search/{query}", documents.searchHandler).Methods(http.MethodGet)
	os.Exit(m.Run())
}

func Test_crawlerDocs_docsHandler(t *testing.T) {
	documents = docs{docs: []crawler.Document{}}
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	req.Header.Add("content-type", "plain/text")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusNoContent {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusNoContent)
	}

	var ul string
	documents.docs = []crawler.Document{
		{
			ID:    0,
			URL:   "https://go.dev/help",
			Title: "Help - The Go Programming Language",
		},
		{
			ID:    1,
			URL:   "https://golang.org/help",
			Title: "Help - The Go Programming Language",
		},
		{
			ID:    2,
			URL:   "https://go.dev/play/",
			Title: "Go Playground - The Go Programming Language",
		},
	}
	req = httptest.NewRequest(http.MethodGet, "/docs", nil)
	req.Header.Add("content-type", "plain/text")
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusOK)
	}
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	got := string(body)
	for _, doc := range documents.docs {
		ul += "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	want := "<html><body><div>" + ul + "</div></body></html>"
	if got != want {
		t.Errorf("invalid body: get %v, want %v", got, want)
	}
}

func Test_crawlerDocs_indexHandler(t *testing.T) {
	documents = docs{docs: []crawler.Document{}}
	id := 2
	req := httptest.NewRequest(http.MethodGet, "/index/"+strconv.Itoa(id), nil)
	req.Header.Add("content-type", "plain/text")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusNoContent {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusNoContent)
	}

	documents.docs = []crawler.Document{
		{
			ID:    0,
			URL:   "https://go.dev/help",
			Title: "Help - The Go Programming Language",
		},
		{
			ID:    1,
			URL:   "https://golang.org/help",
			Title: "Help - The Go Programming Language",
		},
		{
			ID:    2,
			URL:   "https://go.dev/play/",
			Title: "Go Playground - The Go Programming Language",
		},
	}
	req = httptest.NewRequest(http.MethodGet, "/index/"+strconv.Itoa(id), nil)
	req.Header.Add("content-type", "plain/text")
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusOK)
	}
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	got := string(body)
	doc := documents.docs[id]
	ul := "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"" + doc.URL + "\">" + doc.Title + "</a></p>"
	want := "<html><body><div>" + ul + "</div></body></html>"
	if got != want {
		t.Errorf("invalid body: get %v, want %v", got, want)
	}
}

func Test_crawlerDocs_searchHandler(t *testing.T) {
	documents = docs{docs: []crawler.Document{}}
	req := httptest.NewRequest(http.MethodGet, "/search/help", nil)
	req.Header.Add("content-type", "plain/text")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusNoContent {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusNoContent)
	}

	var ul string
	documents.docs = []crawler.Document{
		{
			ID:    0,
			URL:   "https://go.dev/help",
			Title: "Help - The Go Programming Language",
		},
		{
			ID:    1,
			URL:   "https://golang.org/help",
			Title: "Help - The Go Programming Language",
		},
		{
			ID:    2,
			URL:   "https://go.dev/play/",
			Title: "Go Playground - The Go Programming Language",
		},
	}
	req = httptest.NewRequest(http.MethodGet, "/search/help", nil)
	req.Header.Add("content-type", "plain/text")
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusOK)
	}
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	got := string(body)
	for _, doc := range documents.docs {
		if doc.ID == 2 {
			continue
		}
		ul += "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	want := "<html><body><div>" + ul + "</div></body></html>"
	if got != want {
		t.Errorf("invalid body: get %v, want %v", got, want)
	}
}
