package lara

import (
	"LaraGo/lara/render"
	"LaraGo/session"
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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
	l.RootPath = rootPath
	l.Routes = l.routes().(*chi.Mux)

	l.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
		cookie: cookieConfig{
			name:     os.Getenv("COOKIE_NAME"),
			lifetime: os.Getenv("COOKIE_LIFETIME"),
			persist:  os.Getenv("COOKIE_PERSISTS"),
			secure:   os.Getenv("COOKIE_SECURE"),
			domain:   os.Getenv("COOKIE_DOMAIN"),
		},
		sessionType: os.Getenv("SESSION_TYPE"),
	}

	sess := session.Session{
		CookieLifeTime: l.config.cookie.lifetime,
		CookiePersist:  l.config.cookie.persist,
		CookieName:     l.config.cookie.name,
		CookieDomain:   l.config.cookie.domain,
		SessionType:    l.config.sessionType,
	}

	l.Session = sess.InitSession()

	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)
	l.JetViews = views

	l.createRenderer()

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

// ListenAndServe starts the server
func (l *Lara) ListenAndServe() {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", l.config.port),
		ErrorLog:          l.ErrorLog,
		Handler:           l.Routes,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	l.InfoLog.Printf("Starting server on port %s", l.config.port)

	err := srv.ListenAndServe()
	if err != nil {
		l.ErrorLog.Fatal(err)
	}
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

func (l *Lara) createRenderer() {
	myRenderer := render.Render{
		Renderer: l.config.renderer,
		RootPath: l.RootPath,
		Port:     l.config.port,
		JetViews: l.JetViews,
	}
	l.Render = &myRenderer
}

func (l *Lara) BuildDSN() string {
	var dsn string

	switch os.Getenv("DATABASE_TYPE") {
	case "postgres", "postgresql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s "+
			"dbname=%s sslmode=%s timezone=UTC conntect_timeout=5",
			os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_USERNAME"),
			os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_SSL_MODE"))
	default:

	}

	return dsn
}
