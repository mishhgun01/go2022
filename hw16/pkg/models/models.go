package models

// Интерфейс взаимодействия с БД.
type Controller interface {
	Films(id int) ([]Film, error)
	DeleteFilm(item Film) error
	UpdateFilm(item Film) error
	NewFilm(item Film) (int, error)
}

// Структура фильма.
type Film struct {
	ID        int
	Year      int
	Title     string
	BoxOffice int
	StudioID  int
	RateID    int
}
