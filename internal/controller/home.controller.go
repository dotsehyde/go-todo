package controller

import (
	"go-todo/internal/middleware"
	"go-todo/internal/utils"
	"go-todo/internal/views"
	profile_view "go-todo/internal/views/profile"

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
		return utils.Render(c, views.Home("todo"))
	} else {
		return c.Redirect(301, "/auth/")
	}
}

func (s *Controller) Todo(c echo.Context) error {
	return c.String(200, `<h1>Todo</h1>`)
}

func (s *Controller) Profile(c echo.Context) error {
	name, _ := middleware.SessionManager.Get(c.Request().Context(), "name").(string)
	email, _ := middleware.SessionManager.Get(c.Request().Context(), "email").(string)
	data := profile_view.ProfileViewData{
		Name:  name,
		Email: email,
	}
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return utils.Render(c, profile_view.ProfileView(data))

}
