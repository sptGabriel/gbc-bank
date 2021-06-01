package postgres

import (
	"embed"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/sptGabriel/banking/app"
	"net/http"
)

//go:embed migrations
var _migrations embed.FS

func GetMigrationHandler(dbURL string) (*migrate.Migrate, error) {
	// use httpFS until go-migrate implements ioFS
	// (see https://github.com/golang-migrate/migrate/issues/480#issuecomment-731518493)
	source, err := httpfs.New(http.FS(_migrations), "migrations")
	if err != nil {
		return nil, app.NewInternalError("failed to init httpfs", err)
	}

	m, err := migrate.NewWithSourceInstance("httpfs", source, dbURL)
	if err != nil {
		return nil, app.NewInternalError("failed to get migration source", err)
	}

	return m, nil
}

func RunMigrations(dbURL string) error {
	m, err := GetMigrationHandler(dbURL)
	if err != nil {
		return app.NewInternalError("failed to get migration handler", err)
	}
	defer m.Close()

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return app.NewInternalError("failed to execute transactions", err)
	}

	return nil
}
