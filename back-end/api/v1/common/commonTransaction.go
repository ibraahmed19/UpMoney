package common

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID               uint `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	IdBankingAccount uint `gorm:"column:id_banking_account" json:"id_banking_account"`
	// BankingAccount   BankingAccount `gorm:"foreignkey:IdBankingAccount;references:ID" json:"banking_account"`
	IbanSender   string          `gorm:"column:iban_sender;not null" json:"iban_sender"`
	IbanReceiver string          `gorm:"column:iban_receiver;not null" json:"iban_receiver"`
	Type         TransactionType `gorm:"column:type;" json:"type"`
	//Status        TransactionStatus `gorm:"column:status;" json:"status"`
	Amount        float64   `gorm:"column:amount;" json:"amount"`
	ExecutionDate time.Time `gorm:"column:execution_date" json:"execution_date"`
	Description   string    `gorm:"column:description;" json:"description"`
	CreatedBy     uint      `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

type TransactionType string

const (
	Transfer      TransactionType = "Transfer"
	StandingOrder TransactionType = "StandingOrder"
)

type TransactionStatus string

const (
	Created   TransactionStatus = "created"
	Completed TransactionStatus = "completed"
)
