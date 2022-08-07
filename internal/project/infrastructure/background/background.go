package background

import (
	"context"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
	"github.com/zhashkevych/scheduler"
	coreUser "projectname/internal/project/core/user"
	"projectname/internal/project/domain/configuration"
	domainUser "projectname/internal/project/domain/user"
	"time"
)

const ServiceName = `BackgroundServiceName`

func New(ctn di.Container, cfg *viper.Viper) (*scheduler.Scheduler, error) {
	deleteUsersTime := time.Duration(cfg.GetInt(configuration.SyncDeleteUsersTime))

	background := scheduler.NewScheduler()
	background.Add(context.Background(), func(ctx context.Context) {
		go coreUser.DeleteOverdue(ctn, domainUser.DeleteOverdue{TimeRange: cfg.GetInt64(configuration.UsersOverdueTimeInSeconds)})
	}, time.Second*deleteUsersTime)

	return background, nil
}
