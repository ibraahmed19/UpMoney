package common

import (
	insuranceproduct "template_rest_api/api/v1/insuranceProduct"
	"time"

	"gorm.io/gorm"
)

type Client struct {
	//ID                uint `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	User `json:"user"`
	// IdRole            uint                                `gorm:"column:id_role;not null" json:"id_role"`
	// Role              role.Role                           `gorm:"foreignkey:IdRole;references:ID" json:"role"`
	BankingAccounts   []BankingAccount                    `gorm:"foreignKey:ClientID;references:ID"`
	Credits           []Credit                            `gorm:"foreignKey:ClientID;references:ID"`
	InsuranceProducts []insuranceproduct.InsuranceProduct `gorm:"many2many:clientInsuranceProduct"`
	// RoleName          string                              `gorm:"column:role_name;not null" json:"role_name"`
	//Email             string                              `gorm:"column:email;not null;unique" json:"email"`
	InscriptionDate time.Time `gorm:"column:inscription_date" json:"inscription_date"`
	BirthDate       time.Time `gorm:"column:birth_date" json:"birth_date"`
	Gender          string    `gorm:"column:gender;not null" json:"gender"`
	Profession      string    `gorm:"column:profession;not null" json:"profession"`
	// Password          string                              `gorm:"column:password;not null" json:"password"`
	// LastLogin time.Time `gorm:"column:last_login" json:"last_login"`
	// Banks             []common1.Bank                      `gorm:"many2many:banking_accounts" json:"banking_accounts"`
	// Insurances []insurance.Insurance `gorm:"many2many:insurance_policies" json:"insurance_policies"`
	CreatedBy uint `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// check if client exists
func CheckClientExists(db *gorm.DB, id uint) bool {

	// init vars
	client := &Client{}

	// check if row exists
	check := db.Where("id=?", id).First(&client)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// update client balance
func UpdateClientBalance(db *gorm.DB, id uint, amount float64) error {

	// init vars
	client := &Client{}

	// get client's first bank account
	if err := db.Where("client_id=?", id).First(&client).Error; err != nil {
		return err
	}

	// update client's balance
	client.BankingAccounts[0].Balance = client.BankingAccounts[0].Balance + amount

	// update client's balance
	if err := db.Where("id=?", client.BankingAccounts[0].ID).Updates(&client.BankingAccounts[0]).Error; err != nil {
		return err
	}

	return nil
}
