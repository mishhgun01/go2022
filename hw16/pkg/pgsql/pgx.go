package pgsql

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go2022/hw16/pkg/models"
	"golang.org/x/net/context"
)

// Структура БД Postgres.
type Storage struct {
	pool *pgxpool.Pool
}

// Films возвращает фильм с заданным id.
func (s *Storage) Films(ids []int) ([]pkg.Film, error) {
	var data []pkg.Film
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id,
		title,
		year_of_release,
		box_office,
		studio_id,
		rating_id
	FROM films
	WHERE (studio_id=ANY($1) OR array_length($1) is NULL)
	ORDER BY id`,
		intToInt32Array(ids))
	if err != nil {
		return data, err
	}

	defer rows.Close()
	for rows.Next() {
		var item pkg.Film
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
func (s *Storage) DeleteFilm(item pkg.Film) error {
	_, err := s.pool.Exec(context.Background(), `
	DELETE FROM films WHERE id=$1`, item.ID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateFilm обновляет информацию о фильме.
func (s *Storage) UpdateFilm(item pkg.Film) error {
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
		item.RateID)
	if err != nil {
		return err
	}
	return nil
}

// NewFilm добавляет фильм в бд.
func (s *Storage) NewFilm(item pkg.Film) (int, error) {
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
		item.RateID).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}

// Функция перевода слайса int в слайс int32 для корректной работы SQL скрипта.
func intToInt32Array(in []int) []int32 {
	var out []int32
	for _, val := range in {
		out = append(out, int32(val))
	}
	return out
}
