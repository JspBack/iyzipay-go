package iyzipay

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type UpdateSubscriberRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir
	//
	// zorunlu değil.
	Locale string `json:"locale,omitempty" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir. En yaygın kullanış biçimi üye iş yerinin sipariş numarasıdır
	//
	// zorunlu değil.
	ConversationID string `json:"conversationId,omitempty" validate:"omitempty"`

	// Abone adı
	//
	// zorunlu
	Name string `json:"name" validate:"required"`

	// Abone soyadı
	//
	// zorunlu
	Surname string `json:"surname" validate:"required"`

	// Abone e-posta adresi
	//
	// zorunlu
	CustomerReferenceCode string `json:"customerReferenceCode" validate:"required"`

	// Abone fatura adresi
	//
	// zorunlu
	BillingAddress BillingAddress `json:"billingAddress" validate:"required,dive"`

	// Abone kimlik numarası
	//
	// zorunlu.
	IdentityNumber string `json:"identityNumber" validate:"required"`

	// Abone kargo adresi
	//
	// zorunlu değil.
	ShippingAddress ShippingAddress `json:"shippingAddress,omitempty" validate:"omitempty"`
}

type GetSubscriberDetailRequest struct {
	// Abone referans kodu
	CustomerReferenceCode string `json:"customerReferenceCode" validate:"required"`
}

type GetSubscriberListRequest struct {
	// Belirtilen sayfa için tüm aboneleri getirir.
	//
	// zorunlu
	Page int `json:"page" validate:"required"`

	// Bir sayfada kaç adet abone getirileceğini belirler.
	//
	// zorunlu
	Count int `json:"count" validate:"required"`
}

type SubscriberResponse struct {
	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner.
	Status string `json:"status"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int64 `json:"systemTime"`

	// Abone bilgileri
	Data SubscriberData `json:"data"`
}

type GetSubscriberListRequestResponse struct {
	// Yapılan isteğin sonucunu belirtir. İşlem başarılı ise success değeri döner.
	Status string `json:"status"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int64 `json:"systemTime"`

	// Abone bilgileri
	Data PaginatedSubscriberData `json:"data"`
}

type SubscriberData struct {
	// Abone referans kodu
	ReferenceCode string `json:"referenceCode"`

	// Abone oluşturulma tarihi
	CreatedDate int `json:"createdDate"`

	// Abone güncelleme tarihi
	Status string `json:"status"`

	// Abone adı
	Name string `json:"name"`

	// Abone soyadı
	Surname string `json:"surname"`

	// Abone kimlik numarası
	IdentityNumber string `json:"identityNumber"`

	// Abone e-posta adresi
	Email string `json:"email"`

	// Abone telefon numarası
	GsmNumber string `json:"gsmNumber"`

	// Abone fatura adresi
	ContactEmail string `json:"contactEmail"`

	// Abone telefon numarası
	ContactGsmNumber string `json:"contactGsmNumber"`

	// Abone fatura adresi
	BillingAddress BillingAddress `json:"billingAddress"`

	// Abone kargo adresi
	ShippingAddress *ShippingAddress `json:"shippingAddress"`
}

type PaginatedSubscriberData struct {
	// Toplam kaç tane sonuç geldiğini belirtir.
	TotalCount int `json:"totalCount"`

	// Hangi sayfa için listeleme yapıldığını belirtir.
	CurrentPage int `json:"currentPage"`

	// Toplam kaç sayfa sonuç geldiğini belirtir.
	PageCount int `json:"pageCount"`

	// Abone detayları
	Items []SubscriberData `json:"items"`
}

func (r *UpdateSubscriberRequest) validate() error {
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

func (r *GetSubscriberDetailRequest) validate() error {
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

func (r *GetSubscriberListRequest) validate() error {
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
