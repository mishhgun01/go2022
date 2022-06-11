package pkg

import (
	"errors"
	"github.com/segmentio/kafka-go"
	"go2022/project/storage"
)

type Options struct {
	brokers []string
	groupID string
	topic   string
	db      storage.Storage
}

type Producer struct {
	w    *kafka.Writer
	Opts Options
}

type Consumer struct {
	r    *kafka.Reader
	Opts Options
}

type Queue struct {
	W *Producer
	R *Consumer
}

func FillOptions(brokers []string, groupID string, topic string, db storage.Storage) Options {
	return Options{
		brokers: brokers,
		groupID: groupID,
		topic:   topic,
		db:      db,
	}
}

func NewProducer(o Options) (*Producer, error) {
	if o.topic == "" || len(o.brokers) == 0 {
		return nil, errors.New("No parameters for producer in queue")
	}
	return &Producer{
		Opts: o,
		w: &kafka.Writer{
			Topic:    o.topic,
			Addr:     kafka.TCP(o.brokers[0]),
			Balancer: &kafka.LeastBytes{},
		},
	}, nil
}

func NewConsumer(o Options) (*Consumer, error) {
	if o.topic == "" || len(o.brokers) == 0 || o.groupID == "" {
		return nil, errors.New("No parameters for consumer in queue")
	}
	return &Consumer{
		Opts: o,
		r: kafka.NewReader(
			kafka.ReaderConfig{
				Brokers:  o.brokers,
				Topic:    o.topic,
				GroupID:  o.groupID,
				MinBytes: 10e1,
				MaxBytes: 10e6,
			}),
	}, nil
}

func NewQueue(producer *Producer, consumer *Consumer) (*Queue, error) {
	if producer == nil || consumer == nil {
		return nil, errors.New("No Consumer or Producer found")
	}
	return &Queue{
		W: producer,
		R: consumer,
	}, nil
}
