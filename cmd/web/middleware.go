package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// SessionLoad is to maintain the state for the current session ..
func SessionLoad(next http.Handler) http.Handler {
	return appConfig.Session.LoadAndSave(next)
}

// NoSurf is to set base cookies for cross site request forgery
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   appConfig.InProduction,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// FileServerUtil is static file server middleware..
func FileServerUtil(dir string) http.Handler {
	return http.FileServer(http.Dir(dir))
}
