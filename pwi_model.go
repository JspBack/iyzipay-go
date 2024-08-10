package iyzipay

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type InitPWIRequest struct {
	// İstek yapılan dil bilgisidir. İki karakterli dil kodu olarak gönderilir. Örnek: tr, en
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir
	//
	// zorunlu değil
	ConversationID string `json:"conversationId" validate:"omitempty"`

	// Ödeme sepet tutarı. Kırılım tutarlar toplamı, sepet tutarına eşit olmalı.
	//
	// zorunlu
	Price string `json:"price" validate:"required,numeric"`

	// Üye işyeri sepet id’si.
	//
	// zorunlu değil
	BasketId string `json:"basketId" validate:"omitempty"`

	// Ödeme grubu, varsayılan PRODUCT. Geçerli değerler enum içinde sunulmaktadır: PRODUCT, LISTING, SUBSCRIPTION
	//
	// zorunlu değil
	PaymentGroup string `json:"paymentGroup" validate:"omitempty,oneof=PRODUCT LISTING SUBSCRIPTION"`

	// Ödeme akışında üye işyerine başarılı ve hatalı sonucu bildirmek üzere alınan URL adresi. Geçerli bir ssl sertifikasına sahip olmalıdır.
	//
	// zorunlu
	CallbackUrl string `json:"callbackUrl" validate:"required,httpsurl"`

	// Para birimi. Default değeri TL’dir. Kullanılabilen diğer değerler ise USD, EUR, GBP ve IRR’dir.
	//
	// zorunlu
	Currency string `json:"currency" validate:"required,oneof=TRY USD EUR GBP IRR"`

	// İndirim, vergi, taksit komisyonları gibi değerlerin dahil edildiği tutar.
	//
	// zorunlu
	PaidPrice string `json:"paidPrice" validate:"required,numeric"`

	// Taksit bilgisi, tek çekim için 1 gönderilmelidir. Geçerli değerler: 1, 2, 3, 6, 9.
	//
	// zorunlu değil
	EnabledInstallments []string `json:"enabledInstallments" validate:"omitempty,dive,oneof=1 2 3 6 9"`

	// Alıcı bilgileri
	Buyer Buyer `json:"buyer" validate:"required,dive"`

	// Kargo adresi (PHYSICAL product tipi seçilmişse zorunlu)
	ShippingAddress Address `json:"shippingAddress" validate:"omitempty,dive"`

	// Fatura adresi
	BillingAddress Address `json:"billingAddress" validate:"required,dive"`

	// Sepet içerikleri
	BasketItems []BasketItem `json:"basketItems" validate:"required,dive"`
}

type InitPWIResponse struct {
	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeridir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int64 `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationID string `json:"conversationId"`

	// iyzico ile ödeme için oluşturulan tekil değer. Her istek için özel üretilir ve işyerine dönülür.
	// Ödemenin sonucunu öğrenmek için zorunlu bir alandır. Varsayılan olarak 1800 saniye geçerliliği vardır.
	Token string `json:"token"`

	// iyzico ile ödeme için üretilmiş olan token ve link değerinin geçerlilik süresi.
	TokenExpireTime int64 `json:"tokenExpireTime"`

	// iyzico ile ödeme sayfasının eşsiz linki. Son kullanıcının ödeme yapması için yönlendirilmesi gereken adres.
	PayWithIyzicoPageUrl string `json:"payWithIyzicoPageUrl"`

	// İmza bilgisi.
	Signature string `json:"signature"`
}

type CheckPWIRequest struct {
	// İstek yapılan dil bilgisidir. İki karakterli dil kodu olarak gönderilir. Örnek: tr, en
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir
	//
	// zorunlu değil
	ConversationID string `json:"conversationId" validate:"omitempty"`

	// Checkout form için oluşturulan tekil değer. Her istek için özel üretilir ve işyerine dönülür. Ödemenin sonucunu öğrenmek için zorunlu bir alandır.
	//
	// zorunlu
	Token string `json:"token" validate:"required"`
}

type CheckPWIResponse Non3DSPaymentResponse

func (r *InitPWIRequest) validate() error {
	validate := validator.New()

	// https url kontrolü
	validate.RegisterValidation("httpsurl", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		parsedURL, err := url.Parse(field)
		if err != nil {
			return false
		}
		return parsedURL.Scheme == "https"
	})

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

func (r *CheckPWIRequest) validate() error {
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
