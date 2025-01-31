package lara

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (l *Lara) SessionLoad(next http.Handler) http.Handler {
	l.InfoLog.Println("SessionLoad Called")
	return l.Session.LoadAndSave(next)
}

func (l *Lara) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if l.Debug {
		mux.Use(middleware.Logger)
	}

	mux.Use(middleware.Recoverer)
	mux.Use(l.SessionLoad)

	return mux
}
