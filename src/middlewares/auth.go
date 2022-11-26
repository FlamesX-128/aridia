package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

var URL = "https://discord.com/api/oauth2/@me"

type Response struct {
	Expires string `json:"expires"`
	User    struct {
		ID string `json:"id"`
	} `json:"user"`
}

func IsAuthenticated(ctx echo.Context) (ok bool, err error) {
	var cookie *http.Cookie

	if cookie, err = ctx.Cookie("token"); err != nil {
		fmt.Println(err.Error())
		return
	}

	postBody, _ := json.Marshal(map[string]string{
		"token": "Bearer" + cookie.Value,
	})

	req, _ := http.NewRequest("GET", URL, bytes.NewBuffer(postBody))
	var content Response

	if err = json.NewDecoder(req.Body).Decode(&content); err != nil {
		return
	}

	return true, nil
}

// echo.MiddlewareFunc
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ok, err := IsAuthenticated(ctx); err != nil || !ok {
			return ctx.Redirect(http.StatusTemporaryRedirect, "/auth")
		}

		return next(ctx)
	}
}
