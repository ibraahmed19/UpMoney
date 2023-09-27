package insurance

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	"template_rest_api/api/app/common"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new insurance
func (db Database) NewInsurance(ctx *gin.Context) {

	// init vars
	var insurance Insurance
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&insurance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(insurance.Name) || empty_reg.MatchString(insurance.Address) || empty_reg.MatchString(insurance.Type) || empty_reg.MatchString(insurance.Phone) || empty_reg.MatchString(insurance.Email) || empty_reg.MatchString(insurance.Description) || insurance.ManagedBy < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check user exists
	if exists := common1.CheckUserExists(db.DB, insurance.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new insurance
	new_insurance := Insurance{
		Name:        insurance.Name,
		Type:        insurance.Type,
		Address:     insurance.Address,
		Phone:       insurance.Phone,
		Email:       insurance.Email,
		Description: insurance.Description,
		ManagedBy:   insurance.ManagedBy,
		CreatedBy:   session.UserID,
	}

	// create new insurance
	_, err := NewInsurance(db.DB, new_insurance)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all insurances from database
func (db Database) GetInsurances(ctx *gin.Context) {

	// get insurances
	insurances, err := GetInsurances(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, insurances)
}

// get insurance by id
func (db Database) GetInsuranceByID(ctx *gin.Context) {
	//get id value from path
	insurance_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckInsuranceExists(db.DB, uint(insurance_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid insurance id"})
		return
	}
	insuranceId, err := GetInsuranceByID(db.DB, uint(insurance_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, insuranceId)
}

// search insurances from database
func (db Database) SearchInsurances(ctx *gin.Context) {

	// init vars
	var insurance Insurance

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&insurance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search insurances
	insurances, err := SearchInsurances(db.DB, insurance)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, insurances)
}

// update insurance
func (db Database) UpdateInsurance(ctx *gin.Context) {

	// init vars
	var insurance Insurance
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&insurance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(insurance.Name) || empty_reg.MatchString(insurance.Address) || empty_reg.MatchString(insurance.Type) || empty_reg.MatchString(insurance.Phone) || empty_reg.MatchString(insurance.Email) || empty_reg.MatchString(insurance.Description) || insurance.ManagedBy < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check subsidiary_of exists
	if exists := CheckInsuranceExists(db.DB, insurance.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid manager id"})
		return
	}

	// check user exists
	if exists := common1.CheckUserExists(db.DB, insurance.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}

	if insurance.ManagedBy != 0 {

		// check user exists
		if exists := common1.CheckUserExists(db.DB, insurance.ManagedBy); !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
			return
		}
	}

	// get id value from path
	insurance_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// init update user
	update_insurance := Insurance{
		ID:          uint(insurance_id),
		Name:        insurance.Name,
		Type:        insurance.Type,
		Address:     insurance.Address,
		Phone:       insurance.Phone,
		Email:       insurance.Email,
		Description: insurance.Description,
		ManagedBy:   insurance.ManagedBy,
	}

	// update insurance
	if err = UpdateInsurance(db.DB, update_insurance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete insurance from database
func (db Database) DeleteInsurance(ctx *gin.Context) {

	// get id value from path
	insurance_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete insurance
	if err := DeleteInsurance(db.DB, uint(insurance_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
