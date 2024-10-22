package main

import (
	"fmt"
	"log"

	"database/sql"

	"github.com/earthrot/pisvlek/pkg/config"
	"github.com/earthrot/pisvlek/pkg/db"
	"github.com/earthrot/pisvlek/pkg/server"

	_ "modernc.org/sqlite"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg, err := config.New("assets/pisvlek.yml")
	if err != nil {
		log.Fatal("Unable to load config ", err)
	}

	conn, err := sql.Open("sqlite", fmt.Sprintf("%s?_journal_mode=WAL&_synchronous=normal&_busy_timeout=20000", cfg.Database.Filename))
	if err != nil {
		log.Fatal("Unable to open database ", err)
	}

	driver, err := sqlite.WithInstance(conn, &sqlite.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations/",
		"sqlite", driver)
	if err != nil {
		panic(err)
	}

	fmt.Println("Migrating database...")
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	q := db.New(conn)

	server.Run(cfg, q)
}
