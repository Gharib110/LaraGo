package lara

import (
	"github.com/gobuffalo/pop"
	"log"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (l *Lara) PopConnect() (*pop.Connection, error) {
	tx, err := pop.Connect("development")
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (l *Lara) CreatePopMigration(up, down []byte, migrationName, migrationType string) error {
	var migrationPath = l.RootPath + "/migrations"
	err := pop.MigrationCreate(migrationPath, migrationName, migrationType, up, down)
	if err != nil {
		return err
	}
	return nil
}

func (l *Lara) RunPopMigrations(tx *pop.Connection) error {
	var migrationPath = l.RootPath + "/migrations"

	fm, err := pop.NewFileMigrator(migrationPath, tx)
	if err != nil {
		return err
	}

	err = fm.Up()
	if err != nil {
		return err
	}
	return nil
}

func (l *Lara) PopMigrateDown(tx *pop.Connection, steps ...int) error {
	var migrationPath = l.RootPath + "/migrations"

	step := 1
	if len(steps) > 0 {
		step = steps[0]
	}

	fm, err := pop.NewFileMigrator(migrationPath, tx)
	if err != nil {
		return err
	}

	err = fm.Down(step)
	if err != nil {
		return err
	}

	return nil
}

func (l *Lara) PopMigrateReset(tx *pop.Connection) error {
	var migrationPath = l.RootPath + "/migrations"

	fm, err := pop.NewFileMigrator(migrationPath, tx)
	if err != nil {
		return err
	}

	err = fm.Reset()
	if err != nil {
		return err
	}
	return nil
}

func (l *Lara) MigrateUp(dsn string) error {
	m, err := migrate.New("file://"+l.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil {
		log.Println("Error running migration:", err)
		return err
	}
	return nil
}

func (l *Lara) MigrateDownAll(dsn string) error {
	m, err := migrate.New("file://"+l.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Down(); err != nil {
		return err
	}

	return nil
}

func (l *Lara) Steps(n int, dsn string) error {
	m, err := migrate.New("file://"+l.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Steps(n); err != nil {
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

	if err := m.Force(-1); err != nil {
		return err
	}

	return nil
}
