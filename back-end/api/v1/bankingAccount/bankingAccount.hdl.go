package bankingaccount

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"template_rest_api/api/v1/common"
	"template_rest_api/middleware"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new bankingaccount
func (db Database) NewBankingAccount(ctx *gin.Context) {

	// init vars
	var bankingaccount common.BankingAccount
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&bankingaccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(bankingaccount.Type) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check if user exists
	if exists := common.CheckUserExists(db.DB, bankingaccount.UserID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user does not exist"})
		return
	}

	// check if bank exists
	// if exists := bank.CheckBankExists(db.DB, bankingaccount.BankID); !exists {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "bank does not exist"})
	// 	return
	// }

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new bankingaccount
	new_bankingaccount := common.BankingAccount{
		// BankID:       bankingaccount.BankID,
		UserID:       bankingaccount.UserID,
		Balance:      0,
		Type:         bankingaccount.Type,
		OpenningDate: time.Now(),
		// Transactions: bankingaccount.Transactions,
		CreatedBy: session.UserID,
	}

	rand.Seed(time.Now().UnixNano())
	card := "TN"
	for i := 0; i < 14; i++ {
		n := rand.Intn(10)
		card += fmt.Sprint(n)
	}

	new_bankingaccount.Iban = card

	// create new bankingaccount
	_, err := NewBankingAccount(db.DB, new_bankingaccount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all bankingaccounts from database
func (db Database) GetBankingAccounts(ctx *gin.Context) {

	// get bankingaccounts
	bankingaccounts, err := GetBankingAccounts(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bankingaccounts)
}

// get bankingaccount by id
func (db Database) GetBankingAccountByID(ctx *gin.Context) {
	//get id value from path
	bankingaccount_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := common.CheckBankingAccountExists(db.DB, uint(bankingaccount_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid banking account id"})
		return
	}
	bankingaccountId, err := GetBankingAccountByID(db.DB, uint(bankingaccount_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, bankingaccountId)
}

// Get banking accounts By user  ID
func (db Database) GetBankingAccountsByClientId(ctx *gin.Context) {

	// extract client ID from request parameters
	client_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid client ID"})
		return
	}

	// get banking accounts associated with the given cleint ID
	bankingAccounts, err := GetBankingAccountsByUserId(db.DB, uint(client_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// return the transactions as a JSON response
	ctx.JSON(http.StatusOK, bankingAccounts)
}

// search bankingaccounts from database
func (db Database) SearchBankingAccounts(ctx *gin.Context) {

	// init vars
	var bankingaccount common.BankingAccount

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&bankingaccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search bankingaccounts from database
	bankingaccounts, err := SearchBankingAccounts(db.DB, bankingaccount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bankingaccounts)
}

func (db Database) UpdateBankingAccount(ctx *gin.Context) {

	// init vars
	var bankingaccount common.BankingAccount
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&bankingaccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	bankingaccount_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(bankingaccount.Type) || empty_reg.MatchString(bankingaccount.Iban) || bankingaccount.Balance == 0 || bankingaccount.OpenningDate.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	bankingaccount.ID = uint(bankingaccount_id)

	// update bankingaccount
	if err = UpdateBankingAccount(db.DB, bankingaccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}
func UpdateClientBalance(db *gorm.DB, id uint, amount float64) error {

	// init vars
	user := &common.User{}

	// get client's first bank account
	if err := db.Where("id=?", id).First(&user).Error; err != nil {
		return err
	}

	var bankingAccounts []common.BankingAccount
	if err := db.Where("user_id = ?", id).Find(&bankingAccounts).Error; err != nil {
		return nil
	}

	var currentBankingAccount common.BankingAccount
	for _, bankingAccount := range bankingAccounts {

		if bankingAccount.Type == "current" {
			currentBankingAccount = bankingAccount
			break
		}
	}

	currentBankingAccount.Balance = currentBankingAccount.Balance + amount

	// update client's balance
	if err := db.Where("id=?", currentBankingAccount.ID).Updates(currentBankingAccount).Error; err != nil {
		return err
	}

	return nil
}
func (db Database) DeleteBankingAccount(ctx *gin.Context) {

	// get id from path
	bankingaccount_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete bankingaccount
	if err = DeleteBankingAccount(db.DB, uint(bankingaccount_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
