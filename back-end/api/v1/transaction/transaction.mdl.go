package transaction

import (
	"template_rest_api/api/v1/common"

	"gorm.io/gorm"
)

// create new transaction
func NewTransaction(db *gorm.DB, transaction common.Transaction) (common.Transaction, error) {
	return transaction, db.Create(&transaction).Error
}

// get all transactions
func GetTransactions(db *gorm.DB) (transaction []common.Transaction, err error) {
	return transaction, db.Find(&transaction).Error
}

// get all transactions by banking account id
func GetTransactionsByBankingAccountId(db *gorm.DB, IdBankingAccount uint) ([]common.Transaction, error) {
	var transactions []common.Transaction
	if err := db.Where("id_banking_account = ?", IdBankingAccount).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// get transaction by id
func GetTransactionByID(db *gorm.DB, id uint) (transaction common.Transaction, err error) {
	return transaction, db.Find(&transaction, "id=?", id).Error
}

// check transaction exists
func CheckTransactionExists(db *gorm.DB, id uint) bool {

	// init vars
	transaction := &common.Transaction{}

	// check if row exists
	check := db.Where("id=?", id).First(&transaction)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search transactions
func SearchTransactions(db *gorm.DB, transaction common.Transaction) (transactions []common.Transaction, err error) {
	return transactions, db.Where(&transaction).Find(&transactions).Error
}

// update transaction
func UpdateTransaction(db *gorm.DB, transaction common.Transaction) error {
	return db.Where("id=?", transaction.ID).Updates(&transaction).Error
}

// delete transaction
func DeleteTransaction(db *gorm.DB, transaction_id uint) error {
	return db.Where("id=?", transaction_id).Delete(&common.Transaction{}).Error
}
