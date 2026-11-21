//go:build integration

package test

import (
	"context"
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestFinalizeTDSPaymentRequest(t *testing.T) {
	apikey, secretkey := sandboxCredentials(t)
	PaymentID := "1"
	ConversationID := "123456789"

	client, err := iyzipay.New(apikey, secretkey)
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

	response, err := client.FinalizeTDSPayment(context.Background(), FinalizeRequest)
	if err != nil {
		t.Errorf("Error creating payment request: %v", err)
		return
	}

	t.Logf("Payment request created: %v", response)
}
