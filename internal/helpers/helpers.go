package helpers

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// URLParamInt retorna un parámetro de la url como int
func URLParamInt(c echo.Context, key string) (int, error) {
	return strconv.Atoi(c.Param(key))
}

// QueryInt retorna un parámetro query como int
func QueryInt(c echo.Context, key string) (int, error) {
	return strconv.Atoi(c.QueryParam(key))
}
