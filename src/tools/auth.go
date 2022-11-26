package tools

import (
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func GetAuthToken(ctx echo.Context) (auth2 *oauth2.Token, err error) {
	var cookie *http.Cookie

	if cookie, err = ctx.Cookie("token"); err != nil {
		return
	}

	if auth2, err = DeserializeToken([]byte(cookie.Value)); err != nil {
		return
	}

	return auth2, nil
}
