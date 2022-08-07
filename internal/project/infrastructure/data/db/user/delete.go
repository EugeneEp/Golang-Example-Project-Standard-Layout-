package user

import (
	"go.uber.org/zap"
	"projectname/internal/project/domain/user"
	"time"
)

const queryDeleteOverdue = `DELETE FROM "user".users WHERE created_at < $1`

func (c *context) DeleteOverdue(rq user.DeleteOverdue) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	r := time.Duration(rq.TimeRange)
	d := time.Now()
	d.Add(-r)

	if err := c.driver.Query(
		queryDeleteOverdue,
		d,
	); err != nil {
		c.reg.Log.Error(user.ErrNotDeleted.Error(), zap.Error(err))
		return user.ErrNotDeleted
	}

	return nil
}
