package v4

import (
	"github.com/labstack/echo/v4"
	"projectname/internal/project/interfaces/web/echo/v4/user"
)

func Bind(g *echo.Group) {
	route := g.Group("/v4")
	user.Bind(route)
}
