package storage

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

type Analytics struct {
	consumer  *kafka.Reader
	Count     int
	MiddleLen int
	Lengths   []int
	mu        *sync.Mutex
}

type ConsumerOptions struct {
	brokers []string
	groupID string
	topic   string
}

func New(op ConsumerOptions) *Analytics {
	kfk := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  op.brokers,
		Topic:    op.topic,
		GroupID:  op.groupID,
		MinBytes: 10e1,
		MaxBytes: 10e6,
	})
	return &Analytics{
		mu:       &sync.Mutex{},
		consumer: kfk,
	}
}

func (a *Analytics) Update() int {
	a.mu.Lock()
	msg, err := a.consumer.FetchMessage(context.Background())
	if err != nil {
		log.Println(err)
	}
	link := string(msg.Value)
	a.Count += 1
	a.Lengths = append(a.Lengths, len(link))
	a.MiddleLen = sum(a.Lengths) / a.Count
	a.mu.Unlock()
	return a.Count
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
