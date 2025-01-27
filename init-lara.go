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
		Debug:   true,
		Version: "",
	}

	err = la.New(path)
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		App: la,
	}

	return app
}
