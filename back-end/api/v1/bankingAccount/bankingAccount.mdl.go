package bankingaccount

import (
	"template_rest_api/api/v1/common"

	"gorm.io/gorm"
)

// create new bankingaccount
func NewBankingAccount(db *gorm.DB, bankingaccount common.BankingAccount) (common.BankingAccount, error) {
	return bankingaccount, db.Create(&bankingaccount).Error
}

// get all bankingaccounts
func GetBankingAccounts(db *gorm.DB) (bankingaccount []common.BankingAccount, err error) {
	return bankingaccount, db.Find(&bankingaccount).Error
}

// get baking account by id
func GetBankingAccountByID(db *gorm.DB, id uint) (bankingaccount common.BankingAccount, err error) {
	return bankingaccount, db.Find(&bankingaccount, "id=?", id).Error
}

// get baking account by iban
func GetBankingAccountByIban(db *gorm.DB, iban string) (bankingaccount common.BankingAccount, err error) {
	return bankingaccount, db.Find(&bankingaccount, "iban=?", iban).Error
}

// get banking accounts by user id
func GetBankingAccountsByUserId(db *gorm.DB, user_id uint) ([]common.BankingAccount, error) {
	var bankingAccounts []common.BankingAccount
	if err := db.Where("user_id = ?", user_id).Find(&bankingAccounts).Error; err != nil {
		return nil, err
	}
	return bankingAccounts, nil
}

// search bankingaccounts
func SearchBankingAccounts(db *gorm.DB, bankingaccount common.BankingAccount) (bankingaccounts []common.BankingAccount, err error) {
	return bankingaccounts, db.Where(&bankingaccount).Find(&bankingaccounts).Error
}

// update bankingaccount
func UpdateBankingAccount(db *gorm.DB, bankingaccount common.BankingAccount) error {
	return db.Where("id=?", bankingaccount.ID).Updates(&bankingaccount).Error
}

// delete bankingaccount
func DeleteBankingAccount(db *gorm.DB, bankingaccount_id uint) error {
	return db.Where("id=?", bankingaccount_id).Delete(&common.BankingAccount{}).Error
}
