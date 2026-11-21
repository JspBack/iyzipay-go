//go:build integration

package test

import (
	"context"
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestCardStorageNewUser(t *testing.T) {
	apikey, secretkey := sandboxCredentials(t)

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	request := iyzipay.CardStorageNewUserRequest{
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
		// ExternalId: "external id",
	}

	response, err := client.CreateCardStorageNewUser(context.Background(), &request)
	if err != nil {
		t.Errorf("Error creating card: %v", err)
		return
	}

	t.Logf("Card Storage Response: %v", response)
}

func TestCardStorageExUser(t *testing.T) {
	apikey, secretkey := sandboxCredentials(t)

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	request := iyzipay.CardStorageExUserRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		CardAlias:      "testCard",
		CardUserKey:    "f1436803-789d-a7d4-0a15-9a3b14672aa5",
		Card: &iyzipay.CardStorageCard{
			ExpireYear:     "2030",
			ExpireMonth:    "12",
			CardNumber:     "5528790000000008",
			CardHolderName: "John Doe",
		},
		ExternalId: "12345",
	}

	response, err := client.CreateCardStorageExUser(context.Background(), &request)
	if err != nil {
		t.Errorf("Error creating card: %v", err)
		return
	}

	t.Logf("Card Storage Response: %v", response)
}

func TestCardStorageDeleteCard(t *testing.T) {
	apikey, secretkey := sandboxCredentials(t)
	cardUserKey, cardToken := sandboxStoredCard(t, "IYZIPAY_DELETE_CARD")

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	request := iyzipay.CardStorageDeleteRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		CardUserKey:    cardUserKey,
		CardToken:      cardToken,
	}

	response, err := client.DeleteCardStorageCard(context.Background(), &request)
	if err != nil {
		t.Errorf("Error deleting card: %v", err)
		return
	}

	t.Logf("Card Storage Response: %v", response)
}

func TestCardStorageRetrieveCard(t *testing.T) {
	apikey, secretkey := sandboxCredentials(t)

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	request := iyzipay.CardStorageRetrieveRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		CardUserKey:    "f1436803-789d-a7d4-0a15-9a3b14672aa5",
	}

	response, err := client.GetCardStorageCards(context.Background(), &request)
	if err != nil {
		t.Errorf("Error getting cards: %v", err)
		return
	}

	t.Logf("Card Storage Response: %v", response)
}
