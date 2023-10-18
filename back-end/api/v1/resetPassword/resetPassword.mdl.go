package resetPassword

import (
	"template_rest_api/api/app/user"

	"gorm.io/gorm"
)

type resetPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// update user
func UpdatePassword(db *gorm.DB, userID uint, newPassword string) error {
	return db.Model(&user.User{}).Where("id = ?", userID).Update("password", newPassword).Error
}

// check if user exists
func CheckMailExists(db *gorm.DB, email string) bool {

	// init vars
	user := &user.User{}

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
