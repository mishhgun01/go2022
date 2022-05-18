-- Удалить таблицы, если они существуют.
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS directors;
DROP TABLE IF EXISTS ratings;
DROP TABLE IF EXISTS films;
DROP TABLE IF EXISTS studios;
DROP TABLE IF EXISTS films_and_actors;
DROP TABLE IF EXISTS films_and_directors;

-- Актёры.
CREATE TABLE actors(
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL DEFAULT '',
                       surname TEXT NOT NULL DEFAULT '',
                       year_of_birth INTEGER NOT NULL DEFAULT 0,
                       UNIQUE (name, surname, year_of_birth)
);

-- Режиссеры.
CREATE TABLE directors(
                          id SERIAL PRIMARY KEY ,
                          name TEXT NOT NULL DEFAULT '',
                          surname TEXT NOT NULL DEFAULT '',
                          year_of_birth INTEGER NOT NULL DEFAULT 0,
                          UNIQUE (name, surname, year_of_birth)
);

-- Возрастной рейтинг.
CREATE TABLE ratings(
                        id SERIAL PRIMARY KEY ,
                        title TEXT NOT NULL DEFAULT '' UNIQUE
);

-- Фильмы.
CREATE TABLE films(
                      id SERIAL PRIMARY KEY ,
                      title TEXT NOT NULL DEFAULT '',
                      year_of_release INTEGER NOT NULL DEFAULT 0,
                      box_office BIGINT DEFAULT 0,
                      studio_id BIGINT NOT NULL REFERENCES studios(id),
                      rating_id BIGINT NOT NULL REFERENCES ratings(id),
                      UNIQUE (title, year_of_release)
);

-- Студии.
CREATE TABLE studios(
                        id SERIAL PRIMARY KEY ,
                        title TEXT NOT NULL DEFAULT ''  UNIQUE
);

-- Связь фильмов и актёров.
CREATE TABLE films_and_actors(
                                 id SERIAL PRIMARY KEY ,
                                 film_id BIGINT NOT NULL REFERENCES films(id),
                                 actor_id BIGINT NOT NULL REFERENCES actors(id),
                                 UNIQUE (film_id, actor_id)
);

-- Связь режиссёров и актёров.
CREATE TABLE films_and_directors(
                                    id SERIAL PRIMARY KEY ,
                                    film_id BIGINT NOT NULL REFERENCES films(id),
                                    director_id BIGINT NOT NULL REFERENCES directors(id),
                                    UNIQUE (film_id, director_id)
);
