package main

import (
	controller "github.com/Farber98/go-echo-server/internal/controllers"
	"github.com/labstack/echo/v4"
)

func initRouter() *echo.Echo {
	r := echo.New()
	initRoutes(r)
	return r
}

func initRoutes(r *echo.Echo) {
	r.GET("/", controller.TestServer)
}
