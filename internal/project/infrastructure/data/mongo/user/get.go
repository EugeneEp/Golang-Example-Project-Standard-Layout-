package user

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"projectname/internal/project/domain/configuration"
	"projectname/internal/project/domain/user"
	"projectname/internal/project/infrastructure/config"
	common "projectname/internal/project/infrastructure/data/common/user"
)

func (c *context) Get(rq user.Get) (*user.GetResult, error) {
	var cfg *viper.Viper

	if err := c.ctn.Fill(config.ServiceName, &cfg); err != nil {
		return nil, err
	}

	UsersCollection := c.driver.Client.Database(cfg.GetString(configuration.MongoDbName)).Collection(common.CollectionName)
	collection := UsersCollection.FindOne(c.driver.Ctx, bson.M{"id": rq.ID})

	var u *user.Entity
	if err := collection.Decode(&u); err != nil {
		return nil, err
	}

	return &user.GetResult{
		Entity: *u,
	}, nil
}
