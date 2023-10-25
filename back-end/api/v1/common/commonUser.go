package common

import (
	common1 "template_rest_api/api/app/common"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint             `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	FirstName       string           `gorm:"column:first_name;not null" json:"first_name"`
	LastName        string           `gorm:"column:last_name;not null" json:"last_name"`
	Email           string           `gorm:"column:email;not null;unique" json:"email"`
	Username        string           `gorm:"column:username;not null;unique" json:"username"`
	Password        string           `gorm:"column:password;not null" json:"password"`
	Adress          string           `gorm:"column:adress;not null" json:"adress"`
	Country         string           `gorm:"column:country;not null" json:"country"`
	City            string           `gorm:"column:city;" json:"city"`
	ZipCode         uint             `gorm:"column:zip_code;" json:"zip_code"`
	Phone           string           `gorm:"column:phone;not null;unique " json:"phone"`
	Roles           []common1.Role   `gorm:"many2many:UserRole;" json:"roles"`
	LastLogin       time.Time        `gorm:"column:last_login" json:"last_login"`
	InscriptionDate time.Time        `gorm:"column:inscription_date" json:"inscription_date"`
	BirthDate       time.Time        `gorm:"column:birth_date" json:"birth_date"`
	Gender          string           `gorm:"column:gender;not null" json:"gender"`
	Profession      string           `gorm:"column:profession;not null" json:"profession"`
	BankingAccounts []BankingAccount `gorm:"foreignKey:UserID;references:ID"`
	Credits         []Credit         `gorm:"foreignKey:UserID;references:ID"`
	CreatedBy       uint             `gorm:"column:created_by" json:"created_by"`
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
