package iyzipay

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type RefundPaymentV1Request struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir
	//
	// zorunlu değil.
	Locale string `json:"locale,omitempty" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir.
	//
	// zorunlu değil.
	ConversationID string `json:"conversationId,omitempty" validate:"omitempty"`

	// İşlemin gönderildiği ip adresi.
	//
	// zorunlu
	Ip string `json:"ip,omitempty" validate:"omitempty,ip"`

	// İşlemin gönderildiği ip adresi.
	//
	// zorunlu
	Price string `json:"price" validate:"required"`

	// iyzico tarafından işleme verilen benzersiz ödeme kırılım numarası.
	//
	// zorunlu
	PaymentTransactionID string `json:"paymentTransactionId" validate:"required"`
}

type RefundPaymentV2Request struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir.
	Locale string `json:"locale,omitempty" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir.
	ConversationID string `json:"conversationId,omitempty" validate:"omitempty"`

	// İşlemin gönderildiği ip adresi.
	Ip string `json:"ip,omitempty" validate:"omitempty,ip"`

	// İade edilmek istenen tutar.
	Price string `json:"price" validate:"required"`

	// iyzico tarafından işleme verilen benzersiz ödeme numarası.
	PaymentID string `json:"paymentId" validate:"required"`
}

type CancelPaymentRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir.
	Locale string `json:"locale,omitempty" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir.
	ConversationID string `json:"conversationId,omitempty" validate:"omitempty"`

	// İşlemin gönderildiği ip adresi.
	Ip string `json:"ip,omitempty" validate:"omitempty,ip"`

	// iyzico tarafından işleme verilen benzersiz ödeme numarası.
	PaymentID string `json:"paymentId" validate:"required"`
}

type RefundPaymentResponse struct {
	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeridir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int64 `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationID string `json:"conversationId"`

	// İade alıcı referansı(ARN). Tüketicilerin iadeleri takip edebilecekleri bir değer.
	HostReference string `json:"hostReference"`

	// İade alıcı referansı(ARN). Tüketicilerin iadeleri takip edebilecekleri bir değer.
	Price string `json:"price"`

	// İptali yapılan ödeme para birimi.
	Currency string `json:"currency"`

	// İptali yapılan ödeme para birimi.
	PaymentId string `json:"paymentId"`

	// iyzico tarafından işleme verilen benzersiz ödeme kırılım numarası.
	PaymentTransactionID string `json:"paymentTransactionId"`
}

type CancelPaymentResponse struct {
	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeridir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int64 `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationID string `json:"conversationId"`

	// İade alıcı referansı(ARN). Tüketicilerin iadeleri takip edebilecekleri bir değer.
	Price string `json:"price"`

	// İptali yapılan ödeme para birimi.
	Currency string `json:"currency"`

	// İptali yapılan ödeme para birimi.
	PaymentId string `json:"paymentId"`
}

func (r *RefundPaymentV1Request) validate() error {
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

func (r *RefundPaymentV2Request) validate() error {
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

func (r *CancelPaymentRequest) validate() error {
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
