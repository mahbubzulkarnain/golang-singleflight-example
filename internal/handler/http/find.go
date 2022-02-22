package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) Find(ctx echo.Context) error {
	users, err := h.UserService.Find(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, users)
}
