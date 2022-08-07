package user

import (
	"github.com/sarulabs/di/v2"
	"go.uber.org/zap"
	"projectname/internal/project/infrastructure/data/common/user"
	"projectname/internal/project/infrastructure/data/registry"
	"projectname/internal/project/infrastructure/logger"
	"sync"
)

type (
	context struct {
		reg    *registry.Context
		mutex  sync.RWMutex
		logger *zap.Logger
		ctn    di.Container
	}
)

func Context(ctn di.Container) (user.Interface, error) {
	reg, err := registry.Ctx(ctn)

	if err != nil {
		return nil, err
	}

	var log *zap.Logger

	if err = ctn.Fill(logger.BaseServiceName, &log); err != nil {
		return nil, err
	}

	return &context{
		ctn:    ctn,
		reg:    reg,
		mutex:  sync.RWMutex{},
		logger: log,
	}, nil
}
