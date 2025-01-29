package main

import (
	"LaraGo/handlers"
	"LaraGo/lara"
)

type application struct {
	App      *lara.Lara
	Handlers *handlers.Handlers
}

func main() {
	app := initApplication()
	app.App.ListenAndServe()
}
