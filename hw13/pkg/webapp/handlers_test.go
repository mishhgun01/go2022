package webapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go2022/hw13/pkg/crawler"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var documents = Docs{[]crawler.Document{
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
}}
var r *mux.Router

func TestMain(m *testing.M) {
	r = mux.NewRouter()
	api := New(r, documents)
	api.r.HandleFunc("/api/v1/docs", api.docs).Methods(http.MethodGet)
	api.r.HandleFunc("/api/v1/index/{query}", api.index).Methods(http.MethodGet)
	api.r.HandleFunc("/api/v1/newDoc", api.newDocument).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/v1/deleteDoc/{id}", api.deleteDocument).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/v1/updateDoc/{id}", api.updateDocument).Methods(http.MethodPost, http.MethodOptions)
	os.Exit(m.Run())
}

// Test for /docs handler.
func TestAPI_docs(t *testing.T) {

	var ul string

	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs", nil)
	req.Header.Add("content-type", "plain/text")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusOK)
	}
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	got := string(body)
	for _, doc := range documents.Docs {
		ul += "<p><a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	want := "<html><body><div>" + ul + "</div></body></html>"
	if got != want {
		t.Errorf("invalid body: get %v, want %v", got, want)
	}
}

// Test for /index/{query} handler.
func TestAPI_index(t *testing.T) {

	var ul string
	req := httptest.NewRequest(http.MethodGet, "/api/v1/index/help", nil)
	req.Header.Add("content-type", "plain/text")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusOK)
	}
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	got := string(body)
	for _, doc := range documents.Docs {
		if doc.ID == 2 {
			continue
		}
		ul += "<p><a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	want := "<html><body><div>" + ul + "</div></body></html>"
	if got != want {
		t.Errorf("invalid body: get %v, want %v", got, want)
	}
}

// Test for /newDoc handler.
func TestAPI_newDocument(t *testing.T) {
	var d = crawler.Document{
		ID:    3,
		URL:   "https://yandex.ru",
		Title: "Yandex",
	}
	documents.Docs = append(documents.Docs, d)
	data, err := json.Marshal(d)
	if err != nil {
		log.Println(err.Error())
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/newDoc", bytes.NewBuffer(data))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusOK)
	}
	req = httptest.NewRequest(http.MethodGet, "/api/v1/docs", nil)
	req.Header.Add("content-type", "plain/text")
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	got := string(body)
	var ul string
	for _, doc := range documents.Docs {
		ul += "<p><a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	want := "<html><body><div>" + ul + "</div></body></html>"
	if got != want {
		t.Errorf("invalid body: get %v, want %v", got, want)
	}
}

// Test for /deleteDoc handler.
func TestAPI_deleteDoc(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/deleteDoc/2", nil)
	req.Header.Add("content-type", "plain/text")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusOK)
	}
	req = httptest.NewRequest(http.MethodGet, "/api/v1/docs", nil)
	req.Header.Add("content-type", "plain/text")
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	got := string(body)
	var ul string
	for i := range documents.Docs {
		if i == 2 {
			continue
		}
		ul += "<p><a href=\"/index/" + fmt.Sprint(documents.Docs[i].ID) + "\">" + documents.Docs[i].Title + "</a></p>"
	}
	want := "<html><body><div>" + ul + "</div></body></html>"
	if got != want {
		t.Errorf("invalid body: get %v, want %v", got, want)
	}
}

// Test for /updateDoc handler.
func TestAPI_updateDoc(t *testing.T) {
	var d = crawler.Document{
		ID:    2,
		URL:   "https://yandex.ru",
		Title: "Yandex",
	}

	for _, v := range documents.Docs {
		if d.ID == v.ID {
			v = d
		}
	}
	data, err := json.Marshal(d)
	if err != nil {
		log.Println(err.Error())
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/updateDoc/2", bytes.NewBuffer(data))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusOK)
	}
	req = httptest.NewRequest(http.MethodGet, "/api/v1/docs", nil)
	req.Header.Add("content-type", "plain/text")
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	got := string(body)
	var ul string
	for i := range documents.Docs {
		ul += "<p><a href=\"/index/" + fmt.Sprint(documents.Docs[i].ID) + "\">" + documents.Docs[i].Title + "</a></p>"
	}
	want := "<html><body><div>" + ul + "</div></body></html>"
	if got != want {
		t.Errorf("invalid body: get %v, want %v", got, want)
	}
}
