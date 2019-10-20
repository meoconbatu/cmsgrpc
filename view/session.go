package view

import (
	"net/http"
	"time"
)

const cookieName = "_goproj_sess"

// GetSession gets the current session from the cookie
func GetSession(w http.ResponseWriter, r *http.Request) string {
	s, err := r.Cookie(cookieName)
	if err != nil {
		http.Error(w, "Please login to view this page", http.StatusUnauthorized)
		return ""
	}
	return s.Value
}

// SetSession sets the session for the given user
func SetSession(w http.ResponseWriter, user string) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    user,
		Expires:  time.Now().Add(time.Hour * 72),
		HttpOnly: true,
	})
}
