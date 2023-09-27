package simulation

import (
	"time"

	"gorm.io/gorm"
)

type Simulation struct {
	ID             uint      `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Type           string    `gorm:"column:type;" json:"type"`
	Date           time.Time `gorm:"column:date" json:"date"`
	Results        string    `gorm:"column:results;" json:"results"`
	VariablesEntry string    `gorm:"column:variables_entry;" json:"variables_entry"`
	VariablesOut   string    `gorm:"column:variables_out;" json:"variables_out"`
	CreatedBy      uint      `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new simulation
func NewSimulation(db *gorm.DB, simulation Simulation) (Simulation, error) {
	return simulation, db.Create(&simulation).Error
}

// get all simulations
func GetSimulations(db *gorm.DB) (simulation []Simulation, err error) {
	return simulation, db.Find(&simulation).Error
}

// get simulation by id
func GetSimulationByID(db *gorm.DB, id uint) (simulation Simulation, err error) {
	return simulation, db.Find(&simulation, "id=?", id).Error
}

// check simulation exists
func CheckSimulationExists(db *gorm.DB, id uint) bool {

	// init vars
	simulation := &Simulation{}

	// check if row exists
	check := db.Where("id=?", id).First(&simulation)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search simulations
func SearchSimulations(db *gorm.DB, simulation Simulation) (simulations []Simulation, err error) {
	return simulations, db.Where(&simulation).Find(&simulations).Error
}

// update simulation
func UpdateSimulation(db *gorm.DB, simulation Simulation) error {
	return db.Where("id=?", simulation.ID).Updates(&simulation).Error
}

// delete simulation
func DeleteSimulation(db *gorm.DB, simulation_id uint) error {
	return db.Where("id=?", simulation_id).Delete(&Simulation{}).Error
}
