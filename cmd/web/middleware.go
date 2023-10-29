package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// csrf protection for all post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// session loads and saves the session in every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
