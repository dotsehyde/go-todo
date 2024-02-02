package controller

import (
	"context"
	"go-todo/internal/database"
	"go-todo/internal/middleware"
	"go-todo/internal/utils"
	"go-todo/internal/views/auth"
	"time"

	"github.com/jeanphorn/log4go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Controller) Login(c echo.Context) error {

	if middleware.SessionManager.Exists(c.Request().Context(), "logged") {
		return c.Redirect(301, "/")
	}
	return utils.Render(c, auth.Login(""))
}

func (s *Controller) ProcessLogin(c echo.Context) error {
	ctx := context.Background()
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		log4go.LOGGER("error").Error("Error validating input: %v", "Email or password is empty")
		return utils.Render(c, auth.Login("Invalid credentials"))
	}
	user, err := s.Db.GetUserByEmail(ctx, email)
	if err != nil {
		log4go.LOGGER("error").Error("Error getting user: %v", err)
		return utils.Render(c, auth.Login("User not found"))
	}
	if err := verifyPassword(password, user.Password); err != nil {
		log4go.LOGGER("error").Error("Error verifying password: %v", err)
		return utils.Render(c, auth.Login("Invalid credentials"))
	}
	middleware.SessionManager.Put(c.Request().Context(), "logged", true)
	middleware.SessionManager.Put(c.Request().Context(), "name", user.Name)
	middleware.SessionManager.Put(c.Request().Context(), "email", user.Email)
	middleware.SessionManager.Put(c.Request().Context(), "id", user.ID)
	return c.Redirect(301, "/")
}

func (s *Controller) Register(c echo.Context) error {

	if middleware.SessionManager.Exists(c.Request().Context(), "logged") {
		return c.Redirect(301, "/")
	}
	return utils.Render(c, auth.Register(""))
}

func (s *Controller) ProcessRegister(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	input := struct {
		Name     string `form:"name" json:"name" validate:"required"`
		Email    string `form:"email" json:"email" validate:"email"`
		Password string `form:"password" json:"password" validate:"required,min=6"`
	}{}
	if err := c.Bind(&input); err != nil {
		log4go.LOGGER("error").Error("Error binding input: %v", err)
		return utils.Render(c, auth.Register(err.Error()))
	}

	if err := Validate.Struct(input); err != nil {
		log4go.LOGGER("error").Error("Error validating input: %v", err)
		return utils.Render(c, auth.Register(err.Error()))
	}

	hashPassword, err := hashPassword(input.Password)
	if err != nil {
		log4go.LOGGER("error").Error("Error hashing password: %v", err)
		return utils.Render(c, auth.Register(err.Error()))
	}
	input.Password = hashPassword
	// check if user exists
	_, err = s.Db.GetUserByEmail(ctx, input.Email)
	if err == nil {
		log4go.LOGGER("error").Error("Error creating user: %v", err)
		return utils.Render(c, auth.Register("User already exists"))
	}
	user, err := s.Db.CreateUser(ctx, database.CreateUserParams{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashPassword,
	})

	if err != nil {
		log4go.LOGGER("error").Error("Error creating user: %v", err)
		return utils.Render(c, auth.Register(err.Error()))
	}
	log4go.LOGGER("info").Info("User created: %v", user)
	middleware.SessionManager.Put(c.Request().Context(), "logged", true)
	middleware.SessionManager.Put(c.Request().Context(), "name", user.Name)
	middleware.SessionManager.Put(c.Request().Context(), "email", user.Email)
	middleware.SessionManager.Put(c.Request().Context(), "id", user.ID)

	return c.Redirect(301, "/")
}

func (s *Controller) Logout(c echo.Context) error {

	middleware.SessionManager.Destroy(c.Request().Context())
	return c.Redirect(301, "/auth/")
}

func hashPassword(password string) (string, error) {
	// Hashing the password with bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func verifyPassword(inputPassword, hashedPassword string) error {
	// Comparing the input password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err
}
