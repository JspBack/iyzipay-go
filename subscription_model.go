package iyzipay

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type CreateSubscriptionProductRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir. En yaygın kullanış biçimi üye iş yerinin ürün numarasıdır
	//
	// zorunlu değil
	ConversationId string `json:"conversationId" validate:"omitempty"`

	// Ürün adı. Eşsiz bir isim olmalıdır. Mevcut bir isim başka bir ürüne verilemez.
	//
	// zorunlu
	Name string `json:"name" validate:"required"`

	// Ürün açıklaması. Bu açıklama müşterilere gösterilebilir veya tarafınızda bir not olabilir.
	//
	// zorunlu değil
	Description string `json:"description" validate:"omitempty"`
}

type UpdateSubscriptionProductRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// Ürüne ait eşsiz referans kodu
	//
	// zorunlu
	ProductReferenceCode string `json:"productReferenceCode" validate:"required"`

	// Ürüne verilecek yeni isim.
	//
	// zorunlu
	Name string `json:"name" validate:"required"`

	// Ürüne verilecek yeni açıklama.
	Description string `json:"description" validate:"omitempty"`
}

type DeleteSubscriptionProductRequest struct {
	// Ürüne ait eşsiz referans kodu
	//
	// zorunlu
	ProductReferenceCode string `json:"productReferenceCode" validate:"required"`
}

type GetSubscriptionProductDetailRequest struct{}

type GetSubscriptionProductListRequest struct{}

type GetSubscriptionProductDetailResponse struct{}

type GetSubscriptionProductListResponse struct{}

type SubscriptionProductResponse struct {
	// Yapılan isteğin sonucunu belirtir. İşlem başarılı ise success değeri döner.
	Status string `json:"status"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int64 `json:"systemTime"`

	// Oluşturulan ürün bilgilerini içerirr
	Data Subscription `json:"data"`
}

type Subscription struct {
	// Oluşturulan ürüne ait eşsiz referans kodu. Ürün güncellemek veya silmek, ürün detayı görmek ve plan oluşturmak için kullanılır.
	ReferenceCode string `json:"referenceCode"`

	// Ürün oluşturulma tarihi.
	CreatedDate string `json:"createdDate"`

	// Ürün adı. İstek sırasında ürüne verdiğiniz isimdir.
	Name string `json:"name"`

	// Ürün adı. İstek sırasında ürüne verdiğiniz isimdir.
	Description string `json:"description"`

	// Yapılan isteğin sonucunu belirtir. İşlem başarılı ise success değeri döner.
	Status string `json:"status"`

	// Ürün'e bağlı planları gösterir. Ürün oluşturulduğu anda henüz bir plan olmadığı için boş gelmektedir.
	PricingPlans []PricingPlan `json:"pricingPlans"`
}

type PricingPlan struct {
	// Dökümantasyonda belirtilmemiş
	ReferenceCode string `json:"referenceCode"`

	// Dökümantasyonda belirtilmemiş
	CreatedDate string `json:"createdDate"`

	// Dökümantasyonda belirtilmemiş
	Name string `json:"name"`

	// Dökümantasyonda belirtilmemiş
	Price float64 `json:"price"`

	// Dökümantasyonda belirtilmemiş
	PaymentInterval string `json:"paymentInterval"`

	// Dökümantasyonda belirtilmemiş
	PaymentIntervalCount int `json:"paymentIntervalCount"`

	// Dökümantasyonda belirtilmemiş
	TrialPeriodDays int `json:"trialPeriodDays"`

	// Dökümantasyonda belirtilmemiş
	CurrencyCode string `json:"currencyCode"`

	// Dökümantasyonda belirtilmemiş
	ProductReferenceCode string `json:"productReferenceCode"`

	// Dökümantasyonda belirtilmemiş
	PlanPaymentType string `json:"planPaymentType"`

	// Yapılan isteğin sonucunu belirtir. İşlem başarılı ise success değeri döner.
	Status string `json:"status"`
}

func (r CreateSubscriptionProductRequest) validate() error {
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

func (r UpdateSubscriptionProductRequest) validate() error {
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

func (r DeleteSubscriptionProductRequest) validate() error {
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
