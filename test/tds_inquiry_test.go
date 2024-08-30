package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestInquiryRequest(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."
	PaymentId := "1"
	Ip := "85.34.78.112"

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	request := iyzipay.InquiryRequest{
		PaymentId:             PaymentId,
		Ip:                    Ip,
		ConversationID:        "123456789",
		Locale:                "tr",
		PaymentConversationID: "123456789",
	}

	response, err := client.PaymentInquiryRequest(&request)
	if err != nil {
		t.Errorf("Error creating payment: %v", err)
		return
	}

	t.Logf("Payment Inquiry Response: %v", response)
}
