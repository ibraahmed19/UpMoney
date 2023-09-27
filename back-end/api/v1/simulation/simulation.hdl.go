package simulation

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

// create new simulation
func (db Database) NewSimulation(ctx *gin.Context) {

	// init vars
	var simulation Simulation
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check json validity
	if err := ctx.ShouldBindJSON(&simulation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(simulation.Type) || empty_reg.MatchString(simulation.Results) || empty_reg.MatchString(simulation.VariablesEntry) || empty_reg.MatchString(simulation.VariablesOut) || simulation.Date.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new simulation
	new_simulation := Simulation{
		ID:             0,
		Type:           simulation.Type,
		Date:           simulation.Date,
		Results:        simulation.Results,
		VariablesEntry: simulation.VariablesEntry,
		VariablesOut:   simulation.VariablesOut,
		CreatedBy:      session.UserID,
	}

	// create new simulation
	_, err := NewSimulation(db.DB, new_simulation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all simulations from database
func (db Database) GetSimulations(ctx *gin.Context) {

	// get simulations
	simulations, err := GetSimulations(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, simulations)
}

// get simulation by id
func (db Database) GetSimulationByID(ctx *gin.Context) {
	//get id value from path
	simulation_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckSimulationExists(db.DB, uint(simulation_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid simulation id"})
		return
	}
	simulationId, err := GetSimulationByID(db.DB, uint(simulation_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, simulationId)
}

// search simulations from database
func (db Database) SearchSimulations(ctx *gin.Context) {

	// init vars
	var simulation Simulation

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&simulation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search simulations from database
	simulations, err := SearchSimulations(db.DB, simulation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, simulations)
}

func (db Database) UpdateSimulation(ctx *gin.Context) {

	// init vars
	var simulation Simulation
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&simulation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	simulation_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(simulation.Type) || empty_reg.MatchString(simulation.Results) || empty_reg.MatchString(simulation.VariablesEntry) || empty_reg.MatchString(simulation.VariablesOut) || simulation.Date.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	simulation.ID = uint(simulation_id)

	// update simulation
	if err = UpdateSimulation(db.DB, simulation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteSimulation(ctx *gin.Context) {

	// get id from path
	simulation_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete simulation
	if err = DeleteSimulation(db.DB, uint(simulation_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
