package deposit

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DepositRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("deposits", "write", enforcer), baseInstance.NewDeposit)
	router.GET("/all", middleware.Authorize("deposits", "read", enforcer), baseInstance.GetDeposits)
	router.GET("/:id", middleware.Authorize("deposits", "read", enforcer), baseInstance.GetDepositByID)
	router.POST("/search", middleware.Authorize("deposits", "read", enforcer), baseInstance.SearchDeposits)
	router.PUT("/:id", middleware.Authorize("deposits", "write", enforcer), baseInstance.UpdateDeposit)
	router.DELETE("/:id", middleware.Authorize("deposits", "write", enforcer), baseInstance.DeleteDeposit)
}
