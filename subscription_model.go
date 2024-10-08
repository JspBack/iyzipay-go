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

type GetSubscriptionProductDetailRequest struct {
	// Ürüne ait eşsiz referans kodu
	//
	// zorunlu
	ProductReferenceCode string `json:"productReferenceCode" validate:"required"`
}

type GetSubscriptionProductListRequest struct {
	// Belirtilen sayfa için tüm ürünleri getirir
	//
	// zorunlu
	Page int `json:"page" validate:"required"`

	// Sayfa başına kaç ürün listeleneceğini belirtir
	//
	// zorunlu
	Count int `json:"count" validate:"required"`
}

type CreateSubscriptionPlanRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir. En yaygın kullanış biçimi üye iş yerinin ürün numarasıdır
	//
	// zorunlu değil
	ConversationId string `json:"conversationId" validate:"omitempty"`

	// Plan adı. Ödeme planı ile ilgili kısa bir bilgilendirme olacak şekilde değer girilmelidir
	//
	// zorunlu
	Name string `json:"name" validate:"required"`

	// Planın ilişkilendirileceği ürünün referans kodu.
	//
	// zorunlu
	ProductReferenceCode string `json:"productReferenceCode" validate:"required"`

	// Ödeme periyodunun kaç kez tekrarlanacağını belirler. Örneğin ayda bir ödeme alan planınıza bu değeri 12 olarak girerseniz, 12 ay boyunca ödeme alınacaktır.
	//
	// zorunlu değil
	RecurrenceCount int `json:"recurrenceCount" validate:"omitempty"`

	// Abonelik tipini belirtir. RECURRING değeri girilmelidir.
	//
	// zorunlu
	PlanPaymentType string `json:"planPaymentType" validate:"required,oneof=RECURRING"`

	// Deneme süresi veya ücretsiz kullanım süresi olarak kullanılır. Bu değer girilirse, ilk ödeme girilen gün sayısı sonunda alınır.
	//
	// zorunlu değil
	TrialPeriodDays int `json:"trialPeriodDays" validate:"omitempty"`

	// Ödeme periodunun hangi sıklıkta olacağını belirler. Örneğin paymentInterval değeri WEEKLY, paymentIntervalCount değeri 2 olursa, ödemeler 2 hafta bir alınır.
	//
	// zorunlu
	PaymentIntervalCount int `json:"paymentIntervalCount" validate:"required"`

	// Tekrarlı ödemenin alınacağı periodu belirler. DAILY, WEEKLY, MONTHLY, YEARLY değerlerini alabilir.
	//
	// zorunlu
	PaymentInterval string `json:"paymentInterval" validate:"required,oneof=DAILY WEEKLY MONTHLY YEARLY"`

	// Ödemenin alınacağı para birimi. TL, USD, EUR olabilir. TL dışındaki para birimlerinin hesabınıza tanımlandığından emin olunuz.
	//
	// zorunlu
	CurrencyCode string `json:"currencyCode" validate:"required,oneof=TL USD EUR"`

	// Ödeme periyotlarında karttan çekilecek tutar. Plan için geçerli abonelik fiyatı.
	//
	// zorunlu
	Price float64 `json:"price" validate:"required,gt=0"`
}

type UpdateSubscriptionPlanRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir.
	//
	// zorunlu değil
	ConversationId string `json:"conversationId" validate:"omitempty"`

	// Plan adı. Ödeme planı ile ilgili kısa bir bilgilendirme olacak şekilde değer girilmelidir.
	//
	// zorunlu
	Name string `json:"name" validate:"required"`

	// Plana ait eşsiz referans kodu. Plan güncellemek veya silmek, plan detayını görmek, abonelik başlatmak için kullanılır
	//
	// zorunlu
	PricingPlanReferenceCode string `json:"pricingPlanReferenceCode" validate:"required"`

	// Deneme süresi veya ücretsiz kullanım süresi olarak kullanılır. Bu değer girilirse, ilk ödeme girilen gün sayısı sonunda alınır.
	//
	// zorunlu değil
	TrialPeriodDays int `json:"trialPeriodDays" validate:"omitempty"`
}

type DeleteSubscriptionPlanRequest struct {
	// Plana ait referans kodu
	//
	// zorunlu
	PricingPlanReferenceCode string `json:"pricingPlanReferenceCode" validate:"required"`
}

type GetSubscriptionPlanDetailRequest struct {
	// Plana ait referans kodu
	//
	// zorunlu
	PricingPlanReferenceCode string `json:"pricingPlanReferenceCode" validate:"required"`
}

type GetSubscriptionPlanListRequest struct {
	// Belirtilen sayfa için tüm planları getirir.
	//
	// zorunlu
	Page int `json:"page" validate:"required"`

	// Sayfa başına kaç plan listelenceğini belirtir.
	//
	// zorunlu
	Count int `json:"count" validate:"required"`
}

