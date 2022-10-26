package user

import (
	"github.com/sarulabs/di/v2"
	"projectname/internal/project/infrastructure/data/common/user"
	"projectname/internal/project/infrastructure/data/mongo"
	driver "projectname/internal/project/infrastructure/nosql/mongo"
	"sync"
)

type (
	context struct {
		ctn    di.Container
		reg    *mongo.Context
		mutex  sync.RWMutex
		driver *driver.Driver
	}
)

func Context(ctn di.Container) (user.Interface, error) {
	reg, err := mongo.Ctx(ctn)

	if err != nil {
		return nil, err
	}

	var d *driver.Driver

	if err = ctn.Fill(driver.ServiceName, &d); err != nil {
		return nil, err
	}

	return &context{
		ctn:    ctn,
		reg:    reg,
		driver: d,
	}, nil
}
