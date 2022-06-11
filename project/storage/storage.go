package storage

import "go2022/project/storage/models"

type Storage interface {
	NewLink(link string) (string, error)
	GetLink(short string) (models.Link, error)
}
