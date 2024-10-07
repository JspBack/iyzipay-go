package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestCreateIndividualSubMerchant(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	req := &iyzipay.IndividualSubMerchantRequest{
		IndividualSubMerchantUpdateRequest: iyzipay.IndividualSubMerchantUpdateRequest{
			Locale:                "tr",
			ConversationId:        "123456789",
			Name:                  "John's market",
			Email:                 "email@SubMerchantemail.com",
			GsmNumber:             "+905350000000",
			Address:               "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			Iban:                  "TR180006200119000006672315",
			ContactName:           "Jane",
			ContactSurname:        "Doe",
			Currency:              "TRY",
			SubMerchantExternalId: "S49222",
			IdentityNumber:        "31300864726",
		},
		SubMerchantType: "PERSONAL",
	}

	response, err := client.CreateIndividualSubMerchant(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	t.Logf("Response: %v", response)
}

func TestCreatePrivateSubMerchant(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	req := &iyzipay.PrivateSubMerchantRequest{
		PrivateSubMerchantUpdateRequest: iyzipay.PrivateSubMerchantUpdateRequest{
			Locale:                "tr",
			ConversationId:        "123456789",
			Name:                  "Jane's market",
			Email:                 "email@SubMerchantemail.com",
			GsmNumber:             "+905350000000",
			Address:               "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			Iban:                  "TR180006200119000006672315",
			TaxOffice:             "Tax office",
			LegalCompanyTitle:     "Jane Doe inc",
			Currency:              "TRY",
			SubMerchantExternalId: "S49222",
			IdentityNumber:        "31300864726",
		},
		SubMerchantType: "PRIVATE_COMPANY",
	}

	response, err := client.CreatePrivateSubMerchant(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	t.Logf("Response: %v", response)
}

func TestCreateLimitedCompanySubMerchant(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	req := &iyzipay.LimitedCompanySubMerchantRequest{
		LimitedCompanySubMerchantUpdateRequest: iyzipay.LimitedCompanySubMerchantUpdateRequest{},
		SubMerchantType:                        "LIMITED_OR_JOINT_STOCK_COMPANY",
	}

	response, err := client.CreateLimitedCompanySubMerchant(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	t.Logf("Response: %v", response)
}

func TestUpdateIndividualSubMerchant(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	req := &iyzipay.IndividualSubMerchantUpdateRequest{
		Locale:                "tr",
		ConversationId:        "123456789",
		Name:                  "John's market",
		Email:                 "email@SubMerchantemail.com",
		GsmNumber:             "+905350000000",
		Address:               "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		Iban:                  "TR180006200119000006672315",
		ContactName:           "Jane",
		ContactSurname:        "Doe",
		Currency:              "TRY",
		SubMerchantExternalId: "S49222",
		IdentityNumber:        "31300864726",
	}

	response, err := client.UpdateIndividualSubMerchant(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	t.Logf("Response: %v", response)
}

func TestUpdatePrivateSubMerchant(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	req := &iyzipay.PrivateSubMerchantUpdateRequest{
		Locale:                "tr",
		ConversationId:        "123456789",
		Name:                  "Jane's market",
		Email:                 "email@SubMerchantemail.com",
		GsmNumber:             "+905350000000",
		Address:               "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		Iban:                  "TR180006200119000006672315",
		TaxOffice:             "Tax office",
		LegalCompanyTitle:     "Jane Doe inc",
		Currency:              "TRY",
		SubMerchantExternalId: "S49222",
		IdentityNumber:        "31300864726",
	}

	response, err := client.UpdatePrivateSubMerchant(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	t.Logf("Response: %v", response)
}

func TestUpdateLimitedCompanySubMerchant(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	req := &iyzipay.LimitedCompanySubMerchantUpdateRequest{
		Locale:                "tr",
		ConversationId:        "123456789",
		Name:                  "Jane's market",
		Email:                 "email@SubMerchantemail.com",
		GsmNumber:             "+905350000000",
		Address:               "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		Iban:                  "TR180006200119000006672315",
		TaxOffice:             "Tax office",
		TaxNumber:             "31300864726",
		LegalCompanyTitle:     "Jane Doe inc",
		Currency:              "TRY",
		SubMerchantExternalId: "S49222",
	}

	response, err := client.UpdateLimitedCompanySubMerchant(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	t.Logf("Response: %v", response)
}

func TestSubMerchantInquiry(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	req := &iyzipay.SubMerchantInquiryRequest{
		Locale:                "tr",
		ConversationID:        "123456789",
		SubMerchantExternalId: "S49222",
	}

	response, err := client.SubMerchantInquiry(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	t.Logf("Response: %v", response)
}

func TestUpdateSubMerchantProduct(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	req := &iyzipay.SubMerchantProductUpdateRequest{
		Locale:               "tr",
		ConversationId:       "123456789",
		PaymentTransactionId: "1",
		SubMerchantPrice:     "1",
		SubMerchantKey:       "1",
	}

	response, err := client.UpdateSubMerchantProduct(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	t.Logf("Response: %v", response)
}
