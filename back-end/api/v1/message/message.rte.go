package message

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MessageRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.GET("/:id", middleware.Authorize("messages", "read", enforcer), baseInstance.GetMessageByID)
	router.POST("/search", middleware.Authorize("messages", "read", enforcer), baseInstance.SearchTrigger)
}
