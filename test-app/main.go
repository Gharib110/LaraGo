package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/Gharib110/LaraGo"
)

type application struct {
	App        *Lara.Lara
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
