package client

import (
	"net/http"
	"os"
	"regexp"
	"strconv"
	common1 "template_rest_api/api/app/common"
	"template_rest_api/api/v1/common"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new client
func (db Database) NewClient(ctx *gin.Context) {

	// init vars
	var client common.Client
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(client.User.FirstName) || empty_reg.MatchString(client.User.LastName) || empty_reg.MatchString(client.User.Adress) || empty_reg.MatchString(client.User.Country) || empty_reg.MatchString(client.User.Phone) || empty_reg.MatchString(client.User.Email) || empty_reg.MatchString(client.User.Username) || empty_reg.MatchString(client.User.Password) || empty_reg.MatchString(client.Profession) || empty_reg.MatchString(client.Gender) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// // check role exists
	// if _, check := role.GetRoleByID(db.DB, client.IdRole); check != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "this role does not exist"})
	// 	return
	// }

	// check if role exists
	// if exists := role.CheckRoleExists(db.DB, client.IdRole); !exists {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "role does not exist"})
	// 	return
	// }

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// hash password
	HashPassword(&client.User.Password)

	// init new client
	new_client := common.Client{
		User:            client.User,
		InscriptionDate: client.InscriptionDate,
		BirthDate:       client.BirthDate,
		Gender:          client.Gender,
		Profession:      client.Profession,
		CreatedBy:       session.UserID,
	}
	new_role := common1.Role{
		ID:   1,
		Name: "root",
	}

	// create client
	if _, err := NewClient(db.DB, new_client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// permission
	db.Enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(client.User.ID), 10), new_role.Name)

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all clients from database
func (db Database) GetClients(ctx *gin.Context) {

	// get clients
	clients, err := GetClients(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, clients)
}

// get client by id
func (db Database) GetClientByID(ctx *gin.Context) {
	//get id value from path
	client_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := common.CheckClientExists(db.DB, uint(client_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid client id"})
		return
	}
	clientId, err := GetClientByID(db.DB, uint(client_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, clientId)
}

// search clients from database
func (db Database) SearchClients(ctx *gin.Context) {

	// init vars
	var client common.Client

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search clients from database
	clients, err := SearchClients(db.DB, client)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, clients)
}

func (db Database) UpdateClient(ctx *gin.Context) {

	// init vars
	var client common.Client
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	user_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(client.User.FirstName) || empty_reg.MatchString(client.User.LastName) || empty_reg.MatchString(client.User.Adress) || empty_reg.MatchString(client.User.Country) || empty_reg.MatchString(client.User.Phone) || empty_reg.MatchString(client.User.Email) || empty_reg.MatchString(client.User.Username) || empty_reg.MatchString(client.User.Password) || empty_reg.MatchString(client.Profession) || empty_reg.MatchString(client.Gender) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// // check role exists
	// if _, check := role.GetRoleByID(db.DB, client.IdRole); check != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "this role does not exist"})
	// 	return
	// }

	// check if role exists
	// if exists := role.CheckRoleExists(db.DB, client.IdRole); !exists {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "role does not exist"})
	// 	return
	// }

	// hash password
	HashPassword(&client.User.Password)

	// ignore key attributs
	client.User.ID = uint(user_id)
	client.CreatedBy = 0

	// update client
	if err = UpdateClient(db.DB, client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteClient(ctx *gin.Context) {

	// get id from path
	client_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete client
	if err = DeleteClient(db.DB, uint(client_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
