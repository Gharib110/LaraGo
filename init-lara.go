package main

import (
	"LaraGo/handlers"
	"LaraGo/lara"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	la := &lara.Lara{
		AppName: "myTestApp",
	}

	err = la.New(path)
	if err != nil {
		log.Fatal(err)
	}

	la.InfoLog.Println("Debug is set to, ", la.Debug)
	myHandlers := &handlers.Handlers{
		App: la,
	}

	app := &application{
		App:      la,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
