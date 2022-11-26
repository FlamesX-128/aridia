package problems

import (
	"github.com/FlamesX-128/aridia/src/database"
	models "github.com/FlamesX-128/aridia/src/models/database"
	"github.com/labstack/echo"
)

func PostProblem(c echo.Context) (err error) {
	var problem models.PostProblem

	// Deserialize the problem.
	if err = c.Bind(&problem); err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	// Insert the problem.
	if err = database.InsertProblem(problem); err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	return c.JSON(200, map[string]string{
		"message": "Problem successfully created",
	})
}
