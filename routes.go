package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) routes() *chi.Mux {

	a.App.Routes.Get("/", a.Handlers.Home)

	a.App.Routes.Get("/jet", func(w http.ResponseWriter, r *http.Request) {
		err := a.App.Render.JetPage(w, r, "home.jet", nil, nil)
		if err != nil {
			a.App.ErrorLog.Println("error rendering: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error rendering 502"))
			return
		}
	})
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))
	return a.App.Routes
}
