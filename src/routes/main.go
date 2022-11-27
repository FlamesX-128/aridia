package routes

import (
	"log"
	"path"

	"github.com/FlamesX-128/aridia/src/database"
	"github.com/FlamesX-128/aridia/src/middlewares"
	"github.com/FlamesX-128/aridia/src/routes/auth"
	"github.com/FlamesX-128/aridia/src/routes/posts"
	"github.com/FlamesX-128/aridia/src/routes/users"
	"github.com/labstack/echo"
)

func SetupRoutes(dirPath string, rPort string, dUri string) error {
	r := echo.New()

	// Setup the RethinkDB.
	if err := database.Connect(dUri); err != nil {
		return err
	}

	if err := database.Setup(); err != nil {
		return err
	}

	// Api routes.
	api := r.Group("/api", middlewares.JSON, middlewares.Auth)

	api.GET("/post/:id", posts.GetProblem)
	api.PUT("/post/:id", posts.PutProblem)

	api.POST("/posts", posts.PostProblem)
	api.GET("/posts", posts.GetProblems)

	api.GET("/user/:id", users.GetUser)
	api.GET("/users", users.GetUsers)

	// Static files.
	r.Static("/public", path.Join(dirPath, "public"))

	// Auth routes.
	r.GET("/auth/redirect", auth.Redirect)
	r.GET("/auth", auth.GetAuth)

	// User routes.
	r.GET("/", func(ctx echo.Context) error {
		return ctx.File(path.Join(dirPath, "views", "home.html"))
	})

	log.Println("Listening on port ", rPort)
	return r.Start((":" + rPort))
}
