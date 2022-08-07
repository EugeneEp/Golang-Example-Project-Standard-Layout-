package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
	"projectname/internal/project/domain/basic"
)

type Context struct {
	Container di.Container
	echo.Context
}

const (
	KeyUserID = `user_id`
)

// BuildRequest формирует basic.Request на основании данных, полученных из запроса:
// RequestId - идентификатор http запроса для отслеживания в логах
// KeyUserID - идентификатор пользователя
func (c Context) BuildRequest() (req basic.Request) {
	req.RequestId = c.Response().Header().Get(echo.HeaderXRequestID)

	if userId, exists := c.Get(KeyUserID).(string); exists {
		req.CurrentUserId = userId
	}

	return
}

func WebContext(container di.Container) echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			return handlerFunc(&Context{Container: container, Context: context})
		}
	}
}
