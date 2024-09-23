package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

// Error code:100001, Error message:Sistem hatası
//
// Dökümantasyonda YOK.
func TestCreateSubscriptionProduct(t *testing.T) {
	apikey = "sandbox-..."
	secretkey = ""
	client, _ = iyzipay.New(apikey, secretkey)

	req := &iyzipay.CreateSubscriptionProductRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Name:           "John's market",
		Description:    "Prime subscription",
	}

	response, err := client.CreatSubscriptionProduct(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	t.Logf("Response: %v", response)
}

// Error code:100001, Error message:Sistem hatası
//
// Dökümantasyonda YOK.
func TestUpdateSubscriptionProduct(t *testing.T) {
	apikey = "sandbox-H1pFBNawbh1ClEFAKRTc85CM6owLYL0V"
	secretkey = "IQho1I1Swd98yNbecmmCIjLJKN4zypHv"
	client, _ = iyzipay.New(apikey, secretkey)

	req := &iyzipay.UpdateSubscriptionProductRequest{
		Locale:               "tr",
		ProductReferenceCode: "S49222",
		Name:                 "Jane's market",
		Description:          "Prime subscription",
	}

	response, err := client.UpdateSubscriptionProduct(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	t.Logf("Response: %v", response)
}

// Error code:100001, Error message:Sistem hatası
//
// Dökümantasyonda YOK.
func TestDeleteSubscriptionProduct(t *testing.T) {
	apikey = "sandbox-H1pFBNawbh1ClEFAKRTc85CM6owLYL0V"
	secretkey = "IQho1I1Swd98yNbecmmCIjLJKN4zypHv"
	client, _ = iyzipay.New(apikey, secretkey)

	req := &iyzipay.DeleteSubscriptionProductRequest{
		ProductReferenceCode: "S49222",
	}

	response, err := client.DeleteSubscriptionProduct(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	t.Logf("Response: %v", response)
}
