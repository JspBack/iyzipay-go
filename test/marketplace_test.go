package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestApproveProduct(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	request := &iyzipay.MarketplaceProductRequest{
		Locale:               "tr",
		ConversationId:       "123456789",
		PaymentTransactionId: "123456789",
	}

	response, err := client.ApproveProduct(*request)
	if err != nil {
		t.Errorf("Error approving product: %v", err)
		return
	}

	if response.Status != "success" {
		t.Errorf("Error creating payment: %v", response)
		return
	}

	t.Logf("ApproveProduct: %v", response)
}

func TestDisapproveProduct(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	request := &iyzipay.MarketplaceProductRequest{
		Locale:               "tr",
		ConversationId:       "123456789",
		PaymentTransactionId: "123456789",
	}

	response, err := client.DisapproveProduct(*request)
	if err != nil {
		t.Errorf("Error approving product: %v", err)
		return
	}

	if response.Status != "success" {
		t.Errorf("Error creating payment: %v", response)
		return
	}

	t.Logf("ApproveProduct: %v", response)
}
