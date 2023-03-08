package pgsql

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"go2022/hw16/pkg/models"
)

// Структура БД Postgres.
type Storage struct {
	pool *pgxpool.Pool
}

// Films возвращает фильм с заданным id.
func (s *Storage) Films(studioID int) ([]models.Film, error) {
	var data []models.Film
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id,
		title,
		year_of_release,
		box_office,
		studio_id,
		rating_id
	FROM films
	WHERE 
		(studio_id = $1 OR $1 = 0)
	ORDER BY id`,
		studioID,
	)
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		var item models.Film
		err = rows.Scan(
			&item.ID,
			&item.Title,
			&item.Year,
			&item.BoxOffice,
			&item.StudioID,
			&item.RateID,
		)
		if err != nil {
			return data, err
		}
		data = append(data, item)
	}
	return data, nil
}

// DeleteFilm удаляет фильм.
func (s *Storage) DeleteFilm(req models.Request) error {
	_, err := s.pool.Exec(context.Background(), `
	DELETE FROM films WHERE id=$1`, req.ID)
	return err
}

// UpdateFilm обновляет информацию о фильме.
func (s *Storage) UpdateFilm(item models.Film) error {
	_, err := s.pool.Exec(context.Background(), `
	UPDATE films
		SET title = $2,
			year_of_release = $3,
			box_office = $4,
			studio_id=$5,
			rating_id=$6
	WHERE id=$1`,
		item.ID,
		item.Title,
		item.Year,
		item.BoxOffice,
		item.StudioID,
		item.RateID,
	)
	if err != nil {
		return err
	}
	return nil
}

// NewFilm добавляет фильм в бд.
func (s *Storage) NewFilm(item models.Film) (int, error) {
	var id int

	err := s.pool.QueryRow(context.Background(), `
	INSERT INTO films ( 
		title, 
		year_of_release,
		box_office,
		studio_id,
		rating_id)
	VALUES ($1,$2,$3,$4,$5) RETURNING id`,
		item.Title,
		item.Year,
		item.BoxOffice,
		item.StudioID,
		item.RateID,
	).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}
