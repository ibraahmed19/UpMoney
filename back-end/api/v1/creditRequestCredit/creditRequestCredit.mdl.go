package creditrequestcredit

import (
	"gorm.io/gorm"
)

type CreditRequestCredit struct {
	ID              uint `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	CreditId        uint `gorm:"column:credit_id" json:"credit_id"`
	CreditRequestId uint `gorm:"column:credit_request_id" json:"credit_request_id"`
	CreatedBy       uint `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new creditrequestcredit
func NewCreditRequestCredit(db *gorm.DB, creditrequestcredit CreditRequestCredit) (CreditRequestCredit, error) {
	return creditrequestcredit, db.Create(&creditrequestcredit).Error
}

// get all creditrequestscredits
func GetCreditRequestsCredits(db *gorm.DB) (creditrequestcredit []CreditRequestCredit, err error) {
	return creditrequestcredit, db.Find(&creditrequestcredit).Error
}

// get creditrequestcredit by id
func GetCreditRequestCreditByID(db *gorm.DB, id uint) (creditrequestcredit CreditRequestCredit, err error) {
	return creditrequestcredit, db.Find(&creditrequestcredit, "id=?", id).Error
}

// check creditrequestscredits exists
func CheckCreditRequestCreditExists(db *gorm.DB, id uint) bool {

	// init vars
	creditrequestcredit := &CreditRequestCredit{}

	// check if row exists
	check := db.Where("id=?", id).First(&creditrequestcredit)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search creditrequestscredits
func SearchCreditRequestsCredits(db *gorm.DB, creditrequestcredit CreditRequestCredit) (creditrequestscredits []CreditRequestCredit, err error) {
	return creditrequestscredits, db.Where(&creditrequestcredit).Find(&creditrequestscredits).Error
}

// update creditrequestcredit
func UpdateCreditRequestCredit(db *gorm.DB, creditrequestcredit CreditRequestCredit) error {
	return db.Where("id=?", creditrequestcredit.ID).Updates(&creditrequestcredit).Error
}

// delete creditrequestcredit
func DeleteCreditRequestCredit(db *gorm.DB, creditrequestcredit_id uint) error {
	return db.Where("id=?", creditrequestcredit_id).Delete(&CreditRequestCredit{}).Error
}
