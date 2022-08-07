package user

import (
	"go.uber.org/zap"
	"projectname/internal/project/domain/user"
	domain "projectname/internal/project/domain/user"
)

const queryAddValue = `INSERT INTO "user".users (id, display_name) VALUES ($1, $2)`

func (c *context) Create(rq user.Create) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if err := c.driver.Query(
		queryAddValue,
		rq.ID,
		rq.DisplayName,
	); err != nil {
		c.reg.Log.Error(domain.ErrNotCreated.Error(), zap.Error(err))
		return domain.ErrNotCreated
	}

	return nil
}
