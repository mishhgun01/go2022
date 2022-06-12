package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go2022/project/storage/models"
)

type DB struct {
	c          *mongo.Client
	name       string
	collection string
}

func New(URI string, name string, collectionName string) (*DB, error) {
	opts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.Background(), opts)
	defer client.Disconnect(context.Background())
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	DB := DB{
		c:          client,
		name:       name,
		collection: collectionName,
	}
	return &DB, nil
}

func (db *DB) NewLink(link models.Link) (string, error) {
	collection := db.c.Database(db.name).Collection(db.collection)

	_, err := collection.InsertOne(context.Background(), link)
	if err != nil {
		return "", err
	}
	return link.Short, nil
}

func (db *DB) GetLink(link models.Link) (models.Link, error) {
	collection := db.c.Database(db.name).Collection(db.collection)
	filter := bson.D{}
}
