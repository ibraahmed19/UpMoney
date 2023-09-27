package creditrequest

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

// create new creditrequest
func (db Database) NewCreditRequest(ctx *gin.Context) {

	// init vars
	var creditrequest CreditRequest
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&creditrequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(creditrequest.Type) || empty_reg.MatchString(creditrequest.Justification) || empty_reg.MatchString(creditrequest.RequestStatus) || empty_reg.MatchString(creditrequest.Notes) || creditrequest.ClientId == 0 || creditrequest.AmountRequested == 0 || creditrequest.InterestRate == 0 || creditrequest.Duration == 0 || creditrequest.SubmissionDate.IsZero() || creditrequest.ApprovalDate.IsZero() || creditrequest.RejectionDate.IsZero() || creditrequest.ApprovedAmount == 0 || creditrequest.MonthlyPayments == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}
	// check if client exists
	if exists := common.CheckClientExists(db.DB, creditrequest.ClientId); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "client does not exist"})
		return
	}
	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new creditrequest
	new_creditrequest := CreditRequest{
		ClientId:        creditrequest.ClientId,
		Type:            creditrequest.Type,
		AmountRequested: creditrequest.AmountRequested,
		Justification:   creditrequest.Justification,
		RequestStatus:   creditrequest.RequestStatus,
		SubmissionDate:  creditrequest.SubmissionDate,
		ApprovalDate:    creditrequest.ApprovalDate,
		RejectionDate:   creditrequest.RejectionDate,
		ApprovedAmount:  creditrequest.ApprovedAmount,
		InterestRate:    creditrequest.InterestRate,
		Duration:        creditrequest.Duration,
		MonthlyPayments: creditrequest.MonthlyPayments,
		Notes:           creditrequest.Notes,
		CreatedBy:       session.UserID,
	}

	// create new creditrequest
	_, err := NewCreditRequest(db.DB, new_creditrequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all creditrequests from database
func (db Database) GetCreditRequests(ctx *gin.Context) {

	// get creditrequests
	creditrequests, err := GetCreditRequests(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, creditrequests)
}

// get creditrequest by id
func (db Database) GetCreditRequestByID(ctx *gin.Context) {
	//get id value from path
	creditrequest_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckCreditRequestExists(db.DB, uint(creditrequest_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid credit request id"})
		return
	}
	creditrequestId, err := GetCreditRequestByID(db.DB, uint(creditrequest_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, creditrequestId)
}

// search creditrequests from database
func (db Database) SearchCreditRequests(ctx *gin.Context) {

	// init vars
	var creditrequest CreditRequest

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&creditrequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search creditrequests from database
	creditrequests, err := SearchCreditRequests(db.DB, creditrequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, creditrequests)
}

func (db Database) UpdateCreditRequest(ctx *gin.Context) {

	// init vars
	var creditrequest CreditRequest
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&creditrequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	creditrequest_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(creditrequest.Type) || empty_reg.MatchString(creditrequest.Justification) || empty_reg.MatchString(creditrequest.RequestStatus) || empty_reg.MatchString(creditrequest.Notes) || creditrequest.ClientId == 0 || creditrequest.AmountRequested == 0 || creditrequest.InterestRate == 0 || creditrequest.Duration == 0 || creditrequest.SubmissionDate.IsZero() || creditrequest.ApprovalDate.IsZero() || creditrequest.RejectionDate.IsZero() || creditrequest.ApprovedAmount == 0 || creditrequest.MonthlyPayments == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	creditrequest.ID = uint(creditrequest_id)

	// update credit
	if err = UpdateCreditRequest(db.DB, creditrequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteCreditRequest(ctx *gin.Context) {

	// get id from path
	creditrequest_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete creditrequest
	if err = DeleteCreditRequest(db.DB, uint(creditrequest_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
