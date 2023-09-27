package user

import (
	"net/http"
	"os"
	"regexp"
	"strconv"

	common1 "template_rest_api/api/app/common"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new user
func (db Database) NewUser(ctx *gin.Context) {

	// init vars
	var user User
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(user.FirstName) || empty_reg.MatchString(user.LastName) || empty_reg.MatchString(user.Email) || empty_reg.MatchString(user.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check role exists
	if check := common1.CheckRolesExists(db.DB, user.Roles); !check {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "those roles does not exist"})
		return
	}

	// // check if insurance exists
	// if exists := insurance.CheckInsuranceExists(db.DB, user.InsuranceID); !exists {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "insurance does not exist"})
	// 	return
	// }

	// // check if bank exists
	// if exists := bank.CheckBankExists(db.DB, user.BankID); !exists {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "bank does not exist"})
	// 	return
	// }

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// hash password
	HashPassword(&user.Password)

	// init new user
	new_user := User{
		ID:        0,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Adress:    user.Adress,
		Country:   user.Country,
		City:      user.City,
		ZipCode:   user.ZipCode,
		Phone:     user.Phone,
		Roles:     user.Roles,
		LastLogin: user.LastLogin,
		// BankID:      session.BankID,
		// InsuranceID: session.InsuranceID,
		CreatedBy: session.UserID,
	}

	// create user
	if _, err := NewUser(db.DB, new_user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// permission
	db.Enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(user.ID), 10), user.Roles)

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all users from database
func (db Database) GetUsers(ctx *gin.Context) {

	// get users
	users, err := GetUsers(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// get user by id
func (db Database) GetUserByID(ctx *gin.Context) {
	//get id value from path
	user_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckUserExists(db.DB, uint(user_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}
	userId, err := GetUserByID(db.DB, uint(user_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userId)
}

// search users from database
func (db Database) SearchUsers(ctx *gin.Context) {

	// init vars
	var user User

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search users from database
	users, err := SearchUsers(db.DB, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (db Database) UpdateUser(ctx *gin.Context) {

	// init vars
	var user User
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&user); err != nil {
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
	if empty_reg.MatchString(user.FirstName) || empty_reg.MatchString(user.LastName) || empty_reg.MatchString(user.Email) || empty_reg.MatchString(user.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check role exists
	if check := common1.CheckRolesExists(db.DB, user.Roles); !check {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "those roles does not exist"})
		return
	}

	// hash password
	HashPassword(&user.Password)

	// ignore key attributs
	user.ID = uint(user_id)
	user.CreatedBy = 0
	// user.BankID = 0

	// update user
	if err = UpdateUser(db.DB, user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteUser(ctx *gin.Context) {

	// get id from path
	user_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete user
	if err = DeleteUser(db.DB, uint(user_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
