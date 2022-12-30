package mongo

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DB struct {
	client *m.Client
}

func NewConnection(uri string) *DB {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, _ := m.Connect(
		ctx,
		options.Client().ApplyURI(uri),
	)

	return &DB{client: client}
}

func (mon *DB) ListDatabaseNames() ([]string, error) {
	databases, err := mon.client.ListDatabaseNames(
		context.TODO(),
		bson.D{},
	)

	return databases, err
}

func (mon *DB) ListCollections(db string) ([]string, error) {
	collections, err := mon.client.Database(db).
		ListCollectionNames(
			context.TODO(),
			bson.D{{"type", "collection"}},
		)

	return collections, err
}

func (mon *DB) ListViews(db string) ([]string, error) {
	views, err := mon.client.Database(db).
		ListCollectionNames(
			context.TODO(),
			bson.D{{"type", "view"}},
		)

	return views, err
}

func (mon *DB) ListUsers(db string) ([]string, error) {
	var users []string
	var userDecoder struct {
		Users []map[string]interface{}
	}

	_ = mon.client.Database(db).
		RunCommand(
			context.TODO(),
			bson.D{{"usersInfo", 1}},
		).Decode(&userDecoder)

	if len(userDecoder.Users) != 0 {
		users = append(users, userDecoder.Users[0]["user"].(string))
	}
	return users, nil
}

func (mon *DB) RunCommand(database string, command string) string {
	var result bson.M
	mon.client.Database(database).
		RunCommand(
			context.TODO(),
			bson.D{{command, 1}},
		).
		Decode(&result)

	output, _ := json.MarshalIndent(result, "", "    ")
	return string(output)
}
