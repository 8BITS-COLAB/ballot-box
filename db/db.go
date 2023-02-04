package db

import (
	"github.com/8BITS-COLAB/ballot-box/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New() *gorm.DB {
	dsn, _ := env.Get("DB_DSN")
	d, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Discard,
	})
	// sql, _ := d.DB()

	return d
}
