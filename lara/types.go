package lara

import (
	"LaraGo/lara/render"
	"database/sql"
	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
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
	Session  *scs.SessionManager
	DB       DB
}

type config struct {
	port        string
	renderer    string
	cookie      cookieConfig
	sessionType string
	db          dBConfig
}

type cookieConfig struct {
	name     string
	domain   string
	lifetime string
	persist  string
	secure   string
}

type dBConfig struct {
	dsn string
	db  string
}

type DB struct {
	DataType string
	Pool     *sql.DB
}
