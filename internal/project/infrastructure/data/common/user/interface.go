package user

import (
	"projectname/internal/project/domain/user"
)

type Interface interface {
	Get(req user.Get) (*user.GetResult, error)
	Create(req user.Create) error
	DeleteOverdue(req user.DeleteOverdue) error
}
