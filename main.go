package main

import (
	"LaraGo/data"
	"LaraGo/handlers"
	"LaraGo/lara"
)

type application struct {
	App      *lara.Lara
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	app := initApplication()
	app.App.ListenAndServe()
}
