package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/Gharib110/LaraGo"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init Lara
	lara := &Lara.Lara{}
	err = lara.New(path)
	if err != nil {
		log.Fatal(err)
	}

	lara.AppName = "test-app"

	myMiddleware := &middleware.Middleware{
		App: lara,
	}

	myHandlers := &handlers.Handlers{
		App: lara,
	}

	app := &application{
		App:        lara,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
