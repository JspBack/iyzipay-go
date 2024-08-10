package main

import (
	"fmt"
	"os"

	"github.com/JspBack/iyzipay-go"
)

func main() {
	apikey := os.Getenv("IYZIPAY_API_KEY")
	apiSecret := os.Getenv("IYZIPAY_API_SECRET")
	PaymentId := "1"
	Ip := "85.34.78.112"

	client, err := iyzipay.New(apikey, apiSecret)
	if err != nil {
		panic(err)
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
		fmt.Println(err)
		return
	}

	if response.Status == "success" {
		fmt.Println("Payment Inquiry Response: ", response)
		return
	}

	panic("Payment Inquiry Request Failed")
}
