package iyzipay

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type IndividualSubMerchantUpdateRequest struct {
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
	SubMerchantExternalId string `json:"SubMerchantExternalId" validate:"required"`

	// Alt üye işyeri TC Kimlik Numarası.
	//
	// zorunlu
	IdentityNumber string `json:"identityNumber" validate:"required"`
}

type PrivateSubMerchantUpdateRequest struct {
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
	SubMerchantExternalId string `json:"SubMerchantExternalId" validate:"required"`

	// Alt üye işyeri TC Kimlik Numarası.
	//
	// zorunlu değil
	IdentityNumber string `json:"identityNumber" validate:"omitempty"`
}

type LimitedCompanySubMerchantUpdateRequest struct {
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
	SubMerchantExternalId string `json:"SubMerchantExternalId" validate:"required"`
}

type IndividualSubMerchantRequest struct {
	IndividualSubMerchantUpdateRequest
	// Bireysel için PERSONAL değeri gönderilmelidir.
	//
	// zorunlu
	SubMerchantType string `json:"SubMerchantType" validate:"required,oneof=PERSONAL"`
}

type PrivateSubMerchantRequest struct {
	PrivateSubMerchantUpdateRequest
	// Şahıs Şirketi için PRIVATE_COMPANY değeri gönderilmelidir.
	//
	// zorunlu
	SubMerchantType string `json:"SubMerchantType" validate:"required,oneof=PRIVATE_COMPANY"`
}

type LimitedCompanySubMerchantRequest struct {
	LimitedCompanySubMerchantUpdateRequest
	// Limited veya Anonim Şirket için LIMITED_OR_JOINT_STOCK_COMPANY değeri gönderilmelidir.
	//
	// zorunlu
	SubMerchantType string `json:"SubMerchantType" validate:"required,oneof=LIMITED_OR_JOINT_STOCK_COMPANY"`
}

type SubMerchantInquiryRequest struct {

	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir.
	//
	// zorunlu değil
	ConversationID string `json:"conversationId" validate:"omitempty"`

	// Alt üye işyeri tekil dış ID’si.
	//
	// zorunlu
	SubMerchantExternalId string `json:"subMerchantExternalId" validate:"required"`
}

type SubMerchantProductUpdateRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir
	//
	// zorunlu değil
	ConversationId string `json:"conversationId" validate:"omitempty"`

	// Ödeme kırılımına ait id.
	//
	// zorunlu
	PaymentTransactionId string `json:"paymentTransactionId" validate:"required"`

	// Alt üye işyerine IBAN adresine gönderilmesi istenen tutar.
	//
	// zorunlu
	SubMerchantPrice string `json:"subMerchantPrice" validate:"required,numeric"`

	// Alt üye işyeri oluşturma sorgusundan dönen değer.
	//
	// zorunlu
	SubMerchantKey string `json:"subMerchantKey" validate:"required"`
}

type SubMerchantResponse struct {
	UpdateSubMerchantResponse
	// Alt üye işyerini simgeleyen tekil değer.
	SubMerchantKey string `json:"SubMerchantKey"`
}

type UpdateSubMerchantResponse struct {
	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri tr’dir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int64 `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationId string `json:"conversationId"`
}

type SubMerchantInquiryResponse struct {

	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri tr’dir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int64 `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationId string `json:"conversationId"`

	// Alt üye işyeri adı.
	Name string `json:"name"`

	// Alt üye işyeri e-posta adresi.
	Email string `json:"email"`

	// Alt üye işyeri telefon numarası
	GsmNumber string `json:"gsmNumber"`

	// Alt üye işyeri adresi.s
	Address string `json:"address"`

	// Alt üye işyeri IBAN bilgisi. legalCompanyTitle ile uyumlu bir IBAN olmalı
	Iban string `json:"iban"`

	// Alt üye banka ülkesi
	BankCountry string `json:"bankCountry"`

	// Para birimi default TL olarak belirlenmiştir. USD, EUR, GBP, RUB, CHF ve NOK para birimleri de belirlenebilir
	Currency string `json:"currency"`

	// Alt üye işyeri vergi dairesi.
	TaxOffice string `json:"taxOffice"`

	// Alt üye işyeri yasal şirket ünvanı
	LegalCompanyTitle string `json:"legalCompanyTitle"`

	// Alt üye işyeri yasal şirket ünvanı
	SubMerchantExternalId string `json:"subMerchantExternalId"`

	// Alt üye işyeri TC Kimlik Numarası
	IdentityNumber string `json:"identityNumber"`

	// Alt üye tipi.
	SubMerchantType string `json:"subMerchantType"`

	// Alt üye işyerini simgeleyen tekil değer
	SubMerchantKey string `json:"subMerchantKey"`
}

