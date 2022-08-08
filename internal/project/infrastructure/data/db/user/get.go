package user

import (
	"go.uber.org/zap"
	"projectname/internal/project/domain/user"
	domain "projectname/internal/project/domain/user"
	"time"
)

const queryGetValue = `SELECT t1.id, t1.display_name, t1.created_at FROM "user".users t1 WHERE t1.id=$1`

type getValueResult struct {
	ID          string     `json:"id"`
	DisplayName string     `json:"display_name"`
	CreatedAt   *time.Time `json:"created_at"`
}

func (c *context) Get(rq user.Get) (*user.GetResult, error) {
	var res getValueResult

	if err := c.driver.Get(&res,
		queryGetValue,
		rq.ID,
	); err != nil {
		c.reg.Log.Error(domain.ErrNotFound.Error(), zap.Error(err))
		return nil, domain.ErrNotFound
	}

	return &user.GetResult{
		Entity: user.Entity{
			ID:          res.ID,
			DisplayName: res.DisplayName,
			CreatedAt:   res.CreatedAt.Unix(),
		},
	}, nil
}
