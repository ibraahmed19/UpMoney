package insurancerequest

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"template_rest_api/api/app/insurance"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new insurancerequest
func (db Database) NewInsuranceRequest(ctx *gin.Context) {

	// init vars
	var insurancerequest InsuranceRequest
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&insurancerequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(insurancerequest.Type) || empty_reg.MatchString(insurancerequest.Justification) || empty_reg.MatchString(insurancerequest.Status) || empty_reg.MatchString(insurancerequest.Comments) || insurancerequest.Amount < 0 || insurancerequest.SubmissionDate.IsZero() || insurancerequest.ResponseDate.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check if insurance exists
	if exists := insurance.CheckInsuranceExists(db.DB, insurancerequest.IdInsurance); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "insurance does not exist"})
		return
	}

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new insurancerequest
	new_insurancerequest := InsuranceRequest{
		IdInsurance:    insurancerequest.IdInsurance,
		Type:           insurancerequest.Type,
		Amount:         insurancerequest.Amount,
		Justification:  insurancerequest.Justification,
		SubmissionDate: insurancerequest.SubmissionDate,
		ResponseDate:   insurancerequest.ResponseDate,
		Status:         insurancerequest.Status,
		Comments:       insurancerequest.Comments,
		CreatedBy:      session.UserID,
	}

	// create new insurancerequest
	_, err := NewInsuranceRequest(db.DB, new_insurancerequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all insurancerequests from database
func (db Database) GetInsuranceRequests(ctx *gin.Context) {

	// get insurancerequests
	insurancerequests, err := GetInsuranceRequests(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, insurancerequests)
}

// get insurancerequest by id
func (db Database) GetInsuranceRequestByID(ctx *gin.Context) {
	//get id value from path
	insurancerequest_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckInsuranceRequestExists(db.DB, uint(insurancerequest_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid insurance request id"})
		return
	}
	insurancerequestId, err := GetInsuranceRequestByID(db.DB, uint(insurancerequest_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, insurancerequestId)
}

// search insurancerequests from database
func (db Database) SearchInsuranceRequests(ctx *gin.Context) {

	// init vars
	var insurancerequest InsuranceRequest

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&insurancerequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search insurancerequests
	insurancerequests, err := SearchInsuranceRequests(db.DB, insurancerequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, insurancerequests)
}

// update insurancerequests
func (db Database) UpdateInsuranceRequest(ctx *gin.Context) {

	// init vars
	var insurancerequest InsuranceRequest
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&insurancerequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(insurancerequest.Type) || empty_reg.MatchString(insurancerequest.Justification) || empty_reg.MatchString(insurancerequest.Status) || empty_reg.MatchString(insurancerequest.Comments) || insurancerequest.Amount == 0 || insurancerequest.SubmissionDate.IsZero() || insurancerequest.ResponseDate.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// get id value from path
	insurancerequest_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// init update user
	update_insurancerequest := InsuranceRequest{
		ID:             uint(insurancerequest_id),
		IdInsurance:    insurancerequest.IdInsurance,
		Type:           insurancerequest.Type,
		Amount:         insurancerequest.Amount,
		Justification:  insurancerequest.Justification,
		SubmissionDate: insurancerequest.SubmissionDate,
		ResponseDate:   insurancerequest.ResponseDate,
		Status:         insurancerequest.Status,
		Comments:       insurancerequest.Comments,
	}

	// update insurancerequest
	if err = UpdateInsuranceRequest(db.DB, update_insurancerequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete insurancerequest from database
func (db Database) DeleteInsuranceRequest(ctx *gin.Context) {

	// get id value from path
	insurancerequest_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete insurancerequest
	if err := DeleteInsuranceRequest(db.DB, uint(insurancerequest_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
