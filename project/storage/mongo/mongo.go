package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go2022/project/storage/models"
	"math/rand"
	"time"
)

type DB struct {
	c    *mongo.Client
	conn bool
}

func New(URI string) (*DB, error) {
	opts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	DB := DB{
		c:    client,
		conn: true,
	}
	return &DB, nil
}

func (db *DB) NewLink(link string) (string, error) {

}

func (db *DB) GetLink(short string) (models.Link, error) {

}

func longToShort() string {
	rand.Seed(time.Now().UnixNano())
	short := make([]byte, 5)
	for i := range short {
		short[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(short)
}
