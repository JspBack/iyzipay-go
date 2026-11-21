//go:build integration

package test

import (
	"context"
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestCheckPWIPaymentRequest(t *testing.T) {
	apikey, secretkey := sandboxCredentials(t)
	Token := "1"

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	CheckRequest := &iyzipay.PWIInquiryRequest{
		Locale:         "tr",
		ConversationID: "123456789",
		Token:          Token,
	}

	response, err := client.CheckPWIPaymentRequest(context.Background(), CheckRequest)
	if err != nil {
		t.Errorf("Error creating payment request: %v", err)
		return
	}

	t.Logf("Payment request created: %v", response)
}
