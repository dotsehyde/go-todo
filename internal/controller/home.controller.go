package controller

import (
	"go-todo/internal/middleware"
	"go-todo/internal/utils"
	"go-todo/internal/views"

	"github.com/labstack/echo/v4"
)

func (s *Controller) Home(c echo.Context) error {
	// sess, _ := session.Get("session", c)
	// fmt.Println(sess.Values["logged"])
	// if l, ok := c.Get("logged").(bool); ok && l {
	// return utils.Render(c, views.Home())
	// }
	// return c.Redirect(301, "/auth/")
	if middleware.SessionManager.Exists(c.Request().Context(), "logged") {
		return utils.Render(c, views.Home())
	} else {
		return c.Redirect(301, "/auth/")
	}
}
