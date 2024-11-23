package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestIyzilink(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."
	convId := "123456789"
	img1b64 := ""
	img2b64 := ""

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	createReq := &iyzipay.IyzilinkCreateRequest{
		ConversationID:       convId,
		Locale:               "tr",
		Name:                 "Test Iyzilink",
		Price:                "123.1",
		Description:          "Test Iyzilink Description",
		InstallmentRequested: "false",
		AddressIgnorable:     "true",
		CurrencyCode:         "TRY",
		EncodedImageFile:     img1b64,
		ProductBuyerInfo: &iyzipay.ProductBuyerInfo{
			BuyerName:      "Test Buyer",
			BuyerSurname:   "Test Buyer Surname",
			BuyerCity:      "Istanbul",
			BuyerCountry:   "Turkey",
			BuyerGsmNumber: "5555555555",
			BuyerEmail:     "a@a.com",
			BuyerAddress:   "Test Buyer Address",
		},
		StockCount:   1,
		StockEnabled: true,
	}

	cResp, err := client.IyzilinkCreate(createReq)
	if err != nil {
		t.Errorf("Error create iyzilink: %v", err)
		return
	}

	if cResp.Status == "success" {
		t.Logf("Iyzilink Create Response: %v", cResp)
	} else {
		t.Errorf("Iyzilink Create Err Response: %v", cResp)
		return
	}

	detail1Req := &iyzipay.IyzilinkGetDetailRequest{
		ConversationID: convId,
		Locale:         "tr",
		Page:           1,
		Count:          10,
		Token:          "AAC8RQ",
	}

	det1Resp, err := client.IyzilinkGetDetail(detail1Req)
	if err != nil {
		t.Errorf("Error detail check: %v", err)
		return
	}

	if det1Resp.Status == "success" {
		t.Logf("Iyzilink Get Detail Response: %v", det1Resp)
	} else {
		t.Errorf("Iyzilink Get Detail Err Response: %v", det1Resp)
		return
	}

	statusUReq := &iyzipay.IyzilinkStatusUpdateRequest{
		ConversationID: convId,
		Locale:         "tr",
		Token:          "AAC8RQ",
		ProductStatus:  "PASSIVE",
	}

	suResp, err := client.IyzilinkStatusUpdate(statusUReq)
	if err != nil {
		t.Errorf("Error status update check: %v", err)
		return
	}

	if suResp.Status == "success" {
		t.Logf("Iyzilink Status Update Response: %v", suResp)
	} else {
		t.Errorf("Iyzilink Status Update Err Response: %v", suResp)
		return
	}

	updateReq := &iyzipay.IyzilinkUpdateRequest{
		ConversationID:       convId,
		Locale:               "tr",
		Name:                 "Test Iyzilink Updated",
		Price:                "123.1",
		InstallmentRequested: "false",
		SoldLimit:            "10",
		AddressIgnorable:     "true",
		CurrencyCode:         "TRY",
		EncodedImageFile:     img2b64,
		Description:          "Test Iyzilink Description Updated",
		Token:                "AAC8RQ",
	}

	uResp, err := client.IyzilinkUpdate(updateReq)
	if err != nil {
		t.Errorf("Error update check: %v", err)
		return
	}

	if uResp.Status == "success" {
		t.Logf("Iyzilink Update Response: %v", uResp)
	} else {
		t.Errorf("Iyzilink Update Err Response: %v", uResp)
		return
	}

	deleteReq := &iyzipay.IyzilinkDeleteRequest{
		ConversationID: convId,
		Locale:         "tr",
		Token:          "AAC8RQ",
	}

	dResp, err := client.IyzilinkDelete(deleteReq)
	if err != nil {
		t.Errorf("Error delete check: %v", err)
		return
	}

	if dResp.Status == "success" {
		t.Logf("Iyzilink Delete Response: %v", dResp)
	} else {
		t.Errorf("Iyzilink Delete Err Response: %v", dResp)
		return
	}

	c1Req := &iyzipay.IyzilinkCreateRequest{
		ConversationID:       convId,
		Locale:               "tr",
		Name:                 "Test Iyzilink 1",
		Price:                "123.1",
		Description:          "Test Iyzilink 1 Description",
		InstallmentRequested: "false",
		AddressIgnorable:     "true",
		CurrencyCode:         "TRY",
		EncodedImageFile:     img1b64,
		ProductBuyerInfo: &iyzipay.ProductBuyerInfo{
			BuyerName:      "Test Buyer 1",
			BuyerSurname:   "Test Buyer Surname 1",
			BuyerCity:      "Istanbul",
			BuyerCountry:   "Turkey",
			BuyerGsmNumber: "5555555555",
			BuyerEmail:     "a@a.com",
			BuyerAddress:   "Test Buyer Address",
		},
		StockCount:   1,
		StockEnabled: true,
	}

	c1Resp, err := client.IyzilinkCreate(c1Req)
	if err != nil {
		t.Errorf("Error iyzilink create: %v", err)
		return
	}

	if c1Resp.Status != "success" {
		t.Errorf("Iyzilink Create1 Err Response: %v", c1Resp)
		return
	}

	c2Req := &iyzipay.IyzilinkCreateRequest{
		ConversationID:       convId,
		Locale:               "tr",
		Name:                 "Test Iyzilink 2",
		Price:                "321.1",
		Description:          "Test Iyzilink Description 2",
		InstallmentRequested: "false",
		AddressIgnorable:     "true",
		CurrencyCode:         "TRY",
		EncodedImageFile:     img2b64,
		ProductBuyerInfo: &iyzipay.ProductBuyerInfo{
			BuyerName:      "Test Buyer 2",
			BuyerSurname:   "Test Buyer Surname 2",
			BuyerCity:      "Istanbul",
			BuyerCountry:   "Turkey",
			BuyerGsmNumber: "5555555555",
			BuyerEmail:     "a@a.com",
			BuyerAddress:   "Test Buyer Address",
		},
		StockCount:   10,
		StockEnabled: true,
	}

	c2Resp, err := client.IyzilinkCreate(c2Req)
	if err != nil {
		t.Errorf("Error create2 iyzilink: %v", err)
		return
	}

	if c2Resp.Status != "success" {
		t.Errorf("Iyzilink Create2 Err Response: %v", c2Resp)
		return
	}

	listReq := &iyzipay.IyzilinkGetListRequest{
		ConversationID: convId,
		Locale:         "tr",
		Page:           1,
		Count:          10,
	}

	lResp, err := client.IyzilinkGetList(listReq)
	if err != nil {
		t.Errorf("Error list check: %v", err)
		return
	}

	if lResp.Status == "success" {
		t.Logf("Iyzilink Get List Response: %v", lResp)
	} else {
		t.Errorf("Iyzilink Get List Err Response: %v", lResp)
		return
	}

}
