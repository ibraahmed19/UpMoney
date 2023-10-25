package resetPassword

import (
	"template_rest_api/api/v1/common"

	"gorm.io/gorm"
)

type resetPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// update user
func UpdatePassword(db *gorm.DB, userID uint, newPassword string) error {
	return db.Model(&common.User{}).Where("id = ?", userID).Update("password", newPassword).Error
}

// check if user exists
func CheckMailExists(db *gorm.DB, email string) bool {

	// init vars
	user := &common.User{}

	// check if row exists
	check := db.Where("email=?", email).First(&user)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
