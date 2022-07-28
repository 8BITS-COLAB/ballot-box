package db

import (
	"database/sql"
	"log"

	"github.com/8BITS-COLAB/ballot-box/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, *sql.DB) {
	dsn, err := env.Get("DB_DSN")

	if err != nil {
		log.Fatalf("failed to get db dsn: %s", err)
	}

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	sql, err := d.DB()

	if err != nil {
		log.Fatalf("failed to get sql database: %s", err)
	}

	return d, sql
}
