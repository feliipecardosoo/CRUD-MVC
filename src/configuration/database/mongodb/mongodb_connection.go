package mongodb

import (
	"context"
	"crud/src/configuration/logger"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL      = "MONGODB_URL"
	MONGODB_DATABASE = "MONGODB_NOME"
)

func NewMongoDBConnection(
	ctx context.Context,
) (
	*mongo.Database, error,
) {
	mongodb_url := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_url))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("MongoDB conectado")
	return client.Database(mongodb_database), nil
}
