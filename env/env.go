package env

import (
	"database/sql"
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Env struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func New() (*gorm.DB, *sql.DB) {
	d, err := gorm.Open(sqlite.Open("env.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	sql, err := d.DB()

	if err != nil {
		log.Fatalf("failed to get sql database: %s", err)
	}

	d.AutoMigrate(&Env{})

	return d, sql
}

func Set(key, value string) {
	d, sql := New()
	defer sql.Close()

	var e Env

	if err := d.Where("key = ?", key).First(&e).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		d.Create(&Env{Key: key, Value: value})
	}

	if err := d.Model(&e).Where("key = ?", key).UpdateColumn("value", value).Error; err != nil && value != e.Value {
		log.Fatalf("failed to set config: %s", err)
	}

}

func Get(key string) (string, error) {
	d, sql := New()
	defer sql.Close()

	var e Env

	if err := d.Where("key = ?", key).First(&e).Error; err != nil {
		return "", err
	}

	return e.Value, nil
}
