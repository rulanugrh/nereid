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
	conf *App
}

func InitialDB(conf *App) *Connection {
	return &Connection{conf: conf}
}

func (conn *Connection) ToConnectMongoDB() {
	serverApiOptions := options.ServerAPI(options.ServerAPIVersion1)
	dsn := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		conn.conf.Database.User,
		conn.conf.Database.Pass,
		conn.conf.Database.Host,
	)

	clientOptions := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverApiOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	conn.DB = client
}
