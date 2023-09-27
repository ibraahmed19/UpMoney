package creditrequest

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreditRequestRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("creditrequests", "write", enforcer), baseInstance.NewCreditRequest)
	router.GET("/all", middleware.Authorize("creditrequests", "read", enforcer), baseInstance.GetCreditRequests)
	router.GET("/:id", middleware.Authorize("creditrequests", "read", enforcer), baseInstance.GetCreditRequestByID)
	router.POST("/search", middleware.Authorize("creditrequests", "read", enforcer), baseInstance.SearchCreditRequests)
	router.PUT("/:id", middleware.Authorize("creditrequests", "write", enforcer), baseInstance.UpdateCreditRequest)
	router.DELETE("/:id", middleware.Authorize("creditrequests", "write", enforcer), baseInstance.DeleteCreditRequest)
}
