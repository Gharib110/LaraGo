package lara

import (
	"net/http"
	"strconv"

	"github.com/justinas/nosurf"
)

func (l *Lara) SessionLoad(next http.Handler) http.Handler {
	return l.Session.LoadAndSave(next)
}

func (l *Lara) NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	secure, _ := strconv.ParseBool(l.config.cookie.secure)

	csrfHandler.ExemptGlob("/api/*")

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   secure,
		SameSite: http.SameSiteStrictMode,
		Domain:   l.config.cookie.domain,
	})

	return csrfHandler
}
