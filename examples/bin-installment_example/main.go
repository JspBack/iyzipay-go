package main

import (
	"context"
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

	binReq := &iyzipay.BinRequest{
		Locale:         "tr",
		BinNumber:      "454671",
		ConversationId: "123456789",
	}

	binRes, err := client.BinControlRequest(context.Background(), binReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Card association:", binRes.CardAssociation)

	instReq := &iyzipay.InstallmentRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Price:          "123.1",
		BinNumber:      "454671", // yollamazsanız mevcut tüm taksit seçenekleri gelecektir.
	}

	instRes, err := client.InstallmentControlRequest(context.Background(), instReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Installment details:", instRes.InstallmentDetails)
}
