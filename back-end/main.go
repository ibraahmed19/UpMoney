package main

import (
	"fmt"
	"template_rest_api/api/v1/paymee"
	"template_rest_api/server"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Configurez votre point de terminaison de webhook pour écouter les notifications WhatsApp
	//r.POST("/whatsapp-webhook", chatbot.HandleWhatsAppWebhook)
	// Run server on port APP_PORT
	server.RunServer()

	// Create Paymee client
	client := &paymee.PaymeeClient{}
	client.APIKey = "135a85997bf19dc1a63424cc532039a86f57310f"

	// Create a new payment
	payment := paymee.Payment{
		Amount:             56.0,
		InsuranceProductID: 1,
		Note:               "Assurence voiture",
		FirstName:          "John",
		LastName:           "Doe",
		Email:              "john.doe@example.com",
		Phone:              "123456789",
		ReturnUrl:          "https://example.com/return",
		CancelUrl:          "https://example.com/cancel",
		WebhookUrl:         "https://example.com/webhook",
		OrderId:            "123",
	}

	// yetsab fi base de donnes

	paymentURL, err := paymee.InitiatePayment(client, payment)
	if err != nil {
		fmt.Println("Erreur lors de l'initialisation du paiement:", err)
		return
	}

	fmt.Println("URL de paiement:", paymentURL)
	// Run Gin server
	r.Run(":8080")
}
