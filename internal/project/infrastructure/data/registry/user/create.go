package user

import (
	"golang.org/x/sys/windows/registry"
	"projectname/internal/project/domain/user"
	domain "projectname/internal/project/domain/user"
	common "projectname/internal/project/infrastructure/data/common/user"
)

func (c *context) Create(rq user.Create) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var path = common.Path + "\\" + rq.ID
	key, _, err := registry.CreateKey(registry.LOCAL_MACHINE, path, registry.ALL_ACCESS)

	if err != nil {
		return domain.ErrNotCreated
	}

	if err := key.SetStringValue(common.FieldDisplayName, rq.DisplayName); err != nil {
		return err
	}

	rq.SetCreatedAt()

	if err := key.SetQWordValue(common.FieldCreatedAt, uint64(rq.CreatedAt)); err != nil {
		return err
	}

	defer func() {
		if err := key.Close(); err != nil {
			c.logger.Warn(err.Error())
		}
	}()

	return nil
}
