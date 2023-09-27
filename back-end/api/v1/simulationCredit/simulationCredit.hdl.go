package simulationcredit

import (
	"net/http"
	"strconv"
	"template_rest_api/api/v1/common"
	"template_rest_api/api/v1/simulation"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new simulationcredit
func (db Database) NewSimulationCredit(ctx *gin.Context) {

	// init vars
	var simulationcredit SimulationCredit
	// check json validity
	if err := ctx.ShouldBindJSON(&simulationcredit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if simulationcredit.CreditId < 0 || simulationcredit.SimulationId < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check credit exists
	if exists := common.CheckCreditExists(db.DB, simulationcredit.CreditId); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "credit does not exist"})
		return
	}

	// check simulation exists
	if exists := simulation.CheckSimulationExists(db.DB, simulationcredit.SimulationId); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "simulation does not exist"})
		return
	}
	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new simulationcredit
	new_simulationcredit := SimulationCredit{
		CreditId:     simulationcredit.CreditId,
		SimulationId: simulationcredit.SimulationId,
		CreatedBy:    session.UserID,
	}

	// create new simulationcredit
	_, err := NewSimulationCredit(db.DB, new_simulationcredit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all simulationscredits from database
func (db Database) GetSimulationsCredits(ctx *gin.Context) {

	// get simulationscredits
	simulationscredits, err := GetSimulationsCredits(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, simulationscredits)
}

// get simulationcredit by id
func (db Database) GetSimulationCreditByID(ctx *gin.Context) {
	//get id value from path
	simulationcredit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckSimulationCreditExists(db.DB, uint(simulationcredit_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid simulation credit id"})
		return
	}
	simulationcreditId, err := GetSimulationCreditByID(db.DB, uint(simulationcredit_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, simulationcreditId)
}

// search simulationscredits from database
func (db Database) SearchSimulationsCredits(ctx *gin.Context) {

	// init vars
	var simulationcredit SimulationCredit

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&simulationcredit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search simulationscredits
	simulationscredits, err := SearchSimulationsCredits(db.DB, simulationcredit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, simulationscredits)
}

// update simulationscredits
func (db Database) UpdateSimulationCredit(ctx *gin.Context) {

	// init vars
	var simulationcredit SimulationCredit

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&simulationcredit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if simulationcredit.CreditId < 0 || simulationcredit.SimulationId < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// get id value from path
	simulationcredit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// init update simulationcredit
	update_simulationcredit := SimulationCredit{
		ID:           uint(simulationcredit_id),
		CreditId:     simulationcredit.CreditId,
		SimulationId: simulationcredit.SimulationId,
	}

	// update simulationcredit
	if err = UpdateSimulationCredit(db.DB, update_simulationcredit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete simulationcredit from database
func (db Database) DeleteSimulationCredit(ctx *gin.Context) {

	// get id value from path
	simulationcredit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete simulationcredit
	if err := DeleteSimulationCredit(db.DB, uint(simulationcredit_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
