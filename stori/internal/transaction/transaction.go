package transaction

import (
	"log"
	"stori/config"
	"stori/internal/report"
	"stori/internal/transaction/repository"
	"stori/pkg/reader"
)

type Handler interface {
	ProcessAndSave([]reader.Data) error
	GetSummary() (*report.Summary, error)
	SendSummary(*report.Summary) error
}

type TransactionModule struct{}

func (m *TransactionModule) ProvideTransactionModule(repo repository.Transaction) Transaction {
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

	data, err := reader.ReadFile(config.Config.CSV_FILE)
	if err != nil {
		log.Printf("Reading file error: %v", err)
		return err
	}

	// Sanitizing data and saving in the db
	if err := t.Handler.ProcessAndSave(data); err != nil {
		log.Printf("ProcessAndSave method failed, error: %v", err)
		return err
	}

	summary, err := t.Handler.GetSummary()
	if err != nil {
		log.Printf("GetSummary method failed, error: %v", err)
		return err
	}

	if err := t.Handler.SendSummary(summary); err != nil {
		log.Printf("SendSummary method failed, error: %v", err)
		return err
	}

	log.Println("Email sent successfully")

	return nil
}
