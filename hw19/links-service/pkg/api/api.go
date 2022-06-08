package api

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"go2022/hw19/links-service/pkg/links"
	"log"
	"net/http"
)

type API struct {
	db       *links.DB
	r        *mux.Router
	producer *kafka.Writer
}

type Options struct {
	brokers []string
	topic   string
	db      *links.DB
}

func New(op Options) (*API, error) {
	kfk, err := NewKfk(op.brokers, op.topic)
	if err != nil {
		return nil, err
	}
	return &API{
		db:       op.db,
		r:        mux.NewRouter(),
		producer: kfk,
	}, nil
}

func (api *API) Handle() {
	api.r.HandleFunc("/newLink/{query}", api.NewLink).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/link/{query}", api.Link).Methods(http.MethodGet, http.MethodOptions)
}

func (api *API) ListenAndServe(addr string) {
	log.Print("Listen on tcp://" + addr)
	http.ListenAndServe(addr, api.r)
}

func NewKfk(brokers []string, topic string) (*kafka.Writer, error) {
	if len(brokers) == 0 || brokers[0] == "" || topic == "" {
		return nil, errors.New("не указаны параметры подключения к Kafka")
	}

	w := &kafka.Writer{
		Topic:    topic,
		Addr:     kafka.TCP(brokers[0]),
		Balancer: &kafka.LeastBytes{},
	}

	return w, nil
}
