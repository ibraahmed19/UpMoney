package transaction

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("transactions", "write", enforcer), baseInstance.NewTransaction)
	router.GET("/all", middleware.Authorize("transactions", "read", enforcer), baseInstance.GetTransactions)
	router.GET("/all/:id", middleware.Authorize("transactions", "read", enforcer), baseInstance.GetTransactionsByBankingAccountId)
	router.GET("/:id", middleware.Authorize("transactions", "read", enforcer), baseInstance.GetTransactionByID)
	router.POST("/search", middleware.Authorize("transactions", "read", enforcer), baseInstance.SearchTransactions)
	router.PUT("/:id", middleware.Authorize("transactions", "write", enforcer), baseInstance.UpdateTransaction)
	router.DELETE("/:id", middleware.Authorize("transactions", "write", enforcer), baseInstance.DeleteTransaction)
}
