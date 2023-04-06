package repository

import (
	"errors"
	"stori/internal/domain"
	"stori/internal/report"

	"gorm.io/gorm"
)

type RepositoryModule struct{}

func (m *RepositoryModule) ProvideRepositoryModule() Transaction {
	return &repository{}
}

type Transaction interface {
	// Common
	InjectDB(db *gorm.DB)

	// Account model
	GetAccountByName(string) (string, error)
	CreateAccount(data domain.Account) error

	// Transaction model
	CreateTransaction(data domain.Transaction) error
	GetSummary() (*report.AccountSummary, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) InjectDB(db *gorm.DB) {
	r.db = db
}

func (r *repository) GetSummary() (*report.AccountSummary, error) {
	// Get account info
	var account domain.Account
	if err := r.db.Last(&account).Error; err != nil {
		return nil, err
	}

	if account.AccountEmail == "" {
		return nil, errors.New("account not valid")
	}

	// Getting balance and count transactions by month
	/*
		month|debit|credit|subtotal|
		---+-----+------+--------+
		08 |    1|     1|  -10.46|
		07 |    1|     1|   50.20|
	*/

	var balance []report.Balance
	balanceNumbersQuery := `
		SELECT to_char(date,'mon') AS month,
			COUNT(case when txn.debit_amount != 0 then txn.debit_amount end) AS debit,
			COUNT(case when txn.credit_amount != 0 then txn.credit_amount end) AS credit,
			SUM(txn.debit_amount + txn.credit_amount) AS subtotal
		FROM transaction txn GROUP BY 1;
	`

	if err := r.db.Raw(balanceNumbersQuery).Scan(&balance).Error; err != nil {
		return nil, err
	}

	// Average info
	/*
		credit|debit |number_debit|number_credit|
		------+------+-----------+-----------+
		70.50 | -30.76|          2|          2|
	*/
	var average *report.Average
	averageQuery := `
		SELECT 
			SUM(t.credit_amount) AS credit, 
			SUM(t.debit_amount) AS debit,
			COUNT(case when t.debit_amount != 0 then t.debit_amount end) AS number_debit,
			COUNT(case when t.credit_amount != 0 then t.credit_amount end) AS number_credit
		FROM transaction t`

	if err := r.db.Raw(averageQuery).Scan(&average).Error; err != nil {
		return nil, err
	}

	return &report.AccountSummary{
		Email:    account.AccountEmail,
		Name:     account.AccountName,
		Balances: balance,
		Average:  *average,
	}, nil
}

func (r *repository) GetAccountByName(bank string) (string, error) {
	var account domain.Account
	if err := r.db.First(&account, "bank_name", bank).Error; err != nil {
		return "", err
	}

	return account.ID, nil
}

func (r *repository) CreateAccount(data domain.Account) error {
	if err := r.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateTransaction(data domain.Transaction) error {
	if err := r.db.Create(data).Error; err != nil {
		return nil
	}

	return nil
}
