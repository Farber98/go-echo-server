package main

import (
	controllers "github.com/Farber98/go-echo-server/internal/controllers"
	"github.com/labstack/echo/v4"
)

func initRouter() *echo.Echo {
	r := echo.New()
	initRoutes(r)
	return r
}

func initRoutes(r *echo.Echo) {
	r.GET("/", controllers.TestServer)
	r.GET("/cats", controllers.GetCatsByQueryParam)
	r.GET("/cats/:id", controllers.GetCatById)
	r.POST("/catecho", controllers.AddCatByBodyWithEcho)
	r.POST("/catio", controllers.AddCatByBodyWithIo)
}
