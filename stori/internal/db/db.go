package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresModule struct {
	URL string
}

func (m *PostgresModule) ProvidePostgresDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(m.URL), &gorm.Config{})

	return db, err
}
