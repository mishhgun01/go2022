package links

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

const (
	// Набор символов для сокращения ссылки
	shortChars = "abcdefghijklmnopqrstuvwxyz123456789"
	// Длинна короткой ссылки
	urlLen = 6
)

var (
	// Максимально возможное число ссылок для заданного набора символов и длины короткой ссылки
	maxUrls = int(math.Pow(float64(len([]byte(shortChars))), urlLen))
)

var (
	// Зарезервированный адрес
	link = "https://test.ru"
)

type Link struct {
	Link  string
	Short string
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
	short := ""
	for {
		short = randomString(urlLen)
		res := ""
		for _, shortLink := range db.links {
			if shortLink.Short == short {
				res = short
				break
			}
		}
		if res == "" {
			break
		}
	}
	link := Link{
		Link:  url,
		Short: fmt.Sprintf("%s/%s", link, short),
	}
	db.links = append(db.links, link)
	db.mu.Unlock()
	return link.Short
}

func (db *DB) Link(link string) string {
	var s string
	db.mu.Lock()
	defer db.mu.Unlock()
	for _, v := range db.links {
		if v.Short == link {
			s = v.Short
		}
	}
	return s
}

func randomString(n int) string {
	letters := []rune(shortChars)

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
