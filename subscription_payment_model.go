package iyzipay

import (
	"fmt"
	"net/url"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type InitSubscriptionWithCheckoutFormRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir
	//
	// zorunlu değil.
	Locale string `json:"locale,omitempty" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir. En yaygın kullanış biçimi üye iş yerinin sipariş numarasıdır
	//
	// zorunlu değil.
	ConversationId string `json:"conversationId,omitempty" validate:"omitempty"`

	// İşlem sonucunda yönlendirilecek sayfa adresidir.(https olmalıdır.)
	//
	// zorunlu.
	CallbackUrl string `json:"callbackUrl" validate:"required,httpsurl"`

	// Eğer abonelik fiziksel bir ürün ise gönderilmelidir.
	//
	// zorunlu değil.
	ShippingAddress ShippingAddress `json:"shippingAddress,omitempty" validate:"omitempty"`

	// Fatura adresi bilgileri
	//
	// zorunlu.
	BillingAddress BillingAddress `json:"billingAddress" validate:"required,dive"`

	// Abone kimlik bilgileri
	//
	// zorunlu.
	IdentityNumber string `json:"identityNumber" validate:"required"`

	// Abone telefon numarası
	//
	// zorunlu.
	GsmNumber string `json:"gsmNumber" validate:"required"`

	// Abone e-posta adresi
	//
	// zorunlu.
	Email string `json:"email" validate:"required,email"`

	// Abone soyadı
	//
	// zorunlu.
	Surname string `json:"surname" validate:"required"`

	// Abone adı
	//
	// zorunlu.
	Name string `json:"name" validate:"required"`

	// Abonelik başlangıç durumu (PENDING veya ACTIVE)
	//
	// zorunlu değil.
	SubscriptionInitialStatus string `json:"subscriptionInitialStatus" validate:"omitempty,oneof=PENDING ACTIVE"`

	// Abonelikte uygulanacak abonelik planı referans kodu.
	//
	// zorunlu.
	PricingPlanReferenceCode string `json:"pricingPlanReferenceCode" validate:"required"`
}

type InquirySubscriptionWithCheckoutFormRequest struct {
	// Checkout form için oluşturulan tekil değer. Her istek için özel üretilir ve işyerine dönülür.
	//
	// zorunlu
	Token string `json:"token" validate:"required"`
}

type InitSubscriptionWithNonTDSRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir
	//
	// zorunlu değil.
	Locale string `json:"locale,omitempty" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir. En yaygın kullanış biçimi üye iş yerinin sipariş numarasıdır
	//
	// zorunlu değil.
	ConversationID string `json:"conversationId,omitempty" validate:"omitempty"`

	// Eğer abonelik fiziksel bir ürün ise gönderilmelidir.
	//
	// zorunlu değil.
	ShippingAddress ShippingAddress `json:"shippingAddress,omitempty" validate:"omitempty"`

	// Fatura adresi bilgileri
	//
	// zorunlu.
	BillingAddress BillingAddress `json:"billingAddress" validate:"required,dive"`

	// Abone kimlik bilgileri
	//
	// zorunlu.
	IdentityNumber string `json:"identityNumber" validate:"required"`

	// Abone telefon numarası
	//
	// zorunlu.
	GsmNumber string `json:"gsmNumber" validate:"required"`

	// Abone email bilgisi
	//
	// zorunlu.
	Email string `json:"email" validate:"required,email"`

	// Abone soyadı
	//
	// zorunlu.
	Surname string `json:"surname" validate:"required"`

	// Abone adı
	//
	// zorunlu.
	Name string `json:"name" validate:"required"`

	// Abonelik başlangıç durumu (PENDING veya ACTIVE)
	//
	// zorunlu değil.
	SubscriptionInitialStatus string `json:"subscriptionInitialStatus" validate:"omitempty,oneof=PENDING ACTIVE"`

	// Abonelikte uygulanacak plan referans kodu.
	//
	// zorunlu.
	PricingPlanReferenceCode string `json:"pricingPlanReferenceCode" validate:"required"`

	// Abone kart bilgileri
	PaymentCard PaymentCard `json:"paymentCard" validate:"required,dive"`
}

type InitSubscriptionWithCheckoutFormResponse struct {
	// Yapılan isteğin sonucunu belirtir. Aboneliğin başarılı şekilde başlaması durumunda success değeri döner.
	Status string `json:"status"`

	// Checkout form için üretilmiş olan token değerinin geçerlilik süresi.
	SystemTime int64 `json:"systemTime"`

	// Checkout formun gösterilmesi için gerekli javascript html kodu.
	CheckoutFormContent string `json:"checkoutFormContent"`

	// Checkout form için oluşturulan tekil değer. Her istek için özel üretilir ve işyerine dönülür. Abonelik detaylarını öğrenmek için kullanılmalıdır.
	Token string `json:"token"`

	// Checkout form için üretilmiş olan token değerinin geçerlilik süresi.
	TokenExpireTime int64 `json:"tokenExpireTime"`
}

type SubscriptionPaymentResponse struct {
	// Yapılan isteğin sonucunu belirtir. İşlem başarılı ise success değeri döner.
	Status string `json:"status"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int64 `json:"systemTime"`

	Data CfSubscriptionData `json:"data"`
}

type CfSubscriptionData struct {
	// Abonelik için üretilen eşsiz referans kodu.
	ReferenceCode string `json:"referenceCode"`

	// Abonelik güncellemelerinde üye işeyerinin eşleştirme yapılabileceği eşsiz referans kodu.
	ParentReferenceCode string `json:"parentReferenceCode"`

	// Aboneliğe ait plan referans kodu.
	PricingPlanReferenceCode string `json:"pricingPlanReferenceCode"`

	// Müşteri bilgilerine istinaden oluşturulmuş eşsiz müşteri kodu. Bu kod email adresi baz alınarak oluşturulur ve müşteri işlemleri bu kod ile de yapılabilir.
	CustomerReferenceCode string `json:"customerReferenceCode"`

	// Abonelik durumunu gösterir. İstek esnasında gönderilmişse, sonuçta aynen geri iletilir. Gönderilmemişse ACTIVE değeri döner.
	SubscriptionStatus string `json:"subscriptionStatus"`

	// Ödeme planında belirlenen deneme süresidir. Bu süreç boyunca karttan ödeme alınmaz.
	TrialDays int `json:"trialDays"`

	// Deneme süresinin başlangıç tarihini gösteren unix timestamp değeridir.
	TrialStartDate int `json:"trialStartDate"`

	// Deneme süresinin bitiş tarihini gösteren unix timestamp değeridir.
	TrialEndDate int `json:"trialEndDate"`

	// Abonelik oluşturulma tarihinin unix timestamp değeridir.
	CreatedDate int `json:"createdDate"`

	// Abonelik başlangıç tarihinin unix timestamp değeridir.
	StartDate int `json:"startDate"`

	// Abonelik bitiş tarihinin unix timestamp değeridir.
	EndDate int `json:"endDate"`
}

func (r *InitSubscriptionWithCheckoutFormRequest) validate() error {
	validate := validator.New()

	validate.RegisterValidation("httpsurl", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		parsedURL, err := url.Parse(field)
		if err != nil {
			return false
		}
		return parsedURL.Scheme == "https"
	})

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

func (r *InquirySubscriptionWithCheckoutFormRequest) validate() error {
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

func (r *InitSubscriptionWithNonTDSRequest) validate() error {
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
