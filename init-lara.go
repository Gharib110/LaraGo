package main

import (
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
	app := &application{
		App: la,
	}

	return app
}
