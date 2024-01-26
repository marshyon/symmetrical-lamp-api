package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required for migrations
	_ "github.com/lib/pq"
)

func (d *Database) MigrateDB() error {
	fmt.Println("Migrating DB...")
	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create postres driver : %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("could not run up migrate : %w", err)
		}
	}
	fmt.Println("Migrated DB...")
	return nil
}
