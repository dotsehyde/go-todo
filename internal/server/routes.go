package server

import (
	"go-todo/internal/controller"
	"go-todo/internal/middleware"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	em "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Static("/", "public")
	// e.Use(em.Recover())
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerMiddleware)
	e.Use(em.Secure())
	e.Use(em.CORS())
	e.Use(em.Gzip())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))))
	e.Pre(em.AddTrailingSlash())

	//Routes
	authGroup := e.Group("/auth")
	authGroup.GET("/", controller.Login)
	authGroup.POST("/login/", controller.ProcessLogin)
	authGroup.GET("/register/", controller.Register)
	authGroup.POST("/register/", controller.ProcessRegister)

	//Private routes
	// e.Use(middleware.VerifySession)
	// e.Use(echojwt.JWT(em.JWTConfig{
	// 	SigningKey:  []byte(os.Getenv("SESSION_KEY")),
	// 	TokenLookup: "cookie:token",
	// }))
	e.GET("/", controller.Home)
	// e.Any("/*", func(c echo.Context) error {
	// 	return c.String(http.StatusNotFound, "404 Not Found - Unknown Path")
	// })
	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
