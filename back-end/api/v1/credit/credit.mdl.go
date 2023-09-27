package credit

import (
	"template_rest_api/api/v1/common"

	"gorm.io/gorm"
)

// create new credit
func NewCredit(db *gorm.DB, credit common.Credit) (common.Credit, error) {
	return credit, db.Create(&credit).Error
}

// get all credits
func GetCredits(db *gorm.DB) (credit []common.Credit, err error) {
	return credit, db.Find(&credit).Error
}

// get credit by id
func GetCreditByID(db *gorm.DB, id uint) (credit common.Credit, err error) {
	return credit, db.Find(&credit, "id=?", id).Error
}



// search credits
func SearchCredits(db *gorm.DB, credit common.Credit) (credits []common.Credit, err error) {
	return credits, db.Where(&credit).Find(&credits).Error
}

// update credit
func UpdateCredit(db *gorm.DB, credit common.Credit) error {
	return db.Where("id=?", credit.ID).Updates(&credit).Error
}

// delete credit
func DeleteCredit(db *gorm.DB, credit_id uint) error {
	return db.Where("id=?", credit_id).Delete(&common.Credit{}).Error
}
