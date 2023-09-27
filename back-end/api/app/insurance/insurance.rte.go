package insurance

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare insurance routes
func RoutesInsurances(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("insurances", "write", enforcer), baseInstance.NewInsurance)
	router.GET("/all", middleware.Authorize("insurances", "read", enforcer), baseInstance.GetInsurances)
	router.GET("/:id", middleware.Authorize("insurances", "read", enforcer), baseInstance.GetInsuranceByID)
	router.POST("/search", middleware.Authorize("insurances", "read", enforcer), baseInstance.SearchInsurances)
	router.PUT("/:id", middleware.Authorize("insurances", "write", enforcer), baseInstance.UpdateInsurance)
	router.DELETE("/:id", middleware.Authorize("insurances", "write", enforcer), baseInstance.DeleteInsurance)
}
