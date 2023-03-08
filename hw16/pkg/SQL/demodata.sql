-- Демонстрационные данные - фильмы Интерстеллар и Начало
-- Режиссёр.
INSERT INTO directors(id,name, surname, year_of_birth) VALUES (0,'Christopher', 'Nolan', 1970);

-- Актёры.
INSERT INTO actors(id, name, surname, year_of_birth) VALUES (0,'Leonardo', 'DiCaprio', 1974),
                                                            (1,'Joseph','Gordon-Levitt',1981),
                                                            (2,'Elliot', 'Page', 1987),
                                                            (3,'Marion', 'Cotillard', 1975),
                                                            (4,'Tom', 'Hardy', 1977),
                                                            (5,'Cillian', 'Murphy', 1976),
                                                            (6, 'Matthew', 'McConaughey', 1969),
                                                            (7, 'Anne', 'Hathaway', 1982),
                                                            (8, 'Michael', 'Caine', 1933);
ALTER SEQUENCE actors_id_seq RESTART WITH 100;

-- Студии.
INSERT INTO studios(id, title) VALUES (0,'Legendary Pictures'),
                                      (1,'Warner Bros. Pictures');
ALTER SEQUENCE studios_id_seq RESTART WITH 100;
-- Возрастные рейтинги.
INSERT INTO ratings(id, title) VALUES (0,'PG-13');
ALTER SEQUENCE ratings_id_seq RESTART WITH 100;

-- Фильмы.
INSERT INTO films(id, title, year_of_release, box_office, studio_id, rating_id) VALUES
                                                                                    (0, 'Inception', 2010, 160000000, 0, 0),
                                                                                    (1, 'Interstellar', 2014, 165000000,1,0);
-- Начало и актёры.
INSERT INTO films_and_actors(film_id,actor_id) VALUES
                                                   (0,0),
                                                   (0,1),
                                                   (0,2),
                                                   (0,3),
                                                   (0,4),
                                                   (0,5);
-- Начало и режиссёр.
INSERT INTO films_and_directors(film_id,director_id) VALUES (0,0);


--Интерстеллар и актёры.
INSERT INTO films_and_actors(film_id, actor_id) VALUES
                                                    (1,6),
                                                    (1,7),
                                                    (1,8);


-- Интерстеллар и режиссёр.
INSERT INTO films_and_directors(film_id,director_id) VALUES (1,0);
