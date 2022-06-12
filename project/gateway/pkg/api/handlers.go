package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"go2022/project/storage/models"
	"net/http"
)

func (api *API) url(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["short"]
	url := api.cache.Url(short)
	if url != "" {
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}
	url, err := api.db.GetLink(short)
	if err != nil {
		http.Error(w, "No content", http.StatusNoContent)
		return
	}
	if url == "" {
		http.Error(w, "No content", http.StatusNoContent)
		return
	}

}

func (api *API) newUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]
	link := models.MakeLink(url)
	links := make([]models.Link, 1)
	links = append(links, link)
	err := api.cache.Update(links)
	if err != nil {
		http.Error(w, "No content", http.StatusNoContent)
		return
	}
	short, err := api.db.NewLink(link)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	message := kafka.Message{Key: []byte(short), Value: []byte(url)}
	err = api.w.W.WriteMessages(context.Background(), message)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(short))
}
