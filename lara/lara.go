package lara

import (
	"fmt"
	"log"
	"os"
	"strconv"
)
import "github.com/joho/godotenv"

const version = "1.0.0"

func (l *Lara) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"views", "handlers", "migrations", "data", "public", "tmp", "logs", "middleware"},
	}

	err := l.Init(pathConfig)
	if err != nil {
		return err
	}

	err = l.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// read .env file
	err = godotenv.Load(fmt.Sprintf("%s/.env", rootPath))
	if err != nil {
		return err
	}

	infoLog, errLog := l.startLoggers()
	l.InfoLog = infoLog
	l.ErrorLog = errLog
	l.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	l.Version = version

	return nil
}

func (l *Lara) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := l.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lara) checkDotEnv(path string) error {
	err := l.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}

func (l *Lara) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errLog *log.Logger
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errLog
}
