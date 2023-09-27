package creditrequestcredit

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare creditrequestcredit routes
func RoutesCreditRequestsCredits(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("creditrequestscredits", "write", enforcer), baseInstance.NewCreditRequestCredit)
	router.GET("/all", middleware.Authorize("creditrequestscredits", "read", enforcer), baseInstance.GetCreditRequestsCredits)
	router.GET("/:id", middleware.Authorize("creditrequestscredits", "read", enforcer), baseInstance.GetCreditRequestCreditByID)
	router.POST("/search", middleware.Authorize("creditrequestscredits", "read", enforcer), baseInstance.SearchCreditRequestsCredits)
	router.PUT("/:id", middleware.Authorize("creditrequestscredits", "write", enforcer), baseInstance.UpdateCreditRequestCredit)
	router.DELETE("/:id", middleware.Authorize("creditrequestscredits", "write", enforcer), baseInstance.DeleteCreditRequestCredit)
}
