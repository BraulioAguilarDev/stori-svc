package main

import (
	"stori/internal/db"
	"stori/internal/email"
	"stori/internal/transaction"
	"stori/internal/transaction/handler"
	"stori/internal/transaction/repository"

	"github.com/alecthomas/inject"
	"gorm.io/gorm"
)

type Application struct{}

func (app *Application) Start(db *gorm.DB, repository repository.Transaction,
	handler handler.TransactionHandler, transaction transaction.Transaction) {
	// Config db instance and handlers
	repository.InjectDB(db)
	transaction.AddHandler(handler)

	if err := transaction.Run(); err != nil {
		panic(err)
	}
}

func main() {
	app := new(Application)

	inject := inject.New()
	inject.Install(
		&db.PostgresModule{},
		&repository.RepositoryModule{},
		&handler.TransactionModule{},
		&transaction.TransactionModule{},
		&email.EmailModule{},
	)
	inject.Call(app.Start)
}
