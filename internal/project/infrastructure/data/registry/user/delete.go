package user

import (
	"golang.org/x/sys/windows/registry"
	"projectname/internal/project/domain/user"
	common "projectname/internal/project/infrastructure/data/common/user"
	"time"
)

func (c *context) DeleteOverdue(rq user.DeleteOverdue) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var path = common.Path
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.ALL_ACCESS)

	if err == registry.ErrNotExist {
		return user.ErrNotFound
	}

	keys, err := key.ReadSubKeyNames(-1)
	if err != nil {
		return err
	}

	t := time.Now().Unix() - rq.TimeRange

	for _, v := range keys {
		var u *user.GetResult
		u, err = c.Get(user.Get{
			ID: v,
		})

		if err != nil {
			return err
		}

		if u.Entity.CreatedAt < t {
			if err = registry.DeleteKey(registry.LOCAL_MACHINE, path+`\`+u.ID); err != nil {
				return err
			}
		}
	}

	defer func() {
		if err := key.Close(); err != nil {
			c.logger.Warn(err.Error())
		}
	}()

	return nil
}
