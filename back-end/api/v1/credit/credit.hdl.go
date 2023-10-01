package credit

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

// create new credit
func (db Database) NewCredit(ctx *gin.Context) {

	// init vars
	var credit common.Credit
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&credit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// // check fields
	if empty_reg.MatchString(credit.Type) || empty_reg.MatchString(credit.Status) || empty_reg.MatchString(credit.Comments) || credit.ClientID == 0 || credit.Amount <= 0 || credit.InterestRate <= 0 || credit.MonthlyPayments <= 0 || credit.EndDate.IsZero() || credit.StartDate.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check if client exists
	if exists := common.CheckClientExists(db.DB, credit.ClientID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "client does not exist"})
		return
	}

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new credit
	new_credit := common.Credit{
		ClientID:        credit.ClientID,
		Type:            credit.Type,
		Amount:          credit.Amount,
		InterestRate:    credit.InterestRate,
		Duration:        credit.Duration,
		StartDate:       credit.StartDate,
		EndDate:         credit.EndDate,
		MonthlyPayments: credit.MonthlyPayments,
		Status:          credit.Status,
		Comments:        credit.Comments,
		CreatedBy:       session.UserID,
	}

	// create new credit
	_, err := NewCredit(db.DB, new_credit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// update balance
	if err := common.UpdateClientBalance(db.DB, new_credit.ClientID, new_credit.Amount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all credits from database
func (db Database) GetCredits(ctx *gin.Context) {

	// get credits
	credits, err := GetCredits(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, credits)
}

// get credit by id
func (db Database) GetCreditByID(ctx *gin.Context) {
	//get id value from path
	credit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := common.CheckCreditExists(db.DB, uint(credit_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid bill id"})
		return
	}
	creditId, err := GetCreditByID(db.DB, uint(credit_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, creditId)
}

// search credits from database
func (db Database) SearchCredits(ctx *gin.Context) {

	// init vars
	var credit common.Credit

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&credit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search credits from database
	credits, err := SearchCredits(db.DB, credit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, credits)
}

func (db Database) UpdateCredit(ctx *gin.Context) {

	// init vars
	var credit common.Credit
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&credit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	credit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(credit.Type) || empty_reg.MatchString(credit.Status) || empty_reg.MatchString(credit.Comments) || credit.ClientID < 1 || credit.Amount <= 0 || credit.InterestRate <= 0 || credit.Duration <= 0 || credit.StartDate.IsZero() || credit.EndDate.IsZero() || credit.MonthlyPayments <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	credit.ID = uint(credit_id)

	// update credit
	if err = UpdateCredit(db.DB, credit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteCredit(ctx *gin.Context) {

	// get id from path
	credit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete credit
	if err = DeleteCredit(db.DB, uint(credit_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
