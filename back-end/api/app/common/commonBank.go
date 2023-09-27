package common1

import "gorm.io/gorm"

type Bank struct {
	ID        uint   `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Name      string `gorm:"column:name;not null;unique" json:"name"`
	Address   string `gorm:"column:address;" json:"address"`
	City      string `gorm:"column:city;" json:"city"`
	Country   string `gorm:"column:country;" json:"country"`
	ZipCode   uint   `gorm:"column:zip_code;" json:"zip_code"`
	Phone     string `gorm:"column:phone;not null;unique;" json:"phone"`
	Email     string `gorm:"column:email;not null;unique;" json:"email"`
	Contact   string `gorm:"column:contact;not null;unique;" json:"contact"`
	Comments  string `gorm:"column:comments;not null;" json:"comments"`
	ManagedBy uint   `gorm:"column:managed_by;" json:"managed_by"`
	CreatedBy uint   `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}
