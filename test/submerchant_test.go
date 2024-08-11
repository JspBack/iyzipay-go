package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

var (
	apikey    = "sandbox-..."
	secretkey = "sandbox-..."
	client, _ = iyzipay.New(apikey, secretkey)
)

func TestCreateIndividualSubmerchant(t *testing.T) {

	req := &iyzipay.IndividualSubmerchantRequest{
		Locale:                "tr",
		ConversationId:        "123456789",
		Name:                  "John's market",
		Email:                 "email@submerchantemail.com",
		GsmNumber:             "+905350000000",
		Address:               "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		Iban:                  "TR180006200119000006672315",
		ContactName:           "Jane",
		ContactSurname:        "Doe",
		Currency:              "TRY",
		SubmerchantExternalId: "S49222",
		IdentityNumber:        "31300864726",
		SubMerchantType:       "PERSONAL",
	}

	response, err := client.CreateIndividualSubmerchant(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	if response.Status != "success" {
		t.Errorf("Error getting response: %v", response.Status)
		return
	}

	t.Logf("Response: %v", response)
}

func TestCreatePrivateSubmerchant(t *testing.T) {

	req := &iyzipay.PrivateSubmerchantRequest{
		Locale:                "tr",
		ConversationId:        "123456789",
		Name:                  "Jane's market",
		Email:                 "email@submerchantemail.com",
		GsmNumber:             "+905350000000",
		Address:               "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		Iban:                  "TR180006200119000006672315",
		TaxOffice:             "Tax office",
		LegalCompanyTitle:     "Jane Doe inc",
		Currency:              "TRY",
		SubMerchantExternalId: "S49222",
		IdentityNumber:        "31300864726",
		SubMerchantType:       "PRIVATE_COMPANY",
	}

	response, err := client.CreatePrivateSubmerchant(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	if response.Status != "success" {
		t.Errorf("Error getting response: %v", response.Status)
		return
	}

	t.Logf("Response: %v", response)
}

func TestCreateLimitedCompanySubmerchant(t *testing.T) {
	req := &iyzipay.LimitedCompanySubmerchantRequest{
		Locale:                "tr",
		ConversationId:        "123456789",
		Name:                  "Jane's market",
		Email:                 "email@submerchantemail.com",
		GsmNumber:             "+905350000000",
		Address:               "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		Iban:                  "TR180006200119000006672315",
		TaxOffice:             "Tax office",
		TaxNumber:             "31300864726",
		LegalCompanyTitle:     "Jane Doe inc",
		Currency:              "TRY",
		SubMerchantExternalId: "S49222",
		SubMerchantType:       "LIMITED_OR_JOINT_STOCK_COMPANY",
	}

	response, err := client.CreateLimitedCompanySubmerchant(*req)
	if err != nil {
		t.Errorf("Error getting response: %v", err)
		return
	}

	if response.Status != "success" {
		t.Errorf("Error getting response: %v", response.Status)
		return
	}

	t.Logf("Response: %v", response)
}
