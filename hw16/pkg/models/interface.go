package pkg

// Интерфейс взаимодействия с БД.
type Storage interface {
	Films(id []int) ([]Film, error)
	DeleteFilm(item Film) error
	UpdateFilm(item Film) error
	NewFilm(item Film) (int, error)
}
