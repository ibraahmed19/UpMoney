package bank

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	common1 "template_rest_api/api/app/common"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new bank
func (db Database) NewBank(ctx *gin.Context) {

	// init vars
	var bank common1.Bank
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(bank.Name) || empty_reg.MatchString(bank.Address) || empty_reg.MatchString(bank.City) || empty_reg.MatchString(bank.Country) || empty_reg.MatchString(bank.Phone) || empty_reg.MatchString(bank.Email) || empty_reg.MatchString(bank.Contact) || empty_reg.MatchString(bank.Comments) || bank.ZipCode < 1 || bank.ManagedBy < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check user exists
	if exists := common1.CheckUserExists(db.DB, bank.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new bank
	new_bank := common1.Bank{
		Name:      bank.Name,
		Address:   bank.Address,
		City:      bank.City,
		Country:   bank.Country,
		ZipCode:   bank.ZipCode,
		Phone:     bank.Phone,
		Email:     bank.Email,
		Contact:   bank.Contact,
		Comments:  bank.Comments,
		ManagedBy: bank.ManagedBy,
		CreatedBy: session.UserID,
	}

	// create new bank
	_, err := NewBank(db.DB, new_bank)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all banks from database
func (db Database) GetBanks(ctx *gin.Context) {

	// get banks
	banks, err := GetBanks(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, banks)
}

// get bank by id
func (db Database) GetBankByID(ctx *gin.Context) {
	//get id value from path
	bank_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckBankExists(db.DB, uint(bank_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid bank id"})
		return
	}
	bankId, err := GetBankByID(db.DB, uint(bank_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, bankId)
}

// search banks from database
func (db Database) SearchBanks(ctx *gin.Context) {

	// init vars
	var bank common1.Bank

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search banks
	banks, err := SearchBanks(db.DB, bank)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, banks)
}

// update bank
func (db Database) UpdateBank(ctx *gin.Context) {

	// init vars
	var bank common1.Bank
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(bank.Name) || empty_reg.MatchString(bank.Address) || empty_reg.MatchString(bank.City) || empty_reg.MatchString(bank.Country) || empty_reg.MatchString(bank.Phone) || empty_reg.MatchString(bank.Email) || empty_reg.MatchString(bank.Contact) || empty_reg.MatchString(bank.Comments) || bank.ZipCode < 1 || bank.ManagedBy < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check subsidiary_of exists
	if exists := CheckBankExists(db.DB, bank.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid manager id"})
		return
	}

	// check user exists
	if exists := common1.CheckUserExists(db.DB, bank.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}

	if bank.ManagedBy != 0 {

		// check user exists
		if exists := common1.CheckUserExists(db.DB, bank.ManagedBy); !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
			return
		}
	}

	// get id value from path
	bank_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// init update bank
	update_bank := common1.Bank{
		ID:        uint(bank_id),
		Name:      bank.Name,
		Address:   bank.Address,
		City:      bank.City,
		Country:   bank.Country,
		ZipCode:   bank.ZipCode,
		Phone:     bank.Phone,
		Email:     bank.Email,
		Contact:   bank.Contact,
		Comments:  bank.Comments,
		ManagedBy: bank.ManagedBy,
	}

	// update bank
	if err = UpdateBank(db.DB, update_bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete bank from database
func (db Database) DeleteBank(ctx *gin.Context) {

	// get id value from path
	bank_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete bank
	if err := DeleteBank(db.DB, uint(bank_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
