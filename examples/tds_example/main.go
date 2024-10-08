package main

import (
	"fmt"
	"os"

	"github.com/JspBack/iyzipay-go"
)

func main() {
	apikey := os.Getenv("IYZIPAY_API_KEY")
	apiSecret := os.Getenv("IYZIPAY_API_SECRET")
	conversationID := "123456789"

	client, err := iyzipay.New(
		apikey,
		apiSecret,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// If you have saved card, you can use CardUserKey and CardToken inside PaymentCard struct.
	req := &iyzipay.InitTDSRequest{
		PaymentRequest: iyzipay.PaymentRequest{
			Locale:         "tr",
			ConversationID: conversationID,
			Price:          "1.0",
			PaidPrice:      "1.0",
			Installment:    1,
			PaymentChannel: "WEB",
			BasketID:       "B67832",
			PaymentGroup:   "PRODUCT",
			PaymentCard: iyzipay.PaymentCard{
				CardHolderName: "John Doe",
				CardNumber:     "5528790000000008",
				ExpireYear:     "2030",
				ExpireMonth:    "12",
				CVC:            "123",
				RegisterCard:   0,
			},
			// PaymentCard: iyzipay.PaymentCard{
			// 	CardUserKey: "card user key",
			// 	CardToken:   "card token",
			// },
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

	res, decodedHtmlContent, err := client.InitializeTDSPayment(
		req,
		iyzipay.WithBinRequest(false),        // Bin kontrolü yapmak istemiyorsanız false yapabilirsiniz.
		iyzipay.WithHtmlDecodeRequest(false), // Html içeriğini decode etmek istemiyorsanız false yapabilirsiniz.
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	if res.Status == "success" {
		fmt.Println("PaymentId: ", res.PaymentID)
		fmt.Println("DecodedHtmlContent: ", decodedHtmlContent)
	}

	finalizeReq := &iyzipay.TDSPaymentRequest{
		Locale:                "tr",
		ConversationId:        "123456789",
		PaymentId:             res.PaymentID,
		PaymentConversationId: conversationID,
	}

	// Bu örneği direk çalıştırırsanız hata alırsınız. Öncelikle decodedHtmlContent'e gidip işlemi tamamlamanız gerekmektedir.
	finalizeRes, err := client.FinalizeTDSPayment(finalizeReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	if finalizeRes.Status == "success" {
		fmt.Println("Status: ", finalizeRes.Status)
		fmt.Println("Response: ", finalizeRes)
		return
	}

	fmt.Println("TDS Payment Request Failed")
}
