package deposit

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"template_rest_api/api/v1/transaction"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new deposit
func (db Database) NewDeposit(ctx *gin.Context) {

	// init vars
	var deposit Deposit
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&deposit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(deposit.Description) || deposit.IbanAccount == 0 || deposit.Amount == 0 || deposit.Date.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}
	// check if transaction exists
	if exists := transaction.CheckTransactionExists(db.DB, deposit.IdTransaction); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "transaction does not exist"})
		return
	}
	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new deposit
	new_deposit := Deposit{
		IdTransaction: deposit.IdTransaction,
		IbanAccount:   deposit.IbanAccount,
		Amount:        deposit.Amount,
		Date:          deposit.Date,
		Description:   deposit.Description,
		CreatedBy:     session.UserID,
	}

	// create new deposit
	_, err := NewDeposit(db.DB, new_deposit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all deposits from database
func (db Database) GetDeposits(ctx *gin.Context) {

	// get deposits
	deposits, err := GetDeposits(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, deposits)
}

// get deposit by id
func (db Database) GetDepositByID(ctx *gin.Context) {
	//get id value from path
	deposit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckDepositExists(db.DB, uint(deposit_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid deposit id"})
		return
	}
	depositId, err := GetDepositByID(db.DB, uint(deposit_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, depositId)
}

// search deposits	 from database
func (db Database) SearchDeposits(ctx *gin.Context) {

	// init vars
	var deposit Deposit

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&deposit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search deposits from database
	deposits, err := SearchDeposits(db.DB, deposit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, deposits)
}

func (db Database) UpdateDeposit(ctx *gin.Context) {

	// init vars
	var deposit Deposit
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&deposit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	deposit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(deposit.Description) || deposit.IbanAccount == 0 || deposit.Amount == 0 || deposit.Date.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	deposit.ID = uint(deposit_id)

	// update deposit
	if err = UpdateDeposit(db.DB, deposit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteDeposit(ctx *gin.Context) {

	// get id from path
	deposit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete deposit
	if err = DeleteDeposit(db.DB, uint(deposit_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
