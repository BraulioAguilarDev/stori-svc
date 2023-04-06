package domain

import "time"

type Account struct {
	ID           string    `gorm:"primaryKey"`
	BankName     string    `gorm:"not null"`
	Number       int       `gorm:"not null"`
	Currency     string    `gorm:"not null"`
	AccountName  string    `gorm:"not null"`
	AccountEmail string    `gorm:"not null"`
	CreateTs     time.Time `gorm:"autoCreateTime:milli"`
}

func (a *Account) TableName() string {
	return "account"
}
