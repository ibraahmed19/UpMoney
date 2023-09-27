package bank

import (
	"template_rest_api/api/app/common"

	"gorm.io/gorm"
)

// create new bank
func NewBank(db *gorm.DB, bank common1.Bank) (common1.Bank, error) {
	return bank, db.Create(&bank).Error
}

// get all banks
func GetBanks(db *gorm.DB) (bank []common1.Bank, err error) {
	return bank, db.Find(&bank).Error
}

// get bank by id
func GetBankByID(db *gorm.DB, id uint) (bank common1.Bank, err error) {
	return bank, db.Find(&bank, "id=?", id).Error
}

// check bank exists
func CheckBankExists(db *gorm.DB, id uint) bool {

	// init vars
	bank := &common1.Bank{}

	// check if row exists
	check := db.Where("id=?", id).First(&bank)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search banks
func SearchBanks(db *gorm.DB, bank common1.Bank) (banks []common1.Bank, err error) {
	return banks, db.Where(&bank).Find(&banks).Error
}

// update bank
func UpdateBank(db *gorm.DB, bank common1.Bank) error {
	return db.Where("id=?", bank.ID).Updates(&bank).Error
}

// delete bank
func DeleteBank(db *gorm.DB, bank_id uint) error {
	return db.Where("id=?", bank_id).Delete(&common1.Bank{}).Error
}
