package deposit

import (
	"template_rest_api/api/v1/common"
	"time"

	"gorm.io/gorm"
)

type Deposit struct {
	ID            uint               `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	IdTransaction uint               `gorm:"column:id_transaction;not null" json:"id_transaction"`
	Transaction   common.Transaction `gorm:"foreignkey:IdTransaction;references:ID" json:"transaction"`
	IbanAccount   uint               `gorm:"column:iban_account" json:"iban_account"`
	Amount        float64            `gorm:"column:amount;" json:"amount"`
	Date          time.Time          `gorm:"column:date" json:"date"`
	Description   string             `gorm:"column:description;not null" json:"description"`
	CreatedBy     uint               `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new deposit
func NewDeposit(db *gorm.DB, deposit Deposit) (Deposit, error) {
	return deposit, db.Create(&deposit).Error
}

// get all deposits
func GetDeposits(db *gorm.DB) (deposit []Deposit, err error) {
	return deposit, db.Find(&deposit).Error
}

// get deposit by id
func GetDepositByID(db *gorm.DB, id uint) (deposit Deposit, err error) {
	return deposit, db.Find(&deposit, "id=?", id).Error
}

// check deposit exists
func CheckDepositExists(db *gorm.DB, id uint) bool {

	// init vars
	deposit := &Deposit{}

	// check if row exists
	check := db.Where("id=?", id).First(&deposit)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search deposits
func SearchDeposits(db *gorm.DB, deposit Deposit) (deposits []Deposit, err error) {
	return deposits, db.Where(&deposit).Find(&deposits).Error
}

// update deposit
func UpdateDeposit(db *gorm.DB, deposit Deposit) error {
	return db.Where("id=?", deposit.ID).Updates(&deposit).Error
}

// delete deposit
func DeleteDeposit(db *gorm.DB, deposit_id uint) error {
	return db.Where("id=?", deposit_id).Delete(&Deposit{}).Error
}
