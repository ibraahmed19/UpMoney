package simulationcredit

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare simulationcredit routes
func RoutesSimulationsCredits(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("simulationscredits", "write", enforcer), baseInstance.NewSimulationCredit)
	router.GET("/all", middleware.Authorize("simulationscredits", "read", enforcer), baseInstance.GetSimulationsCredits)
	router.GET("/:id", middleware.Authorize("simulationscredits", "read", enforcer), baseInstance.GetSimulationCreditByID)
	router.POST("/search", middleware.Authorize("simulationscredits", "read", enforcer), baseInstance.SearchSimulationsCredits)
	router.PUT("/:id", middleware.Authorize("simulationscredits", "write", enforcer), baseInstance.UpdateSimulationCredit)
	router.DELETE("/:id", middleware.Authorize("simulationscredits", "write", enforcer), baseInstance.DeleteSimulationCredit)
}
