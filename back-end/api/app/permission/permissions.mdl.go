package app_permission

import (
	common1 "template_rest_api/api/app/common"

	"gorm.io/gorm"
)

func GetPermissions(db *gorm.DB) (permissions []common1.CasbinRule, err error) {
	return permissions, db.Table("casbin_rule").Find(&permissions, "ptype=?", "p").Error
}

func GetPermissionByID(db *gorm.DB, id uint) (permission common1.CasbinRule, err error) {
	return permission, db.Table("casbin_rule").Where("id = ? AND ptype = 'p'", id).Find(&permission).Error
}

func CheckRolePermissionExists(db *gorm.DB, role string) (role_exists common1.Role, err error) {
	return role_exists, db.Table("roles").Where("name=?", role).Error
}

func CheckRoleExists(db *gorm.DB, name string) (role common1.Role, err error) {
	return role, db.Table("roles").Where("name = ?", name).Find(&role).Error
}
