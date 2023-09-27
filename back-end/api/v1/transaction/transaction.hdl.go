package transaction

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"template_rest_api/api/v1/common"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new transaction
func (db Database) NewTransaction(ctx *gin.Context) {

	// init vars
	var transaction common.Transaction
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(transaction.IbanSender) || empty_reg.MatchString(transaction.IbanReceiver) || empty_reg.MatchString(string(transaction.Type)) || empty_reg.MatchString(string(transaction.Status)) || empty_reg.MatchString(transaction.Description) { //|| transaction.Amount <= 0 || transaction.Fee <= 0 || transaction.ExecutionDate.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	//check if BankingAccount exists
	if exists := common.CheckBankingAccountExists(db.DB, transaction.IdBankingAccount); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Banking Account does not exist"})
		return
	}

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new transaction
	new_transaction := common.Transaction{
		IdBankingAccount: transaction.IdBankingAccount,
		IbanSender:       transaction.IbanSender,
		IbanReceiver:     transaction.IbanReceiver,
		Type:             transaction.Type,
		Status:           transaction.Status,
		Amount:           transaction.Amount,
		Fee:              transaction.Fee,
		ExecutionDate:    transaction.ExecutionDate,
		Description:      transaction.Description,
		CreatedBy:        session.UserID,
	}

	// create new transaction
	_, err := NewTransaction(db.DB, new_transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all transactions from database
func (db Database) GetTransactions(ctx *gin.Context) {

	// get transactions
	transactions, err := GetTransactions(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

// Get Transactions By Banking account ID retrieves all transactions associated with a given banking account ID
func (db Database) GetTransactionsByBankingAccountId(ctx *gin.Context) {

	// extract banking account ID from request parameters
	IdBankingAccount, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid banking account ID"})
		return
	}

	// get transactions associated with the given banking account ID
	transactions, err := GetTransactionsByBankingAccountId(db.DB, uint(IdBankingAccount))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// return the transactions as a JSON response
	ctx.JSON(http.StatusOK, transactions)
}

// get transaction by id
func (db Database) GetTransactionByID(ctx *gin.Context) {
	//get id value from path
	transaction_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckTransactionExists(db.DB, uint(transaction_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid transaction id"})
		return
	}
	transactionId, err := GetTransactionByID(db.DB, uint(transaction_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transactionId)
}

// search transactions from database
func (db Database) SearchTransactions(ctx *gin.Context) {

	// init vars
	var transaction common.Transaction

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search transactions from database
	transactions, err := SearchTransactions(db.DB, transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

// update transaction
func (db Database) UpdateTransaction(ctx *gin.Context) {

	// init vars
	var transaction common.Transaction
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	transaction_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(transaction.IbanSender) || empty_reg.MatchString(transaction.IbanReceiver) || empty_reg.MatchString(string(transaction.Type)) || empty_reg.MatchString(string(transaction.Status)) || empty_reg.MatchString(transaction.Description) || transaction.Amount == 0 || transaction.Fee == 0 || transaction.ExecutionDate.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	transaction.ID = uint(transaction_id)

	// update transaction
	if err = UpdateTransaction(db.DB, transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteTransaction(ctx *gin.Context) {

	// get id from path
	transaction_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete transaction
	if err = DeleteTransaction(db.DB, uint(transaction_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
