package simulation

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SimulationRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("simulations", "write", enforcer), baseInstance.NewSimulation)
	router.GET("/all", middleware.Authorize("simulations", "read", enforcer), baseInstance.GetSimulations)
	router.GET("/:id", middleware.Authorize("simulations", "read", enforcer), baseInstance.GetSimulationByID)
	router.POST("/search", middleware.Authorize("simulations", "read", enforcer), baseInstance.SearchSimulations)
	router.PUT("/:id", middleware.Authorize("simulations", "write", enforcer), baseInstance.UpdateSimulation)
	router.DELETE("/:id", middleware.Authorize("simulations", "write", enforcer), baseInstance.DeleteSimulation)
}
