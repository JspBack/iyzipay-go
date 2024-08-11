package iyzipay

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type MarketplaceProductRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşmesi yapmak için kullanılabilir.
	//
	// zorunlu değil
	ConversationId string `json:"conversationId" validate:"omitempty"`

	// Ödeme kırılımına ait ürünün iyzico taradından belirlenmiş unique id'si.
	//
	// zorunlu
	PaymentTransactionId string `json:"paymentTransactionId" validate:"required"`
}

type MarketplaceProductResponse struct {
	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri tr'dir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime string `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir
	ConversationId string `json:"conversationId"`

	// Ödeme kırılımına ait ürünün iyzico taradından belirlenmiş unique id'si.
	PaymentTransactionId string `json:"paymentTransactionId"`
}

func (r MarketplaceProductRequest) validate() error {
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
