package database

type Device struct {
	Id   string `json:"id" rethinkdb:"id,omitempty"`
	SSID string `json:"ssid" rethinkdb:"ssid,omitempty"`
	Type string `json:"type" rethinkdb:"type,omitempty"`
}

type Connection struct {
	Name string `json:"name" rethinkdb:"name,omitempty"`
	Url  string `json:"url" rethinkdb:"url,omitempty"`
}

type User struct {
	Id       string `json:"id" rethinkdb:"id,omitempty"`
	Username string `json:"username" rethinkdb:"username,omitempty"`
	Password string `json:"password" rethinkdb:"password,omitempty"`
	Email    string `json:"email" rethinkdb:"email,omitempty"`

	Badges []string `json:"badges" rethinkdb:"badges"`
	Posts  []string `json:"posts" rethinkdb:"posts"`

	Connections []Connection `json:"connections" rethinkdb:"connections"`
	AuthDevices []Device     `json:"auth_devices" rethinkdb:"auth_devices"`

	CreatedAt string `json:"created_at" rethinkdb:"created_at,omitempty"`
}
