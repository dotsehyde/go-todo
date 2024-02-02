package controller

import (
	"context"
	"database/sql"
	"go-todo/internal/database"
	"go-todo/internal/middleware"
	"go-todo/internal/utils"
	todo_view "go-todo/internal/views/todo"
	"strconv"

	"github.com/jeanphorn/log4go"
	"github.com/labstack/echo/v4"
)

func (s *Controller) TodoForm(c echo.Context) error {
	if middleware.SessionManager.Exists(c.Request().Context(), "logged") {
		c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
		return utils.Render(c, todo_view.TodoForm())
	} else {
		return c.Redirect(301, "/auth/")
	}
}

func (s *Controller) CreateTodo(c echo.Context) error {
	if middleware.SessionManager.Exists(c.Request().Context(), "logged") {
		ctx := context.Background()
		title := c.FormValue("title")
		content := c.FormValue("content")
		id, _ := middleware.SessionManager.Get(c.Request().Context(), "id").(int64)
		_, err := s.Db.CreateTodo(ctx, database.CreateTodoParams{
			Title: title,
			Content: sql.NullString{
				String: content,
				Valid:  true,
			},
			OwnerID: id,
		})
		if err != nil {
			log4go.LOGGER("error").Error("Error creating todo: %v", err)
			return err
		}
		data, err := s.Db.ListTodosByOwner(ctx, id)
		if err != nil {
			log4go.LOGGER("error").Error("Error getting todos: %v", err)
			return err
		}
		c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
		return utils.Render(c, todo_view.TodoView(data))
	} else {
		return c.Redirect(301, "/auth/")
	}
}

func (s *Controller) DeleteTodo(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	//convert id to int64
	delId, _ := strconv.ParseInt(id, 10, 64)
	s.Db.DeleteTodo(ctx, delId)

	userId, _ := middleware.SessionManager.Get(c.Request().Context(), "id").(int64)
	data, err := s.Db.ListTodosByOwner(ctx, userId)
	if err != nil {
		log4go.LOGGER("error").Error("Error getting todos: %v", err)
		return err
	}
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return utils.Render(c, todo_view.TodoView(data))
}
