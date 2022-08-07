package user

import (
	"github.com/sarulabs/di/v2"
	"go.uber.org/zap"
	domain "projectname/internal/project/domain/user"
	"projectname/internal/project/infrastructure/data"
	"projectname/internal/project/infrastructure/logger"
)

func Get(ctn di.Container, rq domain.Get) (*domain.GetResult, error) {
	var (
		ctx data.Context
	)

	if err := ctn.Fill(data.ServiceName, &ctx); err != nil {
		return nil, err
	}

	u, err := ctx.User().Get(rq)

	if err != nil {
		return nil, err
	}

	u.ID = rq.ID

	return u, nil
}

func Create(ctn di.Container, rq domain.Create) (*domain.CreateResult, error) {
	var (
		ctx data.Context
	)

	if err := ctn.Fill(data.ServiceName, &ctx); err != nil {
		return nil, err
	}

	rq.GenerateID()

	if err := ctx.User().Create(rq); err != nil {
		return nil, err
	}

	return &domain.CreateResult{Entity: domain.Entity{
		ID:          rq.ID,
		DisplayName: rq.DisplayName,
	}}, nil
}

func DeleteOverdue(ctn di.Container, rq domain.DeleteOverdue) {
	var (
		ctx data.Context
		log *zap.Logger
	)

	if err := ctn.Fill(logger.BaseServiceName, &log); err != nil {
		return
	}

	if err := ctn.Fill(data.ServiceName, &ctx); err != nil {
		log.Error(err.Error())
	}

	if err := ctx.User().DeleteOverdue(rq); err != nil {
		log.Error(err.Error())
	}
}
