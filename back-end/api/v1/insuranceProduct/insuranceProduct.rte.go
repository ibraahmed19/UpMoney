package insuranceproduct

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare insuranceproduct routes
func RoutesInsuranceProducts(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("insuranceproducts", "write", enforcer), baseInstance.NewInsuranceProduct)
	router.GET("/all", middleware.Authorize("insuranceproducts", "read", enforcer), baseInstance.GetInsuranceProducts)
	router.GET("/:id", middleware.Authorize("insuranceproducts", "read", enforcer), baseInstance.GetInsuranceProductByID)
	router.POST("/search", middleware.Authorize("insuranceproducts", "read", enforcer), baseInstance.SearchInsuranceProducts)
	router.PUT("/:id", middleware.Authorize("insuranceproducts", "write", enforcer), baseInstance.UpdateInsuranceProduct)
	router.DELETE("/:id", middleware.Authorize("insuranceproducts", "write", enforcer), baseInstance.DeleteInsuranceProduct)
}
