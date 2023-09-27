package bill

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

// create new bill
func (db Database) NewBill(ctx *gin.Context) {

	// init vars
	var bill Bill
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&bill); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//check fields
	if empty_reg.MatchString(bill.BillNumber) || empty_reg.MatchString(bill.ClientName) || bill.IbanAccount == 0 || bill.Amount == 0 || bill.IssueDate.IsZero() || bill.DueDate.IsZero() || bill.CreationDate.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}
	// check if transaction exists
	if exists := transaction.CheckTransactionExists(db.DB, bill.IdTransaction); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "transaction does not exist"})
		return
	}
	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new bill
	new_bill := Bill{
		ID:            0,
		IdTransaction: bill.IdTransaction,
		BillNumber:    bill.BillNumber,
		ClientName:    bill.ClientName,
		IbanAccount:   bill.IbanAccount,
		Amount:        bill.Amount,
		IssueDate:     bill.IssueDate,
		DueDate:       bill.DueDate,
		CreationDate:  bill.CreationDate,
		CreatedBy:     session.UserID,
	}

	// create new transaction
	_, err := NewBill(db.DB, new_bill)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all bills from database
func (db Database) GetBills(ctx *gin.Context) {

	// get bills
	bills, err := GetBills(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bills)
}

// get bill by id
func (db Database) GetBillByID(ctx *gin.Context) {
	//get id value from path
	bill_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckBillExists(db.DB, uint(bill_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid bill id"})
		return
	}
	billId, err := GetBillByID(db.DB, uint(bill_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, billId)
}

// search bills	 from database
func (db Database) SearchBills(ctx *gin.Context) {

	// init vars
	var bill Bill

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&bill); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search bills from database
	bills, err := SearchBills(db.DB, bill)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bills)
}

func (db Database) UpdateBill(ctx *gin.Context) {

	// init vars
	var bill Bill
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&bill); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	bill_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(bill.BillNumber) || empty_reg.MatchString(bill.ClientName) || bill.IbanAccount == 0 || bill.Amount == 0 || bill.IssueDate.IsZero() || bill.DueDate.IsZero() || bill.CreationDate.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	bill.ID = uint(bill_id)

	// update bill
	if err = UpdateBill(db.DB, bill); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteBill(ctx *gin.Context) {

	// get id from path
	bill_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete bill
	if err = DeleteBill(db.DB, uint(bill_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
