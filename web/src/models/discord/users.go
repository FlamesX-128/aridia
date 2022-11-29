package discord

type Application struct {
	Id string `json:"id"`
}

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int    `json:"public_flags"`
}

type APIUserResponse struct {
	Application `json:"application"`
	Scopes      []string `json:"scopes"`
	Expires     string   `json:"expires"`
	User        `json:"user"`
}
