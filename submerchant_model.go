package iyzipay

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type IndividualSubmerchantRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir
	//
	// zorunlu değil
	ConversationId string `json:"conversationId" validate:"omitempty"`

	// Alt üye işyeri adı.
	//
	// zorunlu değil
	Name string `json:"name" validate:"omitempty"`

	// Alt üye işyeri e-posta adresi.
	//
	// zorunlu
	Email string `json:"email" validate:"required,email"`

	// Alt üye işyeri telefon numarası.
	//
	// zorunlu değil
	GsmNumber string `json:"gsmNumber" validate:"omitempty"`

	// Alt üye işyeri adresi.
	//
	// zorunlu
	Address string `json:"address" validate:"required"`

	// Alt üye işyeri IBAN bilgisi. contactName ve contactSurname ile uyumlu bir IBAN olmalı. Eğer alt üye işyeri ekleme esnasında boş bırakılırsa, ürüne onay vermeden önce mutlaka doldurulmalıdır.
	//
	// zorunlu değil
	Iban string `json:"iban" validate:"omitempty"`

	// Alt üye işyeri adı.
	//
	// zorunlu
	ContactName string `json:"contactName" validate:"required"`

	// Alt üye işyeri soyadı.
	//
	// zorunlu
	ContactSurname string `json:"contactSurname" validate:"required"`

	// Para birimi default TL olarak belirlenmiştir. USD, EUR, GBP, RUB, CHF ve NOK para birimleri de belirlenebilir
	//
	// zorunlu değil
	Currency string `json:"currency" validate:"omitempty,oneof=TRY USD EUR GBP RUB CHF NOK"`

	// Alt üye işyeri tekil dış ID’si, sizin sisteminizdeki ID olabilir.
	//
	// zorunlu
	SubmerchantExternalId string `json:"subMerchantExternalId" validate:"required"`

	// Alt üye işyeri TC Kimlik Numarası.
	//
	// zorunlu
	IdentityNumber string `json:"identityNumber" validate:"required"`

	// Bireysel için PERSONAL değeri gönderilmelidir.
	//
	// zorunlu
	SubMerchantType string `json:"subMerchantType" validate:"required,oneof=PERSONAL"`
}

type PrivateSubmerchantRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir
	//
	// zorunlu değil
	ConversationId string `json:"conversationId" validate:"omitempty"`

	// Alt üye işyeri adı.
	//
	// zorunlu değil
	Name string `json:"name" validate:"omitempty"`

	// Alt üye işyeri e-posta adresi.
	//
	// zorunlu
	Email string `json:"email" validate:"required,email"`

	// Alt üye işyeri telefon numarası.
	//
	// zorunlu değil
	GsmNumber string `json:"gsmNumber" validate:"omitempty"`

	// Alt üye işyeri adresi.
	//
	// zorunlu
	Address string `json:"address" validate:"required"`

	// Alt üye işyeri IBAN bilgisi. legalCompanyTitle ile uyumlu bir IBAN olmalı. Eğer alt üye işyeri ekleme esnasında boş bırakılırsa, ürüne onay vermeden önce mutlaka doldurulmalıdır.
	//
	// zorunlu değil
	Iban string `json:"iban" validate:"omitempty"`

	// Alt üye işyeri vergi dairesi.
	//
	// zorunlu
	TaxOffice string `json:"taxOffice" validate:"required"`

	// Alt üye işyeri yasal şirket ünvanı.
	//
	// zorunlu
	LegalCompanyTitle string `json:"legalCompanyTitle" validate:"required"`

	// Para birimi default TL olarak belirlenmiştir. USD, EUR, GBP, RUB, CHF ve NOK para birimleri de belirlenebilir
	//
	// zorunlu değil
	Currency string `json:"currency" validate:"omitempty,oneof=TRY USD EUR GBP RUB CHF NOK"`

	// Alt üye işyeri tekil dış ID’si, sizin sisteminizdeki ID olabilir
	//
	// zorunlu
	SubMerchantExternalId string `json:"subMerchantExternalId" validate:"required"`

	// Alt üye işyeri TC Kimlik Numarası.
	//
	// zorunlu değil
	IdentityNumber string `json:"identityNumber" validate:"omitempty"`

	// Şahıs Şirketi için PRIVATE_COMPANY değeri gönderilmelidir.
	//
	// zorunlu
	SubMerchantType string `json:"subMerchantType" validate:"required,oneof=PRIVATE_COMPANY"`
}

type LimitedCompanySubmerchantRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir
	//
	// zorunlu değil
	ConversationId string `json:"conversationId" validate:"omitempty"`

	// Alt üye işyeri adı.
	//
	// zorunlu değil
	Name string `json:"name" validate:"omitempty"`

	// Alt üye işyeri e-posta adresi.
	//
	// zorunlu
	Email string `json:"email" validate:"required,email"`

	// Alt üye işyeri telefon numarası.
	//
	// zorunlu değil
	GsmNumber string `json:"gsmNumber" validate:"omitempty"`

	// Alt üye işyeri adresi.
	//
	// zorunlu
	Address string `json:"address" validate:"required"`

	// Alt üye işyeri IBAN bilgisi. legalCompanyTitle ile uyumlu bir IBAN olmalı. Eğer alt üye işyeri ekleme esnasında boş bırakılırsa, ürüne onay vermeden önce mutlaka doldurulmalıdır.
	//
	// zorunlu değil
	Iban string `json:"iban" validate:"omitempty"`

	// Alt üye işyeri vergi dairesi.
	//
	// zorunlu
	TaxOffice string `json:"taxOffice" validate:"required"`

	// Alt üye işyeri vergi numarası.
	//
	// zorunlu değil
	TaxNumber string `json:"taxNumber" validate:"omitempty"`

	// Alt üye işyeri yasal şirket ünvanı.
	//
	// zorunlu
	LegalCompanyTitle string `json:"legalCompanyTitle" validate:"required"`

	// Para birimi default TL olarak belirlenmiştir. USD, EUR, GBP, RUB, CHF ve NOK para birimleri de belirlenebilir
	//
	// zorunlu değil
	Currency string `json:"currency" validate:"omitempty,oneof=TRY USD EUR GBP RUB CHF NOK"`

	// Alt üye işyeri tekil dış ID’si, sizin sisteminizdeki ID olabilir
	//
	// zorunlu
	SubMerchantExternalId string `json:"subMerchantExternalId" validate:"required"`

	// Limited veya Anonim Şirket için LIMITED_OR_JOINT_STOCK_COMPANY değeri gönderilmelidir.
	//
	// zorunlu
	SubMerchantType string `json:"subMerchantType" validate:"required,oneof=LIMITED_OR_JOINT_STOCK_COMPANY"`
}

type SubmerchantResponse struct {
	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri tr’dir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int64 `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationId string `json:"conversationId"`

	// Alt üye işyerini simgeleyen tekil değer.
	SubmerchantKey string `json:"submerchantKey"`
}

func (r *IndividualSubmerchantRequest) validate() error {
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

func (r *PrivateSubmerchantRequest) validate() error {
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

func (r *LimitedCompanySubmerchantRequest) validate() error {
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
