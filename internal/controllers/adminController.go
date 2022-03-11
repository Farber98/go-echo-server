package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MainAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Mensaje string }{"Secret page"})
}
