package resetPassword

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesResetPassword(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/resetpassword", baseInstance.SignUpUser)
}
