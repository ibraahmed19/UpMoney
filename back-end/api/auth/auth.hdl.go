package app_auth

import (
	"net/http"
	"os"
	"regexp"
	common1 "template_rest_api/api/app/common"
	"template_rest_api/api/app/role"
	user "template_rest_api/api/app/user"
	"template_rest_api/api/v1/common"
	"template_rest_api/middleware"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// signup user
func (db Database) SignUpUser(ctx *gin.Context) {

	// init vars
	var userToInsert common.User
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// upmarshal sent json
	if err := ctx.ShouldBindJSON(&userToInsert); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check field validity
	if empty_reg.MatchString(userToInsert.Username) || empty_reg.MatchString(userToInsert.Email) || empty_reg.MatchString(userToInsert.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// hash password
	user.HashPassword(&userToInsert.Password)

	// create new user
	new_user := common.User{
		FirstName:       userToInsert.FirstName,
		LastName:        userToInsert.LastName,
		Email:           userToInsert.Email,
		Username:        userToInsert.Username,
		Password:        userToInsert.Password,
		Adress:          userToInsert.Adress,
		Country:         userToInsert.Country,
		City:            userToInsert.City,
		ZipCode:         userToInsert.ZipCode,
		Phone:           userToInsert.Phone,
		InscriptionDate: time.Now(),
		BirthDate:       userToInsert.BirthDate,
		Gender:          userToInsert.Gender,
		Profession:      userToInsert.Profession,
	}

	// Définir l'ID du rôle par défaut sur 1
	new_user.Roles = []common1.Role{{ID: 1}}

	role_id := new_user.Roles[0].ID
	role, err := role.GetRoleByID(db.DB, uint(role_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	new_user.Roles = append(new_user.Roles, role)

	// create user
	//saved_user,
	_, err = user.NewUser(db.DB, new_user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// signin user
func (db Database) SignInUser(ctx *gin.Context) {

	// init cars
	var user_login UserLogIn
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&user_login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	// check field validity
	if empty_reg.MatchString(user_login.Email) || empty_reg.MatchString(user_login.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check if email exists
	dbUser, err := user.GetUserByEmail(db.DB, user_login.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "no Such User Found"})
		return
	}

	// update last login
	dbUser.LastLogin = time.Now()

	// update user
	if err := user.UpdateUser(db.DB, dbUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// compare password
	if isTrue := user.ComparePassword(dbUser.Password, user_login.Password); isTrue {

		// generate token
		token := middleware.GenerateToken(dbUser.ID, "root")
		ctx.JSON(http.StatusOK, UserLogedIn{Token: token})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"message": "password not matched"})
}
