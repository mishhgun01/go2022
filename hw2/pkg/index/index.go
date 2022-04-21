package index

// Обратный индекс отсканированных документов.

import (
	"go2022/hw2/pkg/crawler"
)

// Interface определяет контракт службы индексирования документов.
type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
