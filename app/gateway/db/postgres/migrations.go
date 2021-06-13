package postgres

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/sptGabriel/banking/app"
	"net/http"
)

//go:embed migrations
var _migrations embed.FS

func GetMigrationHandler(dbURL string) (*migrate.Migrate, error) {
	operation := "Postgres.GetMigrationHandler"
	// use httpFS until go-migrate implements ioFS
	// (see https://github.com/golang-migrate/migrate/issues/480#issuecomment-731518493)
	source, err := httpfs.New(http.FS(_migrations), "migrations")
	if err != nil {
		return nil, app.Err(operation, err)
	}

	m, err := migrate.NewWithSourceInstance("httpfs", source, dbURL)
	if err != nil {
		return nil, app.Err(operation, err)
	}

	return m, nil
}

func RunMigrations(dbURL string) error {
	operation := "Postgres.RunMigrations"

	m, err := GetMigrationHandler(dbURL)
	if err != nil {
		return app.Err(operation, err)
	}
	defer m.Close()

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		return app.Err(operation, err)
	}

	return nil
}
