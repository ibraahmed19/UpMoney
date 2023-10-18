package common

import (
	"template_rest_api/api/v1/simulation"
	"time"

	"gorm.io/gorm"
)

type Credit struct {
	ID     uint `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	UserID uint `gorm:"column:user_id" json:"user_id"`
	// Client          Client                  `gorm:"foreignkey:ClientId;references:ID" json:"client"`
	Type            string                  `gorm:"column:type;" json:"type"`
	Amount          float64                 `gorm:"column:amount;" json:"amount"`
	InterestRate    float64                 `gorm:"column:interest_rate;" json:"interest_rate"`
	Duration        string                  `gorm:"column:duration" json:"duration"`
	StartDate       time.Time               `gorm:"column:start_date;not null" json:"start_date"`
	EndDate         time.Time               `gorm:"column:end_date;not null" json:"end_date"`
	MonthlyPayments float64                 `gorm:"column:monthly_payments;" json:"monthly_payments"`
	Status          string                  `gorm:"column:status;" json:"status"`
	Comments        string                  `gorm:"column:comments;" json:"comments"`
	Simulations     []simulation.Simulation `gorm:"many2many:simulationCredits"`
	//CreditRequests  []creditrequest.CreditRequest `gorm:"many2many:credit_request_credits" json:"credit_request_credits"`
	CreatedBy uint `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// check credit exists
func CheckCreditExists(db *gorm.DB, id uint) bool {

	// init vars
	credit := &Credit{}

	// check if row exists
	check := db.Where("id=?", id).First(&credit)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
