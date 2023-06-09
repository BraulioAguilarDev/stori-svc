package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"math"
	"stori/config"
	"stori/internal/domain"
	"stori/internal/email"
	"stori/internal/report"
	"stori/internal/transaction/repository"
	"stori/pkg/reader"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type TransactionModule struct{}

func (m *TransactionModule) ProvideTransactionModule(repo repository.Transaction, email email.Email) TransactionHandler {
	return &transactionHandler{
		Repo:  repo,
		Email: email,
	}
}

type TransactionHandler interface {
	ProcessAndSave([]reader.Data) error
	GetSummary() (*report.Summary, error)
	SendSummary(*report.Summary) error
}

type transactionHandler struct {
	Repo  repository.Transaction
	Email email.Email
}

func (h *transactionHandler) ProcessAndSave(data []reader.Data) error {
	accounts := make(map[string]domain.Account)
	for _, row := range data[1:] {
		number, err := strconv.Atoi(row.Number)
		if err != nil {
			return err
		}

		accountId, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		if len(row.AccountEmail) == 0 {
			return errors.New("email is required")
		}

		accounts[row.AccountEmail] = domain.Account{
			ID:           accountId.String(),
			BankName:     row.BankName,
			Number:       number,
			Currency:     row.Currency,
			AccountName:  row.AccountName,
			AccountEmail: row.AccountEmail,
			CreateTs:     time.Now(),
		}
	}

	for _, account := range accounts {
		log.Printf("Saving account: %v\n", account)
		if err := h.Repo.CreateAccount(account); err != nil {
			log.Printf("Creating account error: %s\n", account.AccountName)
		}
	}

	var transactions []domain.Transaction
	for _, t := range data[1:] {
		log.Printf("Working the transaction: %v\n", t)
		// Parsing date
		date, err := time.Parse("2006-01-02", t.Date)
		if err != nil {
			return err
		}

		// Check if is debit/credit
		var debit, credit float64
		mode := t.Amount[0:1]
		if mode == "-" {
			debit, _ = strconv.ParseFloat(t.Amount[1:], 64)
		} else {
			credit, _ = strconv.ParseFloat(t.Amount[1:], 64)
		}

		metadataJson, err := json.Marshal(t)
		if err != nil {
			return err
		}

		// Get bank from db
		bank, err := h.Repo.GetAccountByEmail(t.AccountEmail)
		if err != nil {
			return err
		}

		transactionID, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		transactions = append(transactions, domain.Transaction{
			ID:           transactionID.String(),
			AccountID:    bank,
			Date:         date,
			DebitAmount:  -debit,
			CreditAmount: credit,
			Metadata:     metadataJson,
			CreateTs:     time.Now(), // TODO: use autoCreateTime
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

func (h *transactionHandler) GetSummary() (*report.Summary, error) {
	accountSummary, err := h.Repo.GetSummary()
	if err != nil {
		return nil, err
	}

	var txns []report.TransactionMonth
	var total float64

	for _, b := range accountSummary.Balances {
		numberByMon := (b.Credit + b.Debit)
		txns = append(txns, report.TransactionMonth{
			Month:  b.Month,
			Number: numberByMon,
		})
		total += b.Subtotal
	}

	summary := &report.Summary{}
	// account info
	summary.Email = accountSummary.Email
	summary.Name = accountSummary.Name
	// totals
	summary.Total = roundFloat(total, 2)
	summary.Transactions = txns

	summary.AverageCreditAmount = roundFloat((accountSummary.Average.Credit / accountSummary.Average.NumberCredit), 2)
	summary.AverageDebitAmount = roundFloat((accountSummary.Average.Debit / accountSummary.Average.NumberDebit), 2)

	return summary, nil
}

func (h *transactionHandler) SendSummary(r *report.Summary) error {
	tmp := template.Must(template.ParseFiles("internal/email/template/summary.html"))
	var body bytes.Buffer
	if err := tmp.Execute(&body, r); err != nil {
		return err
	}

	if err := h.Email.Send(r.Email, config.Config.SG_SENDER, "Transactions Summary Report", body.String()); err != nil {
		log.Printf("Sending email error %v", err)
		return err
	}

	return nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
