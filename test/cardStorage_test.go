package test

import (
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestCardStorageNewUser(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

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

	response, err := client.CreateCardStorageNewUser(&request)
	if err != nil {
		t.Errorf("Error creating card: %v", err)
		return
	}

	t.Logf("Card Storage Response: %v", response)
}

func TestCardStorageExUser(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

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

	response, err := client.CreateCardStorageExUser(&request)
	if err != nil {
		t.Errorf("Error creating card: %v", err)
		return
	}

	t.Logf("Card Storage Response: %v", response)
}

func TestCardStorageDeleteCard(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

	client, err := iyzipay.New(apikey, secretkey)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
	}

	request := iyzipay.CardStorageDeleteRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		CardUserKey:    "f1436803-789d-a7d4-0a15-9a3b14672aa5",
		CardToken:      "d68ce209-35f1-5631-9481-e003a7c3c75b",
	}

	response, err := client.DeleteCardStorageCard(&request)
	if err != nil {
		t.Errorf("Error deleting card: %v", err)
		return
	}

	t.Logf("Card Storage Response: %v", response)
}

func TestCardStorageRetrieveCard(t *testing.T) {
	apikey := "sandbox-..."
	secretkey := "sandbox-..."

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

	response, err := client.GetCardStorageCards(&request)
	if err != nil {
		t.Errorf("Error getting cards: %v", err)
		return
	}

	t.Logf("Card Storage Response: %v", response)
}
