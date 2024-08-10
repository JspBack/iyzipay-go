package iyzipay

import "gopkg.in/go-playground/validator.v9"

type BinRequest struct {
	// Dil (default: tr).
	//
	// zorunlu değil.
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// Kartın ilk 6 hanesi.
	//
	// zorunlu.
	BinNumber string `json:"binNumber" validate:"required,numeric"`

	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	//
	// zorunlu değil.
	ConversationId string `json:"conversationId" validate:"omitempty"`
}

type BinResponse struct {
	// Kartın ilk 6 hanesi.
	BinNumber string `json:"binNumber"`

	// Kartın tipi.
	CardType string `json:"cardType"`

	// Kartın markası.
	CardAssociation string `json:"cardAssociation"`

	// Kartın ailesi.
	CardFamily string `json:"cardFamily"`

	// Banka adı.
	BankName string `json:"bankName"`

	// Banka kodu.
	BankCode int `json:"bankCode"`

	// ???
	Commercial int `json:"commercial"`

	// İstek sonucu.
	Status string `json:"status"`

	// İstek sonucun dilini belirten parametre.
	Locale string `json:"locale"`

	// Sistem zamanı ?
	SystemTime int `json:"systemTime"`

	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	ConversationId string `json:"conversationId"`
}

func (r BinRequest) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
