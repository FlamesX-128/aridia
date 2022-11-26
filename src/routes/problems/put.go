package problems

import (
	"log"

	"github.com/FlamesX-128/aridia/src/database"
	models "github.com/FlamesX-128/aridia/src/models/database"
	"github.com/labstack/echo"
)

func PutProblem(c echo.Context) (err error) {
	var problem models.PutProblem

	if err = c.Bind(&problem); err != nil {
		log.Println("An error occured while binding the problem: ", err)

		c.JSON(500, map[string]string{
			"message": err.Error(),
		})

		return
	}

	if err = database.UpdateProblem(c.Param("id"), problem); err != nil {
		log.Println("An error occured while updating the problem: ", err)

		c.JSON(500, map[string]string{
			"message": err.Error(),
		})

		return
	}

	return c.JSON(200, map[string]string{
		"message": "Problem successfully updated",
	})
}
