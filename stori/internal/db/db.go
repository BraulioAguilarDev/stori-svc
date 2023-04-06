package db

import (
	"embed"
	"stori/config"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var fs embed.FS

type PostgresModule struct{}

func (m *PostgresModule) ProvidePostgresDB() (*gorm.DB, error) {
	dbPool, err := gorm.Open(postgres.Open(config.Config.DSN_DB), &gorm.Config{})

	db, _ := dbPool.DB()
	goose.SetDialect("postgres")
	goose.SetBaseFS(fs)
	if err := goose.Up(db, "migrations"); err != nil {
		return nil, err
	}

	return dbPool, err
}
