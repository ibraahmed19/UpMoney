package insurancepolicy

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"template_rest_api/api/app/insurance"
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

// create new insurancepolicy
func (db Database) NewInsurancePolicy(ctx *gin.Context) {

	// init vars
	var insurancepolicy InsurancePolicy
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&insurancepolicy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(insurancepolicy.Type) || empty_reg.MatchString(insurancepolicy.WarrantyDetails) || empty_reg.MatchString(insurancepolicy.Franchise) || empty_reg.MatchString(insurancepolicy.PolicyStatus) || insurancepolicy.Amount == 0 || insurancepolicy.CoverageLimit == 0 || insurancepolicy.StartDate.IsZero() || insurancepolicy.EndTime.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check client exists
	if exists := common.CheckClientExists(db.DB, insurancepolicy.ClientId); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "client does not exist"})
		return
	}

	// check insurance exists
	if exists := insurance.CheckInsuranceExists(db.DB, insurancepolicy.InsuranceId); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "insurance does not exist"})
		return
	}
	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new insurancepolicy
	new_insurancepolicy := InsurancePolicy{
		ClientId:        insurancepolicy.ClientId,
		InsuranceId:     insurancepolicy.InsuranceId,
		Type:            insurancepolicy.Type,
		Amount:          insurancepolicy.Amount,
		StartDate:       insurancepolicy.StartDate,
		EndTime:         insurancepolicy.EndTime,
		WarrantyDetails: insurancepolicy.WarrantyDetails,
		Franchise:       insurancepolicy.Franchise,
		CoverageLimit:   insurancepolicy.CoverageLimit,
		PolicyStatus:    insurancepolicy.PolicyStatus,
		CreatedBy:       session.UserID,
	}

	// create new insurancepolicy
	_, err := NewInsurancePolicy(db.DB, new_insurancepolicy)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all insurancepolicies from database
func (db Database) GetInsurancePolicies(ctx *gin.Context) {

	// get insurancepolicies
	insurancepolicies, err := GetInsurancePolicies(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, insurancepolicies)
}

// get insurancepolicy by id
func (db Database) GetInsurancePolicyByID(ctx *gin.Context) {
	//get id value from path
	insurancepolicy_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckInsurancePolicyExists(db.DB, uint(insurancepolicy_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid insurance policy id"})
		return
	}
	insurancepolicyId, err := GetInsurancePolicyByID(db.DB, uint(insurancepolicy_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, insurancepolicyId)
}

// search insurancepolicies from database
func (db Database) SearchInsurancePolicies(ctx *gin.Context) {

	// init vars
	var insurancepolicy InsurancePolicy

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&insurancepolicy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search insurancepolicies
	insurancepolicies, err := SearchInsurancePolicies(db.DB, insurancepolicy)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, insurancepolicies)
}

// update insurancepolicies
func (db Database) UpdateInsurancePolicy(ctx *gin.Context) {

	// init vars
	var insurancepolicy InsurancePolicy
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&insurancepolicy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(insurancepolicy.Type) || empty_reg.MatchString(insurancepolicy.WarrantyDetails) || empty_reg.MatchString(insurancepolicy.Franchise) || empty_reg.MatchString(insurancepolicy.PolicyStatus) || insurancepolicy.Amount == 0 || insurancepolicy.CoverageLimit == 0 || insurancepolicy.StartDate.IsZero() || insurancepolicy.EndTime.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// get id value from path
	insurancepolicy_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// init update user
	update_insurancepolicy := InsurancePolicy{
		ID:              uint(insurancepolicy_id),
		ClientId:        insurancepolicy.ClientId,
		InsuranceId:     insurancepolicy.InsuranceId,
		Type:            insurancepolicy.Type,
		Amount:          insurancepolicy.Amount,
		StartDate:       insurancepolicy.StartDate,
		EndTime:         insurancepolicy.EndTime,
		WarrantyDetails: insurancepolicy.WarrantyDetails,
		Franchise:       insurancepolicy.Franchise,
		CoverageLimit:   insurancepolicy.CoverageLimit,
		PolicyStatus:    insurancepolicy.PolicyStatus,
	}

	// update insurancepolicy
	if err = UpdateInsurancePolicy(db.DB, update_insurancepolicy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete insurancepolicy from database
func (db Database) DeleteInsurancePolicy(ctx *gin.Context) {

	// get id value from path
	insurancepolicy_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete insurancepolicy
	if err := DeleteInsurancePolicy(db.DB, uint(insurancepolicy_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
