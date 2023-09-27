package creditrequest

import (
	"time"

	"gorm.io/gorm"
)

type CreditRequest struct {
	ID              uint      `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	ClientId        uint      `gorm:"column:created_by" json:"client_id"`
	Type            string    `gorm:"column:type;" json:"type"`
	AmountRequested float64   `gorm:"column:amount_requested;" json:"amount_requested"`
	Justification   string    `gorm:"column:justification;" json:"justification"`
	RequestStatus   string    `gorm:"column:request_status;" json:"request_status"`
	SubmissionDate  time.Time `gorm:"column:submission_date;not null" json:"submission_date" `
	ApprovalDate    time.Time `gorm:"column:approval_date;not null" json:"approval_date"`
	RejectionDate   time.Time `gorm:"column:rejection_date;not null" json:"rejection_date"`
	ApprovedAmount  float64   `gorm:"column:approved_amount;" json:"approved_amount"`
	InterestRate    float64   `gorm:"column:interest_rate;" json:"interest_rate"`
	Duration        uint      `gorm:"column:duration" json:"duration"`
	MonthlyPayments float64   `gorm:"column:monthly_payments;" json:"monthly_payments"`
	Notes           string    `gorm:"column:notes;" json:"notes"`
	CreatedBy       uint      `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new creditrequest
func NewCreditRequest(db *gorm.DB, creditrequest CreditRequest) (CreditRequest, error) {
	return creditrequest, db.Create(&creditrequest).Error
}

// get all creditrequests
func GetCreditRequests(db *gorm.DB) (creditrequest []CreditRequest, err error) {
	return creditrequest, db.Find(&creditrequest).Error
}

// get creditrequest by id
func GetCreditRequestByID(db *gorm.DB, id uint) (creditrequest CreditRequest, err error) {
	return creditrequest, db.Find(&creditrequest, "id=?", id).Error
}

// check creditrequest exists
func CheckCreditRequestExists(db *gorm.DB, id uint) bool {

	// init vars
	creditrequest := &CreditRequest{}

	// check if row exists
	check := db.Where("id=?", id).First(&creditrequest)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search creditrequest
func SearchCreditRequests(db *gorm.DB, creditrequest CreditRequest) (creditrequests []CreditRequest, err error) {
	return creditrequests, db.Where(&creditrequest).Find(&creditrequests).Error
}

// update creditrequest
func UpdateCreditRequest(db *gorm.DB, creditrequest CreditRequest) error {
	return db.Where("id=?", creditrequest.ID).Updates(&creditrequest).Error
}

// delete creditrequest
func DeleteCreditRequest(db *gorm.DB, creditrequest_id uint) error {
	return db.Where("id=?", creditrequest_id).Delete(&CreditRequest{}).Error
}
