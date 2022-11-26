package routes

import (
	"log"
	"path"

	"github.com/FlamesX-128/aridia/src/database"
	"github.com/FlamesX-128/aridia/src/routes/problems"
	"github.com/labstack/echo"
)

func SetupRoutes(dirPath string, rPort string, dUri string) error {
	if err := database.Connect(dUri); err != nil {
		return err
	}

	r := echo.New()

	// Api routes.
	r.GET("/api/problems/:id", problems.GetProblem)
	r.PUT("/api/problems/:id", problems.PutProblem)

	r.POST("/api/problems", problems.PostProblem)
	r.GET("/api/problems", problems.GetProblems)

	// Static files.
	r.Static("/public", path.Join(dirPath, "public"))

	// User routes.
	r.GET("/", func(ctx echo.Context) error {
		return ctx.File(path.Join(dirPath, "templates", "home.html"))
	})

	log.Println("Listening on port ", rPort)
	return r.Start((":" + rPort))
}
