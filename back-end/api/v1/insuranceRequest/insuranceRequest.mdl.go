package insurancerequest

import (
	"template_rest_api/api/app/insurance"
	"time"

	"gorm.io/gorm"
)

type InsuranceRequest struct {
	ID             uint                `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	IdInsurance    uint                `gorm:"column:id_insurance;" json:"id_insurance"`
	Insurance      insurance.Insurance `gorm:"foreignkey:IdInsurance;references:ID" json:"insurance"`
	Type           string              `gorm:"column:type;not null;" json:"type"`
	Amount         uint                `gorm:"column:amount;" json:"amount"`
	Justification  string              `gorm:"column:justification;not null;" json:"justification"`
	SubmissionDate time.Time           `gorm:"column:submission_date" json:"submission_date"`
	ResponseDate   time.Time           `gorm:"column:response_date" json:"response_date"`
	Status         string              `gorm:"column:status;not null;" json:"status"`
	Comments       string              `gorm:"column:comments;not null;" json:"comments"`
	CreatedBy      uint                `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new insurancerequest
func NewInsuranceRequest(db *gorm.DB, insurancerequest InsuranceRequest) (InsuranceRequest, error) {
	return insurancerequest, db.Create(&insurancerequest).Error
}

// get all insurancerequests
func GetInsuranceRequests(db *gorm.DB) (insurancerequest []InsuranceRequest, err error) {
	return insurancerequest, db.Find(&insurancerequest).Error
}

// get insurancerequest by id
func GetInsuranceRequestByID(db *gorm.DB, id uint) (insurancerequest InsuranceRequest, err error) {
	return insurancerequest, db.Find(&insurancerequest, "id=?", id).Error
}

// check insurancerequest exists
func CheckInsuranceRequestExists(db *gorm.DB, id uint) bool {

	// init vars
	insurancerequest := &InsuranceRequest{}

	// check if row exists
	check := db.Where("id=?", id).First(&insurancerequest)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search insurancerequests
func SearchInsuranceRequests(db *gorm.DB, insurancerequest InsuranceRequest) (insurancerequests []InsuranceRequest, err error) {
	return insurancerequests, db.Where(&insurancerequest).Find(&insurancerequests).Error
}

// update insurancerequest
func UpdateInsuranceRequest(db *gorm.DB, insurancerequest InsuranceRequest) error {
	return db.Where("id=?", insurancerequest.ID).Updates(&insurancerequest).Error
}

// delete insurancerequest
func DeleteInsuranceRequest(db *gorm.DB, insurancerequest_id uint) error {
	return db.Where("id=?", insurancerequest_id).Delete(&InsuranceRequest{}).Error
}
