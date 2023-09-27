package insurancerequest

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare insurancerequest routes
func RoutesInsuranceRequests(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("insurancerequests", "write", enforcer), baseInstance.NewInsuranceRequest)
	router.GET("/all", middleware.Authorize("insurancerequests", "read", enforcer), baseInstance.GetInsuranceRequests)
	router.GET("/:id", middleware.Authorize("insurancerequests", "read", enforcer), baseInstance.GetInsuranceRequestByID)
	router.POST("/search", middleware.Authorize("insurancerequests", "read", enforcer), baseInstance.SearchInsuranceRequests)
	router.PUT("/:id", middleware.Authorize("insurancerequests", "write", enforcer), baseInstance.UpdateInsuranceRequest)
	router.DELETE("/:id", middleware.Authorize("insurancerequests", "write", enforcer), baseInstance.DeleteInsuranceRequest)
}
