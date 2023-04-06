package handler

import (
	"encoding/json"
	"log"
	"stori/internal/domain"
	"stori/internal/report"
	"stori/internal/transaction/repository"
	"stori/pkg/reader"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type TransactionModule struct{}

func (m *TransactionModule) ProvideTransactionModule(repo repository.Transaction) TransactionHandler {
	return &transactionHandler{
		Repo: repo,
	}
}

type TransactionHandler interface {
	ProcessAndSave([]reader.Data) error
	GetSummary() (*report.Summary, error)
}

type transactionHandler struct {
	Repo repository.Transaction
}

func (h *transactionHandler) GetSummary() (*report.Summary, error) {
	balance, average, err := h.Repo.GetSummary()
	if err != nil {
		return nil, err
	}

	var txns []report.TransactionMonth
	var total float64

	for _, b := range balance {
		numberByMon := (b.Credit + b.Debit)
		txns = append(txns, report.TransactionMonth{
			Month:  b.Month,
			Number: numberByMon,
		})
		total += b.Subtotal
	}

	summary := &report.Summary{}
	summary.Total = total
	summary.Transactions = txns

	summary.AverageCreditAmount = (average.Credit / average.NumberCredit)
	summary.AverageDebitAmount = (average.Debit / average.NumberDebit)

	return summary, nil
}

func (h *transactionHandler) ProcessAndSave(data []reader.Data) error {
	accounts := make(map[string]domain.Account)
	for _, row := range data[1:] {
		number, err := strconv.Atoi(row.Number)
		if err != nil {
			panic(err)
		}

		accountId, err := uuid.NewRandom()
		if err != nil {
			panic(err)
		}

		accounts[row.BankName] = domain.Account{
			ID:       accountId.String(),
			BankName: row.BankName,
			Number:   number,
			Currency: row.Currency,
			CreateTs: time.Now(),
		}
	}

	for _, account := range accounts {
		log.Printf("Saving %s bank...\n", account.BankName)
		if err := h.Repo.CreateAccount(account); err != nil {
			log.Printf("creating bank %s error: %v\n", account.BankName, err)
		}
	}

	var transactions []domain.Transaction
	for _, t := range data[1:] {
		log.Printf("Working the transaction: %v\n", t)
		// Pase date
		date, err := time.Parse("2006-01-02", t.Date)
		if err != nil {
			panic(err)
		}

		// Check if is debit/credit
		var debit, credit float64
		mode := t.Amount[0:1]
		if mode == "-" {
			debit, _ = strconv.ParseFloat(t.Amount[1:], 64)
		} else {
			credit, _ = strconv.ParseFloat(t.Amount[1:], 64)
		}

		number, _ := strconv.Atoi(t.Number)
		metadataJson, err := json.Marshal(t)
		if err != nil {
			panic(err)
		}

		// Get bank from db
		bank, err := h.Repo.GetAccountByName(t.BankName)
		if err != nil {
			panic(err)
		}

		transactionID, err := uuid.NewRandom()
		if err != nil {
			panic(err)
		}

		transactions = append(transactions, domain.Transaction{
			ID:              transactionID.String(),
			AccountID:       bank,
			Date:            date,
			DebitAmount:     -debit,
			CreditAmount:    credit,
			OperationNumber: number,
			Metadata:        metadataJson,
			CreateTs:        time.Now(), // TODO: move to autoCreateTime
		})
	}

	for _, txn := range transactions {
		if err := h.Repo.CreateTransaction(txn); err != nil {
			log.Printf("creating transaction error: %v", err)
			continue
		}
	}

	return nil
}
