package insuranceproduct

import (
	"template_rest_api/api/v1/common"
	"template_rest_api/api/v1/paymee"

	"gorm.io/gorm"
)

type InsuranceProduct struct {
	ID uint `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	// IdInsurance       uint                `gorm:"column:id_insurance;" json:"id_insurance"`
	// Insurance         insurance.Insurance `gorm:"foreignkey:IdInsurance;references:ID" json:"insurance"`
	Payments          []paymee.Payment `gorm:"foreignKey:InsuranceProductID;references:ID"`
	Users             []common.User    `gorm:"many2many:userInsuranceProduct"`
	Name              string           `gorm:"column:name;not null;" json:"name"`
	Type              string           `gorm:"column:type;not null;" json:"type"`
	Description       string           `gorm:"column:description;not null;" json:"description"`
	Amount            uint             `gorm:"column:amount;" json:"amount"`
	CoverageLimit     uint             `gorm:"column:coverage_limit;" json:"coverage_limit"`
	Duration          string           `gorm:"column:duration;not null;" json:"duration"`
	PaymentPeriod     string           `gorm:"column:payment_period;not null;" json:"payment_period"`
	AdditionalOptions string           `gorm:"column:additional_options;not null;" json:"additional_options"`
	Status            string           `gorm:"column:status;not null;" json:"status"`
	CreatedBy         uint             `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new insuranceproduct
func NewInsuranceProduct(db *gorm.DB, insuranceproduct InsuranceProduct) (InsuranceProduct, error) {
	return insuranceproduct, db.Create(&insuranceproduct).Error
}

// get all insuranceproducts
func GetInsuranceProducts(db *gorm.DB) (insuranceproduct []InsuranceProduct, err error) {
	return insuranceproduct, db.Find(&insuranceproduct).Error
}

// get insuranceproduct by id
func GetInsuranceProductByID(db *gorm.DB, id uint) (insuranceproduct InsuranceProduct, err error) {
	return insuranceproduct, db.Find(&insuranceproduct, "id=?", id).Error
}

// check insuranceproduct exists
func CheckInsuranceProductExists(db *gorm.DB, id uint) bool {

	// init vars
	insuranceproduct := &InsuranceProduct{}

	// check if row exists
	check := db.Where("id=?", id).First(&insuranceproduct)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search insuranceproducts
func SearchInsuranceProducts(db *gorm.DB, insuranceproduct InsuranceProduct) (insuranceproducts []InsuranceProduct, err error) {
	return insuranceproducts, db.Where(&insuranceproduct).Find(&insuranceproducts).Error
}

// update insuranceproduct
func UpdateInsuranceProduct(db *gorm.DB, insuranceproduct InsuranceProduct) error {
	return db.Where("id=?", insuranceproduct.ID).Updates(&insuranceproduct).Error
}

// delete insuranceproduct
func DeleteInsuranceProduct(db *gorm.DB, insuranceproduct_id uint) error {
	return db.Where("id=?", insuranceproduct_id).Delete(&InsuranceProduct{}).Error
}
