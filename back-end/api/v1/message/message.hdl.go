package message

import (
	"net/http"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

func (db Database) SearchTrigger(ctx *gin.Context) {

	// init vars
	var message Message

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search credits from database
	messages, err := SearchTrigger(db.DB, message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, messages)
}

func (db Database) GetMessageByID(ctx *gin.Context) {
	//get id value from path
	message_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckMessageExists(db.DB, uint(message_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid bill id"})
		return
	}
	messageId, err := GetMessageByID(db.DB, uint(message_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, messageId)
}
