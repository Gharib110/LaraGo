package lara

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (l *Lara) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if l.Debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)
	mux.Use(l.SessionLoad)
	mux.Use(l.NoSurf)

	return mux
}

// Routes are lara specific routes, which are mounted in the routes file
// in lara applications
func Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/test-c", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("it works!"))
	})
	return r
}
