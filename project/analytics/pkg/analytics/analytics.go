package analytics

import (
	"context"
	"go2022/project/queue/pkg"
	"log"
	"sync"
)

type Analytics struct {
	Count     int
	MiddleLen int
	Lengths   []int
	mu        *sync.Mutex
	r         *pkg.Consumer
}

func New(r *pkg.Consumer) *Analytics {
	return &Analytics{
		Count:     0,
		MiddleLen: 0,
		Lengths:   nil,
		mu:        &sync.Mutex{},
		r:         r,
	}
}

func (a *Analytics) Update() int {
	a.mu.Lock()
	msg, err := a.r.R.FetchMessage(context.Background())
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
