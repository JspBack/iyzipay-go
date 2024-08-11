package iyzipay

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type CFRequest struct {
	InitPWIRequest `validate:"dive"`

	// Ödeme kanalı. Geçerli değerler SHOPIFY, MAGENTO, PRESTASHOP, WOOCOMMERCE, OPENCART
	//
	// zorunlu değil.
	PaymentSource string `json:"paymentSource" validate:"omitempty,oneof=SHOPIFY MAGENTO PRESTASHOP WOOCOMMERCE OPENCART"`
}

type CFResponse struct {
	InitPWIResponse

	// Ödeme sayfası için dönen script tagi
	CheckoutFormContent string `json:"checkoutFormContent"`

	// Ödeme sayfası url bilgisi
	PaymentPageUrl string `json:"paymentPageUrl"`
}

type CFInquiryRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir
	//
	// zorunlu değil.
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// Checkout form için oluşturulan tekil değer. Her istek için özel üretilir ve işyerine dönülür. Ödemenin sonucunu öğrenmek için zorunlu bir alandır.
	//
	// zorunlu
	Token string `json:"token" validate:"required"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir
	//
	// zorunlu değil.
	ConversationId string `json:"conversationId" validate:"omitempty"`
}

type CFInquiryResponse struct {
	Non3DSPaymentResponse

	//Bankadan dönen değerdir. Sadece ödeme başarısız ise ve işlem 3ds ile yapılmışsa bu değer döner. 0,2,3,4,5,6,7 değerlerini alabilir.
	MdStatus int `json:"mdStatus"`

	// Checkout form için oluşturulan tekil değer. Her istek için özel üretilir ve işyerine dönülür. Ödemenin sonucunu öğrenmek için zorunlu bir alandır.
	Token string `json:"token"`

	// Ödeme isteğinin durumunu gösterir. Success ise karttan ilgili tutar çekilmiştir. SUCCESS, FAILURE, INIT_THREEDS, CALLBACK_THREEDS, BKM_POS_SELECTED, CALLBACK_PECCO
	PaymentStatus string `json:"paymentStatus"`

	// Ödeme akışında üye işyerine başarılı ve hatalı sonucu bildirmek üzere alınan URL adresi.
	CallbackUrl string `json:"callbackUrl"`
}

func (r *CFRequest) validate() error {
	validate := validator.New()

	// Tarih formatı doğrulaması
	validate.RegisterValidation("iyziDate", func(fl validator.FieldLevel) bool {
		dateFormat := "2006-01-02 15:04:05"
		dateStr := fl.Field().String()
		if dateStr == "" {
			return true
		}
		_, err := time.Parse(dateFormat, dateStr)
		return err == nil
	})

	// https url kontrolü
	validate.RegisterValidation("httpsurl", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		parsedURL, err := url.Parse(field)
		if err != nil {
			return false
		}
		return parsedURL.Scheme == "https"
	})

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

func (r *CFInquiryRequest) validate() error {
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