type SubMerchantProductUpdateResponse struct {
	// İşlem sonucu başarılı ise success, başarısız ise failure döner.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri tr’dir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int64 `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationId string `json:"conversationId"`

	// Alt üye işyeri adı.
	ItemId string `json:"itemId"`

	// Ödeme kırılımına ait id, üye işyeri tarafından mutlaka saklanmalıdır.
	// Ödeme kırılımının iadesi, onayı, onay geri çekmesi ve iyzico ile iletişimde kullanılır. Tercihen itemId ile ilişkili bir şekilde tutulmalıdır.
	PaymentTransactionId string `json:"paymentTransactionId"`

	// Ödeme kırılımının durumu. Ödeme fraud kontrolünde ise 0 değeri döner, bu durumda fraudStatus değeri de 0’dır.
	// Ödeme, fraud kontrolünden sonra reddedilirse -1 döner. Pazaryeri modelinde ürüne onay verilene dek bu değer 1 olarak döner.
	// Pazaryeri modelinde ürüne onay verilmişse bu değer 2 olur. Geçerli değerler: 0, -1, 1 ve 2.
	TransactionStatus int `json:"transactionStatus"`

	// Toplem tutar
	Price float64 `json:"price"`

	// Müşterinin ödediği tutar
	PaidPrice float64 `json:"paidPrice"`

	// Üye işyerinin uyguladığı vade/komisyon oranı. Örneğin price=100, paidPrice=110 ise üye işyeri vade/komisyon oranı %10’dur. Bilgi amaçlıdır.
	MerchantCommissionRate float64 `json:"merchantCommissionRate"`

	// Üye işyerinin uyguladığı vade/komisyon tutarı. Örneğin price=100, paidPrice=110 ise üye işyeri vade/komisyon tutarı 10’dur. Bilgi amaçlıdır.
	MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`

	// Ödemeye ait iyzico işlem komisyon tutarı.
	IyziCommissionRateAmount float64 `json:"iyziCommissionRateAmount"`

	// Ödemeye ait iyzico işlem ücreti
	IyziCommissionFee float64 `json:"iyziCommissionFee"`

	// Kırılım bazında üye işyeri blokaj oranı. iyzico – üye işyeri anlaşmasına göre, üye işyerine işlem bazında blokaj uygulayabilir.
	// Bu blokaj üye işyeri fraud riskini önlemek içindir, blokaj süresi boyunca para iyzico’da tutulur, bu süre sonrasında üye işyerine gönderilir.
	BlockageRate float64 `json:"blockageRate"`

	// Kırılım bazında üye işyeri blokaj tutarının, üye işyerine yansıyan rakamı. Blokaj tutarı mümkün olduğunca üye işyerine yansıtılır.
	// Eğer blokaj tutarı, üye işyeri tutarından daha büyükse bu durumda alt üye işyerine de yansıtılır.
	BlockageRateAmountMerchant float64 `json:"blockageRateAmountMerchant"`

	// Dökümantasyonda belirtilmemiş
	BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`

	// İşlem bazında blokaj çözülme tarihi. yyyy-MM-dd HH:mm:ss formatındadır, örneğin 2015-10-19 14:36:52
	BlockageResolvedDate string `json:"blockageResolvedDate"`

	// Alt üye işyeri oluşturma sorgusundan dönen değer.
	SubMerchantKey string `json:"subMerchantKey"`

	// Alt üye işyerine IBAN adresine gönderilmiş tutar.
	SubMerchantPrice float64 `json:"subMerchantPrice"`

	// Dökümantasyonda belirtilmemiş
	SubMerchantPayoutRate float64 `json:"subMerchantPayoutRate"`

	// Dökümantasyonda belirtilmemiş
	SubMerchantPayoutAmount float64 `json:"subMerchantPayoutAmount"`

	// Bu kırılım için, iyzico işlem ücreti, komisyon tutarı ve blokajlar düşüldükten sonra üye işyerine gönderilecek tutar.
	MerchantPayoutAmount float64 `json:"merchantPayoutAmount"`

	ConvertedPayout ConvertedPayout `json:"convertedPayout"`
}

func (r *IndividualSubMerchantRequest) validate() error {
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

func (r *PrivateSubMerchantRequest) validate() error {
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

func (r *LimitedCompanySubMerchantRequest) validate() error {
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

func (r *IndividualSubMerchantUpdateRequest) validate() error {
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

func (r *PrivateSubMerchantUpdateRequest) validate() error {
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

func (r *LimitedCompanySubMerchantUpdateRequest) validate() error {
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

func (r *SubMerchantInquiryRequest) validate() error {
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

func (r *SubMerchantProductUpdateRequest) validate() error {
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
