package db

import (
	"stori/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresModule struct{}

func (m *PostgresModule) ProvidePostgresDB() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(config.Config.DSN_DB), &gorm.Config{})

	return db, err
}
