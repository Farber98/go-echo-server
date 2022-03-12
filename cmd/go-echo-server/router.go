package main

import (
	"crypto/subtle"
	"net/http"

	controllers "github.com/Farber98/go-echo-server/internal/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initRouter() *echo.Echo {
	r := echo.New()
	initRoutes(r)
	return r
}

func initRoutes(r *echo.Echo) {
	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
	}))

	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("Jac")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("1234")) == 1 {
			return true, nil
		}
		return false, c.JSON(http.StatusUnauthorized, struct{ Mensaje string }{"Not authorized"})
	}))

	adminGroup.GET("/main", controllers.MainAdmin)

	r.GET("/", controllers.TestServer)

	catGroup := r.Group("/cats")
	catGroup.GET("", controllers.GetCatsByQueryParam)
	catGroup.GET("/:id", controllers.GetCatById)
	catGroup.POST("/add/echo", controllers.AddCatByBodyWithEcho)
	catGroup.POST("/add/io", controllers.AddCatByBodyWithIo)
}
