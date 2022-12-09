package database

import (
	"context"
	"destroyer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context, uri, databaseName string) (Repository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	targetCollection := client.Database(databaseName).Collection("targets")
	return &dataStore{
		targetCollection: targetCollection,
	}, nil
}

type dataStore struct {
	targetCollection *mongo.Collection
}

func (d *dataStore) ListTarget(ctx context.Context) ([]destroyer.Target, error) {
	cursor, err := d.targetCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var result []destroyer.Target
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (d *dataStore) GetTarget(ctx context.Context, id string) (*destroyer.Target, error) {
	var target destroyer.Target
	if err := d.targetCollection.FindOne(ctx, bson.M{"id": id}).Decode(&target); err != nil {
		return nil, err
	}
	return &target, nil
}
