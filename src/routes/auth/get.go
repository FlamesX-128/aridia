package auth

import (
	"net/http"
	"time"

	"github.com/FlamesX-128/aridia/src/misc"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func GetAuth(ctx echo.Context) error {
	return ctx.Redirect(
		http.StatusTemporaryRedirect, misc.AuthConfig.AuthCodeURL("identify"),
	)
}

func Redirect(ctx echo.Context) (err error) {
	var token *oauth2.Token
	var resp *http.Response

	if token, err = misc.AuthConfig.Exchange(ctx.Request().Context(), ctx.FormValue("code")); err != nil {
		return ctx.String(400, err.Error())
	}

	if resp, err = misc.AuthConfig.Client(ctx.Request().Context(), token).Get("https://discord.com/api/users/@me"); err != nil {
		return ctx.String(400, err.Error())
	}

	defer resp.Body.Close()

	ctx.SetCookie(&http.Cookie{
		MaxAge: int(token.Expiry.Sub(time.Now()).Seconds()),
		Value:  token.AccessToken,
		Name:   "token",
		Path:   "/",
	})

	return ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
