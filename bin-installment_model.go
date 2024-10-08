package iyzipay

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type BinRequest struct {
	// Dil (default: tr).
	//
	// zorunlu değil.
	Locale string `json:"locale,omitempty" validate:"omitempty,oneof=tr en"`

	// Kartın ilk 6 hanesi.
	//
	// zorunlu değil
	BinNumber string `json:"binNumber,omitempty" validate:"omitempty,numeric"`

	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	//
	// zorunlu değil.
	ConversationId string `json:"conversationId,omitempty" validate:"omitempty"`
}

type InstallmentRequest struct {
	// Dil (default: tr).
	//
	// zorunlu değil.
	Locale string `json:"locale,omitempty" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir.
	//
	// zorunlu değil.
	ConversationId string `json:"conversationId,omitempty" validate:"omitempty"`

	// Taksite bölünecek tutar.
	//
	// zorunlu.
	Price string `json:"price" validate:"required,numeric"`

	// Kartın ilk 6 hanesi.
	//
	// zorunlu değil.
	BinNumber string `json:"binNumber,omitempty" validate:"omitempty"`
}

type BinResponse struct {
	// İşlem sonucu başarılı ise success, başarısız ise failure döner.
	Status string `json:"status"`

	// İstek sonucun dilini belirten parametre.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int `json:"systemTime"`

	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	ConversationId string `json:"conversationId"`

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
}

type InstallmentResponse struct {
	// İşlem sonucu başarılı ise success, başarısız ise failure döner.
	Status string `json:"status"`

	// İstek sonucun dilini belirten parametre.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int `json:"systemTime"`

	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	ConversationId string `json:"conversationId"`

	// Taksit bilgileri.
	InstallmentDetails []InstallmentDetail `json:"installmentDetails"`
}

type InstallmentDetail struct {
	// Kartın ilk 6 hanesi.
	BinNumber string `json:"binNumber"`

	// Toplam tutar.
	Price float32 `json:"price"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu tipi. Geçerli değerler: CREDIT_CARD, DEBIT_CARD, PREPAID_CARD
	CardType string `json:"cardType"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu kuruluş. Geçerli değerler: TROY, VISA, MASTER_CARD, AMERICAN_EXPRESS.
	CardAssociation string `json:"cardAssociation"`

	// Kartın ailesi.
	CardFamilyName string `json:"cardFamilyName"`

	// İşlemin 3ds yapılmasına gerek olup olmadığını gösterir. 1 veya 0 değerlerini alır. 1 ise işlem 3ds ile yapılmalıdır.
	Force3ds int `json:"force3ds"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu banka adı.
	BankName string `json:"bankName"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu banka kodu.
	BankCode int `json:"bankCode"`

	// ???
	ForceCvc int `json:"forceCvc"`

	// Taksitler
	InstallmentPrices []InstallmentPrice `json:"installmentPrices"`
}

type InstallmentPrice struct {
	// Taksit başına düşen tutar.
	InstallmentPrice float32 `json:"installmentPrice"`

	// Toplam taksitli tutar.
	TotalPrice float32 `json:"totalPrice"`

	// Taksit sayısı.
	InstallmentNumber int `json:"installmentNumber"`
}

func (r BinRequest) validate() error {
	validate := validator.New()

	// Eğer doğrulamada hata oluşursa hata mesajını döndürür.
	err := validate.Struct(r)
	if err != nil {
		var errs []string
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()
			errs = append(errs, fmt.Sprintf("Validation error at field '%s' with tag '%s'", field, tag))
		}
		return fmt.Errorf("validation errors: %s", strings.Join(errs, ", "))
	}

	return nil
}

func (r InstallmentRequest) validate() error {
	validate := validator.New()

	// Eğer doğrulamada hata oluşursa hata mesajını döndürür.
	err := validate.Struct(r)
	if err != nil {
		var errs []string
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()
			errs = append(errs, fmt.Sprintf("Validation error at field '%s' with tag '%s'", field, tag))
		}
		return fmt.Errorf("validation errors: %s", strings.Join(errs, ", "))
	}

	return nil
}
