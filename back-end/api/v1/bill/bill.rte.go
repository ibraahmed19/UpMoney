package bill

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BillRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("bills", "write", enforcer), baseInstance.NewBill)
	router.GET("/all", middleware.Authorize("bills", "read", enforcer), baseInstance.GetBills)
	router.GET("/:id", middleware.Authorize("bills", "read", enforcer), baseInstance.GetBillByID)
	router.POST("/search", middleware.Authorize("bills", "read", enforcer), baseInstance.SearchBills)
	router.PUT("/:id", middleware.Authorize("bills", "write", enforcer), baseInstance.UpdateBill)
	router.DELETE("/:id", middleware.Authorize("bills", "write", enforcer), baseInstance.DeleteBill)
}
