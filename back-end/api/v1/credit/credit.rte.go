package credit

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreditRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("credits", "write", enforcer), baseInstance.NewCredit)
	router.GET("/all", middleware.Authorize("credits", "read", enforcer), baseInstance.GetCredits)
	router.GET("/:id", middleware.Authorize("credits", "read", enforcer), baseInstance.GetCreditByID)
	router.POST("/search", middleware.Authorize("credits", "read", enforcer), baseInstance.SearchCredits)
	router.PUT("/:id", middleware.Authorize("credits", "write", enforcer), baseInstance.UpdateCredit)
	router.DELETE("/:id", middleware.Authorize("credits", "write", enforcer), baseInstance.DeleteCredit)
}
