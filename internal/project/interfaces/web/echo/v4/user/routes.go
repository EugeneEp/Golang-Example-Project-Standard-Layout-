package user

import (
	"github.com/labstack/echo/v4"
	"projectname/internal/project/interfaces/web/echo/middleware/access"
)

func Bind(g *echo.Group) {
	route := g.Group("/user")
	route.Use(access.SetInvoker())
	route.GET("/:id", Get)
	route.POST("", Create)
}
