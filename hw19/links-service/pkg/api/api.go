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
	db *links.DB
	r  *mux.Router
	w  *kafka.Writer
}

type KfkOptions struct {
	Brokers []string
	Topic   string
}

func New(op KfkOptions, db *links.DB) (*API, error) {
	kfk, err := newProducer(op.Brokers, op.Topic)
	if err != nil {
		return nil, err
	}
	return &API{
		db: db,
		r:  mux.NewRouter(),
		w:  kfk,
	}, nil
}

func (api *API) Handle() {
	api.r.HandleFunc("/newLink/{query}", api.NewLink).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/link/{query}", api.Link).Methods(http.MethodGet, http.MethodOptions)
}

func (api *API) ListenAndServe(addr string) error {
	log.Print("Listen on tcp://" + addr)
	err := http.ListenAndServe(addr, api.r)
	return err
}

func newProducer(brokers []string, topic string) (*kafka.Writer, error) {
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
