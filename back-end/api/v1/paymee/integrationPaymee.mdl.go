package paymee

import (
	"gorm.io/gorm"
)

// payment
type Payment struct {
	ID                 uint    `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	InsuranceProductID uint    `gorm:"column:insurance_product_id" json:"insurance_product_id"`
	Amount             float64 `gorm:"column:amount;" json:"amount"`
	Note               string  `gorm:"column:note;" json:"note"`
	FirstName          string  `gorm:"column:first_name;" json:"first_name"`
	LastName           string  `gorm:"column:last_name;" json:"last_name"`
	Email              string  `gorm:"column:email;not null" json:"email"`
	Phone              string  `gorm:"column:phone;not null " json:"phone"`
	ReturnUrl          string  `gorm:"column:return_url;not null " json:"return_url"`
	CancelUrl          string  `gorm:"column:cancel_url;not null " json:"cancel_url"`
	WebhookUrl         string  `gorm:"column:webhook_url;not null " json:"webhook_url"`
	OrderId            string  `gorm:"column:order_id;not null;unique " json:"order_id"`
	CreatedBy          uint    `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new payment
func NewPayment(db *gorm.DB, paymee Payment) (Payment, error) {
	return paymee, db.Create(&paymee).Error
}

// get all payment
func GetPayments(db *gorm.DB) (paymee []Payment, err error) {
	return paymee, db.Find(&paymee).Error
}

// get payment by id
func GetPaymentByID(db *gorm.DB, id uint) (paymee Payment, err error) {
	return paymee, db.Find(&paymee, "id=?", id).Error
}

// search payments
func SearchPayments(db *gorm.DB, payment Payment) (payments []Payment, err error) {
	return payments, db.Where(&payment).Find(&payments).Error
}

// check payment exists
func CheckPaymentExists(db *gorm.DB, id uint) bool {

	// init vars
	paymee := &Payment{}

	// check if row exists
	check := db.Where("id=?", id).First(&paymee)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}

}
