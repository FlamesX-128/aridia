package problems

import (
	"github.com/FlamesX-128/aridia/src/database"
	models "github.com/FlamesX-128/aridia/src/models/database"
	"github.com/labstack/echo"
)

func GetProblems(c echo.Context) (err error) {
	var problems []models.GetProblem

	// Get the problems.
	if problems, err = database.GetProblems(); err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	return c.JSON(200, problems)
}

func GetProblem(c echo.Context) (err error) {
	var problem models.GetProblem

	// Get the problem.
	if problem, err = database.GetProblem(c.Param("id")); err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	return c.JSON(200, problem)
}
