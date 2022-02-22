package http

import (
	"github.com/labstack/echo/v4"

	userService "github.com/mahbubzulkarnain/golang-singleflight-example/pkg/v1/user/service"
)

type Handler struct {
	UserService userService.Service
}

func (h Handler) Route(g *echo.Group) {
	api := g.Group("/users")
	{
		api.GET("/", h.Find)
	}
}
