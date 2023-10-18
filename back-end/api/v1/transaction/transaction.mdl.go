package transaction

import (
	"fmt"
	"template_rest_api/api/v1/common"
	"time"

	"gorm.io/gorm"
)

// create new transaction
func NewTransaction(db *gorm.DB, transaction common.Transaction) (common.Transaction, error) {
	return transaction, db.Create(&transaction).Error
}

func findByDayOfMonth(db *gorm.DB, dayOfMonth int, transactionType string) (transaction []common.Transaction, err error) {
	return transaction, db.Where("type = ? AND EXTRACT(DAY FROM execution_date) > ?", transactionType, dayOfMonth).Find(&transaction).Error

}

func findByTypeAndCurrentDay(db *gorm.DB, transactionType string) (transaction []common.Transaction, err error) {
	return transaction, db.Where("type = ? AND strftime('%d', execution_date) = strftime('%d', CURRENT_TIMESTAMP)", transactionType).Find(&transaction).Error

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


func ProcessStandingOrders(db *gorm.DB) {
	fmt.Println("Executing SO...")
	now := time.Now()
	lastDayOfMonth := time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	var additionalStandingOrders []common.Transaction

	if now.Month() == time.February || now.Month() == time.April || now.Month() == time.June || now.Month() == time.September || now.Month() == time.November {
		if now.Day() == lastDayOfMonth {
			// You may want to replace this with a call to your transaction repository.
			var err error
			additionalStandingOrders, err = findByDayOfMonth(db, lastDayOfMonth, "StandingOrder")
			if err != nil {
				return
			}
		}
	}

	// Replace this with a call to your transaction repository.
	standingOrders, err := findByTypeAndCurrentDay(db, "StandingOrder")
	if err != nil {
		return
	}

	// Combine the standingOrders and additionalStandingOrders slices here.
	standingOrders = append(standingOrders, additionalStandingOrders...)

	// Loop through standingOrders and process each one.
	for _, transaction := range standingOrders {
		transaction.ID = 0// Replace with the appropriate initialization value.
		transaction.Type = "Transfer"
		transaction.ExecutionDate = time.Now()

		// Replace this with the appropriate saveTransaction function.
		NewTransaction(db,transaction)
	}
}







