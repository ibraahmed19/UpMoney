package common

import (
	"time"

	"gorm.io/gorm"
)

type BankingAccount struct {
	ID uint `gorm:"primaryKey;autoIncrement;column:id;unique;" json:"id"`
	// BankID       uint      `gorm:"column:bank_id" json:"bank_id"`
	Transactions []Transaction `gorm:"foreignKey:IdBankingAccount;references:ID"`
	ClientID     uint          `gorm:"column:client_id" json:"client_id"`
	Iban         string        `gorm:"column:iban;not null;unique" json:"iban"`
	Balance      float64       `gorm:"column:balance;" json:"balance"`
	Type         string        `gorm:"column:type;not null" json:"type"`
	OpenningDate time.Time     `gorm:"column:openning_date" json:"openning_date"`
	CreatedBy    uint          `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// check bankingaccount exists
func CheckBankingAccountExists(db *gorm.DB, id uint) bool {

	// init vars
	bankingaccount := &BankingAccount{}

	// check if row exists
	check := db.Where("id=?", id).First(&bankingaccount)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
