package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	DB *mongo.Client
}

func InitialDB() *Connection {
	return &Connection{}
}

func (conn *Connection) ToConnectMongoDB() *mongo.Client {
	conf := GetConfig()

	serverApiOptions := options.ServerAPI(options.ServerAPIVersion1)
	dsn := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Host,
	)

	clientOptions := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverApiOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	conn.DB = client
	return client
}

func CreateCollection(coll string, conn *mongo.Client) *mongo.Collection {
	conf := GetConfig()

	return conn.Database(conf.Database.Name).Collection(coll)
}