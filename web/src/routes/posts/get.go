package posts

import (
	"github.com/FlamesX-128/aridia/src/database"
	models "github.com/FlamesX-128/aridia/src/models/database"
	"github.com/labstack/echo"
)

func GetProblems(c echo.Context) (err error) {
	var problems []models.Post

	filter := make(map[string]interface{})

	// Create the filter of the problems.
	for _, v := range []string{"check", "author_id", "title"} {
		if value := c.QueryParam(v); value != "" {
			filter[v] = value
		}
	}

	// Get the problems.
	if problems, err = database.GetPosts(filter); err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	return c.JSON(200, problems)
}

func GetProblem(c echo.Context) (err error) {
	var problem models.Post

	// Get the problem.
	if problem, err = database.GetPost(c.Param("id")); err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	return c.JSON(200, problem)
}
