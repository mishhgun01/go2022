package models

// Структура фильма.
type Film struct {
	ID        int
	Year      int
	Title     string
	BoxOffice int
	StudioID  int
	RateID    int
}
