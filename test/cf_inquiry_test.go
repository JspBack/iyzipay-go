package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestCFPaymentInquiryRequest(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."
	token := "1"

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	request := &iyzipay.CFInquiryRequest{
		Locale:         "tr",
		Token:          token,
		ConversationId: "123456789",
	}

	response, err := client.CheckoutFormPaymentInquiryRequest(request)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	if response.Status != "success" {
		t.Errorf("Error getting response: %v", response.Status)
		return
	}

	if response.Token != token {
		t.Errorf("Error getting response: %v", response.Token)
		return
	}

	if response.PaymentStatus != "SUCCESS" {
		t.Errorf("Error getting response: %v", response.PaymentStatus)
		return
	}

	t.Logf("Response: %v", response)

}
