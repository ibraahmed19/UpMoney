package insurancepolicy

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare insurancepolicy routes
func RoutesInsurancePolicies(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("insurancepolicies", "write", enforcer), baseInstance.NewInsurancePolicy)
	router.GET("/all", middleware.Authorize("insurancepolicies", "read", enforcer), baseInstance.GetInsurancePolicies)
	router.GET("/:id", middleware.Authorize("insurancepolicies", "read", enforcer), baseInstance.GetInsurancePolicyByID)
	router.POST("/search", middleware.Authorize("insurancepolicies", "read", enforcer), baseInstance.SearchInsurancePolicies)
	router.PUT("/:id", middleware.Authorize("insurancepolicies", "write", enforcer), baseInstance.UpdateInsurancePolicy)
	router.DELETE("/:id", middleware.Authorize("insurancepolicies", "write", enforcer), baseInstance.DeleteInsurancePolicy)
}
