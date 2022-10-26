package data

import (
	"github.com/sarulabs/di/v2"
	"projectname/internal/project/infrastructure/data/mongo/user"
)

func ctxMongo(ctn di.Container) (Context, error) {
	u, err := user.Context(ctn)

	if err != nil {
		return nil, err
	}

	return &context{
		user: u,
	}, nil
}
