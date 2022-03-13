package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Juanito")
		return next(c)
	}
}

func CheckCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			return err
		}
		if cookie.Value == "valorcito" {
			return next(c)
		}
		return c.JSON(http.StatusOK, struct{ Mensaje string }{"TFCookies"})

	}
}
