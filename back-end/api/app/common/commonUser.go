package common1

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint      `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Email       string    `gorm:"column:email;not null;unique" json:"email"`
	Password    string    `gorm:"column:password;not null" json:"password"`
	Role        string    `gorm:"column:role;not null" json:"role"`
	LastLogin   time.Time `gorm:"column:last_login" json:"last_login"`
	BankID      uint      `gorm:"column:bank_id" json:"bank_id"`
	InsuranceID uint      `gorm:"column:insurance_id" json:"insurance_id"`
	CreatedBy   uint      `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// check if user exists
func CheckUserExists(db *gorm.DB, id uint) bool {

	// init vars
	user := &User{}

	// check if row exists
	check := db.Where("id=?", id).First(&user)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
