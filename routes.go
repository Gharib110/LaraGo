package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
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

	a.App.Routes.Get("/create-suer",
		func(w http.ResponseWriter, r *http.Request) {

		})

	a.App.Routes.Get("/get-all-users",
		func(w http.ResponseWriter, r *http.Request) {
			users, err := a.Models.Users.GetAll()
			if err != nil {
				a.App.ErrorLog.Println("error fetching all users: ", err)
				return
			}

			for _, user := range users {
				fmt.Fprintf(w, user.LastName)
			}
		})

	a.App.Routes.Get("/get-user/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			id, _ := strconv.Atoi(chi.URLParam(r, "id"))
			u, _ := a.Models.Users.Get(id)
			fmt.Fprintf(w, "%s %s %s", u.LastName, u.FirstName, u.Email)
		})

	a.App.Routes.Get("/update-user/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			id, _ := strconv.Atoi(chi.URLParam(r, "id"))
			u, _ := a.Models.Users.Get(id)

			u.LastName = a.App.RandomString(10)
			err := u.Update(*u)
			if err != nil {
				a.App.ErrorLog.Println("error updating user: ", err)
				return
			}
			fmt.Fprintf(w, "%s %s %s", u.LastName, u.FirstName, u.Email)
		})

	a.App.Routes.Get("/users/login", a.Handlers.UserLogin)
	a.App.Routes.Post("/users/login", a.Handlers.PostUserLogin)

	return a.App.Routes
}
