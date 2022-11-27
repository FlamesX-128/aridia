package problems

import (
	"github.com/FlamesX-128/aridia/src/database"
	"github.com/FlamesX-128/aridia/src/middlewares"
	mdb "github.com/FlamesX-128/aridia/src/models/database"
	mdc "github.com/FlamesX-128/aridia/src/models/discord"
	"github.com/FlamesX-128/aridia/src/tools"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func PutProblem(ctx echo.Context) (err error) {
	var data mdc.APIUserResponse
	var auth *oauth2.Token

	var new mdb.Post
	var old mdb.Post

	// Deserialize the problem.
	if err = ctx.Bind(&new); err != nil {
		ctx.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	// Get the token from the session.
	if auth, err = tools.GetAuthToken(ctx); err != nil {
		ctx.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	// Get the old problem from the database.
	if old, err = database.GetProblem(new.Id); err != nil {
		ctx.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	// Get the owner of the token.
	if data, _ = middlewares.GetAuthorOfToken(auth, ctx); err != nil {
		ctx.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	// Check if the author is the owner of the problem.
	if data.User.Id != old.AuthorID {
		ctx.JSON(403, map[string]string{"message": "You are not the author of this problem."})

		return
	}

	// Update the problem.
	if err = database.InsertProblem(new); err != nil {
		ctx.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	return ctx.JSON(200, map[string]string{
		"message": "Problem successfully updated",
	})
}
