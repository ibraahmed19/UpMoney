package database

import (
	"fmt"
	"os"
	"strconv"
	common1 "template_rest_api/api/app/common"
	"template_rest_api/api/app/user"
	"template_rest_api/api/v1/common"
	insuranceproduct "template_rest_api/api/v1/insuranceProduct"
	"template_rest_api/api/v1/message"
	"template_rest_api/api/v1/paymee"
	"template_rest_api/api/v1/simulation"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

// auto migrate datbles
func _auto_migrate_tables(db *gorm.DB) {

	// if err :=
	// 	db.Table("banking_accounts").AutoMigrate(&common.BankingAccount{}); err != nil {
	// 	panic(err)
	// }
	// db.SetupJoinTable(&client.Client{}, "banking_account", &common1.Bank{})

	// auto migrate user, role & Bank tables
	if err := db.AutoMigrate(

		&common.User{},
		&common1.Role{},
		// &common1.Bank{},
		// &bill.Bill{},
		&message.Message{},
		&common.Credit{},

		// &creditrequest.CreditRequest{},
		// &deposit.Deposit{},
		// &insurance.Insurance{},
		&insuranceproduct.InsuranceProduct{},
		// &insurancerequest.InsuranceRequest{},
		&simulation.Simulation{},
		&paymee.Payment{},
		&common.BankingAccount{},
		&common.Transaction{},
		//&simulationcredit.SimulationCredit{},
		//&creditrequestcredit.CreditRequestCredit{},
		//&insurancepolicy.InsurancePolicy{},
	); err != nil {
		panic(err)
	}

	// auto migrate casbin table
	if err := db.Table("casbin_rule").AutoMigrate(&common1.CasbinRule{}); err != nil {
		panic(fmt.Sprintf("Error while creating casbin table: %v", err))
	}
	db.SetupJoinTable("casbin_rule", "roles", "role_permission")

	/*if err :=
		db.Table("simulation_credits").AutoMigrate(&simulationcredit.SimulationCredit{}); err != nil {
		panic(err)
	}
	db.SetupJoinTable(&simulationcredit.SimulationCredit{}, "credits", "simulation_credits")

	if err :=
		db.Table("credit_request_credits").AutoMigrate(&creditrequestcredit.CreditRequestCredit{}); err != nil {
		panic(err)
	}
	db.SetupJoinTable(&creditrequestcredit.CreditRequestCredit{}, "credits", "credit_request_credits")*/

	// if err :=
	// 	db.Table("insurance_policies ").AutoMigrate(&insurancepolicy.InsurancePolicy{}); err != nil {
	// 	panic(err)
	// }
	// db.SetupJoinTable(&insurancepolicy.InsurancePolicy{}, "clients", "insurance_policies ")

	// select from clients left join insurance_policy

}

// auto create root user
func _create_root_user(db *gorm.DB, enforcer *casbin.Enforcer) {

	// init vars:
	// root
	var user_id uint
	root_role := common1.Role{}
	root_user := common.User{}
	// default role
	user_role := common1.Role{}

	// create root role
	// check root role exists
	if check := db.Where("name=?", os.Getenv("DEFAULT_ROOT")).Find(&root_role); check.RowsAffected == 0 && check.Error == nil {

		// create role user
		db_role := common1.Role{Name: os.Getenv("DEFAULT_ROOT")}

		if err := db.Create(&db_role).Error; err != nil {
			panic(fmt.Sprintf("[WARNING] error while creating the root role: %v", err))
		}
	}

	// create root user
	// check root user exists
	if check := db.Where("email=?", os.Getenv("DEFAULT_EMAIL")).Find(&root_user); check.RowsAffected == 0 && check.Error == nil {

		// create user
		db_user := common.User{FirstName: os.Getenv("DEFAULT_FIRSTNAME"), LastName: os.Getenv("DEFAULT_LASTNAME"), Email: os.Getenv("DEFAULT_EMAIL"), Username: os.Getenv("DEFAULT_USERNAME"), Password: os.Getenv("DEFAULT_PASSWORD"), Adress: os.Getenv("DEFAULT_ADRESS"), Country: os.Getenv("DEFAULT_COUNTRY"), City: os.Getenv("DEFAULT_CITY"), ZipCode: uint(5025), Phone: os.Getenv("DEFAULT_PHONE")}
		user.HashPassword(&db_user.Password)

		if err := db.Create(&db_user).Error; err != nil {
			panic(fmt.Sprintf("[WARNING] error while creating the root user: %v", err))
		}

		// used to save user id to create Bank
		user_id = db_user.ID
	} else {

		// used to save user id to create Bank
		user_id = root_user.ID
	}

	// add policy
	enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(user_id), 10), os.Getenv("DEFAULT_ROOT"))

	// create user
	if check := db.Where("name=?", os.Getenv("DEFAULT_USER")).Find(&user_role); check.RowsAffected == 0 && check.Error == nil {

		// create role user
		db_role := common1.Role{Name: os.Getenv("DEFAULT_USER")}

		if err := db.Create(&db_role).Error; err != nil {
			panic(fmt.Sprintf("[WARNING] error while creating the user role: %v", err))
		}
	}

	// add policy
	enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(0), 10), os.Getenv("DEFAULT_USER"))

}

func AutoMigrateDatabase(db *gorm.DB, enforcer *casbin.Enforcer) {

	// create tables
	_auto_migrate_tables(db)

	// create root
	//_create_root_user(db, enforcer)
}
