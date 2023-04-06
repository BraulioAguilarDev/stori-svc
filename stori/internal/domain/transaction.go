package domain

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	ID           string    `gorm:"primaryKey"`
	AccountID    string    `gorm:"not null"`
	Date         time.Time `gorm:"not null"`
	DebitAmount  float64
	CreditAmount float64
	Metadata     json.RawMessage `gorm:"not null"`
	CreateTs     time.Time       `gorm:"autoCreateTime:milli"`
}

func (a *Transaction) TableName() string {
	return "transaction"
}
