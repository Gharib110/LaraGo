package lara

import (
	"LaraGo/lara/render"
	"github.com/CloudyKit/jet/v6"
	"github.com/go-chi/chi/v5"
	"log"
)

type initPaths struct {
	rootPath    string
	folderNames []string
}

type Lara struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	config   config
	Routes   *chi.Mux
	Render   *render.Render
	JetViews *jet.Set
}

type config struct {
	port     string
	renderer string
}
