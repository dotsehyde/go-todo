package utils

import (
	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var Validator = validator.New()

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
