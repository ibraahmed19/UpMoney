package bankingaccount

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BankingAccountRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("bankingaccounts", "write", enforcer), baseInstance.NewBankingAccount)
	router.GET("/all", middleware.Authorize("bankingaccounts", "read", enforcer), baseInstance.GetBankingAccounts)
	router.GET("/:id", middleware.Authorize("bankingaccounts", "read", enforcer), baseInstance.GetBankingAccountByID)
	router.GET("/all/:id", middleware.Authorize("bankingaccounts", "read", enforcer), baseInstance.GetBankingAccountsByClientId)
	router.POST("/search", middleware.Authorize("bankingaccounts", "read", enforcer), baseInstance.SearchBankingAccounts)
	router.PUT("/:id", middleware.Authorize("bankingaccounts", "write", enforcer), baseInstance.UpdateBankingAccount)
	router.DELETE("/:id", middleware.Authorize("bankingaccounts", "write", enforcer), baseInstance.DeleteBankingAccount)
}
