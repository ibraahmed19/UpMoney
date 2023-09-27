package bill

import (
	"template_rest_api/api/v1/common"
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	ID            uint               `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	IdTransaction uint               `gorm:"column:id_transaction;not null" json:"id_transaction"`
	Transaction   common.Transaction `gorm:"foreignkey:IdTransaction;references:ID" json:"transaction"`
	BillNumber    string             `gorm:"column:bill_number;not null;unique" json:"bill_number"`
	ClientName    string             `gorm:"column:client_name;" json:"client_name"`
	IbanAccount   uint               `gorm:"column:iban_account" json:"iban_account"`
	Amount        float64            `gorm:"column:amount;" json:"amount"`
	IssueDate     time.Time          `gorm:"column:issue_date" json:"issue_date"`
	DueDate       time.Time          `gorm:"column:due_date" json:"due_date"`
	CreationDate  time.Time          `gorm:"column:creation_date" json:"creation_date"`
	CreatedBy     uint               `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new bill
func NewBill(db *gorm.DB, bill Bill) (Bill, error) {
	return bill, db.Create(&bill).Error
}

// get all bills
func GetBills(db *gorm.DB) (bill []Bill, err error) {
	return bill, db.Find(&bill).Error
}

// get bill by id
func GetBillByID(db *gorm.DB, id uint) (bill Bill, err error) {
	return bill, db.Find(&bill, "id=?", id).Error
}

// check bill exists
func CheckBillExists(db *gorm.DB, id uint) bool {

	// init vars
	bill := &Bill{}

	// check if row exists
	check := db.Where("id=?", id).First(&bill)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search bills
func SearchBills(db *gorm.DB, bill Bill) (bills []Bill, err error) {
	return bills, db.Where(&bill).Find(&bills).Error
}

// update bill
func UpdateBill(db *gorm.DB, bill Bill) error {
	return db.Where("id=?", bill.ID).Updates(&bill).Error
}

// delete bill
func DeleteBill(db *gorm.DB, bill_id uint) error {
	return db.Where("id=?", bill_id).Delete(&Bill{}).Error
}
