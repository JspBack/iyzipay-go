package iyzipay

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type InitTDSRequest struct {
	PaymentRequest `validate:"dive"`

	// 3D Secure işlemi başlatıldığında, banka tarafından kullanıcıya gösterilecek sayfa URL'si.(https olmalıdır.)
	//
	// zorunlu.
	CallbackUrl string `json:"callbackUrl" validate:"required,httpsurl"`
}

type InitTDSResponse struct {
	// Servis yanıt sonucu (başarılı/başarısız)
	Status string `json:"status"`

	// Dil bilgisi
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int64 `json:"systemTime"`

	// Konuşma Id'si
	ConversationID string `json:"conversationId"`

	// 3D Secure işlem için html içeriği (base64 şifreli)
	ThreeDSHtmlContent string `json:"threeDSHtmlContent"`

	// İstek sonrası dönen İşlem Id'si
	PaymentID string `json:"paymentId"`

	// İstek sonrası dönen imza
	Signature string `json:"signature"`
}

type TDSPaymentRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir
	//
	// zorunlu değil.
	Locale string `json:"locale" validate:"omitempty"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir. En yaygın kullanış biçimi üye iş yerinin sipariş numarasıdır
	//
	// zorunlu değil.
	ConversationId string `json:"conversationId" validate:"omitempty"`

	// 3D dönüş bağlantı adresinizden alacağınız paymentid değeridir.
	//
	// zorunlu.
	PaymentId string `json:"paymentId" validate:"required"`

	// 3DInit isteği oluştururken kullandığınız conversationId değeridir.
	//
	// zorunlu.
	PaymentConversationId string `json:"paymentConversationId" validate:"required"`
}

type TDSPaymentResponse struct {
	Non3DSPaymentResponse

	//Bankadan dönen değerdir. Sadece ödeme başarısız ise ve işlem 3ds ile yapılmışsa bu değer döner. 0,2,3,4,5,6,7 değerlerini alabilir.
	MdStatus int `json:"mdStatus"`
}

func (r *InitTDSRequest) validate() error {
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

	// Luhn algoritması ile kredi kartı numarası doğrulaması (https://en.wikipedia.org/wiki/Luhn_algorithm)
	validate.RegisterValidation("creditcard", func(fl validator.FieldLevel) bool {
		cardNumber := fl.Field().String()
		cleanNumber := ""
		for _, ch := range cardNumber {
			if ch >= '0' && ch <= '9' {
				cleanNumber += string(ch)
			}
		}

		if len(cleanNumber) < 13 || len(cleanNumber) > 19 {
			return false
		}

		sum := 0
		shouldDouble := false
		for i := len(cleanNumber) - 1; i >= 0; i-- {
			digit := int(cleanNumber[i] - '0')
			if shouldDouble {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}
			sum += digit
			shouldDouble = !shouldDouble
		}

		return sum%10 == 0
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

func (r TDSPaymentRequest) validate() error {
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
