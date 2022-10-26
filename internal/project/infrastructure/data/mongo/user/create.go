package user

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"projectname/internal/project/domain/configuration"
	"projectname/internal/project/domain/user"
	"projectname/internal/project/infrastructure/config"
	common "projectname/internal/project/infrastructure/data/common/user"
)

func (c *context) Create(rq user.Create) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	var cfg *viper.Viper

	if err := c.ctn.Fill(config.ServiceName, &cfg); err != nil {
		return err
	}

	UsersCollection := c.driver.Client.Database(cfg.GetString(configuration.MongoDbName)).Collection(common.CollectionName)

	rq.SetCreatedAt()
	b, err := bson.Marshal(&rq)
	if err != nil {
		return err
	}

	_, err = UsersCollection.InsertOne(c.driver.Ctx, b)
	if err != nil {
		return err
	}

	return nil
}
