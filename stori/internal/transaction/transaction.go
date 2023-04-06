package transaction

import (
	"log"
	"stori/internal/report"
	"stori/internal/transaction/repository"
	"stori/pkg/reader"
)

type Handler interface {
	ProcessAndSave([]reader.Data) error
	GetSummary() (*report.Summary, error)
}

type TransactionModule struct{}

func (m *TransactionModule) ProvideRabbitClientModule(repo repository.Transaction) Transaction {
	return &transaction{
		Repo: repo,
	}
}

type transaction struct {
	Handler Handler
	Repo    repository.Transaction
}

type Transaction interface {
	AddHandler(Handler)
	Run() error
}

func (t *transaction) AddHandler(handler Handler) {
	t.Handler = handler
}

func (t *transaction) Run() error {
	log.Println("Reading csv file from data directory...")

	// TODO: read file from env or flag params
	data, err := reader.ReadFile("data/txns.csv")
	if err != nil {
		return err
	}

	// Sanitizing data and saving in the db
	if err := t.Handler.ProcessAndSave(data); err != nil {
		return err
	}

	summary, err := t.Handler.GetSummary()
	if err != nil {
		return err
	}

	log.Println(summary)
	// TODO: send email
	return nil
}
