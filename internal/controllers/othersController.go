package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func MainCookie(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Mensaje string }{"Cookies page"})
}

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	if username == "Jack" && password == "1234" {
		cookie := &http.Cookie{}
		cookie.Name = "sessionID"
		cookie.Value = "valorcito"
		cookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, struct{ Mensaje string }{"Logged in"})

	}
	return c.JSON(http.StatusUnauthorized, struct{ Mensaje string }{"Auth error  "})
}
