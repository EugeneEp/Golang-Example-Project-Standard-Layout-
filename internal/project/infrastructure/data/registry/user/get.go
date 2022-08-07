package user

import (
	"golang.org/x/sys/windows/registry"
	"projectname/internal/project/domain/user"
	domain "projectname/internal/project/domain/user"
	common "projectname/internal/project/infrastructure/data/common/user"
)

func (c *context) Get(rq user.Get) (*user.GetResult, error) {
	var u = &user.GetResult{}
	var path = common.Path + `\` + rq.ID
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.ALL_ACCESS)
	if err == registry.ErrNotExist {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	u.DisplayName, _, err = key.GetStringValue(common.FieldDisplayName)
	if err != nil {
		u.DisplayName = common.DefaultDisplayName
	}

	createdAt, _, err := key.GetIntegerValue(common.FieldCreatedAt)
	if err != nil {
		u.CreatedAt = common.DefaultCreatedAt
	} else {
		u.CreatedAt = int64(createdAt)
	}

	u.ID = rq.ID

	defer func() {
		if err := key.Close(); err != nil {
			c.logger.Warn(err.Error())
		}
	}()

	return u, nil
}
