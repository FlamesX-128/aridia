package tools

import (
	"encoding/base64"
	"encoding/json"

	"golang.org/x/oauth2"
)

func SerializeToken(t *oauth2.Token) (str string, err error) {
	var data []byte

	if data, err = json.Marshal(t); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

func DeserializeToken(data []byte) (auth *oauth2.Token, err error) {
	if data, err = base64.StdEncoding.DecodeString(string(data)); err != nil {
		return
	}

	if err = json.Unmarshal(data, &auth); err != nil {
		return
	}

	return
}
