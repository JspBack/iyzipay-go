//go:build integration

package test

import (
	"context"
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestBinControlRequest(t *testing.T) {
	apikey, secretkey := sandboxCredentials(t)

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	binReq := &iyzipay.BinRequest{
		Locale:         "tr",
		BinNumber:      "454360",
		ConversationId: "123456789",
	}

	binResp, err := client.BinControlRequest(context.Background(), binReq)
	if err != nil {
		t.Errorf("Error bin check: %v", err)
		return
	}

	instReq := &iyzipay.InstallmentRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Price:          "123.1",
		// BinNumber:      "454360",
	}

	instResp, err := client.InstallmentControlRequest(context.Background(), instReq)
	if err != nil {
		t.Errorf("Error installment check: %v", err)
		return
	}

	t.Logf("Bin and Installment Response: %v\n%v", binResp, instResp)
}
