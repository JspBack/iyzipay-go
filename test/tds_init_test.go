package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestInitTDSPaymentRequest(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	InitRequest := &iyzipay.InitTDSRequest{
		PaymentRequest: iyzipay.PaymentRequest{
			Locale:         "tr",
			ConversationID: "123456789",
			Price:          "1.0",
			PaidPrice:      "1.0",
			Installment:    1,
			PaymentChannel: "WEB",
			BasketID:       "B67832",
			PaymentGroup:   "PRODUCT",
			// PaymentCard: iyzipay.PaymentCard{
			// 	CardHolderName: "John Doe",
			// 	CardNumber:     "5528790000000008",
			// 	ExpireYear:     "2030",
			// 	ExpireMonth:    "12",
			// 	CVC:            "123",
			// 	RegisterCard:   0,
			// },
			PaymentCard: iyzipay.PaymentCard{
				CardUserKey: "a1083b27-fc26-805a-74ba-80c1460a9a0d",
				CardToken:   "1dbb182a-a5ea-9f6b-bb2e-39b8a5c15cd5",
			},
			Buyer: iyzipay.Buyer{
				ID:                  "BY789",
				Name:                "John",
				Surname:             "Doe",
				IdentityNumber:      "74300864791",
				Email:               "email@email.com",
				GSMNumber:           "+905350000000",
				RegistrationDate:    "2013-04-21 15:12:09",
				LastLoginDate:       "2015-10-05 12:43:35",
				RegistrationAddress: "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
				City:                "Istanbul",
				Country:             "Turkey",
				ZipCode:             "34732",
				IP:                  "85.34.78.112",
			},
			ShippingAddress: iyzipay.ShippingAddress{
				Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
				ContactName: "Jane Doe",
				City:        "Istanbul",
				Country:     "Turkey",
				ZipCode:     "34742",
			},
			BillingAddress: iyzipay.BillingAddress{
				Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
				ContactName: "Jane Doe",
				City:        "Istanbul",
				Country:     "Turkey",
				ZipCode:     "34742",
			},
			BasketItems: []iyzipay.BasketItem{
				{
					ID:        "BI101",
					Name:      "Binocular",
					Category1: "Collectibles",
					Category2: "Hobby",
					ItemType:  "PHYSICAL",
					Price:     "0.3",
				},
				{
					ID:        "BI102",
					Name:      "Game code",
					Category1: "Game",
					Category2: "Online Game Items",
					ItemType:  "VIRTUAL",
					Price:     "0.5",
				},
				{
					ID:        "BI103",
					Name:      "Usb",
					Category1: "Electronics",
					Category2: "Usb",
					ItemType:  "PHYSICAL",
					Price:     "0.2",
				},
			},
			Currency: "TRY",
		},
		CallbackUrl: "https://www.merchant.com/callback",
	}

	response, decodedHtmlContent, err := client.InitializeTDSPayment(InitRequest, iyzipay.WithBinRequest(true), iyzipay.WithHtmlDecodeRequest(true))
	if err != nil {
		t.Errorf("Error creating payment request: %v", err)
		return
	}

	t.Logf("Payment request created: %v", response)
	t.Logf("Decoded html content: %v", decodedHtmlContent)
	t.Logf("Conversation ID: %v", response.ConversationID)
	t.Logf("Payment ID: %v", response.PaymentID)
}
