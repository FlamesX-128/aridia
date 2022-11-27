package problems

import (
	"github.com/FlamesX-128/aridia/src/database"
	"github.com/FlamesX-128/aridia/src/discord"
	models "github.com/FlamesX-128/aridia/src/models/database"
	mdc "github.com/FlamesX-128/aridia/src/models/discord"
	"github.com/FlamesX-128/aridia/src/tools"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func PostProblem(c echo.Context) (err error) {
	var post models.Post

	// Deserialize the problem.
	if err = c.Bind(&post); err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	// Get user
	var resp mdc.APIUserResponse
	var auth *oauth2.Token
	var user models.User

	// Get the token from the cookie.
	if auth, err = tools.GetAuthToken(c); err != nil {
		return
	}

	// Get the user id from the discord with the token.
	if resp, err = discord.GetUser(auth, c); err != nil {
		return
	}

	// Get the user from the database.
	if user, err = database.GetUser(resp.User.Id); err != nil {
		return
	}

	// If the user isn't a admin, he can't set a author id.
	if !user.Admin {
		post.AuthorID = resp.User.Id
	}

	// This fields can't be set.
	post.CreatedAt = tools.GetCurrentTime()
	post.Id = ""

	// Insert the problem.
	if err = database.InsertProblem(post); err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	return c.JSON(200, map[string]string{
		"message": "Problem successfully created",
	})
}
