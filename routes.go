package main

import (
	"fmt"
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
			w.Write([]byte(fmt.Sprintf("error rendering: %s", err)))
			return
		}
	})

	a.App.Routes.Get("/test-session", a.Handlers.SessionTest)
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	// temporary database testing url
	a.App.Routes.Get("test-database", func(w http.ResponseWriter, r *http.Request) {
		query := "select id, first_name from users where id = 1"
		row := a.App.DB.Pool.QueryRowContext(r.Context(), query)

		var id int
		var name string
		err := row.Scan(&id, &name)
		if err != nil {
			a.App.ErrorLog.Println("error scanning: ", err)
			return
		}

		fmt.Fprintf(w, "%d %s", id, name)
	})
	return a.App.Routes
}
