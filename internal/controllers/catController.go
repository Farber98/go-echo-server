package controllers

import (
	"net/http"
	"strconv"

	models "github.com/Farber98/go-echo-server/internal/models"
	"github.com/labstack/echo/v4"
)

func GetCatsByQueryParam(c echo.Context) error {
	cat := &models.Cat{
		Name: c.QueryParam("name"),
		Type: c.QueryParam("type"),
	}
	return c.JSON(http.StatusOK, cat)
}

func GetCatById(c echo.Context) error {
	catId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
	}
	cat := &models.Cat{ID: catId}
	return c.JSON(http.StatusOK, cat)
}
