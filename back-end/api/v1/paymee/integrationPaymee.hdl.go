package paymee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new payment
func (db Database) NewPayment(ctx *gin.Context) {
	// init vars
	var paymee Payment
	emptyReg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// check JSON validity
	if err := ctx.ShouldBindJSON(&paymee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if emptyReg.MatchString(paymee.Note) || emptyReg.MatchString(paymee.FirstName) || emptyReg.MatchString(paymee.LastName) || emptyReg.MatchString(paymee.Email) || emptyReg.MatchString(paymee.Phone) || emptyReg.MatchString(paymee.ReturnUrl) || emptyReg.MatchString(paymee.CancelUrl) || emptyReg.MatchString(paymee.WebhookUrl) || emptyReg.MatchString(paymee.OrderId) || paymee.Amount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new payment
	newPayment := Payment{
		InsuranceProductID: paymee.InsuranceProductID,
		Amount:             paymee.Amount,
		Note:               paymee.Note,
		FirstName:          paymee.FirstName,
		LastName:           paymee.LastName,
		Email:              paymee.Email,
		Phone:              paymee.Phone,
		ReturnUrl:          paymee.ReturnUrl,
		CancelUrl:          paymee.CancelUrl,
		WebhookUrl:         paymee.WebhookUrl,
		OrderId:            paymee.OrderId,
		CreatedBy:          session.UserID,
	}

	// create new payment
	_, err := NewPayment(db.DB, newPayment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// initiate payment on Paymee API
	paymentURL, err := InitiatePayment(&PaymeeClient{}, newPayment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to initiate payment"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created", "payment_url": paymentURL})
}

type PaymeeClient struct {
	APIKey string
}

// initiate payment on Paymee API
func InitiatePayment(client *PaymeeClient, payment Payment) (string, error) {
	client.APIKey = "135a85997bf19dc1a63424cc532039a86f57310f"

	// Prepare request data
	paymentReq := struct {
		Amount             float64 `json:"amount"`
		InsuranceProductID uint    `json:"insurance_product_id"`
		Note               string  `json:"note"`
		FirstName          string  `json:"first_name"`
		LastName           string  `json:"last_name"`
		Email              string  `json:"email"`
		Phone              string  `json:"phone"`
		ReturnURL          string  `json:"return_url"`
		CancelURL          string  `json:"cancel_url"`
		WebhookURL         string  `json:"webhook_url"`
		OrderID            string  `json:"order_id,omitempty"`
	}{
		Amount:             payment.Amount,
		InsuranceProductID: payment.InsuranceProductID,
		Note:               payment.Note,
		FirstName:          payment.FirstName,
		LastName:           payment.LastName,
		Email:              payment.Email,
		Phone:              payment.Phone,
		ReturnURL:          payment.ReturnUrl,
		CancelURL:          payment.CancelUrl,
		WebhookURL:         payment.WebhookUrl,
		OrderID:            payment.OrderId,
	}

	// Convert structure to JSON
	reqBody, err := json.Marshal(paymentReq)
	if err != nil {
		return "", err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", "https://sandbox.paymee.tn/api/v2/payments/create", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	// Add request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token "+client.APIKey)

	// Send HTTP request
	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read HTTP response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	// Parse JSON response
	type PaymentResponse struct {
		Status bool `json:"status"`
		Data   struct {
			Token      string  `json:"token"`
			OrderID    string  `json:"order_id"`
			FirstName  string  `json:"first_name"`
			LastName   string  `json:"last_name"`
			Email      string  `json:"email"`
			Phone      string  `json:"phone"`
			Note       string  `json:"note"`
			Amount     float64 `json:"amount"`
			PaymentURL string  `json:"payment_url"`
		} `json:"data"`
	}

	var paymentResp PaymentResponse
	err = json.Unmarshal(respBody, &paymentResp)
	if err != nil {
		return "", err
	}

	if !paymentResp.Status {
		return "", fmt.Errorf("payment initiation failed")
	}

	return paymentResp.Data.PaymentURL, nil
}

// get all payments from database
func (db Database) GetPayments(ctx *gin.Context) {

	// get payments
	payments, err := GetPayments(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payments)
}

// get payment by id
func (db Database) GetPaymentByID(ctx *gin.Context) {
	//get id value from path
	payment_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if exists := CheckPaymentExists(db.DB, uint(payment_id)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid bill id"})
		return
	}
	paymentId, err := GetPaymentByID(db.DB, uint(payment_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, paymentId)
}

// search payments from database
func (db Database) SearchPayments(ctx *gin.Context) {

	// init vars
	var payment Payment

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search payments from database
	payments, err := SearchPayments(db.DB, payment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payments)
}
