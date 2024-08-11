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

	req := &iyzipay.BinRequest{
		Locale:         "tr",
		BinNumber:      "454671",
		ConversationId: "123456789",
	}

	res, err := client.BinControlRequest(req)
	if err != nil {
		fmt.Println(err)
	}

	if res.Status == "success" {
		fmt.Println("Bin Number: ", res.CardAssociation) // VISA
		return
	}

	fmt.Println("Bin Control Request Failed")
}
