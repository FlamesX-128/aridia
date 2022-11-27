package users

import (
	"github.com/FlamesX-128/aridia/src/database"
	mdb "github.com/FlamesX-128/aridia/src/models/database"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) (err error) {
	var users []mdb.SimplifiedUser

	// Get the problems.
	if users, err = database.GetUsers(); err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	return c.JSON(200, users)
}

func GetUser(c echo.Context) (err error) {
	var user mdb.User

	// Get the problem.
	if user, err = database.GetUser(c.Param("id")); err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	return c.JSON(200, mdb.SimplifiedUser{
		Id:     user.Id,
		Badges: user.Badges,
		Posts:  user.Posts,
	})
}
