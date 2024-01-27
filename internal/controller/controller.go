package controller

import (
	"go-todo/internal/database"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

type Controller struct {
	Db *database.Queries
}
