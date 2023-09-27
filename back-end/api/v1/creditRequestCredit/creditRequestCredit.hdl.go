package creditrequestcredit

import (
	"net/http"
	"strconv"
	"template_rest_api/api/v1/common"
	creditrequest "template_rest_api/api/v1/creditRequest"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new creditrequestcredit
func (db Database) NewCreditRequestCredit(ctx *gin.Context) {

	// init vars
	var creditrequestcredit CreditRequestCredit
	// check json validity
	if err := ctx.ShouldBindJSON(&creditrequestcredit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if creditrequestcredit.CreditId < 0 || creditrequestcredit.CreditRequestId < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check creditrequest exists
	if exists := creditrequest.CheckCreditRequestExists(db.DB, creditrequestcredit.CreditRequestId); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "credit request does not exist"})
		return
	}

	// check credit exists
	if exists := common.CheckCreditExists(db.DB, creditrequestcredit.CreditId); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "credit does not exist"})
		return
	}
	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new creditrequestcredit
	new_creditrequestcredit := CreditRequestCredit{
		CreditId:        creditrequestcredit.CreditId,
		CreditRequestId: creditrequestcredit.CreditRequestId,
		CreatedBy:       session.UserID,
	}

	// create new creditrequestcredit
	_, err := NewCreditRequestCredit(db.DB, new_creditrequestcredit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all creditrequestscredits from database
func (db Database) GetCreditRequestsCredits(ctx *gin.Context) {

	// get creditrequestscredits
	creditrequestscredits, err := GetCreditRequestsCredits(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, creditrequestscredits)
}

// get creditrequestcredit by id
func (db Database) GetCreditRequestCreditByID(ctx *gin.Context) {
	//get id value from path
	creditrequestcredit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckCreditRequestCreditExists(db.DB, uint(creditrequestcredit_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid creditrequestcredit id"})
		return
	}
	creditrequestcreditId, err := GetCreditRequestCreditByID(db.DB, uint(creditrequestcredit_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, creditrequestcreditId)
}

// search creditrequestscredits from database
func (db Database) SearchCreditRequestsCredits(ctx *gin.Context) {

	// init vars
	var creditrequestcredit CreditRequestCredit

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&creditrequestcredit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search creditrequestscredits
	creditrequestscredits, err := SearchCreditRequestsCredits(db.DB, creditrequestcredit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, creditrequestscredits)
}

// update creditrequestscredits
func (db Database) UpdateCreditRequestCredit(ctx *gin.Context) {

	// init vars
	var creditrequestcredit CreditRequestCredit

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&creditrequestcredit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if creditrequestcredit.CreditId < 0 || creditrequestcredit.CreditRequestId < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// get id value from path
	creditrequestcredit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// init update creditrequestcredit
	update_creditrequestcredit := CreditRequestCredit{
		ID:              uint(creditrequestcredit_id),
		CreditId:        creditrequestcredit.CreditId,
		CreditRequestId: creditrequestcredit.CreditRequestId,
	}

	// update creditrequestcredit
	if err = UpdateCreditRequestCredit(db.DB, update_creditrequestcredit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete creditrequestcredit from database
func (db Database) DeleteCreditRequestCredit(ctx *gin.Context) {

	// get id value from path
	creditrequestcredit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete creditrequestcredit
	if err := DeleteCreditRequestCredit(db.DB, uint(creditrequestcredit_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
