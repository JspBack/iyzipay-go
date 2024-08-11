package main

import (
	"fmt"
	"os"

	"github.com/JspBack/iyzipay-go"
)

func main() {
	apikey := os.Getenv("IYZIPAY_API_KEY")
	apiSecret := os.Getenv("IYZIPAY_API_SECRET")

	client, err := iyzipay.New(apikey, apiSecret)
	if err != nil {
		fmt.Println(err)
		return
	}

	req := &iyzipay.CFRequest{
		InitPWIRequest: iyzipay.InitPWIRequest{
			Locale:              "tr",
			ConversationID:      "123456789",
			Price:               "1.0",
			BasketId:            "B67832",
			PaymentGroup:        "PRODUCT",
			CallbackUrl:         "https://www.merchant.com/callback",
			Currency:            "TRY",
			PaidPrice:           "1.0",
			EnabledInstallments: []string{"2", "3", "6", "9"},
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
				IP:                  "85.34.78.112"},
			ShippingAddress: iyzipay.Address{
				Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
				ContactName: "Jane Doe",
				City:        "Istanbul",
				Country:     "Turkey",
				ZipCode:     "34742",
			},
			BillingAddress: iyzipay.Address{
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
		},
		PaymentSource: "SHOPIFY",
		// CardUserKey:  "cardUserKey", // Bu keyin ne yaptığını henüz çözemedim dökümantasyona da eklenmemiş.
	}

	resp, err := client.CheckoutFormPaymentRequest(req, "iframe")
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.Status != "success" {
		fmt.Println("Checkout Form Payment Request Failed, Response: ", resp)
		return
	}

	InquiryReq := &iyzipay.CFInquiryRequest{
		Locale:         "tr",
		Token:          resp.Token,
		ConversationId: "123456789",
	}

	InquiryResp, err := client.CheckoutFormPaymentInquiryRequest(InquiryReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	if InquiryResp.Status == "success" {
		fmt.Println("PaymentId: ", InquiryResp.PaymentID)
		return
	}

	fmt.Println("Checkout Form Payment Inquiry Request Failed")
}