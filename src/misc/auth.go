package misc

import (
	"golang.org/x/oauth2"
)

var AuthConfig = oauth2.Config{
	ClientSecret: "",
	ClientID:     "",
	RedirectURL:  "http://localhost:8000/auth/redirect",
	Scopes:       []string{"identify"},
	Endpoint: oauth2.Endpoint{
		AuthURL:   "https://discord.com/api/oauth2/authorize",
		TokenURL:  "https://discord.com/api/oauth2/token",
		AuthStyle: oauth2.AuthStyleInParams,
	},
}
