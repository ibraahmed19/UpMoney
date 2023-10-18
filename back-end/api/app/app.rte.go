package app

import (
	permission "template_rest_api/api/app/permission"
	"template_rest_api/api/app/role"
	"template_rest_api/api/app/user"
	bankingaccount "template_rest_api/api/v1/bankingAccount"
	"template_rest_api/api/v1/bill"
	"template_rest_api/api/v1/credit"
	"template_rest_api/api/v1/deposit"
	insuranceproduct "template_rest_api/api/v1/insuranceProduct"
	"template_rest_api/api/v1/message"
	"template_rest_api/api/v1/paymee"
	"template_rest_api/api/v1/resetPassword"
	"template_rest_api/api/v1/simulation"
	"template_rest_api/api/v1/transaction"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare app routes
func RoutesApps(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	// user routes
	user.UserRoutes(router.Group("/user"), db, enforcer)

	// role routes
	role.RoutesRoles(router.Group("/role"), db, enforcer)

	// permission routes
	permission.RoutesPermissions(router.Group("/permission"), db, enforcer)

	// bill routes
	bill.BillRoutes(router.Group("/bill"), db, enforcer)

	// credit routes
	credit.CreditRoutes(router.Group("/credit"), db, enforcer)

	// deposit routes
	deposit.DepositRoutes(router.Group("/deposit"), db, enforcer)

	// simulation routes
	simulation.SimulationRoutes(router.Group("/simulation"), db, enforcer)

	// transaction routes
	transaction.TransactionRoutes(router.Group("/transaction"), db, enforcer)

	// bankingaccount routes
	bankingaccount.BankingAccountRoutes(router.Group("/bankingaccount"), db, enforcer)

	// insuranceproduct routes
	insuranceproduct.RoutesInsuranceProducts(router.Group("/insuranceproduct"), db, enforcer)

	// creditrequestcredit routes
	paymee.PaymeeRoutes(router.Group("/paymee"), db, enforcer)

	//chatbot routes
	//chatbot.ChatbotRoutes(router.Group("/webhook"), db, enforcer)

	message.MessageRoutes(router.Group("/message"), db, enforcer)

	resetPassword.RoutesResetPassword(router.Group("/resetpassword"), db, enforcer)

}
