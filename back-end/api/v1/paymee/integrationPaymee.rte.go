package paymee

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaymeeRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("payments", "write", enforcer), baseInstance.NewPayment)
	router.GET("/all", middleware.Authorize("payments", "read", enforcer), baseInstance.GetPayments)
	router.GET("/:id", middleware.Authorize("payments", "read", enforcer), baseInstance.GetPaymentByID)
	router.POST("/search", middleware.Authorize("payments", "read", enforcer), baseInstance.SearchPayments)

}
