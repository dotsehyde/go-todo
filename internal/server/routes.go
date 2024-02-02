package server

import (
	"go-todo/internal/middleware"
	"net/http"

	em "github.com/labstack/echo/v4/middleware"
	scs "github.com/spazzymoto/echo-scs-session"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Static("/", "public")
	e.Use(middleware.LoggerMiddleware)
	e.Use(em.Secure())
	e.Use(em.CORS())
	e.Use(em.Gzip())
	e.Use(scs.LoadAndSave(middleware.SessionManager))
	e.Pre(em.AddTrailingSlash())

	//Routes
	authGroup := e.Group("/auth")
	authGroup.GET("/", s.controller.Login)
	authGroup.POST("/login/", s.controller.ProcessLogin)
	authGroup.GET("/register/", s.controller.Register)
	authGroup.POST("/register/", s.controller.ProcessRegister)
	authGroup.GET("/logout/", s.controller.Logout)

	//Private routes
	// e.Use(middleware.VerifySession)
	// e.Use(echojwt.JWT(em.JWTConfig{
	// 	SigningKey:  []byte(os.Getenv("SESSION_KEY")),
	// 	TokenLookup: "cookie:token",
	// }))
	e.GET("/", s.controller.Home)
	e.GET("/todo/", s.controller.Todo)
	e.GET("/profile/", s.controller.Profile)
	// e.Any("/*", func(c echo.Context) error {
	// 	return c.String(http.StatusNotFound, "404 Not Found - Unknown Path")
	// })

	todo := e.Group("/todo")
	todo.GET("/form/", s.controller.TodoForm)
	todo.POST("/create/", s.controller.CreateTodo)
	todo.DELETE("/delete/:id/", s.controller.DeleteTodo)
	return e
}
