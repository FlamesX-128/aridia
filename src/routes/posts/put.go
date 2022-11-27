package posts

import (
	"github.com/FlamesX-128/aridia/src/database"
	"github.com/FlamesX-128/aridia/src/discord"
	mdb "github.com/FlamesX-128/aridia/src/models/database"
	mdc "github.com/FlamesX-128/aridia/src/models/discord"
	"github.com/FlamesX-128/aridia/src/tools"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func PutProblem(ctx echo.Context) (err error) {
	var auth *oauth2.Token

	var new mdb.Post
	var old mdb.Post

	new.Id = ctx.Param("id")

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
	if old, err = database.GetPost(new.Id); err != nil {
		ctx.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	var resp mdc.APIUserResponse
	var user mdb.User

	// Get the user id from the discord with the token.
	if resp, err = discord.GetUser(auth, ctx); err != nil {
		return
	}

	// Get the user from the database.
	if user, err = database.GetUser(resp.User.Id); err != nil {
		return
	}

	// Check if the author is the owner of the problem.
	if resp.User.Id != old.AuthorID && !user.Admin {
		ctx.JSON(403, map[string]string{"message": "You are not the author of this problem."})

		return
	}

	// Update old post with the new values.
	old.Title = new.Title
	old.Desc = new.Desc

	old.Answer = new.Answer
	old.Steps = new.Steps

	// Only remove check if the user isn't a admin.
	if !user.Admin {
		old.Check = mdb.NONE
	}

	// Update the post from the database.
	if _, err = database.InsertPost(old); err != nil {
		ctx.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	return ctx.JSON(200, map[string]string{
		"message": "Post successfully updated",
	})
}
