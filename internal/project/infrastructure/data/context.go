package data

import (
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
	"projectname/internal/project/domain/configuration"
	"projectname/internal/project/infrastructure/config"
	"projectname/internal/project/infrastructure/data/common/user"
)

type (
	Context interface {
		User() user.Interface
	}

	context struct {
		user user.Interface
	}
)

func (c context) User() user.Interface { return c.user }

const ServiceName = `DataStorageService`

// New Инициализирует сервис хранения данных, в зависимости от настроек конфигурации
func New(ctn di.Container) (Context, error) {
	var cfg *viper.Viper

	if err := ctn.Fill(config.ServiceName, &cfg); err != nil {
		return nil, err
	}

	if cfg.GetString(configuration.DataStoreType) == configuration.DataStoreTypeReg {
		return ctxREG(ctn)
	} else {
		return ctxDB(ctn)
	}
}
