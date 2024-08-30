package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestBinControlRequest(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	request := iyzipay.BinRequest{
		Locale:         "tr",
		BinNumber:      "454360",
		ConversationId: "123456789",
	}

	response, err := client.BinControlRequest(&request)
	if err != nil {
		t.Errorf("Error creating payment: %v", err)
		return
	}

	t.Logf("Bin Control Response: %v", response)
}
