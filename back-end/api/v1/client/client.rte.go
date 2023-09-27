package client

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ClientRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("clients", "write", enforcer), baseInstance.NewClient)
	router.GET("/all", middleware.Authorize("clients", "read", enforcer), baseInstance.GetClients)
	router.GET("/:id", middleware.Authorize("clients", "read", enforcer), baseInstance.GetClientByID)
	router.POST("/search", middleware.Authorize("clients", "read", enforcer), baseInstance.SearchClients)
	router.PUT("/:id", middleware.Authorize("clients", "write", enforcer), baseInstance.UpdateClient)
	router.DELETE("/:id", middleware.Authorize("clients", "write", enforcer), baseInstance.DeleteClient)
}
