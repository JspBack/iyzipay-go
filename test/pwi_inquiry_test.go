package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestCheckPWIPaymentRequest(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."
	Token := "1"

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	CheckRequest := &iyzipay.CheckPWIRequest{
		Locale:         "tr",
		ConversationID: "123456789",
		Token:          Token,
	}

	response, err := client.CheckPWIPaymentRequest(CheckRequest)
	if err != nil {
		t.Errorf("Error creating payment request: %v", err)
		return
	}

	if response.Status != "success" {
		t.Errorf("Error creating payment request: %v", response)
		return
	}

	t.Logf("Payment request created: %v", response)
}
