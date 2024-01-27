package middleware

import (
	"github.com/jeanphorn/log4go"
	"github.com/labstack/echo/v4"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()

		log4go.LOGGER("info").Info("REQ [%s] %s %s", req.Method, req.RequestURI, req.RemoteAddr)

		err := next(c)

		if res.Status >= 400 {
			log4go.LOGGER("error").Error("RES [%s] %s %s %d", req.Method, req.RequestURI, req.RemoteAddr, res.Status)
		} else {
			log4go.LOGGER("info").Info("RES [%s] %s %s %d", req.Method, req.RequestURI, req.RemoteAddr, res.Status)
		}
		return err
	}
}
