package tools

import (
	"net/http"
	"time"
)

func RemoveCookie(w http.ResponseWriter, name string) {
	cookie := http.Cookie{
		Name:   name,
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, &cookie)
}

func RemoveCookieIfExpired(w http.ResponseWriter, name string, expires string) bool {
	end, _ := time.Parse(time.RFC3339, expires)

	if time.Now().After(end) {
		RemoveCookie(w, name)

		return true
	}

	return false
}
