package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"sync"
	"time"
)

var once sync.Once

var (
	Client   *mongo.Client
	Database *mongo.Database
	Context  context.Context
	Cancel   context.CancelFunc
)

func Connect(url string) (*mongo.Client, *mongo.Database) {
	once.Do(func() {
		Context, Cancel = context.WithTimeout(context.Background(), 10*time.Second)
		client, err := mongo.Connect(options.Client().ApplyURI(url))

		if err != nil {
			panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
		}

		err = client.Ping(Context, readpref.Primary())
		if err != nil {
			panic(fmt.Sprintf("Mongo DB ping issue %s", err))
		}

		database := client.Database("mongo-golang-test")

		Client = client
		Database = database
	})
	return Client, Database
}

func CloseConnection() {
	closeConnection(Client, Context, Cancel)
}

func closeConnection(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
	defer func() {
		cancel()
		if err := client.Disconnect(context); err != nil {
			panic(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}()
}
