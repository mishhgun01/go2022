package links

import (
	"strings"
	"sync"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLNMOPQRTTUVWXYZ"

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
	link := Link{
		Link:  url,
		Short: longToShort(url),
	}
	db.links = append(db.links, link)
	db.mu.Unlock()
	return link.Short
}

func (db *DB) Link(link string) string {
	var s string
	for _, v := range db.links {
		if v.Short == link {
			s = v.Short
		}
	}
	return s
}

// Формирование короткой ссылки путём формирования случайного слова из 5 символов.
func longToShort(long string) string {
	s := strings.Split(long, "/")
	short := s[0] + "//" + s[2] + "/"
	return short
}
