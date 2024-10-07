package iyzipay

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type CardStorageCard struct {
	// Saklanacak kartın son kullanma yılı.
	//
	// zorunlu
	ExpireYear string `json:"expireYear" validate:"required,numeric,len=4"`

	// Saklanacak kartın son kullanma ayı.
	//
	// zorunlu
	ExpireMonth string `json:"expireMonth" validate:"required,numeric,len=2,oneof=01 02 03 04 05 06 07 08 09 10 11 12"`

	// Saklanacak kartın kart numarası.
	//
	// zorunlu
	CardNumber string `json:"cardNumber" validate:"required,numeric"`

	// Saklanacak kartın sahibinin adı soyadı.
	//
	// zorunlu
	CardHolderName string `json:"cardHolderName" validate:"required"`
}

type CardStorageNewUserRequest struct {
	// Saklanacak karta verilecek isim
	//
	// zorunlu
	CardAlias string `json:"cardAlias" validate:"required"`

	// Saklanacak kart sahibinin email adresi.
	//
	// zorunlu
	Email string `json:"email" validate:"required,email"`

	Card *CardStorageCard `json:"card" validate:"dive"`

	// Saklanacak kart için özel Id.
	//
	// zorunlu değil.
	ExternalId string `json:"externalId" validate:"omitempty"`

	// Iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır
	//
	// zorunlu değil.
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir
	//
	// zorunlu değil.
	ConversationId string `json:"conversationId" validate:"omitempty"`
}

type CardStorageExUserRequest struct {
	// Saklanacak karta verilecek isim
	//
	// zorunlu
	CardAlias string `json:"cardAlias" validate:"required"`

	// Saklanacak kartın sahibinin kart saklama servisindeki kullanıcı kimliği.
	//
	// zorunlu
	CardUserKey string `json:"cardUserKey" validate:"required"`

	Card *CardStorageCard `json:"card" validate:"dive"`

	// Iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır
	//
	// zorunlu değil.
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir
	//
	// zorunlu değil.
	ConversationId string `json:"conversationId" validate:"omitempty"`
}

type CardStorageDeleteRequest struct {
	// Kart kaydedildikten sonra dönen saklı karda ait token.
	//
	// zorunlu
	CardToken string `json:"cardToken" validate:"required"`

	// Saklanacak kartın sahibinin kart saklama servisindeki kullanıcı kimliği.
	//
	// zorunlu
	CardUserKey string `json:"cardUserKey" validate:"required"`

	// Iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır
	//
	// zorunlu değil.
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir
	//
	// zorunlu değil.
	ConversationId string `json:"conversationId" validate:"omitempty"`
}

type CardStorageResponse struct {
	// Üye işyeri tarafındaki sepetteki ürüne ait tip. Geçerli enum değerler: PHYSICAL ve VIRTUAL.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri tr’dir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int64 `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletililir.
	ConversationID string `json:"conversationId"`

	// Eğer istekte verildiyse özel Id'yi çevirir
	ExternalId *string `json:"externalId"`

	// Saklanan kartın sahibinin kart saklama servisindeki kullanıcı kimliği
	CardUserKey string `json:"cardUserKey"`

	// Saklanan kartın sahibinin kart saklama servisindeki kartın tokeni
	CardToken string `json:"cardToken"`

	// Saklanana kartın adı
	CardAlias string `json:"cardAlias"`

	// Saklana kartın ilk 6 hanesi
	BinNumber string `json:"binNumber"`

	// Saklanan kartın son 4 hanesi
	LastFourDigits string `json:"lastFourDigits"`

	// Saklanan kartın ait olduğu kart tipi. Geçerli değerler: CREDIT_CARD, DEBIT_CARD, PREPAID_CARD
	CardType string `json:"cardType"`

	// Saklanan kartın ait olduğu kuruluş. Geçerli değerler: VISA, MASTER_CARD, AMERICAN_EXPRESS, TROY
	CardAssociation string `json:"cardAssociation"`

	// Saklanan kartın ait olduğu aile. Geçerli değerler: Bonus, Axess, World, Maximum, Paraf, CardFinans, Advantage
	CardFamily string `json:"cardFamily"`

	// Saklanan kartın ait olduğu bankayı temsil eden kod.
	CardBankCode int `json:"cardBankCode"`

	// Saklanan kartın ait olduğu banka adı
	CardBankName string `json:"cardBankName"`
}

type CardStorageNewUserResponse struct {
	CardStorageResponse

	// Oluşturulan kullanıcının emaili
	Email string `json:"email"`
}

type CardStorageExUserResponse struct {
	CardStorageResponse
}

type CardStorageDeleteResponse struct {
	// Üye işyeri tarafındaki sepetteki ürüne ait tip. Geçerli enum değerler: PHYSICAL ve VIRTUAL.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri tr’dir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int64 `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletililir.
	ConversationID string `json:"conversationId"`
}

func (r *CardStorageNewUserRequest) validate() error {
	validate := validator.New()

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

func (r *CardStorageExUserRequest) validate() error {
	validate := validator.New()

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

func (r *CardStorageDeleteRequest) validate() error {
	validate := validator.New()

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
