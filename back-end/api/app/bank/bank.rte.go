package bank

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare bank routes
func RoutesBanks(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("banks", "write", enforcer), baseInstance.NewBank)
	router.GET("/all", middleware.Authorize("banks", "read", enforcer), baseInstance.GetBanks)
	router.GET("/:id", middleware.Authorize("banks", "read", enforcer), baseInstance.GetBankByID)
	router.POST("/search", middleware.Authorize("banks", "read", enforcer), baseInstance.SearchBanks)
	router.PUT("/:id", middleware.Authorize("banks", "write", enforcer), baseInstance.UpdateBank)
	router.DELETE("/:id", middleware.Authorize("banks", "write", enforcer), baseInstance.DeleteBank)
}
