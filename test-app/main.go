package main

import (
	"test-app/data"
	"test-app/handlers"
	"test-app/middleware"

	"github.com/Gharib110/LaraGo"
)

type application struct {
	App        *lara.Lara
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	l := initApplication()
	l.App.ListenAndServe()
}
