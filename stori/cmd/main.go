package main

import (
	"stori/config"
	"stori/internal/db"
	"stori/internal/transaction"
	"stori/internal/transaction/handler"
	"stori/internal/transaction/repository"

	"github.com/alecthomas/inject"
	"gorm.io/gorm"
)

type Application struct{}

func (app *Application) Start(db *gorm.DB, repo repository.Transaction,
	handler handler.TransactionHandler, txns transaction.Transaction) {
	// Config db instance and handlers
	repo.InjectDB(db)
	txns.AddHandler(handler)

	if err := txns.Run(); err != nil {
		panic(err)
	}
}

func main() {
	app := new(Application)

	inject := inject.New()
	inject.Install(
		&db.PostgresModule{
			URL: config.Config.DSN_DB,
		},
		&repository.RepositoryModule{},
		&handler.TransactionModule{},
		&transaction.TransactionModule{},
	)
	inject.Call(app.Start)
}
