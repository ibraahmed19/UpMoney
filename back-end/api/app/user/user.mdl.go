package user

import (
	common1 "template_rest_api/api/app/common"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	FirstName string         `gorm:"column:first_name;not null" json:"first_name"`
	LastName  string         `gorm:"column:last_name;not null" json:"last_name"`
	Email     string         `gorm:"column:email;not null;unique" json:"email"`
	Username  string         `gorm:"column:username;not null;unique" json:"username"`
	Password  string         `gorm:"column:password;not null" json:"password"`
	Adress    string         `gorm:"column:adress;not null" json:"adress"`
	Country   string         `gorm:"column:country;not null" json:"country"`
	City      string         `gorm:"column:city;" json:"city"`
	ZipCode   uint           `gorm:"column:zip_code;" json:"zip_code"`
	Phone     string         `gorm:"column:phone;not null;unique " json:"phone"`
	Roles     []common1.Role `gorm:"many2many:UserRole"`
	LastLogin time.Time      `gorm:"column:last_login" json:"last_login"`
	// BankID    uint      `gorm:"column:bank_id" json:"bank_id"`
	// //Bank        bank.Bank           `gorm:"foreignkey:BankID" json:"bank"`
	// InsuranceID uint `gorm:"column:insurance_id" json:"insurance_id"`
	// // Insurance   insurance.Insurance `gorm:"foreignkey:InsuranceID" json:"insurance"`
	CreatedBy uint `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// hash password
func HashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

// create new user
func NewUser(db *gorm.DB, user User) (User, error) {
	return user, db.Create(&user).Error
}

// get all users
func GetUsers(db *gorm.DB) (users []User, err error) {
	return users, db.Find(&users).Error
}

// search users
func SearchUsers(db *gorm.DB, user User) (users []User, err error) {
	return users, db.Where(&user).Find(&users).Error
}

// update user
func UpdateUser(db *gorm.DB, user User) error {
	return db.Where("id=?", user.ID).Updates(&user).Error
}

// delete user
func DeleteUser(db *gorm.DB, user_id uint) error {
	return db.Where("id=?", user_id).Delete(&User{}).Error
}

// get user by email
func GetUserByEmail(db *gorm.DB, email string) (user User, err error) {
	return user, db.First(&user, "email=?", email).Error
}

// get user by id
func GetUserByID(db *gorm.DB, id uint) (user User, err error) {
	return user, db.Find(&user, "id=?", id).Error
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

// compare two passwords
func ComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}
