package routes

import (
	"log"
	"path"

	"github.com/FlamesX-128/aridia/src/database"
	"github.com/FlamesX-128/aridia/src/middlewares"
	"github.com/FlamesX-128/aridia/src/routes/auth"
	"github.com/FlamesX-128/aridia/src/routes/problems"
	"github.com/labstack/echo"
)

func SetupRoutes(dirPath string, rPort string, dUri string) error {
	r := echo.New()

	if err := database.Connect(dUri); err != nil {
		return err
	}

	if err := database.Setup(); err != nil {
		return err
	}

	// Api routes.
	api := r.Group("/api", middlewares.JSON, middlewares.Auth)

	api.GET("/problems/:id", problems.GetProblem)
	api.PUT("/problems/:id", problems.PutProblem)

	api.POST("/problems", problems.PostProblem)
	api.GET("/problems", problems.GetProblems)

	// Static files.
	r.Static("/public", path.Join(dirPath, "public"))

	// Auth routes.
	r.GET("/auth/redirect", auth.Redirect)
	r.GET("/auth", auth.GetAuth)

	// User routes.
	r.GET("/", func(ctx echo.Context) error {
		return ctx.File(path.Join(dirPath, "templates", "home.html"))
	})

	log.Println("Listening on port ", rPort)
	return r.Start((":" + rPort))
}
