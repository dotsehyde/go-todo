package controller

import (
	"context"
	"go-todo/internal/middleware"
	"go-todo/internal/utils"
	"go-todo/internal/views"
	profile_view "go-todo/internal/views/profile"
	todo_view "go-todo/internal/views/todo"

	"github.com/jeanphorn/log4go"
	"github.com/labstack/echo/v4"
)

func (s *Controller) Home(c echo.Context) error {
	if middleware.SessionManager.Exists(c.Request().Context(), "logged") {
		id, _ := middleware.SessionManager.Get(c.Request().Context(), "id").(int64)
		data, err := s.Db.ListTodosByOwner(context.Background(), id)
		if err != nil {
			log4go.LOGGER("error").Error("Error getting todos: %v", err)
			return err
		}
		c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
		return utils.Render(c, views.Home(data))
	} else {
		return c.Redirect(301, "/auth/")
	}
}

func (s *Controller) Todo(c echo.Context) error {
	ctx := context.Background()
	id, _ := middleware.SessionManager.Get(c.Request().Context(), "id").(int64)
	todos, err := s.Db.ListTodosByOwner(ctx, id)
	if err != nil {
		return err
	}
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return utils.Render(c, todo_view.TodoView(todos))

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
