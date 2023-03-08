package links

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	// Набор символов для сокращения ссылки
	allChars = "abcdefghijklmnopqrstuvwxyz123456789"
	// Длина короткой ссылки
	shortUrlLen = 6
)

var (
	// Зарезервированный адрес
	link = "https://test.ru"
)

type Link struct {
	link  string
	short string
}

type DB struct {
	mu    sync.Mutex
	links []Link
}

func New(links []Link) *DB {
	return &DB{
		mu:    sync.Mutex{},
		links: links,
	}
}

func (db *DB) NewLink(url string) string {
	db.mu.Lock()
	defer db.mu.Unlock()

	short := ""
	for {
		short = randomString(shortUrlLen)
		res := ""
		for _, shortLink := range db.links {
			if shortLink.short == short {
				res = short
				break
			}
		}
		if res == "" {
			break
		}
	}

	link := Link{
		link:  url,
		short: fmt.Sprintf("%s/%s", link, short),
	}

	db.links = append(db.links, link)
	return link.short
}

func (db *DB) Link(link string) string {
	var s string
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, v := range db.links {
		if v.short == link {
			s = v.short
		}
	}

	return s
}

func randomString(n int) string {
	letters := []rune(allChars)
	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
