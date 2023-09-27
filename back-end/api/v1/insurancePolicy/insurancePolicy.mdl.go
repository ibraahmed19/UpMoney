package insurancepolicy

import (
	"time"

	"gorm.io/gorm"
)

type InsurancePolicy struct {
	ID              uint      `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	ClientId        uint      `gorm:"column:client_id;" json:"client_id"`
	InsuranceId     uint      `gorm:"column:insurance_id" json:"insurance_id"`
	Type            string    `gorm:"column:type;not null;" json:"type"`
	Amount          uint      `gorm:"column:amount;" json:"amount"`
	StartDate       time.Time `gorm:"column:start_date" json:"start_date"`
	EndTime         time.Time `gorm:"column:end_date" json:"end_date"`
	WarrantyDetails string    `gorm:"column:warranty_details;not null;" json:"warranty_details"`
	Franchise       string    `gorm:"column:franchise;not null;" json:"franchise"`
	CoverageLimit   uint      `gorm:"column:coverage_limit;" json:"coverage_limit"`
	PolicyStatus    string    `gorm:"column:policy_status;not null;" json:"policy_status"`
	CreatedBy       uint      `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new insurancepolicy
func NewInsurancePolicy(db *gorm.DB, insurancepolicy InsurancePolicy) (InsurancePolicy, error) {
	return insurancepolicy, db.Create(&insurancepolicy).Error
}

// get all insurancepolicies
func GetInsurancePolicies(db *gorm.DB) (insurancepolicy []InsurancePolicy, err error) {
	return insurancepolicy, db.Find(&insurancepolicy).Error
}

// get insurancepolicy by id
func GetInsurancePolicyByID(db *gorm.DB, id uint) (insurancepolicy InsurancePolicy, err error) {
	return insurancepolicy, db.Find(&insurancepolicy, "id=?", id).Error
}

// check insurancepolicy exists
func CheckInsurancePolicyExists(db *gorm.DB, id uint) bool {

	// init vars
	insurancepolicy := &InsurancePolicy{}

	// check if row exists
	check := db.Where("id=?", id).First(&insurancepolicy)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search insurancepolicies
func SearchInsurancePolicies(db *gorm.DB, insurancepolicy InsurancePolicy) (insurancepolicies []InsurancePolicy, err error) {
	return insurancepolicies, db.Where(&insurancepolicy).Find(&insurancepolicies).Error
}

// update insurancepolicy
func UpdateInsurancePolicy(db *gorm.DB, insurancepolicy InsurancePolicy) error {
	return db.Where("id=?", insurancepolicy.ID).Updates(&insurancepolicy).Error
}

// delete insurancepolicy
func DeleteInsurancePolicy(db *gorm.DB, insurancepolicy_id uint) error {
	return db.Where("id=?", insurancepolicy_id).Delete(&InsurancePolicy{}).Error
}
