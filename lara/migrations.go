package lara

import (
	"github.com/golang-migrate/migrate/v4"
	"log"
)

func (l *Lara) Migrate(dsn string) error {
	m, err := migrate.New("file://"+l.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Up(); err != nil {
		log.Println("Migration failed: ", err)
		return err
	}

	return nil
}

func (l *Lara) Rollback(dsn string) error {
	m, err := migrate.New("file://"+l.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}

	defer m.Close()

	if err = m.Down(); err != nil {
		log.Println("Migration failed: ", err)
	}

	return nil
}

func (l *Lara) Steps(n int, dsn string) error {
	m, err := migrate.New("file://"+l.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}

	defer m.Close()

	if err = m.Steps(n); err != nil {
		return err
	}

	return nil
}

func (l *Lara) MigrateForce(dsn string) error {
	m, err := migrate.New("file://"+l.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Force(-1); err != nil {
		return err
	}

	return nil
}
