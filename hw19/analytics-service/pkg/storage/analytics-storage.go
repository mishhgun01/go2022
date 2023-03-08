package storage

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

type Analytics struct {
	r         *kafka.Reader
	count     int
	middleLen int
	lengths   []int
	mu        *sync.Mutex
}

type ConsumerOptions struct {
	Brokers []string
	GroupID string
	Topic   string
}

func New(op ConsumerOptions) (*Analytics, error) {
	kfk, err := newConsumer(op.Brokers, op.GroupID, op.Topic)
	if err != nil {
		return nil, err
	}
	return &Analytics{
		mu: &sync.Mutex{},
		r:  kfk,
	}, nil
}

func (a *Analytics) Update() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	msg, err := a.r.FetchMessage(context.Background())

	if err != nil {
		log.Println(err)
	}

	link := string(msg.Value)
	a.count += 1
	a.lengths = append(a.lengths, len(link))
	a.middleLen = sum(a.lengths) / a.count
	return a.count
}

func sum(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	var result int
	for _, el := range slice {
		result += el
	}

	return result
}

func newConsumer(brokers []string, groupID string, topic string) (*kafka.Reader, error) {
	if len(brokers) == 0 || brokers[0] == "" || topic == "" {
		return nil, errors.New("не указаны параметры подключения к Kafka")
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e1,
		MaxBytes: 10e6,
	})

	return r, nil
}
