package controller

import (
	"go-todo/internal/utils"
	"go-todo/internal/views"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return utils.Render(c, views.Home())
}
