package lara

import "log"

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
}

type config struct {
	port     string
	renderer string
}
