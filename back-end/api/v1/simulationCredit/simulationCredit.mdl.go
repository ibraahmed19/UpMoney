package simulationcredit

import (
	"gorm.io/gorm"
)

type SimulationCredit struct {
	ID           uint `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	CreditId     uint `gorm:"column:credit_id" json:"credit_id"`
	SimulationId uint `gorm:"column:simulation_id" json:"simulation_id"`
	CreatedBy    uint `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new simulationcredit
func NewSimulationCredit(db *gorm.DB, simulationcredit SimulationCredit) (SimulationCredit, error) {
	return simulationcredit, db.Create(&simulationcredit).Error
}

// get all simulationscredits
func GetSimulationsCredits(db *gorm.DB) (simulationcredit []SimulationCredit, err error) {
	return simulationcredit, db.Find(&simulationcredit).Error
}

// get simulationcredit by id
func GetSimulationCreditByID(db *gorm.DB, id uint) (simulationcredit SimulationCredit, err error) {
	return simulationcredit, db.Find(&simulationcredit, "id=?", id).Error
}

// check simulationscredits exists
func CheckSimulationCreditExists(db *gorm.DB, id uint) bool {

	// init vars
	simulationcredit := &SimulationCredit{}

	// check if row exists
	check := db.Where("id=?", id).First(&simulationcredit)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search simulationscredits
func SearchSimulationsCredits(db *gorm.DB, simulationcredit SimulationCredit) (simulationscredits []SimulationCredit, err error) {
	return simulationscredits, db.Where(&simulationcredit).Find(&simulationscredits).Error
}

// update simulationcredit
func UpdateSimulationCredit(db *gorm.DB, simulationcredit SimulationCredit) error {
	return db.Where("id=?", simulationcredit.ID).Updates(&simulationcredit).Error
}

// delete simulationcredit
func DeleteSimulationCredit(db *gorm.DB, simulationcredit_id uint) error {
	return db.Where("id=?", simulationcredit_id).Delete(&SimulationCredit{}).Error
}
