package main

import (
	"fmt"
	"os"

	"github.com/JspBack/iyzipay-go"
)

func main() {
	apiKey := os.Getenv("IYZIPAY_API_KEY")
	apiSecret := os.Getenv("IYZIPAY_API_SECRET")

	client, err := iyzipay.New(apiKey, apiSecret)
	if err != nil {
		fmt.Println(err)
	}

	createNew := &iyzipay.CardStorageNewUserRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		CardAlias:      "testCard",
		Email:          "email@email.com",
		Card: &iyzipay.CardStorageCard{
			ExpireYear:     "2030",
			ExpireMonth:    "12",
			CardNumber:     "5528790000000008",
			CardHolderName: "John Doe",
		},
		ExternalId: "12345",
	}

	// Store card for a new user
	newres, err := client.CreateCardStorageNewUser(createNew)
	if err != nil {
		fmt.Println(err)
	}

	if newres.Status == "success" {
		fmt.Printf("CardUserKey: %s\n CardToken: %s\n", newres.CardUserKey, newres.CardToken)
	}

	createOld := &iyzipay.CardStorageExUserRequest{
		CardAlias:   "testCard",
		CardUserKey: newres.CardUserKey,
		Card: &iyzipay.CardStorageCard{
			ExpireYear:     "2030",
			ExpireMonth:    "12",
			CardNumber:     "5528790000000008",
			CardHolderName: "John Doe",
		},
		Locale:         "tr",
		ConversationId: "123456789",
		ExternalId:     "12345",
	}

	// Store card for an existing user
	exres, err := client.CreateCardStorageExUser(createOld)
	if err != nil {
		fmt.Println(err)
	}

	if exres.Status == "success" {
		fmt.Printf("CardUserKey: %s\n CardToken: %s\n", exres.CardUserKey, exres.CardToken)
	}

	deleteReq := &iyzipay.CardStorageDeleteRequest{
		CardToken:      exres.CardToken,
		CardUserKey:    exres.CardUserKey,
		Locale:         "tr",
		ConversationId: "123456789",
	}

	// Delete the stored card
	delres, err := client.DeleteCardStorageCard(deleteReq)
	if err != nil {
		fmt.Println(err)
	}

	if delres.Status == "success" {
		fmt.Println(" Newest Card Deleted")
	}

	retrieveReq := &iyzipay.CardStorageRetrieveRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		CardUserKey:    exres.CardUserKey,
	}

	// Retrieve stored cards
	retrieveres, err := client.GetCardStorageCards(retrieveReq)
	if err != nil {
		fmt.Println(err)
	}

	// In this case ,only first created card will be retrieved (if you run this code only once)
	if retrieveres.Status == "success" {
		fmt.Printf("Cards: %v\n", retrieveres.CardDetails)
		return
	}

	fmt.Println("Bin Control Request Failed")
}
