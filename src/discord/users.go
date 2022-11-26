package discord

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/FlamesX-128/aridia/src/misc"
	models "github.com/FlamesX-128/aridia/src/models/discord"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

const (
	APIUsersMe = "https://discord.com/api/oauth2/@me"
)

func GetUser(token *oauth2.Token, ctx echo.Context) (user models.APIUserResponse, err error) {
	var resp *http.Response

	if resp, err = misc.AuthConfig.Client(ctx.Request().Context(), token).Get(APIUsersMe); err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return user, errors.New("something went wrong while fetching the user: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err = json.Unmarshal(body, &user); err != nil {
		return
	}

	return
}
