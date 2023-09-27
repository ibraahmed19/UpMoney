package insurance

import (
	"gorm.io/gorm"
)

type Insurance struct {
	ID          uint   `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Name        string `gorm:"column:name;not null;unique" json:"name"`
	Type        string `gorm:"column:type;not null" json:"type"`
	Address     string `gorm:"column:address" json:"address"`
	Phone       string `gorm:"column:phone;not null;unique;" json:"phone"`
	Email       string `gorm:"column:email;not null;unique;" json:"email"`
	Description string `gorm:"column:description;not null;" json:"description"`
	ManagedBy   uint   `gorm:"column:managed_by;" json:"managed_by"`
	CreatedBy   uint   `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new insurance
func NewInsurance(db *gorm.DB, insurance Insurance) (Insurance, error) {
	return insurance, db.Create(&insurance).Error
}

// get all insurances
func GetInsurances(db *gorm.DB) (insurance []Insurance, err error) {
	return insurance, db.Find(&insurance).Error
}

// get insurance by id
func GetInsuranceByID(db *gorm.DB, id uint) (insurance Insurance, err error) {
	return insurance, db.Find(&insurance, "id=?", id).Error
}

// check insurance exists
func CheckInsuranceExists(db *gorm.DB, id uint) bool {

	// init vars
	insurance := &Insurance{}

	// check if row exists
	check := db.Where("id=?", id).First(&insurance)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search insurances
func SearchInsurances(db *gorm.DB, insurance Insurance) (insurances []Insurance, err error) {
	return insurances, db.Where(&insurance).Find(&insurances).Error
}

// update insurance
func UpdateInsurance(db *gorm.DB, insurance Insurance) error {
	return db.Where("id=?", insurance.ID).Updates(&insurance).Error
}

// delete insurance
func DeleteInsurance(db *gorm.DB, insurance_id uint) error {
	return db.Where("id=?", insurance_id).Delete(&Insurance{}).Error
}
