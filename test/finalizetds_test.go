package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestFinalizeTDSPaymentRequest(t *testing.T) {
	ApiKey := "sandbox-..."
	SecretKey := "sandbox-..."
	PaymentID := "1"
	ConversationID := "123456789"

	client, err := iyzipay.New(ApiKey, SecretKey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	FinalizeRequest := &iyzipay.TDSPaymentRequest{
		Locale:                "tr",
		ConversationId:        "123456789",
		PaymentId:             PaymentID,
		PaymentConversationId: ConversationID,
	}

	response, err := client.FinalizeTDSPayment(*FinalizeRequest)
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
