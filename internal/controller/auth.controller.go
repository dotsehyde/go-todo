package controller

import (
	"fmt"
	"go-todo/internal/database"
	"go-todo/internal/models"
	"go-todo/internal/utils"
	"go-todo/internal/views/auth"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var dbClient = database.Client

func Login(c echo.Context) error {
	sess, _ := session.Get("session", c)
	fmt.Println(sess.Values["logged"])
	if l, ok := sess.Values["logged"].(bool); ok && l {
		fmt.Println(l)
		return c.Redirect(301, "/")
	}
	return utils.Render(c, auth.Login(""))
}

func ProcessLogin(c echo.Context) error {
	// ctx := context.Background()
	email := c.FormValue("email")
	// password := c.FormValue("password")

	// user, err := dbClient.User.FindMany(
	// 	db.User.ID.Equals(email),
	// ).Exec(ctx)
	// if errors.Is(err, db.ErrNotFound) {
	// 	fmt.Println(err.Error())
	// 	return utils.Render(c, auth.Login("Invalid credentials"))
	// } else if err != nil {
	// 	fmt.Println(err.Error())
	// 	return utils.Render(c, auth.Login(err.Error()))
	// }
	// if err := verifyPassword(password, user[0].Password); err != nil {
	// 	fmt.Println(err.Error())
	// 	return utils.Render(c, auth.Login(err.Error()))
	// }
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   0,
		HttpOnly: false,
		Secure:   false, //chane
	}
	sess.Values["logged"] = true
	sess.Values["data"] = models.SessionData{Username: "guest", Email: email, Id: uuid.NewString()}
	sess.Save(c.Request(), c.Response())
	return c.Redirect(301, "/")
}

func Register(c echo.Context) error {
	sess, _ := session.Get("session", c)
	if l, ok := sess.Values["logged"].(bool); ok && l {
		return c.Redirect(301, "/")
	}
	return utils.Render(c, auth.Register(""))
}

func ProcessRegister(c echo.Context) error {
	input := struct {
		Name     string `form:"name" json:"name" validate:"required"`
		Email    string `form:"email" json:"email" validate:"email"`
		Password string `form:"password" json:"password" validate:"required,min=6"`
	}{}
	if err := c.Bind(&input); err != nil {
		return utils.Render(c, auth.Register(err.Error()))
	}
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return utils.Render(c, auth.Register(err.Error()))
	}
	fmt.Println(input.Name)
	// hashPassword, err := hashPassword(input.Password)
	// if err != nil {
	// 	return utils.Render(c, auth.Register(err.Error()))
	// }

	// res, err := dbClient.User.CreateOne(
	// 	db.User.Name.Set(input.Name),
	// 	db.User.Email.Set(input.Email),
	// 	db.User.Password.Set(hashPassword),
	// ).Exec(context.Background())

	// if err != nil {
	// 	return utils.Render(c, auth.Register(err.Error()))
	// }
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   0,
		HttpOnly: false,
		Secure:   false, //chane
	}
	sess.Values["logged"] = true
	sess.Values["data"] = models.SessionData{Username: "guest", Email: input.Email, Id: uuid.NewString()}
	sess.Save(c.Request(), c.Response())
	return c.Redirect(301, "/")
}

func hashPassword(password string) (string, error) {
	// Hashing the password with bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
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
