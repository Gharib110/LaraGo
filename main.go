package main

import "LaraGo/lara"

type application struct {
	App *lara.Lara
}

func main() {
	app := initApplication()
	app.App.ListenAndServe()
}
