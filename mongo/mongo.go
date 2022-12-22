package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Mongo struct {
	client *m.Client
}

func Interface(client *m.Client) *Mongo {
	return &Mongo{client}
}

func CreateMongoDBConnection(uri string) (*m.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := m.Connect(
		ctx,
		options.Client().ApplyURI(uri),
	)

	return client, err
}

func (mon *Mongo) ListDatabaseNames() ([]string, error) {
	databases, err := mon.client.ListDatabaseNames(
		context.TODO(),
		bson.D{},
	)

	return databases, err
}

func (mon *Mongo) ListCollections(db string) ([]string, error) {
	collections, err := mon.client.Database(db).
		ListCollectionNames(
			context.TODO(),
			bson.D{{"type", "collection"}},
		)

	return collections, err
}

func (mon *Mongo) ListViews(db string) ([]string, error) {
	views, err := mon.client.Database(db).
		ListCollectionNames(
			context.TODO(),
			bson.D{{"type", "view"}},
		)

	return views, err
}

func (mon *Mongo) ListUsers(db string) ([]string, error) {
	var users struct {
		Users interface{}
	}

	singleResult := mon.client.Database(db).
		RunCommand(
			context.TODO(),
			bson.D{{"usersInfo", 1}},
		).Decode(&users)

	fmt.Println(users.Users, singleResult)
	var yes interface{}
	json.Unmarshal(users.Users, &yes)
	return []string{""}, nil
}
