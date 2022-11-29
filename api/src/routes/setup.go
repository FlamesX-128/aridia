package routes

import (
	"github.com/FlamesX-128/aridia/api/src/database"
	"github.com/FlamesX-128/aridia/api/src/middlewares"
	"github.com/labstack/echo"
)

func SetupRoutes(db *database.Database, port string) error {
	r := echo.New()

	r.Use(middlewares.JSON, middlewares.Auth)

	r.DELETE("/users/:id", func(c echo.Context) error {
		return nil
	})

	r.PUT("/users/:id", func(c echo.Context) error {
		return nil
	})

	r.GET("/users/:id", func(c echo.Context) error {
		return nil
	})

	r.POST("/users", func(c echo.Context) error {
		return nil
	})

	r.GET("/users", func(c echo.Context) error {
		return nil
	})

	r.DELETE("/posts/:id", func(c echo.Context) error {
		return nil
	})

	r.PUT("/posts/:id", func(c echo.Context) error {
		return nil
	})

	r.GET("/posts/:id", func(c echo.Context) error {
		return nil
	})

	r.POST("/posts", func(c echo.Context) error {
		return nil
	})

	r.GET("/posts", func(c echo.Context) error {
		return nil
	})

	return r.Start((":" + port))
}
