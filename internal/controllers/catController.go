package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	helpers "github.com/Farber98/go-echo-server/internal/helpers"
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
	catId, err := helpers.URLParamInt(c, c.Param("id"))
	if err != nil {
		c.Error(err)
	}
	cat := &models.Cat{ID: catId}
	return c.JSON(http.StatusOK, cat)
}

func AddCatByBodyWithEcho(c echo.Context) error {
	cat := models.Cat{}
	if err := c.Bind(&cat); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	fmt.Println("Cat: ", cat)
	return c.JSON(http.StatusOK, struct{ Mensaje string }{"OK"})

}
func AddCatByBodyWithIo(c echo.Context) error {
	cat := models.Cat{}

	defer c.Request().Body.Close()
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	err = json.Unmarshal(b, &cat)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	fmt.Println("Cat: ", cat)
	return c.JSON(http.StatusOK, struct{ Mensaje string }{"OK"})

}