type SubscriptionPlanResponse struct {
	// Yapılan isteğin sonucunu belirtir. İşlem başarılı ise success değeri döner.
	Status string `json:"status"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int64 `json:"systemTime"`

	// Ürün bilgilerini içerirr
	Data SubscriptionPlan `json:"data"`
}

type DeleteSubscriptionPlanResponse struct {
	// Yapılan isteğin sonucunu belirtir. İşlem başarılı ise success değeri döner.
	Status string `json:"status"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int64 `json:"systemTime"`
}

type GetSubscriptionPlanListResponse struct {
	// Yapılan isteğin sonucunu belirtir. İşlem başarılı ise success değeri döner.
	Status string `json:"status"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int64 `json:"systemTime"`

	Data PaginatedSubscriptionPlan `json:"data"`
}

type GetSubscriptionProductListResponse struct {
	// Yapılan isteğin sonucunu belirtir. İşlem başarılı ise success değeri döner.
	Status string `json:"status"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int64 `json:"systemTime"`

	Data PaginatedSubscription `json:"data"`
}

type SubscriptionProductResponse struct {
	// Yapılan isteğin sonucunu belirtir. İşlem başarılı ise success değeri döner.
	Status string `json:"status"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int64 `json:"systemTime"`

	// Ürün bilgilerini içerirr
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

type SubscriptionPlan struct {
	// Oluşturulan ürüne ait eşsiz referans kodu. Ürün güncellemek veya silmek, ürün detayı görmek ve plan oluşturmak için kullanılır.
	ReferenceCode string `json:"referenceCode"`

	// Ürün oluşturulma tarihi.
	CreatedDate string `json:"createdDate"`

	// Ürün adı. İstek sırasında ürüne verdiğiniz isimdir.
	Name string `json:"name"`

	// İstekte gönderdiğiniz değer geri iletilir. Ödeme periyotlarında karttan çekilecek tutardır.
	Price float64 `json:"price"`

	// İstekte gönderdiğiniz değer geri iletilir. Tekrarlı ödemenin alınacağı periodu belirler.
	PaymentInterval string `json:"paymentInterval"`

	// İstekte gönderdiğiniz değer geri iletilir. Ödeme periodunun hangi sıklıkta olacağını belirler.
	PaymentIntervalCount int `json:"paymentIntervalCount"`

	// İstekte gönderdiğiniz değer geri iletilir. Bu değer gönderilmediyse 0 olarak döner.
	TrialPeriodDays int `json:"trialPeriodDays"`

	// İstekte gönderdiğiniz değer geri iletilir. Ödemenin alınacağı para birimi.
	CurrencyCode string `json:"currencyCode"`

	// Plana bağlı ürünün referans kodu. İstekte gönderdiğiniz ürün bilgisi geri iletilir.
	ProductReferenceCode string `json:"productReferenceCode"`

	// RECURRING değerini döner ve ödemenin tekrarlı olduğunu belirtir
	PlanPaymentType string `json:"planPaymentType"`

	// Planın durumunu gösterir. ACTIVE değerini alır.
	Status string `json:"status"`

	// İstekte gönderdiğiniz değer geri iletilir. Ödeme periyodunun kaç kez tekrarlanacağını belirler.
	RecurrenceCount *int `json:"recurrenceCount"`
}

type PaginatedSubscription struct {
	// Toplam kaç tane sonuç geldiğini belirtir.
	TotalCount int `json:"totalCount"`

	// Hangi sayfa için listeleme yapıldığını belirtir.
	CurrentPage int `json:"currentPage"`

	// Toplam kaç sayfa sonuç geldiğini belirtir.
	PageCount int `json:"pageCount"`

	// Ürün detayları
	Items []Subscription `json:"items"`
}

type PaginatedSubscriptionPlan struct {
	// Toplam kaç tane sonuç geldiğini belirtir.
	TotalCount int `json:"totalCount"`

	// Hangi sayfa için listeleme yapıldığını belirtir.
	CurrentPage int `json:"currentPage"`

	// Toplam kaç sayfa sonuç geldiğini belirtir.
	PageCount int `json:"pageCount"`

	// Plan detayları
	Items []PricingPlan `json:"items"`
}

func (r *CreateSubscriptionProductRequest) validate() error {
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

func (r *UpdateSubscriptionProductRequest) validate() error {
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

func (r *DeleteSubscriptionProductRequest) validate() error {
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

func (r *GetSubscriptionProductDetailRequest) validate() error {
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

func (r *GetSubscriptionProductListRequest) validate() error {
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

func (r *CreateSubscriptionPlanRequest) validate() error {
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

func (r *UpdateSubscriptionPlanRequest) validate() error {
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

func (r *DeleteSubscriptionPlanRequest) validate() error {
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

func (r *GetSubscriptionPlanDetailRequest) validate() error {
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

func (r *GetSubscriptionPlanListRequest) validate() error {
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
