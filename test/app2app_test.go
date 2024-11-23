package test

import (
	"strconv"
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestApp2AppRequests(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."
	a2AapiKey := ""
	a2Asecret := ""
	merchantId := ""

	a2acl, err := iyzipay.NewApp2AppClient(apikey, secretkey, a2AapiKey, a2Asecret, merchantId)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	lsRsp, err := a2acl.App2AppListUsers()
	if err != nil {
		t.Errorf("Error listing users: %v", err)
		return
	}

	if lsRsp.Status == "success" {
		for _, user := range lsRsp.UserInfoList {
			t.Logf("User: %v", user)
		}
	}

	createReq := &iyzipay.A2ACreatePaymentRequest{
		Amount:        "1.0",
		Email:         "kullanici1@iyzico.com",
		PaymentSource: "IYZICO",
		CallBackUrl:   "https://callback.url",
	}

	createRsp, err := a2acl.App2CreatePayment(createReq)
	if err != nil {
		t.Errorf("Error creating payment: %v", err)
		return
	}

	if createRsp.Status == "success" {
		t.Logf("Payment created: %v", createRsp)
	}

	pId, err := strconv.Atoi(createRsp.PaymentId)
	if err != nil {
		t.Errorf("Error converting payment id: %v", err)
		return
	}

	refundReq := &iyzipay.A2ARefundPaymentRequest{
		RefundAmount: "1.0",
		PaymentId:    pId,
		Email:        "one of user's email",
		CallBackUrl:  "https://callback.url",
	}

	refundRsp, err := a2acl.App2RefundPayment(refundReq)
	if err != nil {
		t.Errorf("Error refunding payment: %v", err)
		return
	}

	if refundRsp.Status == "success" {
		t.Logf("Payment refunded: %v", refundRsp)
	}

	inquiryReq := &iyzipay.A2AInquiryPaymentRequest{
		Data:                "data",
		PaymentSessionToken: "token",
	}

	inquiryRsp, err := a2acl.App2InquiryPayment(inquiryReq)
	if err != nil {
		t.Errorf("Error inquiring payment: %v", err)
		return
	}

	if inquiryRsp.Status == "success" {
		t.Logf("Payment inquired: %v", inquiryRsp)
	}

	t.Logf("All tests passed")
}
