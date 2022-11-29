package auth

import (
	"net/http"
	"time"

	"github.com/FlamesX-128/aridia/src/database"
	"github.com/FlamesX-128/aridia/src/discord"
	"github.com/FlamesX-128/aridia/src/misc"
	mdc "github.com/FlamesX-128/aridia/src/models/discord"
	"github.com/FlamesX-128/aridia/src/tools"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func GetAuth(ctx echo.Context) error {
	return ctx.Redirect(
		http.StatusTemporaryRedirect, misc.AuthConfig.AuthCodeURL("identify"),
	)
}

func Redirect(ctx echo.Context) (err error) {
	var auth *oauth2.Token
	var resp *http.Response

	if auth, err = misc.AuthConfig.Exchange(ctx.Request().Context(), ctx.FormValue("code")); err != nil {
		return ctx.String(400, err.Error())
	}

	if resp, err = misc.AuthConfig.Client(ctx.Request().Context(), auth).Get("https://discord.com/api/users/@me"); err != nil {
		return ctx.String(400, err.Error())
	}

	defer resp.Body.Close()
	var data string

	if data, err = tools.SerializeToken(auth); err != nil {
		return ctx.String(400, err.Error())
	}

	var response mdc.APIUserResponse

	if response, err = discord.GetUser(auth, ctx); err != nil {
		return
	}

	if err = database.CreateUser(response.User.Id, *auth); err != nil {
		return
	}

	ctx.SetCookie(&http.Cookie{
		MaxAge: int(time.Until(auth.Expiry).Milliseconds()),
		Value:  data,
		Name:   "token",
		Path:   "/",
	})

	return ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
