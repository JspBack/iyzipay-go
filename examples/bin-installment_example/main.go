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
	}

	binReq := &iyzipay.BinRequest{
		Locale:         "tr",
		BinNumber:      "454671",
		ConversationId: "123456789",
	}

	binRes, err := client.BinControlRequest(binReq)
	if err != nil {
		fmt.Println(err)
	}

	if binRes.Status == "success" {
		fmt.Println("Bin Number: ", binRes.CardAssociation) // VISA
		return
	}

	instReq := &iyzipay.InstallmentRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Price:          "123.1",
		BinNumber:      "454671", // yollamazsanız mevcut tüm taksit seçenekleri gelecektir.
	}

	instRes, err := client.InstallmentControlRequest(instReq)
	if err != nil {
		fmt.Println(err)
	}

	if instRes.Status == "success" {
		fmt.Println("Installment Response: ", instRes.InstallmentDetails)
		return
	}

	fmt.Println("Bin Control Request Failed")
}
