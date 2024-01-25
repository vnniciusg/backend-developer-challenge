package utils

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/viper/configs"
)

func RunMigrations(db *sql.DB) {

	migrationPath := filepath.Join("database", "migrations")

	configs, err := configs.LoadConfig()

	if err != nil {
		log.Println(err)
		return
	}

	if err != nil {
		log.Fatalf("ERROR ON GET CONNECTION: %s", err)
	}

	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatalf("ERROR ON GET DRIVER: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationPath, configs.DBDriver, driver)

	if err != nil {
		log.Fatalf("ERROR ON LOAD MIGRATIONS: %s", err)
		return
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("ERROR ON UPDATE MIGRATIONS: %v", err)
	}

	fmt.Println("Migrations completed successfully.")

}
