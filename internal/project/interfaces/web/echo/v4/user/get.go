package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"projectname/internal/project/core/user"
	domain "projectname/internal/project/domain/user"
	"projectname/internal/project/interfaces/web/echo/middleware"
)

func Get(context echo.Context) error {
	var (
		ctx = context.(*middleware.Context)
		req = domain.Get{
			Request: ctx.BuildRequest(),
		}
	)

	req.ID = context.Param(`id`)

	res, err := user.Get(ctx.Container, req)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, res)
}
