package insuranceproduct

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new insuranceproduct
func (db Database) NewInsuranceProduct(ctx *gin.Context) {

	// init vars
	var insuranceproduct InsuranceProduct
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&insuranceproduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(insuranceproduct.Name) || empty_reg.MatchString(insuranceproduct.Type) || empty_reg.MatchString(insuranceproduct.Description) || insuranceproduct.Amount == 0 || insuranceproduct.CoverageLimit == 0 || empty_reg.MatchString(insuranceproduct.Duration) || empty_reg.MatchString(insuranceproduct.PaymentPeriod) || empty_reg.MatchString(insuranceproduct.AdditionalOptions) || empty_reg.MatchString(insuranceproduct.Status) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// // check if insurance exists
	// if exists := insurance.CheckInsuranceExists(db.DB, insuranceproduct.IdInsurance); !exists {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "insurance does not exist"})
	// 	return
	// }

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new insuranceproduct
	new_insuranceproduct := InsuranceProduct{
		// IdInsurance:       insuranceproduct.IdInsurance,
		Name:              insuranceproduct.Name,
		Type:              insuranceproduct.Type,
		Description:       insuranceproduct.Description,
		Amount:            insuranceproduct.Amount,
		CoverageLimit:     insuranceproduct.CoverageLimit,
		Duration:          insuranceproduct.Duration,
		PaymentPeriod:     insuranceproduct.PaymentPeriod,
		AdditionalOptions: insuranceproduct.AdditionalOptions,
		Status:            insuranceproduct.Status,
		CreatedBy:         session.UserID,
	}

	// create new insuranceproduct
	_, err := NewInsuranceProduct(db.DB, new_insuranceproduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all insuranceproducts from database
func (db Database) GetInsuranceProducts(ctx *gin.Context) {

	// get insuranceproducts
	insuranceproducts, err := GetInsuranceProducts(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, insuranceproducts)
}

// get insuranceproduct by id
func (db Database) GetInsuranceProductByID(ctx *gin.Context) {
	//get id value from path
	insuranceproduct_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckInsuranceProductExists(db.DB, uint(insuranceproduct_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid insurance product id"})
		return
	}
	insuranceproductId, err := GetInsuranceProductByID(db.DB, uint(insuranceproduct_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, insuranceproductId)
}

// search insuranceproducts from database
func (db Database) SearchInsuranceProducts(ctx *gin.Context) {

	// init vars
	var insuranceproduct InsuranceProduct

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&insuranceproduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search insuranceproducts
	insuranceproducts, err := SearchInsuranceProducts(db.DB, insuranceproduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, insuranceproducts)
}

// update insuranceproducts
func (db Database) UpdateInsuranceProduct(ctx *gin.Context) {

	// init vars
	var insuranceproduct InsuranceProduct
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&insuranceproduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(insuranceproduct.Name) || empty_reg.MatchString(insuranceproduct.Type) || empty_reg.MatchString(insuranceproduct.Description) || insuranceproduct.Amount == 0 || insuranceproduct.CoverageLimit == 0 || empty_reg.MatchString(insuranceproduct.Duration) || empty_reg.MatchString(insuranceproduct.PaymentPeriod) || empty_reg.MatchString(insuranceproduct.AdditionalOptions) || empty_reg.MatchString(insuranceproduct.Status) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// get id value from path
	insuranceproduct_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// init update user
	update_insuranceproduct := InsuranceProduct{
		ID: uint(insuranceproduct_id),
		// IdInsurance:       insuranceproduct.IdInsurance,
		Name:              insuranceproduct.Name,
		Type:              insuranceproduct.Type,
		Description:       insuranceproduct.Description,
		Amount:            insuranceproduct.Amount,
		CoverageLimit:     insuranceproduct.CoverageLimit,
		Duration:          insuranceproduct.Duration,
		PaymentPeriod:     insuranceproduct.PaymentPeriod,
		AdditionalOptions: insuranceproduct.AdditionalOptions,
		Status:            insuranceproduct.Status,
	}

	// update insuranceproduct
	if err = UpdateInsuranceProduct(db.DB, update_insuranceproduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete insuranceproduct from database
func (db Database) DeleteInsuranceProduct(ctx *gin.Context) {

	// get id value from path
	insuranceproduct_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete insuranceproduct
	if err := DeleteInsuranceProduct(db.DB, uint(insuranceproduct_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
