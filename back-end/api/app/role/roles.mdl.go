package role

import (
	common1 "template_rest_api/api/app/common"

	"gorm.io/gorm"
)

func NewRole(db *gorm.DB, role common1.Role) error {
	return db.Create(&role).Error
}

func GetRoles(db *gorm.DB) (role []common1.Role, err error) {
	return role, db.Find(&role).Error
}

func GetRole(db *gorm.DB, name string) error {
	role := common1.Role{}
	return db.Where("name=?", name).First(&role).Error
}

// // get role by ID
// func GetRoleByID(db *gorm.DB, id uint) (Role, error) {
// 	role := Role{}
// 	return role, db.Where("id=?", id).First(&role).Error
// }

// check if role exists
func CheckRoleExists(db *gorm.DB, id uint) bool {

	// init vars
	role := &common1.Role{}

	// check if row exists
	check := db.Where("id=?", id).First(&role)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func SearchRoles(db *gorm.DB, role common1.Role) (roles []common1.Role, err error) {
	return roles, db.Where(&role).Find(&roles).Error
}

func UpdateRole(db *gorm.DB, role common1.Role) error {
	return db.Where("id=?", role.ID).Updates(&role).Error
}

func DeleteRole(db *gorm.DB, role_id uint) error {
	return db.Where("id=?", role_id).Delete(&common1.Role{}).Error
}
