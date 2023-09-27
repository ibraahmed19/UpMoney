package app

import (
	"template_rest_api/api/app/bank"
	"template_rest_api/api/app/insurance"
	permission "template_rest_api/api/app/permission"
	"template_rest_api/api/app/role"
	"template_rest_api/api/app/user"
	bankingaccount "template_rest_api/api/v1/bankingAccount"
	"template_rest_api/api/v1/bill"
	"template_rest_api/api/v1/client"
	"template_rest_api/api/v1/credit"
	creditrequest "template_rest_api/api/v1/creditRequest"
	creditrequestcredit "template_rest_api/api/v1/creditRequestCredit"
	"template_rest_api/api/v1/deposit"
	insurancepolicy "template_rest_api/api/v1/insurancePolicy"
	insuranceproduct "template_rest_api/api/v1/insuranceProduct"
	insurancerequest "template_rest_api/api/v1/insuranceRequest"
	"template_rest_api/api/v1/message"
	"template_rest_api/api/v1/paymee"
	"template_rest_api/api/v1/simulation"
	simulationcredit "template_rest_api/api/v1/simulationCredit"
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

	// bank routes
	bank.RoutesBanks(router.Group("/bank"), db, enforcer)

	// bill routes
	bill.BillRoutes(router.Group("/bill"), db, enforcer)

	// client routes
	client.ClientRoutes(router.Group("/client"), db, enforcer)

	// credit routes
	credit.CreditRoutes(router.Group("/credit"), db, enforcer)

	// deposit routes
	deposit.DepositRoutes(router.Group("/deposit"), db, enforcer)

	// insurance routes
	insurance.RoutesInsurances(router.Group("/insurance"), db, enforcer)

	// simulation routes
	simulation.SimulationRoutes(router.Group("/simulation"), db, enforcer)

	// transaction routes
	transaction.TransactionRoutes(router.Group("/transaction"), db, enforcer)

	// bankingaccount routes
	bankingaccount.BankingAccountRoutes(router.Group("/bankingaccount"), db, enforcer)

	// creditrequest routes
	creditrequest.CreditRequestRoutes(router.Group("/creditrequest"), db, enforcer)

	// insurancepolicy routes
	insurancepolicy.RoutesInsurancePolicies(router.Group("/insurancepolicy"), db, enforcer)

	// insuranceproduct routes
	insuranceproduct.RoutesInsuranceProducts(router.Group("/insuranceproduct"), db, enforcer)

	// insurancerequest routes
	insurancerequest.RoutesInsuranceRequests(router.Group("/insurancerequest"), db, enforcer)

	// simulationcredit routes
	simulationcredit.RoutesSimulationsCredits(router.Group("/simulationcredit"), db, enforcer)

	// creditrequestcredit routes
	creditrequestcredit.RoutesCreditRequestsCredits(router.Group("/creditrequestcredit"), db, enforcer)

	// creditrequestcredit routes
	paymee.PaymeeRoutes(router.Group("/paymee"), db, enforcer)

	//chatbot routes
	//chatbot.ChatbotRoutes(router.Group("/webhook"), db, enforcer)

	message.MessageRoutes(router.Group("/message"), db, enforcer)

}
