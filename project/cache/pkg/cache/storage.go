package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go2022/project/storage/models"
	"sync"
	"time"
)

type Storage struct {
	mu     *sync.Mutex
	client *redis.Client
	dur    time.Duration
}

func New(conn string, dur time.Duration) *Storage {
	client := redis.NewClient(
		&redis.Options{
			Addr:     conn,
			Password: "",
			DB:       0,
		})
	return &Storage{
		client: client,
		mu:     &sync.Mutex{},
		dur:    dur,
	}
}

func (s *Storage) Url(short string) string {
	s.mu.Lock()
	key := short
	url, err := s.client.Get(context.Background(), key).Result()
	if err != nil {
		return ""
	}
	s.mu.Unlock()
	return url
}

func (s *Storage) Update(urls []models.Link) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, item := range urls {
		key := item.Short
		err := s.client.Set(context.Background(), key, item.Long, s.dur).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) Flush() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.client.FlushAll(context.Background()).Err()
}

func (s *Storage) Close() error {
	return s.client.Close()
}
