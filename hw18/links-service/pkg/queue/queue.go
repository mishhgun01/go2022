package queue

type Storage interface {
	NewLink(url string) string
	Link(link string) string
}
