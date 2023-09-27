package common1

import "gorm.io/gorm"

type Role struct {
	ID        uint   `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Name      string `gorm:"column:name;not null;unique" json:"name"`
	CreatedBy uint   `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// check if roles exists
func CheckRolesExists(db *gorm.DB, roles []Role) bool {

	// init vars
	var exist bool
	// check if row exists
	for _, role := range roles {
		check := db.Where("name=?", role.Name).First(&role)
		if check.Error != nil {
			exist = true
			break
		}
		if check.RowsAffected == 0 {
			exist = false
			break
		} else {
			exist = true
		}
	}
	return exist
}
