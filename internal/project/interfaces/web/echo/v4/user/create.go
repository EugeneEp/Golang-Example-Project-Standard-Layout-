package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"projectname/internal/project/core/user"
	domain "projectname/internal/project/domain/user"
	"projectname/internal/project/interfaces/web/echo/middleware"
	"projectname/internal/project/interfaces/web/echo/response"
)

func Create(context echo.Context) error {
	var (
		ctx = context.(*middleware.Context)
		req domain.Create
	)

	if err := context.Bind(&req); err != nil {
		return response.ErrBadRequest(err)
	}

	res, err := user.Create(ctx.Container, req)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, res)
}
