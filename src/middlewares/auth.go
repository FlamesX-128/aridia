package middlewares

import (
	"errors"
	"net/http"

	"github.com/FlamesX-128/aridia/src/database"
	"github.com/FlamesX-128/aridia/src/discord"
	mdb "github.com/FlamesX-128/aridia/src/models/database"
	mdc "github.com/FlamesX-128/aridia/src/models/discord"
	"github.com/FlamesX-128/aridia/src/tools"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func GetAuthorOfToken(auth *oauth2.Token, ctx echo.Context) (author mdc.APIUserResponse, err error) {
	// Try to get the user from the token, if it fails, probably the token is invalid.
	if author, err = discord.GetUser(auth, ctx); err != nil {
		return
	}

	return
}

func IsAuthorOfProblem(ctx echo.Context, id string) (isAuthor bool, err error) {
	var auth *oauth2.Token
	var resp mdc.APIUserResponse
	var old mdb.GetProblem

	// Get the token from the session.
	if auth, err = tools.GetAuthToken(ctx); err != nil {
		ctx.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	// Try to get the user from the token, if it fails, probably the token is invalid.
	if resp, err = GetAuthorOfToken(auth, ctx); err != nil {
		ctx.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	// Get the problem from the database.
	if old, err = database.GetProblem(id); err != nil {
		ctx.JSON(500, map[string]string{"message": err.Error()})

		return
	}

	// Check if the author is the owner of the problem.
	if resp.User.Id != old.AuthorID {
		ctx.JSON(403, map[string]string{"message": "You are not the author of this problem."})

		return
	}

	return true, nil
}

func IsAuthenticated(ctx echo.Context) (ok bool, err error) {
	var resp mdc.APIUserResponse
	var auth2 *oauth2.Token

	// Get the token from the session.
	if auth2, err = tools.GetAuthToken(ctx); err != nil {
		return
	}

	// Try to get the user from the token, if it fails, probably the token is invalid.
	if resp, err = discord.GetUser(auth2, ctx); err != nil {
		return
	}

	// If the token is expired, delete it from the session.
	if ok = tools.RemoveCookieIfExpired(ctx.Response(), "token", resp.Expires); ok {
		return false, errors.New("the token has expired")
	}

	return true, nil
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ok, err := IsAuthenticated(ctx); err != nil || !ok {
			return ctx.Redirect(http.StatusTemporaryRedirect, "/auth")
		}

		return next(ctx)
	}
}
