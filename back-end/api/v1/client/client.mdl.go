package client

import (
	"template_rest_api/api/v1/common"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// hash password
func HashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

// create new client
func NewClient(db *gorm.DB, client common.Client) (common.Client, error) {
	return client, db.Create(&client).Error
}

// get all clients
func GetClients(db *gorm.DB) (clients []common.Client, err error) {
	return clients, db.Find(&clients).Error
}

// search clients
func SearchClients(db *gorm.DB, client common.Client) (clients []common.Client, err error) {
	return clients, db.Where(&client).Find(&clients).Error
}

// update client
func UpdateClient(db *gorm.DB, client common.Client) error {
	return db.Where("id=?", client.User.ID).Updates(&client).Error
}

// delete client
func DeleteClient(db *gorm.DB, client_id uint) error {
	return db.Where("id=?", client_id).Delete(&common.Client{}).Error
}

// get client by email
func GetClientByEmail(db *gorm.DB, email string) (client common.Client, err error) {
	return client, db.First(&client, "email=?", email).Error
}

// get client by id
func GetClientByID(db *gorm.DB, id uint) (client common.Client, err error) {
	return client, db.Find(&client, "id=?", id).Error
}

// compare two passwords
func ComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}
