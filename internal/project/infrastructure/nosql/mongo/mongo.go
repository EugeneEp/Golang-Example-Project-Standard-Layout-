package mongo

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"projectname/internal/project/domain/configuration"
)

const ServiceName = "MongoDatabaseService"

type Driver struct {
	Client *mongo.Client
	Ctx    context.Context
}

func New(ctx context.Context, cfg *viper.Viper) (*Driver, error) {
	connString := cfg.GetString(configuration.MongoConnection)

	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &Driver{
		Client: client,
		Ctx:    ctx,
	}, nil
}
